package dockers

import (
	"context"
	"errors"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
)

/*
创建容器
**/

/*
imageName     string   = "my-gin:latest"                      //镜像名称
containerName string   = "mygin-latest"                       //容器名称
indexName     string   = "/" + containerName                  //容器索引名称，用于检查该容器是否存在是使用
networkname   string   = "example_bytn"                       //容器网络
cmd           []string = []string{"./ginDocker2"}                       //运行的cmd命令，用于启动container中的程序
env           []string = []string{"name=test"}                       //容器环境变量
workDir       string   = "/go/src/ginDocker2"                 //container工作目录
hostPort      []string = []string{"7070"}                               //container映射到宿主机的端口
containerDir  []string = []string{"/go/src/ginDocker2"}                 //容器挂在目录
hostDir       []string = []string{"/home/youngblood/Go/src/ginDocker2"} //容器挂在到宿主机的目录
**/
func CreateContainer(imageName string, containerName string, networkname string, cmd string, env []string, workDir string, hostPort []string, containerDir []string, hostDir []string) (string, error) {

	if len(containerDir) != len(hostDir) {
		return "", errors.New("containerDir is not length hostDir")
	}

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return "", err
	}

	exposePort := make(map[nat.Port]struct{})
	bindPort := make(map[nat.Port][]nat.PortBinding)
	mounts := make([]mount.Mount, len(containerDir))

	for i := 0; i < len(hostPort); i++ {
		exposePort[nat.Port(hostPort[i])] = struct{}{}

		bindPort[nat.Port(hostPort[i])] = []nat.PortBinding{nat.PortBinding{
			HostIP:   "0.0.0.0",   //docker容器映射的宿主机的ip
			HostPort: hostPort[i], //docker 容器映射到宿主机的端口
		}}
	}

	for i := 0; i < len(containerDir); i++ {
		mounts[i] = mount.Mount{
			Type:   mount.TypeBind,
			Source: hostDir[i],
			Target: containerDir[i],
		}
	}

	//创建容器
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,     //镜像名称
		Tty:          true,          //docker run命令中的-t选项
		OpenStdin:    true,          //docker run命令中的-i选项
		Cmd:          []string{cmd}, //docker 容器中执行的命令
		Env:          env,           //环境变量
		WorkingDir:   workDir,       //docker容器中的工作目录
		ExposedPorts: exposePort,    //docker容器对外开放的端口

	}, &container.HostConfig{
		PortBindings: bindPort,
		Mounts:       mounts, //docker 容器目录挂在到宿主机目录
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkname: &network.EndpointSettings{},
		},
	}, containerName)
	if err == nil {
		log.Printf("success create container:%s\n", resp.ID)
	} else {
		log.Println("failed to create container!!!!!!!!!!!!!")
		return "", err
	}

	return resp.ID, nil
}

//创建运行容器
func RunContainer(imageName string, containerName string, networkname string, cmd string, env []string, workDir string, hostPort []string, containerDir []string, hostDir []string) (string, error) {

	if len(containerDir) != len(hostDir) {
		return "", errors.New("containerDir is not length hostDir")
	}

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return "", err
	}

	exposePort := make(map[nat.Port]struct{})
	bindPort := make(map[nat.Port][]nat.PortBinding)
	mounts := make([]mount.Mount, len(containerDir))

	for i := 0; i < len(hostPort); i++ {
		exposePort[nat.Port(hostPort[i])] = struct{}{}

		bindPort[nat.Port(hostPort[i])] = []nat.PortBinding{nat.PortBinding{
			HostIP:   "0.0.0.0",   //docker容器映射的宿主机的ip
			HostPort: hostPort[i], //docker 容器映射到宿主机的端口
		}}
	}

	for i := 0; i < len(containerDir); i++ {
		mounts[i] = mount.Mount{
			Type:   mount.TypeBind,
			Source: hostDir[i],
			Target: containerDir[i],
		}
	}

	//创建容器
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,     //镜像名称
		Tty:          true,          //docker run命令中的-t选项
		OpenStdin:    true,          //docker run命令中的-i选项
		Cmd:          []string{cmd}, //docker 容器中执行的命令
		Env:          env,           //环境变量
		WorkingDir:   workDir,       //docker容器中的工作目录
		ExposedPorts: exposePort,    //docker容器对外开放的端口

	}, &container.HostConfig{
		PortBindings: bindPort,
		Mounts:       mounts, //docker 容器目录挂在到宿主机目录
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkname: &network.EndpointSettings{},
		},
	}, containerName)
	if err == nil {
		log.Printf("success create container:%s\n", resp.ID)
	} else {
		log.Println("failed to create container!!!!!!!!!!!!!")
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	log.Printf("success run container:%s\n", resp.ID)

	return resp.ID, nil
}
