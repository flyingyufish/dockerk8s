package dockers

import (
	"context"

	"github.com/docker/docker/api/types"
)

/*
获取容器
**/
func ContainerList(all bool) ([]types.Container, error) {
	return getContainer(all)
}

func getContainer(all bool) ([]types.Container, error) {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all})
	if err != nil {
		return nil, err
	}

	return containerList, nil
}
