package utils

import (
	"log"
	"os/exec"
	"strings"

	"github.com/prometheus/procfs"
)

func getInterface() string {
	ipExec, err := exec.LookPath("ip")
	if err != nil {
		log.Fatalf("ip not found in $PATH: %s", err)
	}
	cmd := exec.Command(ipExec, "-o", "route", "show", "to", "default")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	ipInterface := strings.Split(string(output), " ")[4]
	return ipInterface
}

// GetBytes retrives the current byte count from procfs
func GetBytes() (uint64, uint64) {
	p, err := procfs.Self()
	if err != nil {
		log.Fatalf("could not get process: %s", err)
	}
	ipInterface := getInterface()
	stat, err := p.NetDev()
	if err != nil {
		log.Fatalf("could not fetch net stats: %s", err)
	}
	lo := stat[ipInterface]
	return lo.RxBytes, lo.TxBytes
}
