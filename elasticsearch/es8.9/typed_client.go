package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"log"
)

type FullyTypedClient struct {
	*elasticsearch.TypedClient
}

func NewTypedClient() *FullyTypedClient {
	// todo: parse config from app.toml
	// 配置包含用户名和密码
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es-cn-vo93o4tqp0004pwgj.elasticsearch.aliyuncs.com:9200", // 替换为你的 Elasticsearch 地址
		},
		Username: "elastic",    // 替换为你的用户名
		Password: "0b12C4ccLb", // 替换为你的密码
	}

	// 使用自定义配置初始化客户端
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return &FullyTypedClient{
		TypedClient: es,
	}
}

// IndexDoc 索引文档, doc is struct object
func (es *FullyTypedClient) IndexDoc(ctx context.Context, indexName string, doc any, id string) error {
	res, err := es.Index(indexName).
		Id(id).
		Request(doc).
		Refresh(refresh.True).
		Do(ctx)
	//Do(context.TODO())

	fmt.Println("-----FullyTypedClient: ", res)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	return err
}
