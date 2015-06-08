package client

// Parse for `fleetctl list-machines` output
import (
	"bufio"
	"strings"

	execPkg "os/exec"
)

type Machine struct {
	Name string
	IP   string
}

func (this *ClientCLI) Machines() ([]Machine, error) {
	var cmd *execPkg.Cmd

	if this.driver != "" {
		cmd = execPkg.Command(FLEETCTL, this.driver, ENDPOINT_OPTION, this.etcdPeer, ETCD_PREFIX_OPTION, this.etcdPrefix, "list-machines", "--full", "--fields=machine,ip", "--no-legend")
	} else {
		cmd = execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "list-machines", "--full", "--fields=machine,ip", "--no-legend")
	}
	stdout, err := exec(cmd)
	if err != nil {
		return []Machine{}, err
	}

	return parseMachineOutput(stdout)
}

func parseMachineOutput(output string) ([]Machine, error) {
	result := make([]Machine, 0)

	scanner := bufio.NewScanner(strings.NewReader(output))
	// Scan each line of input.
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		words := filterEmpty(strings.Split(line, "\t"))
		machine := Machine{
			Name: words[0],
			IP:   words[1],
		}
		result = append(result, machine)
	}

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		return result, scanner.Err()
	}
	return result, nil
}
