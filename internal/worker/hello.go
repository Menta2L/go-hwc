package worker

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	v1 "github.com/menta2l/go-hwc/api/hardware/v1"
	"github.com/menta2l/go-hwc/internal/utils/netstat"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func HelloWork(worker *Worker) {
	worker.h.Infof("hello work !!!!")
	info := &v1.HardwareInfo{}
	host, err := host.Info()
	if err == nil {
		info.Host = &v1.Host{
			HostID:               host.HostID,
			Hostname:             host.Hostname,
			KernelArch:           host.KernelArch,
			KernelVersion:        host.KernelVersion,
			OS:                   host.OS,
			Platform:             host.Platform,
			PlatformFamily:       host.PlatformFamily,
			PlatformVersion:      host.PlatformVersion,
			VirtualizationRole:   host.VirtualizationRole,
			VirtualizationSystem: host.VirtualizationSystem,
		}
	}
	cpu_info, err := cpu.Info()
	if err == nil {
		info.Cpu = make([]*v1.Cpu, 0)
		for k, _ := range cpu_info {
			info.Cpu = append(info.Cpu, &v1.Cpu{
				CPU:       cpu_info[k].CPU,
				Family:    cpu_info[k].Family,
				Model:     cpu_info[k].Model,
				ModelName: cpu_info[k].ModelName,
				VendorID:  cpu_info[k].VendorID,
			})
		}
	}
	parts, err := disk.Partitions(true)
	if err == nil {
		info.DiskPartition = make([]*v1.DiskPartition, 0)
		for _, part := range parts {
			//diskInfo, err := disk.Usage(part.Mountpoint)
			if err == nil {
				info.DiskPartition = append(info.DiskPartition, &v1.DiskPartition{
					Device:     part.Device,
					Mountpoint: part.Mountpoint,
					Fstype:     part.Fstype,
					Opts:       part.Opts,
				})
			}

		}
	}
	mem_info, err := mem.VirtualMemory()
	if err == nil {
		info.Memory = &v1.Memory{Total: mem_info.Total}
	}
	net_info, err := net.Interfaces()
	if err == nil {
		info.NetworkInterfaces = make([]*v1.NetworkInterfaces, 0)
		for _, ni := range net_info {
			var addrs []string
			for _, addr := range ni.Addrs {
				addrs = append(addrs, addr.Addr)
			}
			info.NetworkInterfaces = append(info.NetworkInterfaces, &v1.NetworkInterfaces{
				Name:         ni.Name,
				Index:        int64(ni.Index),
				HardwareAddr: ni.HardwareAddr,
				MTU:          int64(ni.MTU),
				Flags:        ni.Flags,
				Addrs:        addrs,
			})
		}
	}
	fn := func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	}
	tabs, err := netstat.TCPSocks(fn)
	if err == nil {
		info.Netstat = make([]*v1.Netstat, 0)
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
			info.Netstat = append(info.Netstat, &v1.Netstat{
				Addr:    tmp[0],
				Port:    uint64(port),
				Process: p,
			})
			//fmt.Printf("%-5s %-23.23s %-23.23s %-12s %-16s\n", "tcp", saddr, daddr, e.State, p)
		}
	}
	// using smbios to get additional information from the hardware ?
	// https://github.com/digitalocean/go-smbios
	// https://github.com/siderolabs/go-smbios
	worker.client.Send(context.Background(), &v1.SendRequest{Info: info})
	time.Sleep(10 * time.Second)
}
