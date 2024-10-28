// Copyright (c) 2024 Michael D Henderson. All rights reserved.

// Package main implements the web server for GoBBS.
package main

import (
	"fmt"
	"github.com/mdhender/gobbs/internal/config"
	"github.com/mdhender/gobbs/internal/dot"
	"github.com/mdhender/semver"
	"github.com/spf13/cobra"
	"log"
)

var (
	version = semver.Version{Major: 0, Minor: 0, Patch: 1}

	cmdRoot = &cobra.Command{
		Use:   "gobbs",
		Short: "Root command for our application",
	}

	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of this application",
		Long:  `Display the GoBBS version information.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", version.String())
		},
	}
)

func main() {
	log.SetFlags(log.Lshortfile)

	if err := dot.Load("GOBBS", false, false); err != nil {
		log.Fatalf("main: %+v\n", err)
	}

	cfg := config.Default()
	if err := cfg.Load(); err != nil {
		log.Fatalf("%+v\n", err)
	}

	cmdRoot.AddCommand(cmdServe)
	cmdServe.Flags().StringVar(&argsServe.paths.assets, "assets", "assets", "path to the assets directory")
	cmdServe.Flags().StringVar(&argsServe.paths.components, "components", "components", "path to the components directory")
	cmdServe.Flags().StringVar(&argsServe.paths.database, "database", "gobbs.sqlite", "path to the database file")
	cmdServe.Flags().StringVar(&argsServe.server.host, "host", "localhost", "host to serve on")
	cmdServe.Flags().StringVar(&argsServe.server.port, "port", "29631", "port to bind to")

	cmdRoot.AddCommand(cmdVersion)

	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}
