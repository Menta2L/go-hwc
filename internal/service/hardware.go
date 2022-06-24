package service

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	v1 "github.com/menta2l/go-hwc/api/hardware/v1"
	"github.com/menta2l/go-hwc/internal/biz"
)

// GreeterService is a greeter service.
type HardwareService struct {
	v1.UnimplementedHardwareServer

	uc *biz.HardwareUsecase
}

// NewGreeterService new a greeter service.
func NewHardwareService(uc *biz.HardwareUsecase) *HardwareService {
	return &HardwareService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *HardwareService) Send(ctx context.Context, in *v1.SendRequest) (*v1.SendReply, error) {
	model := &biz.Hardware{}
	err := copier.Copy(&model, in.Info)
	if err != nil {
		return nil, err
	}
	_, err = s.uc.CreateHardware(ctx, model)
	if err != nil {
		return nil, err
	}
	return &v1.SendReply{}, nil
}
func (s *HardwareService) GetHardware(ctx context.Context, in *v1.GetHardwareRequest) (*v1.GetHardwareReply, error) {
	info := &v1.HardwareInfo{}
	switch in.Filter.(type) {
	case *v1.GetHardwareRequest_Id:
		h, err := s.uc.GetByID(ctx, in.GetId().Id)
		if err != nil {
			return nil, err
		}
		err = copier.Copy(&info, &h)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Info %v\n", info)
		return &v1.GetHardwareReply{Info: info}, nil
	case *v1.GetHardwareRequest_Host:
	default:
		panic(fmt.Sprintf("unknown req.Filter type: %T", in.Filter))
	}
	return nil, nil
}
