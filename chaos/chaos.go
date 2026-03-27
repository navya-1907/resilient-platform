package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cli := getDockerClient()

	for {
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if len(containers) > 0 {
			target := containers[rand.Intn(len(containers))]

			fmt.Println("Killing container:", target.Names[0])

			err := cli.ContainerKill(context.Background(), target.ID, "SIGKILL")
			if err != nil {
				fmt.Println("Kill error:", err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func getDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}
