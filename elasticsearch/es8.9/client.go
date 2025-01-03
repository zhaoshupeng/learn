package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type EsLowLevelClient struct {
	*elasticsearch.Client
}

type UpdateObj struct {
	Doc any `json:"doc"`
}

type UpdateString struct {
	Doc string `json:"doc"`
}

type BulkBody struct {
	Id  string `json:"id"`
	Doc any    `json:"doc"`
}

func NewClient() *EsLowLevelClient {
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
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return &EsLowLevelClient{
		Client: es,
	}
}

// IndexDoc 索引文档; doc is json string
func (es *EsLowLevelClient) IndexDoc(ctx context.Context, indexName string, doc string, id string) error {
	res, err := es.Index(
		indexName,
		strings.NewReader(doc),
		es.Index.WithDocumentID(id),
		es.Index.WithRefresh("true"),
		es.Index.WithContext(ctx),
	)
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
		return err
	}
	if res.IsError() {
		return errors.New(res.String())
	}
	return nil
}

// BulkIndexDocs 批量索引文档; d
func (es *EsLowLevelClient) BulkIndexDocs(ctx context.Context, indexName string, docs []*BulkBody) error {
	// 构建 Bulk 请求体
	var buf bytes.Buffer
	for _, doc := range docs {
		// Action line
		action := map[string]map[string]string{
			"index": {"_index": indexName, "_id": doc.Id},
		}
		actionBytes, _ := json.Marshal(action)
		buf.Write(actionBytes)
		buf.WriteByte('\n')

		// Document line
		docBytes, _ := json.Marshal(doc.Doc)
		fmt.Println("-----------BulkIndexDocs: ", doc.Doc)
		buf.Write(docBytes)
		buf.WriteByte('\n')
	}

	// 执行 Bulk 请求
	res, err := es.Bulk(
		bytes.NewReader(buf.Bytes()),
		es.Bulk.WithRefresh("true"),
		es.Bulk.WithContext(ctx),
	)
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Error executing bulk request: %s", err)
		return err
	}

	// 检查响应
	if res.IsError() {
		log.Fatalf("Error BulkIndexDocs response from Elasticsearch: %s", res.String())
		return errors.New(res.String())
	}
	return nil
}

// UpdateObjById 更新文档; doc is  struct object
func (es *EsLowLevelClient) UpdateObjById(ctx context.Context, indexName string, doc any, id string) error {
	uDoc := UpdateObj{Doc: doc}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(uDoc); err != nil {
		log.Fatalf("Error encoding document: %s", err)
		return err
	}
	res, err := es.Update(
		indexName,
		id,
		&buf,
		es.Update.WithRefresh("true"),
		es.Update.WithContext(ctx),
	)
	defer res.Body.Close()

	if err != nil {
		log.Fatalf("Error update document: %s", err)
		return err
	}

	if res.IsError() {
		return errors.New(res.String())
	}
	return nil
}

// UpdateStringById 更新文档; doc is json string
func (es *EsLowLevelClient) UpdateStringById(ctx context.Context, indexName string, doc string, id string) error {
	wholeDoc := `{"doc":` + doc + `}`
	res, err := es.Update(
		indexName,
		id,
		strings.NewReader(wholeDoc),
		es.Update.WithRefresh("true"),
		es.Update.WithContext(ctx),
	)
	defer res.Body.Close()

	if err != nil {
		log.Fatalf("Error update document: %s", err)
		return err
	}
	if res.IsError() {
		return errors.New(res.String())
	}
	return nil
}
