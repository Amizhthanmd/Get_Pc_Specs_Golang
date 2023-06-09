package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)
func main() {
	dd()
	// Retrieve host information
	hostInfo, err := host.Info()
	if err !=nil{
		log.Printf("edwed")
	}
		fmt.Printf("Hostname: %v\n", hostInfo.Hostname)
		fmt.Printf("OS Name: %v\n", hostInfo.Platform)
		fmt.Printf("OS Version: %v %v\n", hostInfo.Platform, hostInfo.PlatformVersion)
		fmt.Printf("OS Kernel Version: %v\n", hostInfo.KernelVersion)
		fmt.Printf("OS Arch: %v\n", runtime.GOARCH)
		numCPU := runtime.NumCPU()
	numCPUCores := int32(numCPU)
	fmt.Printf("Number of CPU Cores: %d\n", numCPUCores)
	
	// Retrieve disk information for "/System/Volumes/Data" mount point
	partitions, err := disk.Partitions(false)
	if err == nil {
		
		for _, partition := range partitions {
			if partition.Mountpoint == "/System/Volumes/Data" {
				usage, err := disk.Usage(partition.Mountpoint)
				if err == nil {
					totalSpaceGB := float64(usage.Total) / (1024 * 1024 * 1024)
					usedSpaceGB := float64(usage.Used) / (1024 * 1024 * 1024)
					freeSpaceGB := float64(usage.Free) / (1024 * 1024 * 1024)
					fmt.Printf("Mount Point: %v\n", partition.Mountpoint)
					fmt.Printf("Total Space: %.2f GB\n", totalSpaceGB)
					fmt.Printf("Used Space: %.2f GB\n", usedSpaceGB)
					fmt.Printf("Free Space: %.2f GB\n", freeSpaceGB)
				}
			}
		}
	}
}

func dd(){
	fmt.Println(time.ANSIC)
}

/* package main

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

func main() {
	// Retrieve host information
	hostInfo, err := host.Info()
	if err == nil {
		fmt.Printf("Hostname: %v\n", hostInfo.Hostname)
		fmt.Printf("OS Name: %v\n", hostInfo.Platform)
		fmt.Printf("OS Version : %v %v\n", hostInfo.Platform, hostInfo.PlatformVersion)
		fmt.Printf("OS KernelVersion : %v\n", hostInfo.KernelVersion)
		fmt.Printf("OS Arch: %v\n", runtime.GOARCH)
		numCPU := runtime.NumCPU()
		fmt.Printf("Number of CPU Cores: %d\n", numCPU)
	}

	// Retrieve disk information
	partitions, err := disk.Partitions(false)
	if err == nil {
		for _, partition := range partitions {
			usage, err := disk.Usage(partition.Mountpoint)
			if err == nil {
				totalSpaceGB := float64(usage.Total) / (1024 * 1024 * 1024)
				fmt.Printf("Device: %v\n", partition.Device)
				fmt.Printf("Mount Point: %v\n", partition.Mountpoint)
				fmt.Printf("Total Space: %.2f GB\n", totalSpaceGB)
			}
		}
	}
}
*/
