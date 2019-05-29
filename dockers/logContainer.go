package dockers

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/*
打印日志
***/
func LogContainer(ctx context.Context, cli *client.Client, id string) {
	//将容器的标准输出显示出来
	out, err := cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	//容器内部的运行状态
	status, err := cli.ContainerStats(ctx, id, true)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, status.Body)
}

/*
打印运行日志
***/
func LogStatContainer(ctx context.Context, cli *client.Client, id string) {
	//容器内部的运行状态
	status, err := cli.ContainerStats(ctx, id, true)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, status.Body)
}
