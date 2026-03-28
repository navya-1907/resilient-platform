package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	http.HandleFunc("/containers", getContainers)

	go monitorAndHeal()

	fmt.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// API: list containers
func getContainers(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	cli := getDockerClient()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, c := range containers {
		fmt.Fprintf(w, "%s - %s\n", c.Names[0], c.State)
	}
}

// Self-healing loop
func monitorAndHeal() {
	for {
		cli := getDockerClient()

		containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			log.Println("Error fetching containers:", err)
			continue
		}

		for _, c := range containers {
			if c.State != "running" {
				log.Println("Restarting:", c.Names[0])
				restartContainer(cli, c.ID)
			}
		}

		time.Sleep(5 * time.Second)
	}
}

// restart logic
func restartContainer(cli *client.Client, id string) {
	timeout := 5
	err := cli.ContainerRestart(context.Background(), id, container.StopOptions{
		Timeout: &timeout,
	})

	if err != nil {
		log.Println("Restart failed:", err)
	}
}

// docker client
func getDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
