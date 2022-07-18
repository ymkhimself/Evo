/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/13 19:43
 * 描述     ：关于容器的操作
 */
package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func StartContainer(image, name, ip, sshpwd string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, &container.HostConfig{}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"": &network.EndpointSettings{
				IPAddress: ip,
			},
		},
	}, nil, name)
	if err != nil {
		return err
	}
	
}
