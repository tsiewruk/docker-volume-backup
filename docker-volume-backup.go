package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func help() {
	fmt.Println("Usage:")
	fmt.Println("Create volume backup:")
	fmt.Println("docker-volume-back create <volume_name> <destination_path/backup_name.tar>")
	fmt.Println("Example:")
	fmt.Println("docker-volume-backup create prometheus_volume /opt/backup/prometheus.tar")
	fmt.Println()
	fmt.Println("Restore volume backup:")
	fmt.Println("docker-volume-backup restore <destination_path/backup_name.tar> <volume_name>")
	fmt.Println("Example:")
	fmt.Println("docker-volume-backup restore /opt/backup/prometheus.tar prometheus_volume")
}

func runDockerCommand(args []string) error {
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 4 {
		help()
		os.Exit(1)
	}

	action := os.Args[1]
	source := os.Args[2]
	destination := os.Args[3]

	var directory, filename string

	switch action {
	case "create":
		directory = filepath.Dir(destination)
		if directory == "." {
			directory, _ = os.Getwd()
		}
		filename = filepath.Base(destination)
		err := runDockerCommand([]string{"run", "--rm", "-v", fmt.Sprintf("%s:/source", source), "-v", fmt.Sprintf("%s:/dest", directory), "busybox", "tar", "cvaf", fmt.Sprintf("/dest/%s", filename), "-C", "/source", "."})
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
	case "restore":
		directory = filepath.Dir(source)
		if directory == "." {
			directory, _ = os.Getwd()
		}
		filename = filepath.Base(source)
		err := runDockerCommand([]string{"run", "--rm", "-v", fmt.Sprintf("%s:/dest", destination), "-v", fmt.Sprintf("%s:/source", directory), "busybox", "tar", "xvf", fmt.Sprintf("/source/%s", filename), "-C", "/dest"})
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(3)
		}
	default:
		help()
		os.Exit(4)
	}
}
