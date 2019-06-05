package dockers

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
)

/*
启动容器
**/

func StartContainer(containerID string) error {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err == nil {
		log.Printf("success start container:%s\n", containerID)
	} else {
		log.Printf("failed to start container:%s!!!!!!!!!!!!!\n", containerID)
	}
	return err
}
