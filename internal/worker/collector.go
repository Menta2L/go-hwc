package worker

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/jinzhu/copier"
	v1 "github.com/menta2l/go-hwc/api/hardware/v1"
	"github.com/menta2l/go-hwc/internal/biz"
	"github.com/menta2l/go-hwc/internal/utils"
	"google.golang.org/protobuf/proto"
)

func CollectorWork(worker *Worker) {

	utils.NewTimedExecutor(2*time.Second, time.Minute).Start(func() {
		worker.log.Infof("hello work !!!!")
		b := &biz.Hardware{}
		err := b.Collect()
		if err != nil {
			fmt.Println(err)
		}
		reply, err := worker.client.GetHardware(context.Background(), &v1.GetHardwareRequest{Filter: &v1.GetHardwareRequest_Id{Id: &v1.GetHardwareRequest_ByID{Id: b.Host.HostID}}})
		if err != nil {
			fmt.Println(err)
			return
		}
		info := &v1.HardwareInfo{}
		err = copier.Copy(info, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		sort.Slice(reply.Info.Cpu, func(i, j int) bool {
			return reply.Info.Cpu[i].CPU < reply.Info.Cpu[j].CPU
		})
		sort.Slice(reply.Info.NetworkInterfaces, func(i, j int) bool {
			return reply.Info.NetworkInterfaces[i].Index < reply.Info.NetworkInterfaces[j].Index
		})

		m1, _ := proto.Marshal(reply.Info)
		m2, _ := proto.Marshal(info)
		if !bytes.Equal(m1, m2) {
			worker.log.Info("Detecting change in coputer configuration. Sending update")
			worker.client.Send(context.Background(), &v1.SendRequest{Info: info})
		}
	}, true)

}
