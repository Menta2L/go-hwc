package biz

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/menta2l/go-hwc/internal/utils/netstat"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
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
	GetByID(context.Context, string) (*Hardware, error)
	GetByHostname(context.Context, string) (*Hardware, error)
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
func (uc *HardwareUsecase) GetByID(ctx context.Context, id string) (*Hardware, error) {
	return uc.repo.GetByID(ctx, id)
}

// CreateHardware creates a Hardware.
func (uc *HardwareUsecase) GetByHostname(ctx context.Context, hostname string) (*Hardware, error) {
	return uc.repo.GetByHostname(ctx, hostname)
}

// CreateHardware creates a Hardware.
func (uc *HardwareUsecase) CreateHardware(ctx context.Context, g *Hardware) (*Hardware, error) {
	uc.log.WithContext(ctx).Infof("CreateHardware:")
	return uc.repo.Save(ctx, g)
}

func (h *Host) Collect() error {
	stats, err := host.Info()
	if err != nil {
		return err
	}
	h.HostID = stats.HostID
	err = copier.Copy(h, stats)
	if err != nil {
		return err
	}
	return nil
}
func (h *Hardware) Collect() error {
	h.Host = &Host{}
	err := h.Host.Collect()
	if err != nil {
		return err
	}
	part, err := disk.Partitions(false)
	if err != nil {
		return err
	}
	h.DiskPartition = make([]*DiskPartition, len(part))

	err = copier.Copy(&h.DiskPartition, part)
	if err != nil {
		return err
	}
	cpuinfo, err := cpu.Info()
	if err != nil {
		return err
	}
	h.Cpu = make([]*Cpu, len(cpuinfo))

	err = copier.Copy(&h.Cpu, cpuinfo)
	if err != nil {
		return err
	}
	net_info, err := net.Interfaces()
	if err != nil {
		return err
	}
	h.NetworkInterfaces = make([]*NetworkInterfaces, 0)
	for _, ni := range net_info {
		var addrs []string
		for _, addr := range ni.Addrs {
			addrs = append(addrs, addr.Addr)
		}
		h.NetworkInterfaces = append(h.NetworkInterfaces, &NetworkInterfaces{
			Name:         ni.Name,
			Index:        int64(ni.Index),
			HardwareAddr: ni.HardwareAddr,
			MTU:          int64(ni.MTU),
			Flags:        ni.Flags,
			Addrs:        addrs,
		})
	}
	fn := func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	}
	tabs, err := netstat.TCPSocks(fn)
	if err == nil {
		h.Netstat = make([]*Netstat, 0)
		lookup := func(skaddr *netstat.SockAddr) string {
			const IPv4Strlen = 17
			addr := skaddr.IP.String()
			if len(addr) > IPv4Strlen {
				addr = addr[:IPv4Strlen]
			}
			return fmt.Sprintf("%s:%d", addr, skaddr.Port)
		}
		for _, e := range tabs {
			p := ""
			if e.Process != nil {
				p = e.Process.String()
			}
			saddr := lookup(e.LocalAddr)
			tmp := strings.Split(saddr, ":")
			port, _ := strconv.Atoi(tmp[1])
			h.Netstat = append(h.Netstat, &Netstat{
				Addr:    tmp[0],
				Port:    uint64(port),
				Process: p,
			})
			//fmt.Printf("%-5s %-23.23s %-23.23s %-12s %-16s\n", "tcp", saddr, daddr, e.State, p)
		}
	}
	return nil
}
