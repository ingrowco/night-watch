package statistics

import (
	"fmt"

	"github.com/c9s/goprocinfo/linux"
)

func diskPopulate(s Stats, path string) error {
	disks, err := linux.ReadDiskStats(path)
	if err != nil {
		return err
	}

	for i := range disks {
		key := fmt.Sprintf("disk_%s_%%s", disks[i].Name)

		s[fillKey(key, "read_ios")] = disks[i].ReadIOs           // number of read I/Os processed
		s[fillKey(key, "read_merges")] = disks[i].ReadMerges     // number of read I/Os merged with in-queue I/O
		s[fillKey(key, "read_sectors")] = disks[i].ReadSectors   // number of 512 byte sectors read
		s[fillKey(key, "read_ticks")] = disks[i].ReadTicks       // total wait time for read requests in milliseconds
		s[fillKey(key, "write_ios")] = disks[i].WriteIOs         // number of write I/Os processed
		s[fillKey(key, "write_merges")] = disks[i].WriteMerges   // number of write I/Os merged with in-queue I/O
		s[fillKey(key, "write_sectors")] = disks[i].WriteSectors // number of 512 byte sectors written
		s[fillKey(key, "write_ticks")] = disks[i].WriteTicks     // total wait time for write requests in milliseconds
		s[fillKey(key, "in_flight")] = disks[i].InFlight         // number of I/Os currently in flight
		s[fillKey(key, "io_ticks")] = disks[i].IOTicks           // total time this block device has been active in milliseconds
		s[fillKey(key, "time_in_queue")] = disks[i].TimeInQueue  // total wait time for all requests in milliseconds
	}

	return nil
}
