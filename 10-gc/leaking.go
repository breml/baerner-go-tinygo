//go:build gc.leaking
// +build gc.leaking

package main

func printMemStats() {
	println("leaking, no stats available")
}
