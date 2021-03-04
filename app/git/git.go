package git

import (
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

func Clone(url, tag, path string) error {
	// Clone the given repository to the given directory
	log.Print("git clone https://github.com/src-d/go-git")

	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}
