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
## Apply DaemonSet
```bash
# apply
$ kubectl apply -f ./deploy/k8s-softroce-device-plugin-ds.yaml

# check
$ kubectl get node worker-node -o json | jq '.status.capacity' | grep 'rdma/soft-roce'
  "rdma/soft-roce": "1"
```

## Run Test Pods
```bash
$ kubectl apply -f ./deploy/k8s-test-pod.yaml
$ kubectl exec -it rdma-demo-1 -- bash
root@rdma-demo-1:~# ibv_devices
    device          	   node GUID
    ------          	----------------
    ib0             	505400fffe506c61

# show IP Address
root@rdma-demo-1:~# ip addr ls
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: eth0@if2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 16:76:46:4a:5e:6e brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.244.153.100/24 brd 10.244.153.255 scope global eth0
       valid_lft forever preferred_lft forever
       
root@rdma-demo-1:~# ib_send_bw -d ib0

************************************
* Waiting for client to connect... *
************************************
```

and then open another terminal

```bash
# Open another terminal
$ kubectl exec -it rdma-demo-2 -- bash
root@rdma-demo-2:~# ibv_devices
    device          	   node GUID
    ------          	----------------
    ib0             	505400fffe71a5dc

root@rdma-demo-2:~# ib_send_bw -d ib0 10.244.153.100
---------------------------------------------------------------------------------------
                    Send BW Test
 Dual-port       : OFF		Device         : ib0
 Number of qps   : 1		Transport type : IB
 Connection type : RC		Using SRQ      : OFF
 PCIe relax order: ON
 ibv_wr* API     : OFF
 TX depth        : 128
 CQ Moderation   : 1
 Mtu             : 1024[B]
 Link type       : Ethernet
 GID index       : 1
 Max inline data : 0[B]
 rdma_cm QPs	 : OFF
 Data ex. method : Ethernet
---------------------------------------------------------------------------------------
 local address: LID 0000 QPN 0x0014 PSN 0xb68107
 GID: 00:00:00:00:00:00:00:00:00:00:255:255:192:168:64:47
 remote address: LID 0000 QPN 0x0012 PSN 0xb6c957
 GID: 00:00:00:00:00:00:00:00:00:00:255:255:192:168:64:48
---------------------------------------------------------------------------------------
 #bytes     #iterations    BW peak[MB/sec]    BW average[MB/sec]   MsgRate[Mpps]
 65536      1000             194.84             194.80 		   0.003117
---------------------------------------------------------------------------------------
```