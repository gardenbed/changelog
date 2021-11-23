package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gardenbed/charm/flagit"
	"github.com/gardenbed/charm/ui"

	"github.com/gardenbed/changelog/generate"
	"github.com/gardenbed/changelog/internal/git"
	"github.com/gardenbed/changelog/metadata"
	"github.com/gardenbed/changelog/spec"
)

func main() {
	// We will change the verbosity level once it is known
	u := ui.New(ui.None)

	// READING SPEC

	s, err := spec.Default().FromFile()
	if err != nil {
		u.Errorf(ui.Red, "%s", err)
		os.Exit(1)
	}

	if err := flagit.Parse(&s, false); err != nil {
		u.Errorf(ui.Red, "%s", err)
		os.Exit(1)
	}

	// Update verbosity level
	if s.General.Verbose {
		u.SetLevel(ui.Debug)
	} else if !s.General.Print {
		u.SetLevel(ui.Info)
	}

	u.Debugf(ui.Cyan, "%s", s)

	// RUNNING COMMANDS

	switch {
	case s.Help:
		if err := s.PrintHelp(); err != nil {
			u.Errorf(ui.Red, "%s", err)
			os.Exit(1)
		}

	case s.Version:
		fmt.Println(metadata.String())

	default:
		// Retrieve git repo informatin

		gitRepo, err := git.NewRepo(u, ".")
		if err != nil {
			u.Errorf(ui.Red, "%s", err)
			os.Exit(1)
		}

		domain, path, err := gitRepo.GetRemote()
		if err != nil {
			u.Errorf(ui.Red, "%s", err)
			os.Exit(1)
		}
		s = s.WithRepo(domain, path)

		g, err := generate.New(s, u)
		if err != nil {
			u.Errorf(ui.Red, "%s", err)
			os.Exit(1)
		}

		ctx := context.Background()

		if _, err := g.Generate(ctx, s); err != nil {
			u.Errorf(ui.Red, "%s", err)
			os.Exit(1)
		}
	}
}
