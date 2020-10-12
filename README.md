Night Watch
===========
![Clipart from https://www.pikpng.com/pngvi/bmhxxi_bleed-area-may-not-be-visible-jon-snow-clipart](https://user-images.githubusercontent.com/1416085/95684837-07c84c80-0c01-11eb-80ba-2d82db7ae33a.png)

[![Maintainability](https://api.codeclimate.com/v1/badges/a8daa678e976f8ef9e26/maintainability)](https://codeclimate.com/github/ingrowco/night-watch/maintainability)
![Go](https://github.com/ingrowco/night-watch/workflows/Go/badge.svg)
![Build Release](https://github.com/ingrowco/night-watch/workflows/Build%20Release/badge.svg)

Night Watch (`NiW`) is a command-line tool for tracking a variety range of Linux-based system statistics that integrated into the [Ingrow](https://ingrow.co) system. The statistics -based on your choice- includes information about the OS, CPU, disks, uptime, memory, load average, and networks.

 
### Data Collection
Collected information on each plugin showed in this table.

| Plugin                        | Collected Information                                                                                                                                                                                                                                 |
|-------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| OS                            | `hostname`                                                                                                                                                                                                                                            |
| CPU                           | `cpu_user`, `cpu_nice`, `cpu_system`, `cpu_idle`, `cpu_io_wait`, `cpu_irq`, `cpu_soft_irq`, `cpu_steal`, `cpu_guest`, `cpu_guest_nice`, `cpu_interrupts`, `cpu_context_switches`, `cpu_boot_time`, `cpu_processes`, `cpu_processes_running`, `cpu_processes_blocked` |
| Disks (name separated)        | `read_ios`, `read_merges`, `read_sectors`, `read_ticks`, `write_ios`, `write_merges`, `write_sectors`, `write_ticks`, `in_flight`, `io_ticks`, `time_in_queue`                                                                                                  |
| Uptime                        | `uptime_total`, `uptime_idle`                                                                                                                                                                                                                          |
| Memory                        | `mem_total`, `mem_free`, `mem_available`, `mem_buffers`, `mem_cached`, `mem_swap_cached`, `mem_active`, `mem_inactive`, `mem_swap_total`, `mem_swap_free`, `mem_dirty`                                                                                          |
| Load Average                  | `loadavg_1min`, `loadavg_5min`, `loadavg_15min`                                                                                                                                                                                                         |
| Network (interface separated) | `received_bytes`, `received_errs`, `received_drop`, `received_packets`, `transmitted_bytes`, `transmitted_errs`, `transmitted_drop`, `transmitted_packets`                                                                                                   | 

## Installation
### Build From Source
To build, you need go 1.14+ installed in your system. The build process can be done with running `go build -v` on `cmd` 
directory.

### Install Precompiled File
TODO

## Configurations
### Environmental Variables
NiW can be configured with `env` variables. For this purpose, you can use this table to configure NiW.

| Variable                 | Description                                                          | Allowed Values                                              | Default Value                               |
|--------------------------|----------------------------------------------------------------------|-------------------------------------------------------------|---------------------------------------------|
| `NIW_LOG_LEVEL`          | Level of log verbosity                                               | `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace`       | `info`                                      |
| `NIW_INGROW_PROJECT`     | UID of project (required)                                            |                                                             |                                             |
| `NIW_INGROW_STREAM`      | The Stream name (required)                                           |                                                             |                                             |
| `NIW_INGROW_APIKEY`      | API key of the project (required)                                    |                                                             |                                             |
| `NIW_INGROW_URL`         | API endpoint URL of Ingrow system (required)                         |                                                             |                                             |
| `NIW_MAIN_INTERVAL`      | Intervals of collecting and sending events                           | \[number\]\[unit:s,m,h\] (s: seconds, m: minutes, h: hours) | `30s`                                       |
| `NIW_MAIN_PLUGINS`       | List of plugins which must be used to collect data (comma separated) | `os`, `cpu`, `disk`, `uptime`, `memory`, `loadavg`, `network`     | `os,cpu,disk,uptime,memory,loadavg,network` |
| `NIW_LINUX_PATH_CPU`     | Path of cpu stats proc file (see Note 1)                             |                                                             | `/proc/stat`                                |
| `NIW_LINUX_PATH_MEMORY`  | Path of memory stats proc file (see Note 1)                          |                                                             | `/proc/meminfo`                             |
| `NIW_LINUX_PATH_NETWORK` | Path of network stats proc file (see Note 1)                         |                                                             | `/proc/net/dev`                             |
| `NIW_LINUX_PATH_UPTIME`  | Path of uptime stats proc file (see Note 1)                          |                                                             | `/proc/uptime`                              |
| `NIW_LINUX_PATH_DISK`    | Path of disk stats proc file (see Note 1)                            |                                                             | `/proc/diskstats`                           |
| `NIW_LINUX_PATH_LOADAVG` | Path of load average stats proc file (see Note 1)                    |                                                             | `/proc/loadavg`                             |

Note 1: See [proc.txt](https://www.mjmwired.net/kernel/Documentation/filesystems/proc.txt) for a guide of finding paths
and interpretations about values. 

### YAML Configuration File
A file named `config.yaml`, located on one of `/etc/niw`, `$HOME/.config/niw` and current directory (`pwd`) can be used to set
configurations. A configuration file has these contents:

```yaml
log:
  level: info

main:
  interval: 30s
  plugins:
    - cpu
    - disk
    - loadavg
    - memory
    - network
    - os
    - uptime

linux:
  path:
    cpu: /proc/stat
    loadavg: /proc/loadavg
    disk: /proc/diskstats
    memory: /proc/meminfo
    network: /proc/net/dev
    uptime: /proc/uptime

ingrow:
  project:
  stream:
  apikey:
  url:
``` 