package main

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type rammem struct {
	MemTotal     uint64 `json:"mem_total"`
	MemFree      uint64 `json:"mem_free"`
	MemAvailable uint64 `json:"mem_available"`
}

func main() {
	for {
		//var ram_data string
		ram_data := ram_status()
		println(ram_data)
		ram_data = ram_status()
		// println("data type: %t", ram_data)
		var mminfo rammem
		mminfo.MemTotal = ram_data.MemTotal
		mminfo.MemFree = ram_data.MemFree
		mminfo.MemAvailable = ram_data.MemAvailable
		println("Total Ram:"+"Free Ram"+"Available Ram", mminfo.MemTotal, mminfo.MemFree, mminfo.MemAvailable)
	}
}

func ram_status() *linuxproc.MemInfo {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("mem read fail")
	}

	return stat
}
