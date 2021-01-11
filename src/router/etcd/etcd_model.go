package etcd

import (
	"context"
	"easy-etcd/src/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// ErrorOnConnect 连接ETCD失败
var ErrorOnConnect = errors.New("连接ETCD失败")

var timeout = 10 * time.Second

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

// Snapshot 建立快照
func Snapshot(machines []string, auth Auth) error {
	endpoints := strings.SplitN(machines[0], "http://", 2)
	endpoint := endpoints[1]
	now := time.Now().Unix()
	res, err := GetByPrefix(machines, auth, "")
	if err != nil {
		return err
	}

	newRes := UtilProcessGetRes(res)

	kvs, _ := json.Marshal(newRes)

	return ioutil.WriteFile(fmt.Sprintf("%s%s_%d_%d", config.BackupDir, endpoint, now, newRes.Count), kvs, 0644)
}

// SnapshotDetail 快照详情
type SnapshotDetail struct {
	CreateAt string `json:"createAt"`
	Endpoint string `json:"endpoint"`
	DocCount int64  `json:"docCount"`
	ID       string `json:"id"`
}

// GetAllLocalSnapshot 获取已存储快照
func GetAllLocalSnapshot(endpointURL string) ([]SnapshotDetail, error) {
	strArr := strings.SplitN(endpointURL, "http://", 2)
	endpoint := ""
	if len(strArr) == 2 {
		endpoint = strArr[1]
	}

	files, err := ioutil.ReadDir(config.BackupDir)
	if err != nil {
		log.Fatal(err)
	}

	snapshotMap := make([]SnapshotDetail, 0, len(files))

	for _, file := range files {
		fileName := file.Name()
		pair := strings.SplitN(fileName, "_", 3)
		if len(pair) != 3 {
			continue
		}
		createAt, err1 := strconv.ParseInt(pair[1], 10, 64)
		count, err2 := strconv.ParseInt(pair[2], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		if endpoint != "" && endpoint != pair[0] {
			continue
		}

		snapshotMap = append(snapshotMap, SnapshotDetail{
			CreateAt: time.Unix(createAt, 0).Format("2006-01-02 15:04:05"),
			Endpoint: pair[0],
			DocCount: count,
			ID:       fileName,
		})
	}

	return snapshotMap, nil
}

// DeleteLocalSnapshot 删除快照
func DeleteLocalSnapshot(snapshotID string) error {
	return os.Remove(fmt.Sprintf("%s%s", config.BackupDir, snapshotID))
}

// RestoreBySnapshot 恢复
func RestoreBySnapshot(machines []string, auth Auth, snapshotID string, clear bool) error {
	client, err := Connect(machines, auth)
	if err != nil {
		return err
	}
	defer client.Close()

	content, err := ioutil.ReadFile(fmt.Sprintf("%s%s", config.BackupDir, snapshotID))
	if err != nil {
		return err
	}

	snapshot := &Res{}

	err = json.Unmarshal(content, snapshot)
	if err != nil {
		return err
	}

	if clear {
		ctx2, cancel2 := context.WithTimeout(context.Background(), timeout)
		_, err = client.Delete(ctx2, "", clientv3.WithPrefix())
		cancel2()
		if err != nil {
			return err
		}
	}

	kvs := snapshot.Kvs

	for _, kv := range kvs {
		ctx3, cancel3 := context.WithTimeout(context.Background(), timeout)
		_, err = client.Put(ctx3, kv.Key, kv.Value)
		cancel3()
		if err != nil {
			return err
		}
	}

	return nil
}
