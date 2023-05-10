package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultMunInstanceCost is initially set the same as in wasmd
	DefaultInstanceCost uint64 = 60_000
	// DefaultMunCompileCost set to a large number for testing
	DefaultCompileCost uint64 = 100
)

// MunGasRegisterConfig is defaults plus a custom compile amount
func GasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultInstanceCost
	gasConfig.CompileCost = DefaultCompileCost

	return gasConfig
}

func NewCascadiaWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(GasRegisterConfig())
}
