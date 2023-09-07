# k8s-SoftRoCE-device-plugin
Kubernetes Device Plugin for SoftRoCE

## Create Soft-RoCE Link
```bash
# Create soft RoCE link
$ modprobe rdma_rxe
$ rdma link add ib0 type rxe netdev enp0s1

# Show IB Devices
$ ibv_devices
    device          	   node GUID
    ------          	----------------
    ib0             	505400fffe71a5dc
```
## 