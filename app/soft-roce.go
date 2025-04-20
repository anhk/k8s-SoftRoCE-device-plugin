package app

import (
	"context"
	"fmt"
	"k8s-softroce-device-plugin/pkg/log"
	"k8s-softroce-device-plugin/pkg/utils"
	"net"
	"os"
	"path"
	"time"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type SoftRoceDevicePlugin struct {
}

func NewSoftRoceDevicePlugin() *SoftRoceDevicePlugin {
	return &SoftRoceDevicePlugin{}
}

func (m *SoftRoceDevicePlugin) GetDevicePluginOptions(_ context.Context, _ *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	return nil, nil
}

func (m *SoftRoceDevicePlugin) ListAndWatch(_ *pluginapi.Empty, server pluginapi.DevicePlugin_ListAndWatchServer) error {
	for i := 0; i < 256; i++ {
		utils.Must(server.Send(&pluginapi.ListAndWatchResponse{
			Devices: []*pluginapi.Device{{
				ID:       fmt.Sprintf("ib%d", i),
				Health:   pluginapi.Healthy,
				Topology: &pluginapi.TopologyInfo{Nodes: []*pluginapi.NUMANode{{ID: 999999}}},
			}},
		}))
	}

	// TODO: ctx->Done()
	select {}
}

func (m *SoftRoceDevicePlugin) GetPreferredAllocation(_ context.Context, _ *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
	return nil, nil
}

func (m *SoftRoceDevicePlugin) Allocate(ctx context.Context, request *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {

	// 查看当前IB卡的数量
	files, err := os.ReadDir("/sys/class/infiniband")
	if err != nil {
		fmt.Println("未检测到 InfiniBand 网卡或没有权限读取 /sys/class/infiniband")
		return nil, err
	}

	response := &pluginapi.AllocateResponse{}
	for i := 0; i < len(files); i++ {
		response.ContainerResponses = append(response.ContainerResponses, &pluginapi.ContainerAllocateResponse{
			Devices: []*pluginapi.DeviceSpec{{
				ContainerPath: fmt.Sprintf("/dev/infiniband/uverbs%d", i),
				HostPath:      fmt.Sprintf("/dev/infiniband/uverbs%d", i),
				Permissions:   "rw",
			}},
		})
	}

	return response, nil
}

func (m *SoftRoceDevicePlugin) PreStartContainer(ctx context.Context, request *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return nil, nil
}

func register(endpoint, resourceName string) {
	conn, err := unixDial(endpoint, 5*time.Second)
	utils.Must(err)
	defer conn.Close()

	client := pluginapi.NewRegistrationClient(conn)
	req := &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		Endpoint:     path.Base(ServerSock),
		ResourceName: resourceName,
	}

	_, err = client.Register(context.Background(), req)
	utils.Must(err)
}

func unixDial(endpoint string, timeout time.Duration) (*grpc.ClientConn, error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c, err := grpc.DialContext(timeoutCtx, endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return net.DialTimeout("unix", endpoint, timeout)
		}))
	return c, err
}

func (m *SoftRoceDevicePlugin) Start() {
	utils.Must(os.MkdirAll(pluginapi.DevicePluginPath, 0755))
	_ = unix.Unlink(ServerSock)

	sock, err := net.Listen("unix", ServerSock)
	utils.Must(err)

	server := grpc.NewServer([]grpc.ServerOption{}...)
	pluginapi.RegisterDevicePluginServer(server, m)

	go func() { utils.Must(server.Serve(sock)) }()
	// Wait for server to start by launching a blocking connection
	conn, err := unixDial(ServerSock, 5*time.Second)
	utils.Must(err)
	utils.Must(conn.Close())
	log.Info("test sock ok")

	register(pluginapi.KubeletSocket, ResourceName)
	log.Info("register device plugin ok")
}

func (m *SoftRoceDevicePlugin) Stop() {
}
