package statistics

import (
	"github.com/c9s/goprocinfo/linux"
)

func memoryPopulate(s Stats, path string) error {
	mem, err := linux.ReadMemInfo(path)
	if err != nil {
		return err
	}

	s["mem_total"] = mem.MemTotal
	s["mem_free"] = mem.MemFree
	s["mem_available"] = mem.MemAvailable
	s["mem_buffers"] = mem.Buffers
	s["mem_cached"] = mem.Cached
	s["mem_swap_cached"] = mem.SwapCached
	s["mem_active"] = mem.Active
	s["mem_inactive"] = mem.Inactive
	s["mem_swap_total"] = mem.SwapTotal
	s["mem_swap_free"] = mem.SwapFree
	s["mem_dirty"] = mem.Dirty

	return nil
}
