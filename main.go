package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rizalfakhri/disker/channel/rabbitmq"
	"github.com/rizalfakhri/disker/system/cpu"
	"github.com/rizalfakhri/disker/system/disk"
)

type Config struct {
	data    string
	channel string
}

var config Config

func main() {
	flag.StringVar(&config.data, "data", "disk", "The data to be published")
	flag.StringVar(&config.channel, "channel", "rabbitmq", "To where the data will be published")

	flag.Parse()

	switch config.data {
	case "disk":
		publishDiskUsages()
		break
	case "cpu":
		publishCpuInfo()
		break
	default:
		fmt.Printf("%s is not supported data type", config.data)
	}

}

func publishDiskUsages() {

	// This will publish the disk usage to selected channel every 10 seconds

	for {
		disks, err := json.Marshal(disk.Get())

		logError(err, "Unable to get disk data")

		// The only channel available right now is rabbitmq, will implement later
		switch config.channel {
		case "rabbitmq":
			rabbitmq.Dispatch(disks)
			break
		default:
			fmt.Printf("Channel %s does not supported.", config.channel)
		}

		time.Sleep(10 * time.Second)
	}
}

func publishCpuInfo() {

	// This will publish the CPU info to selected channel every 10 seconds

	for {
		cpus, err := json.Marshal(cpu.Get())

		logError(err, "Unable to get cpu info")

		// The only channel available right now is rabbitmq, will implement later
		switch config.channel {
		case "rabbitmq":
			rabbitmq.Dispatch(cpus)
			break
		default:
			fmt.Printf("Channel %s does not supported.", config.channel)
		}

		time.Sleep(10 * time.Second)
	}

}

func logError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
