package main

import (
	"encoding/json"
	"os/exec"
)

// drbdStatus is for parsing relevant data we want to convert to metrics
type drbdStatus struct {
	Name    string `json:"name"`
	Role    string `json:"role"`
	Devices []struct {
		Volume    int    `json:"volume"`
		DiskState string `json:"disk-state"`
	} `json:"devices"`
	Connections []struct {
		PeerNodeID  int    `json:"peer-node-id"`
		PeerRole    string `json:"peer-role"`
		PeerDevices []struct {
			Volume        int    `json:"volume"`
			PeerDiskState string `json:"peer-disk-state"`
		} `json:"peer_devices"`
	} `json:"connections"`
}

// return drbd status in byte raw json
func getDrbdInfo() ([]byte, error) {
	drbdStatusRaw, err := exec.Command("/sbin/drbdsetup", "status", "--json").Output()
	return drbdStatusRaw, err
}

func parseDrbdStatus(statusRaw []byte) ([]drbdStatus, error) {
	var drbdDevs []drbdStatus
	err := json.Unmarshal(statusRaw, &drbdDevs)
	if err != nil {
		return drbdDevs, err
	}
	return drbdDevs, nil
}
