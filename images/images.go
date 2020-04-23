package images

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func Images() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%-20s%-20s%-20s%-20s%-20s\n",
		"REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE")
	for _, image := range images {
		repo := strings.Split(image.RepoDigests[0], "@")[0]
		tag := strings.Split(image.RepoTags[0], ":")[1]
		id := image.ID[7:19]
		created := image.Created
		size := image.Size

		fmt.Printf("%-20s%-20s%-20s%-20d%-20d\n",
			repo, tag, id, created, size)
	}
}