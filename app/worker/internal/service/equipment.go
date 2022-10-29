package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "harvest/api/equipment/v1"
	"harvest/app/worker/internal/biz"
)

type EquipmentService struct {
	v1.UnimplementedEquipmentServer
	ec  *biz.EquipUseCase
	log *log.Helper
}

func NewEquipmentService(ec *biz.EquipUseCase, logger log.Logger) *EquipmentService {
	return &EquipmentService{
		ec:  ec,
		log: log.NewHelper(log.With(logger, "module", "service/equip")),
	}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, req *v1.CreateEquipmentRequest) (*v1.CreateEquipmentReply, error) {
	return &v1.CreateEquipmentReply{}, nil
}
func (s *EquipmentService) UpdateEquipment(ctx context.Context, req *v1.UpdateEquipmentRequest) (*v1.UpdateEquipmentReply, error) {
	return &v1.UpdateEquipmentReply{}, nil
}
func (s *EquipmentService) DeleteEquipment(ctx context.Context, req *v1.DeleteEquipmentRequest) (*v1.DeleteEquipmentReply, error) {
	return &v1.DeleteEquipmentReply{}, nil
}
func (s *EquipmentService) GetEquipment(ctx context.Context, req *v1.GetEquipmentRequest) (*v1.GetEquipmentReply, error) {
	e, err := s.ec.GetEquip(ctx, req.EquipId)
	if err != nil {
		return nil, err
	}
	return &v1.GetEquipmentReply{
		EquipId:   e.Id,
		EquipName: e.Name,
	}, nil
}
func (s *EquipmentService) ListEquipment(ctx context.Context, req *v1.ListEquipmentRequest) (*v1.ListEquipmentReply, error) {
	return &v1.ListEquipmentReply{}, nil
}
