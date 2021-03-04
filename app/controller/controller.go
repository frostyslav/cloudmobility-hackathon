package controller

import (
	"log"
	"strings"

	"github.com/frostyslav/cloudmobility-hackathon/app/git"
	"github.com/frostyslav/cloudmobility-hackathon/app/image"
	"github.com/frostyslav/cloudmobility-hackathon/app/model"
)

const registryURL = ""

func Run(hashmap *model.Hash, url, tag, path, id string) error {
	log.Printf("repo url: %s, repo tag: %s, directory: %s, id: %s", url, tag, path, id)

	path, err := git.Clone(url, tag)
	if err != nil {
		hashmap.SetStatus(id, "git failed")
		log.Print(err)
		return err
	}
	log.Print(path)

	hashmap.SetStatus(id, "git done")

	splitted := strings.Split(url, "/")
	repoName := splitted[len(splitted)-1]

	buildName, err := image.Build(path, registryURL, repoName)
	if err != nil {
		hashmap.SetStatus(id, "build failed")
		log.Print(err)
		return err
	}

	log.Print(buildName)
	hashmap.SetStatus(id, "build done")

	return nil
}
