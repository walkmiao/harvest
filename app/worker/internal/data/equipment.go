package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"harvest/app/worker/internal/biz"
	"harvest/app/worker/internal/data/ent/cfgequipment"
)

var _ biz.EquipRepo = (*equipRepo)(nil)

type equipRepo struct {
	data *Data
	log  *log.Helper
}

func (p equipRepo) GetEquip(ctx context.Context, eid int64) (*biz.Equip, error) {
	eItem, err := p.data.db.CfgEquipment.Query().Where(cfgequipment.Equipid(int(eid))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Equip{
		Id:   int64(eItem.Equipid),
		Name: eItem.Equipname,
	}, nil
}

func NewEquipRepo(data *Data, logger log.Logger) biz.EquipRepo {
	return &equipRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/equip")),
	}
}
