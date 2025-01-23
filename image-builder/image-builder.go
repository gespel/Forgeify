package image_builder

import (
	//"github.com/docker/docker/api/types"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"time"
	//"github.com/docker/docker/pkg/archive"
)

type RepoWorker struct {
	name string
	cli  *client.Client
}

func NewRepoWorker(name string) RepoWorker {
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return RepoWorker{name: name, cli: cli}
}

func (repoW RepoWorker) BuildImage() {
	tar, err := archive.TarWithOptions(repoW.name+"/", &archive.TarOptions{})
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	_, err = repoW.cli.ImageBuild(
		ctx,
		tar,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{"" + repoW.name},
			Remove:     true,
		})
	if err != nil {
		fmt.Println("Error while building image!")
		fmt.Println(err)
		return
	}
}
