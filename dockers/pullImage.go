package dockers

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
)

/*
下载镜像
**/

func PullImage(imageName string) error {

	ctx := context.Background()
	cli, err := GetClient()
	if err != nil {
		return err
	}

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(os.Stdout, out)
	return nil
}

// /*
// 根据oauth下载镜像
// **/
// func PullWithOauthImage(ctx context.Context, cli *client.Client, all bool) types.Container {

// 	authConfig := types.AuthConfig{
// 		Username: "username",
// 		Password: "password",
// 	}
// 	encodedJSON, err := json.Marshal(authConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

// 	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{RegistryAuth: authStr})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer out.Close()
// }
