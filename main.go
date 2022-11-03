package main

import (
	"context"
	"dagger-pac/config"
	"dagger-pac/pkg/logger"
	"dagger.io/dagger"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return
	}
	l := logger.LogurusSetup(cfg)

	if len(os.Args) < 2 {
		l.Error("Must pass in a Git repository to build")
		os.Exit(1)
	}
	repo := os.Args[1]
	if err := build(repo, l); err != nil {
		l.Error(err)
	}
}

// Language: go
func build(repoUrl string, l logger.Interface) error {
	l.Info("Building %s\n", repoUrl)

	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		l.Error("failed to connect to Dagger: %v", err)
		return err
	}
	defer func(client *dagger.Client) {
		err := client.Close()
		if err != nil {
			l.Error("failed to close Dagger client: %v", err)
			return
		}
	}(client)

	// clone repository with Dagger
	repo := client.Git(repoUrl)
	src, err := repo.Branch("main").Tree().ID(ctx)
	if err != nil {
		l.Error("failed to get source tree: %v", err)
		return err
	}

	// get `golang` image
	golang := client.Container().From("golang:1.18.8-alpine3.15")

	// mount cloned repository into `golang` image
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	path := "build/"
	golang = golang.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "build", "-o", path},
	})

	// get reference to build output directory in container
	output, err := golang.Directory(path).ID(ctx)
	if err != nil {
		l.Error("failed to get output directory: %v", err)
		return err
	}

	// create build/ directory on host
	outpath := filepath.Join(".", path)
	err = os.MkdirAll(outpath, os.ModePerm)
	if err != nil {
		l.Error("failed to create output directory: %v", err)
		return err
	}

	// get reference to current working directory on the host
	workdir := client.Host().Workdir()

	// write contents of container build/ directory
	// to the host working directory
	_, err = workdir.Write(ctx, output, dagger.HostDirectoryWriteOpts{Path: path})
	if err != nil {
		l.Error("failed writing to output directory: %v", err)
		return err
	}

	return nil
}
