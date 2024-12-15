package main

import (
	"fmt"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	// Retrieve CPU information
	cpu := cpuid.CPU

	// Display Vendor ID
	fmt.Printf("Vendor ID: %s\n", cpu.VendorString)

	// Display CPU model name
	fmt.Printf("Model Name: %s\n", cpu.BrandName)

	// Display model number and family number
	fmt.Printf("Family Number: %d\n", cpu.Family)
	fmt.Printf("Model Number: %d\n", cpu.Model)

	// Display core count and thread count
	fmt.Printf("Core Count: %d\n", cpu.PhysicalCores)
	fmt.Printf("Thread Count: %d\n\n", cpu.ThreadsPerCore*cpu.PhysicalCores)

	fmt.Println("Recommended YaneuraOu exe on this PC:")
	// Check and display specific features
	// AMD
	if cpu.Family == 25 && cpu.Model >= 1 && cpu.Model <= 127 {
		// 'ZEN3+' is also being treated as 'ZEN3'.
		fmt.Println("ZEN3")
	} else if cpu.Family == 23 && cpu.Model >= 49 && cpu.Model <= 255 {
		fmt.Println("ZEN2")
	} else if cpu.Family == 23 && cpu.Model >= 1 && cpu.Model <= 8 {
		fmt.Println("ZEN1")
	} else if cpu.Family == 23 && cpu.Model >= 16 && cpu.Model <= 23 {
		// 'ZEN+' is also being treated as 'ZEN1'.
		fmt.Println("ZEN1")
		// Intel or AMD ZEN other
	} else if cpu.Supports(cpuid.AVX512VNNI) {
		fmt.Println("AVX512VNNI")
	} else if cpu.Supports(cpuid.AVX512F) {
		fmt.Println("AVX512")
	} else if cpu.Supports(cpuid.AVXVNNI) {
		fmt.Println("AVXVNNI")
	} else if cpu.Supports(cpuid.AVX2) {
		fmt.Println("AVX2")
	} else if cpu.Supports(cpuid.SSE42) {
		fmt.Println("SSE42")
	} else if cpu.Supports(cpuid.SSE4) {
		fmt.Println("SSE41")
	} else {
		fmt.Println("CPU architecture not detected. Try 'SSE42' or 'AVX2'.")
	}
}
