package statistics

import (
	"os"
)

func osPopulate(s Stats, path string) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	s["hostname"] = hostname

	return nil
}
