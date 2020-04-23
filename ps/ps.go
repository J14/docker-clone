package ps

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Ps(all bool) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%-20s%-20s%-20s%-20s%-20s%-20s%s\t\n",
		"CONTAINER ID", "IMAGE", "COMMAND", "CREATED",
		"STATUS", "PORTS", "NAMES")
	for _, container := range containers {
		id := container.ID[:12]
		image := container.Image
		var command string
		if len(container.Command) > 19 {
			command = fmt.Sprintf("\"%s...\"", container.Command[:19])
		} else {
			command = container.Command
		}
		created := container.Created
		status := container.Status
		var port string
		for _, p := range container.Ports {
			port = fmt.Sprintf("%s:%d->%d/%s",
				p.IP, p.PrivatePort, p.PublicPort, p.Type)
		}
		name := container.Names[0][1:]

		fmt.Printf("%-20s%-20s%-20s%-20d%-20s%-20s%-20s\n",
			id, image, command, created, status, port, name)
	}
}