/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/12 21:07
 * 描述     ：这个文件里是管理镜像的相关代码
 */
package docker

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ErrRead = errors.New("读取buildResponse错误")
)

func BuildImage(path string, dockerfile string, name string) ([]byte, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	buildContext, err := os.Open(path) //这里必须接受一个打包的文件
	if err != nil {
		return nil, err
	}
	defer buildContext.Close()
	buildResponse, err := cli.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
		Tags:        []string{name},
		Remove:      true,
		ForceRemove: true,
		Dockerfile:  dockerfile,
	})
	if err != nil {
		return nil, err
	}
	response, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		return nil, errors.New("读取buildResponse错误")
	}
	buildResponse.Body.Close()
	return response, nil
}

// 包括id和tags
type Image struct {
	ID   string `json:"Id"`
	Tags []string
}

func ListImage() ([]Image, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	res := make([]Image, len(images))
	for k, img := range images {
		image := Image{
			ID:   img.ID,
			Tags: img.RepoTags,
		}
		res[k] = image
	}
	return res, nil
}

func RemoveImage(id string) ([]types.ImageDeleteResponseItem, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	resp, err := cli.ImageRemove(ctx, id, types.ImageRemoveOptions{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
