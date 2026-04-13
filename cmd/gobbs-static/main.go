package main

import (
	"flag"
	"log"

	"github.com/mdhender/gobbs/internal/forumsite"
)

func main() {
	var cfg forumsite.Config
	var outDir string

	flag.StringVar(&cfg.SQLitePath, "sqlite-path", "mybb.sqlite3", "path to the SQLite archive")
	flag.StringVar(&cfg.SetupPath, "setup-file", "setup.json", "path to setup.json")
	flag.StringVar(&cfg.TablePrefix, "table-prefix", "", "MyBB table prefix, auto-detected when empty")
	flag.StringVar(&cfg.SiteTitle, "site-title", "PlayByMail Forums Archive", "site title for generated pages")
	flag.StringVar(&cfg.BaseURL, "base-url", "/", "base URL prefix for generated links")
	flag.StringVar(&outDir, "out", "public", "directory for generated static files")
	flag.Parse()

	cfg.LiveTemplate = false

	renderer, err := forumsite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Close()

	if err := renderer.Build(outDir); err != nil {
		log.Fatal(err)
	}
	log.Printf("generated static site in %s", outDir)
}
