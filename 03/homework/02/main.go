package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Port struct {

	// Host IP address that the container's port is mapped to
	IP string `json:"IP,omitempty"`

	// Port on the container
	// Required: true
	PrivatePort uint16 `json:"PrivatePort"`

	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`

	// type
	// Required: true
	Type string `json:"Type"`
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	cli.ContainerList(context.Background(), types.ContainerListOptions{})

	fmt.Printf("%-15s %-20s %-30s %-40s %-10s\n", "CONTAINER ID", "NAME", "IMAGE", "PORTS", "STATUS")
	for _, container := range containers {
		parsedContainerPorts := containerPorts2String(container.Ports)
		parsedContainerPorts = parsedContainerPorts[:len(parsedContainerPorts)-1]
		fmt.Printf("%-15s %-20s %-30s %-40s %-10s\n", container.ID[:10], container.Names[0][1:], container.Image, parsedContainerPorts, container.Status)
	}
}

func containerPorts2String(ports []types.Port) (parsedString string) {
	parsedString = ""
	for _, port := range ports {
		parsedString += fmt.Sprintf("%s:%d->%d/%s,", port.IP, port.PrivatePort, port.PublicPort, port.Type)
	}
	return
}
