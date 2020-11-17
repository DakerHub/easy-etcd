package etcd

import (
	"context"
	"errors"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// ErrorOnConnect 连接ETCD失败
var ErrorOnConnect = errors.New("连接ETCD失败")

// Connect 连接etcd
func Connect(machines []string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   machines,
		DialTimeout: 5 * time.Second,
	})
}

// Put KV存储
func Put(machines []string, key string, value string) (*clientv3.PutResponse, error) {
	client, err := Connect(machines)
	defer client.Close()
	if err != nil {
		return nil, ErrorOnConnect
	}
	res, error := client.Put(context.TODO(), key, value)
	return res, error
}

// GetByPrefix 获取KV
func GetByPrefix(machines []string, key string) (*clientv3.GetResponse, error) {
	client, err := Connect(machines)
	defer client.Close()
	if err != nil {
		return nil, ErrorOnConnect
	}

	return client.Get(context.TODO(), key, clientv3.WithPrefix())
}

// GetSingle 获取单个KV
func GetSingle(machines []string, key string) (*clientv3.GetResponse, error) {
	client, err := Connect(machines)
	defer client.Close()
	if err != nil {
		return nil, ErrorOnConnect
	}

	return client.Get(context.TODO(), key)
}

// Del 删除Key
func Del(machines []string, key string) error {
	client, err := Connect(machines)
	defer client.Close()
	if err != nil {
		return ErrorOnConnect
	}
	_, error := client.Delete(context.TODO(), key)

	return error
}
