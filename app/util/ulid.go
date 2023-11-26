package util

import "github.com/oklog/ulid/v2"

func NewUlid() string {
	return ulid.Make().String()
}

func ParseUlid(s string) (ulid.ULID, error) {
	return ulid.Parse(s)
}
