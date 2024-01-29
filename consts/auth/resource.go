package auth

import (
	"github.com/pkg/errors"
)

type ResType int

const (
	ResourceItem ResType = iota
	ResourceGift
	UnknownResType
)

func (r ResType) Validate() error {
	if r < ResourceItem || r >= UnknownResType {
		return errors.New("res_type unsupported")
	}
	return nil
}

type AuthType int

const (
	Query AuthType = iota
	Add
	Update
	Delete
	Sync
	Release
	Admin
	UnknownAuthType
)

func (a AuthType) Validate() error {
	if a < Query || a >= UnknownAuthType {
		return errors.New("auth_type unsupported")
	}
	return nil
}
