package dockers

import "github.com/docker/docker/client"

var (
	dockerclient *client.Client
	clienterr    error
)

func init() {
	dockerclient, clienterr = client.NewEnvClient()
	if clienterr != nil {
		panic(clienterr)
	}
}

func GetClient() (*client.Client, error) {
	if dockerclient == nil {
		return nil, clienterr
	}
	return dockerclient, nil
}
