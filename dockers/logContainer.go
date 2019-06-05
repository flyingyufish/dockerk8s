package dockers

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
)

/*
打印日志
***/
func LogContainer(id string) error {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return err
	}

	//将容器的标准输出显示出来
	out, err := cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true, Tail: "5"})
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	defer out.Close()
	content, err := ioutil.ReadAll(out)
	if err != nil {
		return err
	}
	fmt.Println(content)
	io.Copy(os.Stdout, out)

	return nil
}

/*
打印运行日志
***/
func StatContainer(id string) error {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return err
	}

	//容器内部的运行状态
	status, err := cli.ContainerStats(ctx, id, true)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, status.Body)

	return nil
}
