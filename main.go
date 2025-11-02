package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	sdk "github.com/danielvollbro/ecac-plugin-sdk"
)

type Config struct {
	Packages []string `hcl:"packages,optional"`
	Update   bool     `hcl:"update,optional"`
}

type AptPlugin struct {
	cfg Config
}

func (a *AptPlugin) Schema() any {
	return &a.cfg
}

func (a *AptPlugin) Validate(ctx context.Context) error {
	if len(a.cfg.Packages) == 0 {
		return errors.New("at least one package required")
	}
	return nil
}

func (a *AptPlugin) Run(params map[string]any) (string, error) {
	for _, p := range a.cfg.Packages {
		fmt.Println("Installing: ", p)
	}
	return "APT run completed", nil
}

type AptRunner struct {
	inner *AptPlugin
}

func (a *AptRunner) Run(params map[string]any) (string, error) {
	return a.inner.Run(params)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		plugin := &AptPlugin{}
		sdk.Serve(&AptRunner{inner: plugin})
		return
	}
	fmt.Println("ECAC APT plugin binary")
}
