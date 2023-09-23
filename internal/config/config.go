/*
 * gobbs - threaded forum server
 *
 * Copyright (c) 2021 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package config implements a configuration for GoBBS servers
package config

import (
	"flag"
	"os"
	"path"
	"time"

	"github.com/peterbourgon/ff/v3"
)

type Config struct {
	Debug bool
	App   struct {
		Root            string
		TimestampFormat string
	}
	Cookies struct {
		HttpOnly bool
		Secure   bool
	}
	DB struct {
		Host   string
		Port   int
		Name   string
		User   string
		Secret string
	}
	Data struct {
		Path string
	}
	FileName string
	Server   struct {
		Scheme  string
		Host    string
		Port    string
		Timeout struct {
			Idle  time.Duration
			Read  time.Duration
			Write time.Duration
		}
		Key     string
		Salt    string
		WebRoot string
	}
}

// Default returns a default configuration.
// These are the values without loading the environment, configuration file, or command line.
func Default() *Config {
	var cfg Config
	cfg.App.Root = "D:/GoLand/gobbs/"
	cfg.App.TimestampFormat = "2006-01-02T15:04:05.99999999Z"
	cfg.DB.Port = 3306
	cfg.Data.Path = cfg.App.Root + "test/data/"
	cfg.Server.Host = "localhost"
	cfg.Server.Key = "curry.aka.yrruc"
	cfg.Server.Port = "3000"
	cfg.Server.Salt = "pepper"
	cfg.Server.Scheme = "http"
	cfg.Server.Timeout.Idle = 10 * time.Second
	cfg.Server.Timeout.Read = 5 * time.Second
	cfg.Server.Timeout.Write = 10 * time.Second
	cfg.Server.WebRoot = cfg.App.Root + "web/"
	return &cfg
}

// Load updates the values in a Config in this order:
//  1. It will load a configuration file if one is given on the
//     command line via the `-config` flag. If provided, the file
//     must contain a valid JSON object.
//  2. Environment variables, using the prefix `GOBBS`
//  3. Command line flags
func (cfg *Config) Load() error {
	fs := flag.NewFlagSet("Server", flag.ExitOnError)

	fs.BoolVar(&cfg.Cookies.HttpOnly, "cookies-http-only", cfg.Cookies.HttpOnly, "set HttpOnly flag on cookies")
	fs.BoolVar(&cfg.Cookies.Secure, "cookies-secure", cfg.Cookies.Secure, "set Secure flag on cookies")
	fs.BoolVar(&cfg.Debug, "debug", cfg.Debug, "log debug information (optional)")
	fs.DurationVar(&cfg.Server.Timeout.Idle, "idle-timeout", cfg.Server.Timeout.Idle, "http idle timeout")
	fs.DurationVar(&cfg.Server.Timeout.Read, "read-timeout", cfg.Server.Timeout.Read, "http read timeout")
	fs.DurationVar(&cfg.Server.Timeout.Write, "write-timeout", cfg.Server.Timeout.Write, "http write timeout")
	fs.IntVar(&cfg.DB.Port, "db-port", cfg.DB.Port, "port of mysql database")
	fs.StringVar(&cfg.App.Root, "root", cfg.App.Root, "path to treat as root for relative file references")
	fs.StringVar(&cfg.DB.Host, "db-host", cfg.DB.Host, "host of mysql database")
	fs.StringVar(&cfg.DB.Name, "db-name", cfg.DB.Name, "name of mysql database")
	fs.StringVar(&cfg.DB.Secret, "db-secret", cfg.DB.Secret, "secret for mysql database")
	fs.StringVar(&cfg.DB.User, "db-user", cfg.DB.User, "user in mysql database")
	fs.StringVar(&cfg.Data.Path, "data-path", cfg.Data.Path, "path containing data files")
	fs.StringVar(&cfg.FileName, "config", cfg.FileName, "config file (optional)")
	fs.StringVar(&cfg.Server.Host, "host", cfg.Server.Host, "host name (or IP) to listen on")
	fs.StringVar(&cfg.Server.Key, "key", cfg.Server.Key, "set key for signing tokens")
	fs.StringVar(&cfg.Server.Port, "port", cfg.Server.Port, "port to listen on")
	fs.StringVar(&cfg.Server.Salt, "salt", cfg.Server.Salt, "set salt for hashing passwords")
	fs.StringVar(&cfg.Server.Scheme, "scheme", cfg.Server.Scheme, "http scheme, either 'http' or 'https'")
	fs.StringVar(&cfg.Server.WebRoot, "web-root", cfg.Server.WebRoot, "path to serve web assets from")

	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("GOBBS"), ff.WithConfigFileFlag("config"), ff.WithConfigFileParser(ff.JSONParser)); err != nil {
		return err
	}

	cfg.App.Root = path.Clean(cfg.App.Root)
	cfg.Data.Path = path.Clean(cfg.Data.Path)
	cfg.Server.WebRoot = path.Clean(cfg.Server.WebRoot)

	return nil
}
