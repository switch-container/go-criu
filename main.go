package criu

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/checkpoint-restore/go-criu/rpc"
	"github.com/golang/protobuf/proto"
)

// Criu struct
type Criu struct {
	swrkCmd *exec.Cmd
	swrkSk  *os.File
}

// MakeCriu returns the Criu object required for most operations
func MakeCriu() *Criu {
	return &Criu{}
}

// Prepare sets up everything for the RPC communication to CRIU
func (c *Criu) Prepare() error {
	fds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_SEQPACKET, 0)
	if err != nil {
		return err
	}

	cln := os.NewFile(uintptr(fds[0]), "criu-xprt-cln")
	syscall.CloseOnExec(fds[0])
	srv := os.NewFile(uintptr(fds[1]), "criu-xprt-srv")
	defer srv.Close()

	args := []string{"swrk", strconv.Itoa(fds[1])}
	cmd := exec.Command("criu", args...)

	err = cmd.Start()
	if err != nil {
		cln.Close()
		return err
	}

	c.swrkCmd = cmd
	c.swrkSk = cln

	return nil
}

// Cleanup cleans up
func (c *Criu) Cleanup() {
	if c.swrkCmd != nil {
		c.swrkSk.Close()
		c.swrkSk = nil
		c.swrkCmd.Wait()
		c.swrkCmd = nil
	}
}

func (c *Criu) sendAndRecv(reqB []byte) ([]byte, int, error) {
	cln := c.swrkSk
	_, err := cln.Write(reqB)
	if err != nil {
		return nil, 0, err
	}

	respB := make([]byte, 2*4096)
	n, err := cln.Read(respB)
	if err != nil {
		return nil, 0, err
	}

	return respB, n, nil
}

func (c *Criu) doSwrk(reqType rpc.CriuReqType, opts *rpc.CriuOpts, nfy Notify) error {
	resp, err := c.doSwrkWithResp(reqType, opts, nfy)
	if err != nil {
		return err
	}
	respType := resp.GetType()
	if respType != reqType {
		return errors.New("unexpected responce")
	}

	return nil
}

func (c *Criu) doSwrkWithResp(reqType rpc.CriuReqType, opts *rpc.CriuOpts, nfy Notify) (*rpc.CriuResp, error) {
	var resp *rpc.CriuResp

	req := rpc.CriuReq{
		Type: &reqType,
		Opts: opts,
	}

	if nfy != nil {
		opts.NotifyScripts = proto.Bool(true)
	}

	if c.swrkCmd == nil {
		err := c.Prepare()
		if err != nil {
			return nil, err
		}

		defer c.Cleanup()
	}

	for {
		reqB, err := proto.Marshal(&req)
		if err != nil {
			return nil, err
		}

		respB, respS, err := c.sendAndRecv(reqB)
		if err != nil {
			return nil, err
		}

		resp = &rpc.CriuResp{}
		err = proto.Unmarshal(respB[:respS], resp)
		if err != nil {
			return nil, err
		}

		if !resp.GetSuccess() {
			return resp, fmt.Errorf("operation failed (msg:%s err:%d)",
				resp.GetCrErrmsg(), resp.GetCrErrno())
		}

		respType := resp.GetType()
		if respType != rpc.CriuReqType_NOTIFY {
			break
		}
		if nfy == nil {
			return resp, errors.New("unexpected notify")
		}

		notify := resp.GetNotify()
		switch notify.GetScript() {
		case "pre-dump":
			err = nfy.PreDump()
		case "post-dump":
			err = nfy.PostDump()
		case "pre-restore":
			err = nfy.PreRestore()
		case "post-restore":
			err = nfy.PostRestore(notify.GetPid())
		case "network-lock":
			err = nfy.NetworkLock()
		case "network-unlock":
			err = nfy.NetworkUnlock()
		case "setup-namespaces":
			err = nfy.SetupNamespaces(notify.GetPid())
		case "post-setup-namespaces":
			err = nfy.PostSetupNamespaces()
		case "post-resume":
			err = nfy.PostResume()
		default:
			err = nil
		}

		if err != nil {
			return resp, err
		}

		req = rpc.CriuReq{
			Type:          &respType,
			NotifySuccess: proto.Bool(true),
		}
	}

	return resp, nil
}

// Dump dumps a process
func (c *Criu) Dump(opts rpc.CriuOpts, nfy Notify) error {
	return c.doSwrk(rpc.CriuReqType_DUMP, &opts, nfy)
}

// Restore restores a process
func (c *Criu) Restore(opts rpc.CriuOpts, nfy Notify) error {
	return c.doSwrk(rpc.CriuReqType_RESTORE, &opts, nfy)
}

// PreDump does a pre-dump
func (c *Criu) PreDump(opts rpc.CriuOpts, nfy Notify) error {
	return c.doSwrk(rpc.CriuReqType_PRE_DUMP, &opts, nfy)
}

// StartPageServer starts the page server
func (c *Criu) StartPageServer(opts rpc.CriuOpts) error {
	return c.doSwrk(rpc.CriuReqType_PAGE_SERVER, &opts, nil)
}

// StartPageServerChld starts the page server and returns PID and port
func (c *Criu) StartPageServerChld(opts rpc.CriuOpts) (int, int, error) {
	resp, err := c.doSwrkWithResp(rpc.CriuReqType_PAGE_SERVER_CHLD, &opts, nil)
	if err != nil {
		return 0, 0, err
	}

	return int(resp.Ps.GetPid()), int(resp.Ps.GetPort()), nil
}
