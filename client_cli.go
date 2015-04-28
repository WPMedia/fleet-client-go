package client

import (
	"github.com/coreos/fleet/schema"
	"github.com/juju/errgo"

	"fmt"
	execPkg "os/exec"
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
	return NewClientCLIWithPeer(ENDPOINT_VALUE, "")
}

func NewClientCLIWithPeer(etcdPeer, driver string) FleetClient {
	if driver != "" {
		driver = fmt.Sprintf(" --driver=%s ", driver)
	}
	return &ClientCLI{
		etcdPeer: etcdPeer,
		driver:   driver,
	}
}

func (this *ClientCLI) Submit(name, filePath string) error {
	cmd := execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, "submit", filePath)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Unit(name string) (*schema.Unit, error) {
	return nil, fmt.Errorf("Method not implemented: ClientCLI.Unit")
}

func (this *ClientCLI) Start(name string) error {
	cmd := execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, "start", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Stop(name string) error {
	cmd := execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, "stop", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Load(name string) error {
	cmd := execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, "load", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Destroy(name string) error {
	cmd := execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, "destroy", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
