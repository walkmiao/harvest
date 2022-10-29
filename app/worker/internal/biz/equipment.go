package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Equip struct {
	Id   int64
	Name string
}

type EquipRepo interface {
	GetEquip(ctx context.Context, eid int64) (*Equip, error)
}

type EquipUseCase struct {
	repo EquipRepo
	log  *log.Helper
}

func NewEquipUseCase(repo EquipRepo, logger log.Logger) *EquipUseCase {
	return &EquipUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/equip"))}
}

func (ec *EquipUseCase) GetEquip(ctx context.Context, eid int64) (*Equip, error) {
	return ec.repo.GetEquip(ctx, eid)
}
