package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
)

func main() {
	cid, options := ParseFlags()

	err := run(cid, options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(id string, options Options) error {
	container, err := DockerInspect(id)
	if err != nil {
		return err
	}

	output := Fork(container, options)
	fmt.Println(output)
	return nil
}

func DockerInspect(cid string) (types.ContainerJSON, error) {
	cmd := exec.Command("docker", "inspect", cid)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err)
		os.Exit(2)
	}
	return unmarshalContainer(out)
}

func unmarshalContainer(raw []byte) (types.ContainerJSON, error) {
	var containers []types.ContainerJSON
	err := json.Unmarshal(raw, &containers)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	if len(containers) == 0 {
		return types.ContainerJSON{}, fmt.Errorf("inspect data is empty: %v", string(raw))
	}
	return containers[0], nil
}
