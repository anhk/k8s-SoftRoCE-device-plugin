package app

import (
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	ResourceName = "rdma/soft-roce"
	ServerSock   = pluginapi.DevicePluginPath + "soft-roce.sock"
)

var (
	RequiredRdmaDevices = []string{"issm", "rdma_cm", "umad", "uverbs"}
)
