package cpu

import (
	"fmt"

	cpuUtil "github.com/shirou/gopsutil/cpu"
)

type Cpus struct {
	Type  string `json:"type"`
	Total int    `json:"total_cpu"`
	Cpus  []Cpu  `json:"cpus"`
}

type Cpu struct {
	Vendor     string  `json:"vendor"`
	Family     string  `json:"family"`
	Model      string  `json:"model"`
	ModelName  string  `json:"model_name"`
	Stepping   int32   `json:"stepping"`
	PhysicalID string  `json:"physical_id"`
	CoreID     string  `json:"core_id"`
	TotalCores int32   `json:"total_cores"`
	Mhz        float64 `json:"mhz"`
	CacheSize  int32   `json:"cache_size"`
}

func Get() Cpus {

	cpus, _ := cpuUtil.Info()

	// Get only the physical CPU
	numberOfCpu, _ := cpuUtil.Counts(false)

	cpuSummary := make([]Cpu, 0)

	fmt.Println(cpus)

	for _, cpu := range cpus {

		infoStat := cpuUtil.InfoStat(cpu)

		cpuSummary = append(cpuSummary, Cpu{
			Vendor:     infoStat.VendorID,
			Family:     infoStat.Family,
			Model:      infoStat.Model,
			ModelName:  infoStat.ModelName,
			Stepping:   infoStat.Stepping,
			PhysicalID: infoStat.PhysicalID,
			CoreID:     infoStat.CoreID,
			TotalCores: infoStat.Cores,
			Mhz:        infoStat.Mhz,
			CacheSize:  infoStat.CacheSize,
		})

	}

	totalCpus := Cpus{
		Type:  "cpu",
		Total: numberOfCpu,
		Cpus:  cpuSummary,
	}

	return totalCpus
}
