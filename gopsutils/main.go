package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func Info() (*host.InfoStat, error) {
	return &host.InfoStat{}, nil
}

func main() {
	// 主机
	info1, _ := host.Info()
	fmt.Println("=================主机信息====================")
	JsonPrint(info1)

	fmt.Println("=================CPU信息====================")
	info2, _ := cpu.Info()
	JsonPrint(info2)

	fmt.Println("=================CPU核数====================")
	cores, _ := cpu.Counts(true)
	JsonPrint(cores)

	fmt.Println("=================CPU使用率====================")
	percents, _ := cpu.Percent(time.Second*5, true)
	JsonPrint(percents)

	fmt.Println("=================CPU时间片====================")
	times, _ := cpu.Times(true)
	JsonPrint(times)

	fmt.Println("=================磁盘使用率====================")
	use, _ := disk.Usage("/")
	JsonPrint(use)

	fmt.Println("=================磁盘分区情况====================")
	part1, _ := disk.Partitions(true)
	JsonPrint(part1)

	fmt.Println("=================磁盘IO====================")
	part2, _ := disk.IOCounters("/")
	JsonPrint(part2)

	fmt.Println("=================内存信息====================")
	memory, _ := mem.VirtualMemory()
	JsonPrint(memory)

	fmt.Println("=================交换内存====================")
	swaf, _ := mem.SwapMemory()
	JsonPrint(swaf)

	fmt.Println("=================内存交换设备====================")
	devices, _ := mem.SwapDevices()
	JsonPrint(devices)

	// 获取所有网络接口的统计信息
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting network interfaces: %v\n", err)
		return
	}
	for _, interf := range interfaces {
		fmt.Printf("Interface: %v, MTU: %v, Flags: %v\n", interf.Name, interf.MTU, interf.Flags)
	}

	// 获取网络连接信息
	connections, err := net.Connections("all")
	if err != nil {
		fmt.Printf("Error getting network connections: %v\n", err)
		return
	}
	for _, conn := range connections {
		fmt.Printf("Connection: %v\n", conn)
	}

	// 获取网络统计信息
	netStats, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("Error getting network stats: %v\n", err)
		return
	}
	for _, netStat := range netStats {
		fmt.Printf("Network Interface: %v, Bytes Sent: %v, Bytes Received: %v\n", netStat.Name, netStat.BytesSent, netStat.BytesRecv)
	}
	// 获取当前进程的信息
	selfProcess, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		fmt.Printf("Error getting self process: %v\n", err)
		return
	}

	// 获取进程名称
	name, err := selfProcess.Name()
	if err != nil {
		fmt.Printf("Error getting process name: %v\n", err)
		return
	}
	fmt.Printf("Process Name: %v\n", name)

	// 获取进程的内存使用信息
	memInfo, err := selfProcess.MemoryInfo()
	if err != nil {
		fmt.Printf("Error getting process memory info: %v\n", err)
		return
	}
	fmt.Printf("Memory Info: %v\n", memInfo)

	// 获取进程的CPU时间
	cpuTimes, err := selfProcess.Times()
	if err != nil {
		fmt.Printf("Error getting process CPU times: %v\n", err)
		return
	}
	fmt.Printf("CPU Times: %v\n", cpuTimes)

	// 获取所有进程列表
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Error getting all processes: %v\n", err)
		return
	}
	for _, p := range processes {
		pid := p.Pid
		name, err := p.Name()
		if err != nil {
			fmt.Printf("Error getting process name for PID %v: %v\n", pid, err)
			continue
		}
		fmt.Printf("Process PID: %v, Name: %v\n", pid, name)
	}
}

func JsonPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(string(b))
}
