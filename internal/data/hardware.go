package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/menta2l/go-hwc/internal/biz"
	"github.com/menta2l/go-hwc/internal/data/ent"
	enthost "github.com/menta2l/go-hwc/internal/data/ent/host"
)

type hardwareRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewHardwareRepo(data *Data, logger log.Logger) biz.HardwareRepo {
	return &hardwareRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *hardwareRepo) Save(ctx context.Context, g *biz.Hardware) (*biz.Hardware, error) {
	host, err := r.data.db.Host.Get(ctx, g.Host.HostID)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	if host != nil {
		err := r.data.db.Host.DeleteOne(host).Exec(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, err
		}
	}
	hc := tx.Host.Create().
		SetID(g.Host.HostID).
		SetHostname(g.Host.Hostname).
		SetKernelArch(g.Host.KernelArch).
		SetKernelVersion(g.Host.KernelVersion).
		SetOs(g.Host.OS).
		SetPlatform(g.Host.Platform).
		SetPlatformFamily(g.Host.PlatformFamily).
		SetPlatformVersion(g.Host.PlatformVersion).
		SetVirtualizationRole(g.Host.VirtualizationRole).
		SetVirtualizationSystem(g.Host.VirtualizationSystem)
	for _, cpu := range g.Cpu {
		c, err := tx.Cpu.Create().SetFamily(cpu.Family).SetModel(cpu.Model).SetModelName(cpu.ModelName).SetVendorID(cpu.VendorID).SetCPU(int(cpu.CPU)).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
		hc.AddCPU(c)

	}
	hi, err := hc.Save(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, rollback(tx, err)
	}
	for _, disk := range g.DiskPartition {
		_, err := tx.Disk.Create().SetDevice(disk.Device).SetMountpoint(disk.Mountpoint).SetOpts(disk.Opts).SetFstype(disk.Fstype).SetHostID(hi).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
	}
	for _, ni := range g.NetworkInterfaces {
		_, err := tx.Network.Create().SetHostID(hi).SetHardwareAddr(ni.HardwareAddr).SetIndex(int(ni.Index)).SetMTU(int(ni.MTU)).SetName(ni.Name).SetAddrs(ni.Addrs).SetFlags(ni.Flags).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
	}
	for _, ns := range g.Netstat {
		_, err := tx.Netstat.Create().SetHostID(hi).SetAddr(ns.Addr).SetPort(ns.Port).SetProcess(ns.Process).SetProto(ns.Proto).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return g, nil
}
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
func (r *hardwareRepo) Update(ctx context.Context, g *biz.Hardware) (*biz.Hardware, error) {
	return g, nil
}

func (r *hardwareRepo) GetByID(ctx context.Context, id string) (*biz.Hardware, error) {
	h, err := r.data.db.Host.Query().WithDisk().WithCPU().WithNetstat().WithNetwork().Where(enthost.ID(id)).First(ctx)

	//.Order(ent.Asc(entcpu.FieldCPU))
	if err != nil {
		return nil, err
	}
	ret := &biz.Hardware{
		Host: &biz.Host{
			Hostname:             h.Hostname,
			OS:                   h.Os,
			Platform:             h.Platform,
			PlatformFamily:       h.PlatformFamily,
			PlatformVersion:      h.PlatformVersion,
			KernelVersion:        h.KernelVersion,
			KernelArch:           h.KernelArch,
			VirtualizationSystem: h.VirtualizationSystem,
			VirtualizationRole:   h.VirtualizationRole,
			HostID:               h.ID,
		},
		NetworkInterfaces: make([]*biz.NetworkInterfaces, 0),
		Netstat:           make([]*biz.Netstat, 0),
		DiskPartition:     make([]*biz.DiskPartition, 0),
		Cpu:               make([]*biz.Cpu, 0),
	}
	err = copier.Copy(&ret.NetworkInterfaces, h.Edges.Network)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&ret.Netstat, h.Edges.Netstat)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&ret.DiskPartition, h.Edges.Disk)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&ret.Cpu, h.Edges.CPU)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (r *hardwareRepo) GetByHostname(context.Context, string) (*biz.Hardware, error) {
	return nil, nil
}

func (r *hardwareRepo) ListAll(context.Context) ([]*biz.Hardware, error) {
	return nil, nil
}
