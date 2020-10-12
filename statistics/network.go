package statistics

import (
	"fmt"

	"github.com/c9s/goprocinfo/linux"
)

func networkPopulate(s Stats, path string) error {
	nets, err := linux.ReadNetworkStat(path)
	if err != nil {
		return err
	}

	for i := range nets {
		key := fmt.Sprintf("network_%s_%%s", nets[i].Iface)

		// Received
		s[fillKey(key, "received_bytes")] = nets[i].RxBytes
		s[fillKey(key, "received_errs")] = nets[i].RxErrs
		s[fillKey(key, "received_drop")] = nets[i].RxDrop
		s[fillKey(key, "received_packets")] = nets[i].RxPackets

		// Transmitted
		s[fillKey(key, "transmitted_bytes")] = nets[i].TxBytes
		s[fillKey(key, "transmitted_errs")] = nets[i].TxErrs
		s[fillKey(key, "transmitted_drop")] = nets[i].TxDrop
		s[fillKey(key, "transmitted_packets")] = nets[i].TxPackets
	}

	return nil
}
