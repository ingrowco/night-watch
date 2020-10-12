package statistics

import (
	"time"

	"github.com/c9s/goprocinfo/linux"
)

func cpuPopulate(s Stats, path string) error {
	cpu, err := linux.ReadStat(path)
	if err != nil {
		return err
	}

	s["cpu_user"] = cpu.CPUStatAll.User
	s["cpu_nice"] = cpu.CPUStatAll.Nice
	s["cpu_system"] = cpu.CPUStatAll.System
	s["cpu_idle"] = cpu.CPUStatAll.Idle
	s["cpu_io_wait"] = cpu.CPUStatAll.IOWait
	s["cpu_irq"] = cpu.CPUStatAll.IRQ
	s["cpu_soft_irq"] = cpu.CPUStatAll.SoftIRQ
	s["cpu_steal"] = cpu.CPUStatAll.Steal
	s["cpu_guest"] = cpu.CPUStatAll.Guest
	s["cpu_guest_nice"] = cpu.CPUStatAll.GuestNice
	s["cpu_interrupts"] = cpu.Interrupts
	s["cpu_context_switches"] = cpu.ContextSwitches
	s["cpu_boot_time"] = cpu.BootTime.In(time.UTC).Format(time.RFC3339)
	s["cpu_processes"] = cpu.Processes
	s["cpu_processes_running"] = cpu.ProcsRunning
	s["cpu_processes_blocked"] = cpu.ProcsBlocked

	return nil
}
