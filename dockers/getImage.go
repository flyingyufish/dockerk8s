package dockers

import (
	"context"

	"github.com/docker/docker/api/types"
)

/*
镜像列表
**/

func GetImage(all bool) ([]types.ImageSummary, error) {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}
	return cli.ImageList(ctx, types.ImageListOptions{All: all})
}
