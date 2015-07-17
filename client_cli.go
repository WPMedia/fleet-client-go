package client

import (
	"fmt"
	"github.com/coreos/fleet/schema"
	"github.com/juju/errgo"
	execPkg "os/exec"
	"strings"
)

const (
	FLEETCTL            = "fleetctl"
	ENDPOINT_OPTION     = "--endpoint"
	ENDPOINT_VALUE      = "http://172.17.42.1:4001"
	ETCD_PREFIX_OPTION  = "--etcd-key-prefix"
	DEFAULT_ETCD_PREFIX = "/_coreos.com/fleet/"
)

type ClientCLI struct {
	etcdPeer   string
	driver     string
	etcdPrefix string
}

func NewClientCLI() FleetClient {
	return NewClientCLIWithPeer(ENDPOINT_VALUE)
}

func NewClientCLIWithPeer(etcdPeer string) FleetClient {
	return &ClientCLI{
		etcdPeer:   etcdPeer,
		driver:     "--driver=etcd",
		etcdPrefix: DEFAULT_ETCD_PREFIX,
	}
}

func NewClientCLIWithPeerAndPrefix(etcdPeer, etcdPrefix string) FleetClient {
	client := NewClientCLIWithPeer(etcdPeer)
	if etcdPrefix == "" {
		etcdPrefix = DEFAULT_ETCD_PREFIX
	}
	return &ClientCLI{
		etcdPeer:   etcdPeer,
		driver:     getDriver(),
		etcdPrefix: etcdPrefix,
	}

	return client
}

func args(extras []string, required ...string) []string {
	return append(required, extras...)
}

func (this *ClientCLI) Submit(filePath ...string) error {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, args(filePath, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "submit")...)
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
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "start", "--no-block=true")...)
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
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "stop", "--no-block=true")...)
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
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "load", "--no-block=true")...)
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
		cmd = execPkg.Command(FLEETCTL, args(name, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "destroy")...)
	} else {
		cmd = execPkg.Command(FLEETCTL, args(name, ENDPOINT_OPTION, this.etcdPeer, "destroy")...)
	}
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
