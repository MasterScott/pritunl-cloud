package node

import (
	"strings"
	"time"

	"github.com/pritunl/pritunl-cloud/utils"
)

var (
	netIfaces         = []string{}
	netLastIfacesSync time.Time
	defaultIface      = ""
	defaultIfaceSync  time.Time
)

func GetInterfaces() (ifaces []string, err error) {
	if time.Since(netLastIfacesSync) < 15*time.Second {
		ifaces = netIfaces
		return
	}

	ifacesNew := []string{}
	allIfaces, err := utils.GetInterfaces()
	if err != nil {
		return
	}

	for _, iface := range allIfaces {
		if len(iface) == 14 || iface == "lo" ||
			strings.Contains(iface, "br") ||
			iface == "pritunlhost0" {

			continue
		}
		ifacesNew = append(ifacesNew, iface)
	}

	ifaces = ifacesNew
	netLastIfacesSync = time.Now()
	netIfaces = ifacesNew

	return
}

func getDefaultIface() (iface string, err error) {
	if time.Since(defaultIfaceSync) < 900*time.Second {
		iface = defaultIface
		return
	}

	output, err := utils.ExecCombinedOutput("", "route", "-n")
	if err != nil {
		return
	}

	outputLines := strings.Split(output, "\n")
	for _, line := range outputLines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		if fields[0] == "0.0.0.0" {
			iface = strings.TrimSpace(fields[len(fields)-1])
			_ = strings.TrimSpace(fields[1])
		}
	}

	defaultIface = iface
	defaultIfaceSync = time.Now()

	return
}
