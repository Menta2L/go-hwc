package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Host struct {
	Hostname             string
	OS                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	KernelVersion        string
	KernelArch           string
	VirtualizationSystem string
	VirtualizationRole   string
	HostID               string
}
type Memory struct {
	Total uint64
}
type Cpu struct {
	CPU       int32
	VendorID  string
	Family    string
	Model     string
	ModelName string
}
type NetworkInterfaces struct {
	Index        int64
	MTU          int64
	Name         string
	HardwareAddr string
	Flags        []string
	Addrs        []string
}
type DiskPartition struct {
	Device     string
	Mountpoint string
	Fstype     string
	Opts       []string
}
type Netstat struct {
	Addr    string
	Port    uint64
	Proto   string
	Process string
}

// Greeter is a Greeter model.
type Hardware struct {
	Host              *Host
	Memory            *Memory
	NetworkInterfaces []*NetworkInterfaces
	DiskPartition     []*DiskPartition
	Netstat           []*Netstat
	Cpu               []*Cpu
}

// HardwareRepo is a Greater repo.
type HardwareRepo interface {
	Save(context.Context, *Hardware) (*Hardware, error)
	Update(context.Context, *Hardware) (*Hardware, error)
	FindByID(context.Context, int64) (*Hardware, error)
	ListByHello(context.Context, string) ([]*Hardware, error)
	ListAll(context.Context) ([]*Hardware, error)
}

// HardwareUsecase is a Greeter usecase.
type HardwareUsecase struct {
	repo HardwareRepo
	log  *log.Helper
}

// NewHardwareUsecase new a Greeter usecase.
func NewHardwareUsecase(repo HardwareRepo, logger log.Logger) *HardwareUsecase {
	return &HardwareUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateHardware creates a Hardware.
func (uc *HardwareUsecase) CreateHardware(ctx context.Context, g *Hardware) (*Hardware, error) {
	uc.log.WithContext(ctx).Infof("CreateHardware:")
	return uc.repo.Save(ctx, g)
}
