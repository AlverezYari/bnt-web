package main

import (
	"context"
	"os"
)

type DaggerPipelines struct {
}

func (m *DaggerPipelines) CreateContainer(ctx context.Context, dir *Directory) *Container {
	// Check env to see if we are running locally, or in a GitHub action

	buildEnv := os.Getenv("BUILD_ENV")
	if buildEnv == "dev" {
		buildContainer := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt")
		return buildContainer
	} else if buildEnv == "staging" {
		// Mocking the remote container for now
		buildcontainer := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt")
		return buildcontainer
	} else if buildEnv == "prod" {
		// Mocking the remote container for now
		container := dag.Container().
			From("golang:1.22.2").WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt")
		return container
	}

	return nil
}

func (m *DaggerPipelines) TestContainer(ctx context.Context, ctr *Container) *Container {
	testContainer := ctr.WithExec([]string{"go", "test", "./cmd/web"})
	if testContainer == nil {
		return nil
	}

	return ctr
}

func (m *DaggerPipelines) Build(ctx context.Context, dir *Directory) *Directory {
	// Check env to see if we are running locally, or in a GitHub action

	buildEnv := os.Getenv("BUILD_ENV")
	if buildEnv == "local" {
		goBuildDirectory := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt").
			WithExec([]string{"mkdir", "./out"}).
			WithExec([]string{"go", "build", "-o", "./out/webserver.go", "./cmd/web"}).
			Directory("/tmp")
		return goBuildDirectory
	} else if buildEnv == "github" {
		goBuildDirectory := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt").
			WithExec([]string{"go", "build", "-o", "server.go", "./cmd/web"}).
			Directory("/tmp")
		return goBuildDirectory
	}

	return nil
}

func (m *DaggerPipelines) Test(ctx context.Context, dir *Directory) *Directory {
	// Check env to see if we are running locally, or in a GitHub action

	buildEnv := os.Getenv("BUILD_ENV")
	if buildEnv == "local" {
		goTestDirectory := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt").
			WithExec([]string{"go", "test", "./..."}).
			Directory("/tmp")
		return goTestDirectory
	} else if buildEnv == "github" {
		goTestDirectory := dag.Container().
			From("golang:1.22.2").
			WithEnvVariable("GOARCH", "arm64").
			WithEnvVariable("GOOS", "darwin").
			WithMountedDirectory("/mnt", dir).
			WithWorkdir("/mnt").
			WithExec([]string{"go", "test", "./..."}).
			Directory("/tmp")
		return goTestDirectory
	}

	return nil
}
