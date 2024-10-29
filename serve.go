// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"context"
	"fmt"
	"github.com/mdhender/gobbs/app"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var (
	argsServe struct {
		paths struct {
			assets     string // path to the assets directory
			components string // path to the components directory
			database   string // path to the database file
		}
		server struct {
			host string
			port string
		}
	}

	cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "serve the web application",
		Long:  `Serve the web application.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if argsServe.paths.assets == "" {
				return fmt.Errorf("error: assets: path is required\n")
			} else if ok, err := isDirExists(argsServe.paths.assets); err != nil {
				return fmt.Errorf("assets: %v\n", err)
			} else if !ok {
				return fmt.Errorf("assets: %s: not a directory\n", argsServe.paths.assets)
			} else if argsServe.paths.assets, err = filepath.Abs(argsServe.paths.assets); err != nil {
				return fmt.Errorf("assets: %v\n", err)
			}
			if argsServe.paths.database == "" {
				return fmt.Errorf("error: database: path is required\n")
			} else if ok, err := isFileExists(argsServe.paths.database); err != nil {
				return fmt.Errorf("database: %v\n", err)
			} else if !ok {
				return fmt.Errorf("database: %s: not a file\n", argsServe.paths.database)
			} else if argsServe.paths.database, err = filepath.Abs(argsServe.paths.database); err != nil {
				return fmt.Errorf("database: %v\n", err)
			}
			if argsServe.paths.components == "" {
				return fmt.Errorf("error: components: path is required\n")
			} else if ok, err := isDirExists(argsServe.paths.components); err != nil {
				return fmt.Errorf("components: %v\n", err)
			} else if !ok {
				return fmt.Errorf("components: %s: not a directory\n", argsServe.paths.components)
			} else if argsServe.paths.components, err = filepath.Abs(argsServe.paths.components); err != nil {
				return fmt.Errorf("components: %v\n", err)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("host      : %s\n", argsServe.server.host)
			log.Printf("port      : %s\n", argsServe.server.port)
			log.Printf("assets    : %s\n", argsServe.paths.assets)
			log.Printf("database  : %s\n", argsServe.paths.database)
			log.Printf("components: %v\n", argsServe.paths.components)

			s, err := newServer(
				withAssets(argsServe.paths.assets),
				withComponents(argsServe.paths.components),
				withContext(context.Background()),
				withDatabase(argsServe.paths.database),
				withHost(argsServe.server.host),
				withPort(argsServe.server.port),
			)
			if err != nil {
				log.Fatalf("error: %v\n", err)
			}
			app.PrintAdminRoutes(s.BaseURL(), s.admin.keys.shutdown)
			if err = s.Serve(); err != nil {
				log.Fatalf("server: %v\n", err)
			}
		},
	}
)
