package knative

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Create(id, imageName string) (string, error) {
	out, err := exec.Command("kn", "service", "create", id, "--image", imageName, "--kubeconfig", "/etc/kubernetes/admin.conf",
		"--env", "TARGET=\"Go Sample v1\"", "--user", "1000").Output()
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("exec: %s", err)
	}

	fmt.Printf("Output %s\n", out)
	myString := string(out[:])
	splitted := strings.Split(myString, "\n")

	return splitted[len(splitted)-1], nil
}
