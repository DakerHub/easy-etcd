package etcd

import (
	"context"
	"errors"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// ErrorOnConnect 连接ETCD失败
var ErrorOnConnect = errors.New("连接ETCD失败")

var timeout = 3 * time.Second

type Auth struct {
	Password string
	Username string
}

// Connect 连接etcd
func Connect(machines []string, auth Auth) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:            machines,
		DialTimeout:          1 * time.Second,
		DialKeepAliveTimeout: 1 * time.Second,
		Username:             auth.Username,
		Password:             auth.Password,
	})
}

// Put KV存储
func Put(machines []string, auth Auth, key string, value string) (*clientv3.PutResponse, error) {
	client, err := Connect(machines, auth)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, error := client.Put(ctx, key, value)
	return res, error
}

// GetByPrefix 获取KV
func GetByPrefix(machines []string, auth Auth, key string) (*clientv3.GetResponse, error) {
	client, err := Connect(machines, auth)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return client.Get(ctx, key, clientv3.WithPrefix())
}

// GetSingle 获取单个KV
func GetSingle(machines []string, auth Auth, key string) (*clientv3.GetResponse, error) {
	client, err := Connect(machines, auth)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return client.Get(ctx, key)
}

// Del 删除Key
func Del(machines []string, auth Auth, key string) error {
	client, err := Connect(machines, auth)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, error := client.Delete(ctx, key)

	return error
}
