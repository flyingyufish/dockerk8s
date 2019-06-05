package dockers

import (
	"context"
	"log"
	"time"
)

/*
停止容器
**/
func StopContainer(containerID string) error {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return err
	}
	timeout := 5 * time.Second
	err = cli.ContainerStop(ctx, containerID, &timeout)
	if err == nil {
		log.Printf("success stop container:%s\n", containerID)
	} else {
		log.Printf("failed to stop container:%s!!!!!!!!!!!!!\n", containerID)
	}
	return err
}
