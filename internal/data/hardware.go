package data

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackc/pgtype"
	"github.com/menta2l/go-hwc/internal/biz"
	"github.com/menta2l/go-hwc/internal/data/ent"
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
		c, err := tx.Cpu.Create().SetFamily(cpu.Family).SetModel(cpu.Model).SetModelName(cpu.ModelName).SetVendorID(cpu.VendorID).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
		hc.AddCPUID(c)

	}
	hi, err := hc.Save(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, rollback(tx, err)
	}
	for _, disk := range g.DiskPartition {
		ta := pgtype.TextArray{}
		ta.Set(disk.Opts)
		_, err := tx.Disk.Create().SetDevice(disk.Device).SetMount(disk.Mountpoint).SetOpts(&ta).SetFsType(disk.Fstype).SetHostID(hi).Save(ctx)
		if err != nil {
			r.log.Error(err)
			return nil, rollback(tx, err)
		}
	}
	for _, ni := range g.NetworkInterfaces {
		addr := pgtype.TextArray{}
		addr.Set(ni.Addrs)
		flags := pgtype.TextArray{}
		flags.Set(ni.Flags)
		_, err := tx.Network.Create().SetHostID(hi).SetMAC(ni.HardwareAddr).SetIdx(int(ni.Index)).SetMtu(int(ni.MTU)).SetName(ni.Name).SetAddrs(&addr).SetFlags(&flags).Save(ctx)
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
	b, err := json.MarshalIndent(g, "", "    ")
	if err == nil {
		fmt.Print(string(b))
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

func (r *hardwareRepo) FindByID(context.Context, int64) (*biz.Hardware, error) {
	return nil, nil
}

func (r *hardwareRepo) ListByHello(context.Context, string) ([]*biz.Hardware, error) {
	return nil, nil
}

func (r *hardwareRepo) ListAll(context.Context) ([]*biz.Hardware, error) {
	return nil, nil
}
