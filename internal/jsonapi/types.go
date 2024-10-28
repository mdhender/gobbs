// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

// Package jsonapi implements a REST-ish interface.
// Types are split into three categories, Query, Request, and Response.
package jsonapi

import "time"

type Author struct {
	Id     string
	Name   string
	Roles  map[string]bool
	Secret string
}

type Message struct {
	Id        string
	Author    *Author
	Subject   string
	CreatedAt time.Time
	Body      string
}

type Session struct {
	Id        string
	Author    *Author
	ExpiresAt time.Time
}

type Sessions map[string]*Session

type Thread struct {
	Id       string
	Messages []*Message
}
