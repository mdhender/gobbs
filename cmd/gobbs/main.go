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

// Package main implements a GoBBS server.
package main

import (
	"fmt"
	"github.com/mdhender/gobbs/internal/config"
	"github.com/mdhender/gobbs/internal/server"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // force logs to be UTC
	log.Println("[main] entered")

	cfg := config.Default()
	if err := cfg.Load(); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(2)
	}

	if err := run(cfg); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(2)
	}
}

func run(cfg *config.Config) error {
	srv, err := server.New(cfg)
	if err != nil {
		return err
	} else if srv == nil {
		return fmt.Errorf("assert(srv != nil)")
	}
	return nil
}
