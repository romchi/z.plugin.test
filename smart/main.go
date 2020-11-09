// This is the name of your Go package. All files in a package must use the same name.
// It implements 1 metric, called smart, which returns the expternal IP address of the
// host where Agent is running.
package smart

// Packages we will use. The last one is a must.
// Note, the name "zabbix.com" changed from "zabbix" in 4.4.2 .
import (
	"errors"
	"zabbix.com/pkg/plugin"
)

// Plugin must define structure and embed plugin.Base structure.
type Plugin struct {
	plugin.Base
}

var impl Plugin

// Plugin must implement one or several plugin interfaces.
func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	if len(params) > 0 {
		p.Debugf("received %d parameters while expected none", len(params))
		return nil, errors.New("Too many parameters")
	}

	res := ""

	switch key {
		case "smart.disc.discovery":
			res = "smart.disc.discovery"
		case "smart.disk.type":
			res = "smart.disk.type"
		case "smart.disk.smartstate":
			res = "smart.disk.smartstate"
		case "smart.hwraid":
			res = check_hw_raid()
		case "smart.disk.stats":
			res = "smart.disk.stats"
		default:
			/* SHOULD_NEVER_HAPPEN */
			return 0, plugin.UnsupportedMetricError
	}

	return res, nil
}

func init() {
	// Register our metric, specifying the plugin and metric details.
	// 1 - a pointer to plugin implementation
	// 2 - plugin name
	// 3 - metric name (item key)
	// 4 - metric description
	//
	// NB! The metric description must end with period, otherwise the agent won't start!
	plugin.RegisterMetrics(&impl, "smart",
		"smart.disc.discovery", "Return the external IP of the host where agent is running.",
		"smart.disk.type", "Bla bla.",
		"smart.disk.smartstate", "Bla bla.",
		"smart.hwraid", "Bla bla.",
		"smart.disk.stats", "Bla bla.")
}
