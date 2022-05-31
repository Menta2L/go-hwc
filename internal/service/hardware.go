package service

import (
	"context"

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
	if in.GetInfo().GetHost() != nil {
		model.Host = &biz.Host{
			Hostname:             in.Info.Host.Hostname,
			OS:                   in.Info.Host.OS,
			Platform:             in.Info.Host.Platform,
			PlatformFamily:       in.Info.Host.PlatformFamily,
			PlatformVersion:      in.Info.Host.PlatformVersion,
			KernelVersion:        in.Info.Host.KernelVersion,
			KernelArch:           in.Info.Host.KernelArch,
			VirtualizationSystem: in.Info.Host.VirtualizationSystem,
			VirtualizationRole:   in.Info.Host.VirtualizationRole,
			HostID:               in.Info.Host.HostID,
		}
	}
	if in.GetInfo().GetCpu() != nil {
		cpus := make([]*biz.Cpu, 0)
		for k, _ := range in.Info.Cpu {
			cpus = append(cpus, &biz.Cpu{
				CPU:       in.Info.Cpu[k].CPU,
				Family:    in.Info.Cpu[k].Family,
				Model:     in.Info.Cpu[k].Model,
				ModelName: in.Info.Cpu[k].ModelName,
				VendorID:  in.Info.Cpu[k].VendorID,
			})
		}
		model.Cpu = cpus

	}
	if in.GetInfo().DiskPartition != nil {
		dp := make([]*biz.DiskPartition, 0)
		for k, _ := range in.Info.Cpu {
			dp = append(dp, &biz.DiskPartition{
				Device:     in.Info.DiskPartition[k].Device,
				Mountpoint: in.Info.DiskPartition[k].Mountpoint,
				Fstype:     in.Info.DiskPartition[k].Fstype,
				Opts:       in.Info.DiskPartition[k].Opts,
			})
		}
		model.DiskPartition = dp
	}
	if in.GetInfo().NetworkInterfaces != nil {
		ni := make([]*biz.NetworkInterfaces, 0)
		for k, _ := range in.Info.NetworkInterfaces {
			ni = append(ni, &biz.NetworkInterfaces{
				Name:         in.Info.NetworkInterfaces[k].Name,
				Index:        in.Info.NetworkInterfaces[k].Index,
				HardwareAddr: in.Info.NetworkInterfaces[k].HardwareAddr,
				MTU:          in.Info.NetworkInterfaces[k].MTU,
				Flags:        in.Info.NetworkInterfaces[k].Flags,
				Addrs:        in.Info.NetworkInterfaces[k].Addrs,
			})
		}
		model.NetworkInterfaces = ni
	}
	if in.GetInfo().Netstat != nil {
		ns := make([]*biz.Netstat, 0)
		for k, _ := range in.Info.Netstat {
			ns = append(ns, &biz.Netstat{
				Addr:    in.Info.Netstat[k].Addr,
				Port:    in.Info.Netstat[k].Port,
				Process: in.Info.Netstat[k].Process,
			})
		}
		model.Netstat = ns
	}
	if in.GetInfo().Memory != nil {
		model.Memory = &biz.Memory{Total: in.Info.Memory.GetTotal()}
	}

	_, err := s.uc.CreateHardware(ctx, model)
	if err != nil {
		return nil, err
	}
	return &v1.SendReply{}, nil
}
