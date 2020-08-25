// +build ferret_r0

package machine

import (
	"device/stm32"
	"runtime/interrupt"
)

// https://github.com/fullyautomated/ferret
const (
	PA0  = portA + 0
	PA1  = portA + 1
	PA2  = portA + 2
	PA3  = portA + 3
	PA4  = portA + 4
	PA5  = portA + 5
	PA6  = portA + 6
	PA7  = portA + 7
	PA8  = portA + 8
	PA9  = portA + 9
	PA10 = portA + 10
	PA11 = portA + 11
	PA12 = portA + 12
	PA13 = portA + 13
	PA14 = portA + 14
	PA15 = portA + 15
	PB0  = portB + 0
	PB1  = portB + 1
	PB2  = portB + 2
	PB3  = portB + 3
	PB4  = portB + 4
	PB5  = portB + 5
	PB6  = portB + 6
	PB7  = portB + 7
	PB8  = portB + 8
	PB9  = portB + 9
	PB10 = portB + 10
	PB11 = portB + 11
	PB12 = portB + 12
	PB13 = portB + 13
	PB14 = portB + 14
	PB15 = portB + 15
	PC13 = portC + 13
	PC14 = portC + 14
	PC15 = portC + 15
)

const (
	LED = PA4
)

// UART pins
const (
	UART_TX_PIN     = PA9
	UART_RX_PIN     = PA10
	UART_ALT_TX_PIN = PB6
	UART_ALT_RX_PIN = PB7
)

var (
	// USART1 is the first hardware serial port on the STM32.
	// Both UART0 and UART1 refer to USART1.
	UART0 = UART{
		Buffer: NewRingBuffer(),
		Bus:    stm32.USART1,
	}
	UART1 = &UART0
)

func init() {
	UART0.Interrupt = interrupt.New(stm32.IRQ_USART1, UART0.handleInterrupt)
}

// SPI pins
const (
	SPI0_SCK_PIN = NoPin
	SPI0_SDO_PIN = NoPin
	SPI0_SDI_PIN = NoPin
)

// I2C pins
const (
	SDA_PIN = PB9
	SCL_PIN = PB8
)

// LEDs
const (
	LED_BL1 = PA2
	LED_BL2 = PA7
	LED_PK1 = PA3
	LED_PK2 = PA6
)

// USB pins
const (
	USB_DM = PA11
	USB_DP = PA12
)

// RS485 pins
const (
	RS485_nRE = PB12
	RS485_DE  = PB15
)

// charger pins
const (
	CHG_nCE    = PB0
	CHG_EN_SNK = PB1
	INT_nCHG   = PB2
	INT_nPMC   = PB3
)

// SBU pins (downstream device UART)
const (
	SBU1_A = PA0
	SBU2_A = PA1
	SBU1_B = PB11
	SBU2_B = PB10
)
