package marshal

import (
	"bufio"
	"strings"

	"github.com/ercole-io/ercole-agent-virtualization/model"
)

// Clusters returns a list of Clusters entries extracted
// from the clusters fetcher command output.
func Clusters(cmdOutput []byte) []model.ClusterInfo {
	//This is a true determistic algorithm. I should prove it!
	scanner := bufio.NewScanner(strings.NewReader(string(cmdOutput)))
	clusters := []model.ClusterInfo{}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ",")
		if len(splitted) == 3 && splitted[0] == "Name" && splitted[1] == "NumCPU" && splitted[2] == "NumSockets" {
			continue
		}

		//Check if the line is not the header line
		clusterInfo := model.ClusterInfo{
			Name:    strings.TrimSpace(splitted[0]),
			CPU:     parseInt(splitted[1]),
			Sockets: parseInt(splitted[2]),
			VMs:     []model.VMInfo{},
		}

		clusters = append(clusters, clusterInfo)
	}

	return clusters
}
