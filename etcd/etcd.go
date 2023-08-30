package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
	"sync"
	"time"
)

var (
	client *clientv3.Client
	err    error
	ctx    = context.Background()

	wg sync.WaitGroup
)

func getEtcdAdd() []string {
	return []string{"192.168.245.22:2379"}
}
func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
func GetEtcdClient() (*clientv3.Client, error) {
	client, err = clientv3.New(clientv3.Config{Endpoints: getEtcdAdd(), DialTimeout: 5 * time.Second})
	return client, err
}

func EtcdConnDis(serviceName, address string) error {
	client, err = GetEtcdClient()
	if err != nil {
		return err
	}
	ip, piper := getLocalIP()
	if piper != nil {
		return piper
	}
	allAdd := fmt.Sprintf("%s:%s", ip, address)

	grant, granter := client.Grant(ctx, 6)
	if granter != nil {
		return granter
	}
	_, err = client.Put(ctx, serviceName, allAdd, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}

	alive, aliverr := client.KeepAlive(ctx, grant.ID)
	if aliverr != nil {
		return aliverr
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range alive {
			fmt.Printf("i.id is %v,  i.ttl is %v\n", i.ID, i.TTL)
		}
	}() //携程读取etcd回调参数

	//client.g
	wg.Wait()
	return nil
}
