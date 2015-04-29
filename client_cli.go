package client

import (
	"fmt"
	"github.com/coreos/fleet/schema"
	"github.com/juju/errgo"
	execPkg "os/exec"
	"strings"
)

const (
	FLEETCTL        = "fleetctl"
	ENDPOINT_OPTION = "--endpoint"
	ENDPOINT_VALUE  = "http://172.17.42.1:4001"
)

type ClientCLI struct {
	etcdPeer string
	driver   string
}

func NewClientCLI() FleetClient {
	return NewClientCLIWithPeer(ENDPOINT_VALUE)
}

func NewClientCLIWithPeer(etcdPeer string) FleetClient {
	driver := ""
	cmd := execPkg.Command(FLEETCTL, "--version")
	output, err := exec(cmd)
	if err != nil {
		return nil
	}

	if strings.Contains(output, "0.10") {
		fmt.Printf("Adding driver option for version 0.10\n")
		driver = "--driver=etcd"
	} else {
		fmt.Printf("Not adding driver option: %s\n", output)
	}

	return &ClientCLI{
		etcdPeer: etcdPeer,
		driver:   driver,
	}
}

func args(extras []string, required ...string) []string {
	return append(required, extras...)
}

func (this *ClientCLI) Submit(filePath ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(filePath, this.driver, ENDPOINT_OPTION, this.etcdPeer, "submit")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(filePath, ENDPOINT_OPTION, this.etcdPeer, "submit")...)
	}
	output, err := exec(cmd)

	if err != nil {
		fmt.Printf("Error in submit: %s %s\n", err, output)
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Unit(name string) (*schema.Unit, error) {
	return nil, fmt.Errorf("Method not implemented: ClientCLI.Unit")
}

func (this *ClientCLI) Start(name ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, "start", "--no-block=true")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(name, ENDPOINT_OPTION, this.etcdPeer, "start", "--no-block=true")...)
	}
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Stop(name ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, "stop", "--no-block=true")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(name, ENDPOINT_OPTION, this.etcdPeer, "stop", "--no-block=true")...)
	}
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Load(name ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, "load", "--no-block=true")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(name, ENDPOINT_OPTION, this.etcdPeer, "load", "--no-block=true")...)
	}
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Destroy(name ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, "destroy")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(name, ENDPOINT_OPTION, this.etcdPeer, "destroy")...)
	}
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
