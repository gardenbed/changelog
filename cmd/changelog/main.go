package main

import (
	"context"
	"fmt"

	"github.com/gardenbed/flagit"

	"github.com/gardenbed/changelog/generate"
	"github.com/gardenbed/changelog/internal/git"
	"github.com/gardenbed/changelog/log"
	"github.com/gardenbed/changelog/spec"
	"github.com/gardenbed/changelog/version"
)

func main() {
	// We cannot enable the logger until the verbosity is known
	logger := log.New(log.None)

	// READING SPEC

	s, err := spec.Default().FromFile()
	if err != nil {
		logger.Fatal(err)
	}

	if err := flagit.Parse(&s, false); err != nil {
		logger.Fatal(err)
	}

	// Update logger verbosity
	if s.General.Verbose {
		logger.ChangeVerbosity(log.Debug)
	} else if !s.General.Print {
		logger.ChangeVerbosity(log.Info)
	}

	logger.Debug(s)

	// RUNNING COMMANDS

	switch {
	case s.Help:
		if err := s.PrintHelp(); err != nil {
			logger.Fatal(err)
		}

	case s.Version:
		fmt.Println(version.String())

	default:
		// Retrieve git repo informatin

		gitRepo, err := git.NewRepo(logger, ".")
		if err != nil {
			logger.Fatal(err)
		}

		domain, path, err := gitRepo.GetRemote()
		if err != nil {
			logger.Fatal(err)
		}
		s = s.WithRepo(domain, path)

		g, err := generate.New(s, logger)
		if err != nil {
			logger.Fatal(err)
		}

		ctx := context.Background()

		if _, err := g.Generate(ctx, s); err != nil {
			logger.Fatal(err)
		}
	}
}
