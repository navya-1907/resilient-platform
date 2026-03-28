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
		// get running containers
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			fmt.Println("Error getting containers:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// filter only worker containers
		var workers []container.Summary
		for _, c := range containers {
			if c.Names[0] == "/worker1" || c.Names[0] == "/worker2" {
				workers = append(workers, c)
			}
		}

		// if workers exist, randomly kill one
		if len(workers) > 0 {
			target := workers[rand.Intn(len(workers))]

			fmt.Println("Killing container:", target.Names[0])

			err := cli.ContainerKill(context.Background(), target.ID, "SIGKILL")
			if err != nil {
				fmt.Println("Kill error:", err)
			}
		}

		// wait before next chaos action
		time.Sleep(10 * time.Second)
	}
}

func getDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	return cli
}
