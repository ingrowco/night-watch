package statistics

import (
	"github.com/c9s/goprocinfo/linux"
)

func uptimePopulate(s Stats, path string) error {
	uptime, err := linux.ReadUptime(path)
	if err != nil {
		return err
	}

	s["uptime_total"] = uptime.Total
	s["uptime_idle"] = uptime.Idle

	return nil
}
