// +build bluepill nucleof103rb stm32f4disco ferret_r0

package machine

// Peripheral abstraction layer for the stm32.

const (
	portA Pin = iota * 16
	portB
	portC
	portD
	portE
	portF
	portG
	portH
)
