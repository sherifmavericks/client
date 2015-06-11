package client

import (
	"fmt"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/engine"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/protocol/go"
	"github.com/maxtaco/go-framed-msgpack-rpc/rpc2"
)

// CmdDeviceAdd is the 'device add' command.  It is used for
// device provisioning to enter a secret phrase on an existing
// device.
type CmdDeviceAdd struct {
	phrase    string
	sessionID int
}

// NewCmdDeviceAdd creates a new cli.Command.
func NewCmdDeviceAdd(cl *libcmdline.CommandLine) cli.Command {
	return cli.Command{
		Name:        "add",
		Usage:       "keybase device add \"secret phrase\"",
		Description: "Authorize a new device",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdDeviceAdd{}, "add", c)
		},
	}
}

// RunClient runs the command in client/server mode.
func (c *CmdDeviceAdd) RunClient() error {
	var err error
	c.sessionID, err = libkb.RandInt()
	if err != nil {
		return err
	}
	cli, err := GetDeviceClient()
	if err != nil {
		return err
	}
	protocols := []rpc2.Protocol{
		NewSecretUIProtocol(),
		NewLocksmithUIProtocol(),
	}
	if err := RegisterProtocols(protocols); err != nil {
		return err
	}

	return cli.DeviceAdd(keybase1.DeviceAddArg{SecretPhrase: c.phrase, SessionID: c.sessionID})
}

// Run runs the command in standalone mode.
func (c *CmdDeviceAdd) Run() error {
	ctx := &engine.Context{SecretUI: GlobUI.GetSecretUI(), LocksmithUI: GlobUI.GetLocksmithUI()}
	eng := engine.NewKexSib(G, c.phrase)
	return engine.RunEngine(eng, ctx)
}

// ParseArgv gets the secret phrase from the command args.
func (c *CmdDeviceAdd) ParseArgv(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		return fmt.Errorf("device add takes one arg: the secret phrase")
	}
	c.phrase = ctx.Args()[0]
	return nil
}

// GetUsage says what this command needs to operate.
func (c *CmdDeviceAdd) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:    true,
		KbKeyring: true,
		API:       true,
	}
}
