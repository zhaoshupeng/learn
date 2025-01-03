package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {

	ctx := context.Background()
	// ------low-level api
	es := NewClient()
	// 索引文档

	testIndex := "test-index"
	err := es.IndexDoc(ctx, testIndex, `{"grade" : 4,"name":"liming"}`, "1")
	fmt.Println("-----------index:", err)

	newName := struct {
		Name string `json:"name"`
	}{
		"zhangsan",
	}

	err = es.UpdateStringById(ctx, testIndex, `{"doc":{"name":"张三1"}}`, "3")
	err = es.UpdateObjById(ctx, testIndex, newName, "3")
	fmt.Println("-----------update:", err)

	type Document struct {
		ID      string `json:"-"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// 定义批量插入的文档
	documents := []Document{
		{ID: "1", Title: "Document 1", Content: "Content for document 1"},
		{ID: "2", Title: "Document 2", Content: "Content for document 2"},
		{ID: "3", Title: "Document 3", Content: "Content for document 3"},
	}
	bulkDocs := make([]*BulkBody, 0)
	for _, doc := range documents {
		tmp := &BulkBody{
			Id:  doc.ID,
			Doc: doc,
		}
		bulkDocs = append(bulkDocs, tmp)
	}

	err = es.BulkIndexDocs(ctx, testIndex, bulkDocs)
	fmt.Println("-----------BulkIndexDocs:", err)

	//// ----- full-typed api
	//testFull := "test-full"
	//cli := NewTypedClient()
	//res, err := cli.Info().Do(ctx)
	//
	//document := struct {
	//	Name string `json:"name"`
	//}{
	//	"go-elasticsearch",
	//}
	//fmt.Println("--------------testFull", res, err)
	//err = cli.IndexDoc(ctx, testFull, document, "1")
	//
	//fmt.Println("--------------testFull IndexDoc", err)

}

func main1() {
	es := NewClient()

	// 打印 Elasticsearch 信息
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("--------es info: ", res)

	// 示例：索引文档
	indexRes, err := es.Index(
		"test-index",
		strings.NewReader(`{"title" : "Test with auth"}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer indexRes.Body.Close()
	fmt.Println("----------indexRes: ", indexRes)

	// 示例：搜索文档
	searchReq := esapi.SearchRequest{
		Index: []string{"test-index"},
		Body:  strings.NewReader(`{"query": {"match_all": {}}}`),
	}

	searchRes, err := searchReq.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer searchRes.Body.Close()
	fmt.Println(searchRes)
}
