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
	//
	//newName := struct {
	//	Name string `json:"name"`
	//}{
	//	"zhangsan",
	//}
	//
	//err = es.UpdateStringById(ctx, testIndex, `{"doc":{"name":"张三1"}}`, "3")
	//err = es.UpdateObjById(ctx, testIndex, newName, "3")
	//fmt.Println("-----------update:", err)
	//
	//type Document struct {
	//	//ID      string `json:"-"`
	//	ID      string `json:"id"`
	//	Title   string `json:"title"`
	//	Content string `json:"content"`
	//	Count   int    `json:"count,omitempty"`
	//}
	//
	//// 定义批量插入的文档
	//documents := []Document{
	//	{ID: "1", Title: "Document 1", Content: "Content for document 1", Count: 5},
	//	{ID: "2", Title: "Document 2", Content: "Content for document 2"},
	//	{ID: "3", Title: "Document 3", Content: "Content for document 3"},
	//	{ID: "4", Title: "文档4", Content: "文档4的内容", Count: 10},
	//}
	//bulkDocs := make([]*BulkBody, 0)
	//for _, doc := range documents {
	//	tmp := &BulkBody{
	//		Id:  doc.ID,
	//		Doc: doc,
	//	}
	//	bulkDocs = append(bulkDocs, tmp)
	//}
	//
	//err = es.BulkIndexDocs(ctx, testIndex, bulkDocs)
	//fmt.Println("-----------BulkIndexDocs:", err)

	// 定义复杂查询条件
	query := QueryCondition{
		Match: map[string]string{
			"title.keyword": "Updated Title 5",
			//"count": "10",
		},

		Term: map[string]interface{}{
			"count": 12, // 查询 count 为 12
		},
		//Range: map[string]map[string]interface{}{
		//	"date": {
		//		"gte": "2023-01-01", // 查询日期在 2023-01-01 之后
		//		"lte": "2023-12-31", // 查询日期在 2023-12-31 之前
		//	},
		//},
		//Bool: &BoolQuery{
		//	Must: []map[string]interface{}{
		//		{"match": map[string]interface{}{"category": "science"}}, // 必须匹配 category 为 "science"
		//	},
		//	Should: []map[string]interface{}{
		//		{"term": map[string]interface{}{"priority": "high"}}, // 或者优先级是 "high"
		//	},
		//	MustNot: []map[string]interface{}{
		//		{"term": map[string]interface{}{"status": "archived"}}, // 且不能是 "archived"
		//	},
		//},
	}

	// 定义动态更新字段
	//updateFields := map[string]interface{}{
	//	"test": "test",
	//	//"status":     "reviewed",
	//}
	//
	//// 执行批量更新
	//err = es.UpdateDocumentsByQueryAssign(testIndex, query, updateFields)
	//if err != nil {
	//	log.Fatalf("Failed to update documents: %s", err)
	//}

	// 定义字段更新操作
	operations := map[string]UpdateOperation{
		"count": {
			Operation: "add", // 加操作
			Value:     10,
		},
		"content": {
			Operation: "=", // 乘操作
			Value:     "content1",
		},
	}
	// 执行批量更新
	err = es.UpdateDocumentsByQueryMulti(testIndex, query, operations)
	if err != nil {
		log.Fatalf("Failed to update documents: %s", err)
	}

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
