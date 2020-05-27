package repository

import (
	"network/repository/start"
)

//repo公共处理
type IRepo interface {
	GetRepoModel() *Repo
}

type Repo struct {
	Start start.Repo
}

func (repo *Repo) GetRepoModel() *Repo {
	return repo
}
