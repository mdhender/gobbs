// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

package keyvalue

type Getter interface {
	Get(key string) (string, bool)
}

type Setter interface {
	Set(key, value string)
}

type Store interface {
	Getter
	Setter
}
