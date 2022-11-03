package main

import (
	"dagger-pac/config"
	"dagger-pac/pkg/logger"
	"fmt"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return
	}
	l := logger.LogurusSetup(cfg)
	l.Debug("Hello Logger")
	l.Info("Hello Logger")

	//if len(os.Args) < 2 {
	//	fmt.Println("Must pass in a Git repository to build")
	//	os.Exit(1)
	//}
	//repo := os.Args[1]
	//if err := build(repo); err != nil {
	//	fmt.Println(err)
	//}
}

// Language: go
//func build(repoUrl string) error {
//	fmt.Printf("Building %s\n", repoUrl)
//
//	ctx := context.Background()
//
//	// initialize Dagger client
//	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
//	if err != nil {
//		return err
//	}
//	defer func(client *dagger.Client) {
//		err := client.Close()
//		if err != nil {
//			_ = fmt.Errorf("error: %v", err)
//			return
//		}
//	}(client)
//
//	// clone repository with Dagger
//	repo := client.Git(repoUrl)
//	src, err := repo.Branch("main").Tree().ID(ctx)
//	if err != nil {
//		return err
//	}
//
//	// get `golang` image
//	golang := client.Container().From("golang:latest")
//
//	// mount cloned repository into `golang` image
//	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")
//
//	// define the application build command
//	path := "build/"
//	golang = golang.Exec(dagger.ContainerExecOpts{
//		Args: []string{"go", "build", "-o", path},
//	})
//
//	// get reference to build output directory in container
//	output, err := golang.Directory(path).ID(ctx)
//	if err != nil {
//		return err
//	}
//
//	// create build/ directory on host
//	outpath := filepath.Join(".", path)
//	err = os.MkdirAll(outpath, os.ModePerm)
//	if err != nil {
//		return err
//	}
//
//	// get reference to current working directory on the host
//	workdir := client.Host().Workdir()
//
//	// write contents of container build/ directory
//	// to the host working directory
//	_, err = workdir.Write(ctx, output, dagger.HostDirectoryWriteOpts{Path: path})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
