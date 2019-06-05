package dockers

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

/*
获取容器
**/
func ContainerList(all bool, name string) ([]types.Container, error) {
	filter := filters.NewArgs()
	filter.Add("name", fmt.Sprintf(".%s.*", name))
	return getContainer(all, filter)

}

func getContainer(all bool, filter filters.Args) ([]types.Container, error) {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all, Filters: filter})
	if err != nil {
		return nil, err
	}

	return containerList, nil
}
