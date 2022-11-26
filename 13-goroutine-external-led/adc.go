package main

import (
	"device/nrf"
	"machine"
)

type adc machine.ADC

// Get returns the current value of a ADC pin in the range 0..0xffff.
func (a adc) Get() uint16 {
	var value uint32

	adcPin := a.getADCPin()

	// Enable ADC.
	nrf.ADC.SetENABLE(nrf.ADC_ENABLE_ENABLE_Enabled)

	// Set pin to read.
	nrf.ADC.SetCONFIG_PSEL(adcPin)

	// config ADC
	nrf.ADC.SetCONFIG_RES(nrf.ADC_CONFIG_RES_10bit)
	nrf.ADC.SetCONFIG_INPSEL(nrf.ADC_CONFIG_INPSEL_AnalogInputOneThirdPrescaling)
	nrf.ADC.SetCONFIG_REFSEL(nrf.ADC_CONFIG_REFSEL_SupplyOneThirdPrescaling)

	// Start tasks.
	nrf.ADC.TASKS_START.Set(1)

	// Wait until the sample task is done.
	for nrf.ADC.EVENTS_END.Get() == 0 {
	}
	nrf.ADC.EVENTS_END.Set(0x00)

	value = nrf.ADC.GetRESULT()

	// Stop the ADC
	nrf.ADC.TASKS_STOP.Set(1)

	// Disable ADC.
	nrf.ADC.SetENABLE(nrf.ADC_ENABLE_ENABLE_Disabled)

	if value < 0 {
		value = 0
	}

	// Return 16-bit result from 10-bit value.
	return uint16(value << 6)
}

func (a adc) getADCPin() uint32 {
	switch a.Pin {
	case 1:
		return nrf.ADC_CONFIG_PSEL_AnalogInput2

	case 2:
		return nrf.ADC_CONFIG_PSEL_AnalogInput3

	case 3:
		return nrf.ADC_CONFIG_PSEL_AnalogInput4

	case 4:
		return nrf.ADC_CONFIG_PSEL_AnalogInput5

	case 5:
		return nrf.ADC_CONFIG_PSEL_AnalogInput6

	case 6:
		return nrf.ADC_CONFIG_PSEL_AnalogInput7

	case 26:
		return nrf.ADC_CONFIG_PSEL_AnalogInput0

	case 27:
		return nrf.ADC_CONFIG_PSEL_AnalogInput1

	default:
		return 0
	}
}
