package statistics

import (
	"github.com/c9s/goprocinfo/linux"
)

func loadAvgPopulate(s Stats, path string) error {
	load, err := linux.ReadLoadAvg(path)
	if err != nil {
		return err
	}

	s["loadavg_1min"] = load.Last1Min
	s["loadavg_5min"] = load.Last5Min
	s["loadavg_15min"] = load.Last15Min

	return nil
}
