package etcd

import (
	"go.etcd.io/etcd/clientv3"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Res struct {
	Count int64      `json:"count"`
	Kvs   []KeyValue `json:"kvs"`
}

func UtilProcessGetRes(res *clientv3.GetResponse) *Res {
	newRes := &Res{
		Count: res.Count,
		Kvs:   make([]KeyValue, len(res.Kvs)),
	}

	for idx, kv := range res.Kvs {
		newRes.Kvs[idx] = KeyValue{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
	}

	return newRes
}
