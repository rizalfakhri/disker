package disk

import (
	"log"

	diskUtil "github.com/shirou/gopsutil/disk"
)

type Disks struct {
	Type  string `json:"type"`
	Disks []Disk `json:"disks"`
}

type Disk struct {
	MountPoint  string  `json:"mount_point"`
	FSType      string  `json:"fs_type"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func Get() Disks {

	diskSummary := make([]Disk, 0)

	disks, err := diskUtil.Partitions(true)

	if err != nil {
		log.Fatalf("Unable to get the disk usages due to: %s", err)
	}

	for _, disk := range disks {

		diskUsage, _ := diskUtil.Usage(disk.Mountpoint)

		diskSummary = append(diskSummary, Disk{
			MountPoint:  diskUsage.Path,
			FSType:      diskUsage.Fstype,
			Total:       diskUsage.Total,
			Free:        diskUsage.Free,
			Used:        diskUsage.Used,
			UsedPercent: diskUsage.UsedPercent,
		})
	}

	totalDisks := Disks{
		Type:  "disk",
		Disks: diskSummary,
	}

	return totalDisks
}
