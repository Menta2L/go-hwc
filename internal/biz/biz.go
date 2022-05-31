package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewHardwareUsecase)
var WorkerProviderSet = wire.NewSet(NewHardwareUsecase)
