// Copyright (c) 2023-2024 Michael D Henderson. All rights reserved.

package config

import "fmt"

type DBConfig struct {
	Host   string
	Port   int
	Name   string
	User   string
	Secret string
}

func (db DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True", db.User, db.Secret, db.Host, db.Port, db.Name)
}
