package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
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

// QueryCondition 定义通用查询条件结构体
type QueryCondition struct {
	Match map[string]string                 `json:"match,omitempty"`
	Term  map[string]interface{}            `json:"term,omitempty"`
	Range map[string]map[string]interface{} `json:"range,omitempty"`
	Bool  *BoolQuery                        `json:"bool,omitempty"`
}

// BoolQuery 定义布尔查询条件
type BoolQuery struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	Should  []map[string]interface{} `json:"should,omitempty"`
	MustNot []map[string]interface{} `json:"must_not,omitempty"`
}

// UpdateDocumentsByQueryAssign 通用方法：查询并更新赋值多条数据，使用结构体传参
func (es *EsLowLevelClient) UpdateDocumentsByQueryAssign(index string, query QueryCondition, updateFields map[string]interface{}) error {

	// 检查查询条件是否为空
	if query.Match == nil && query.Term == nil && query.Range == nil && query.Bool == nil {
		return errors.New("query condition is empty")
	}

	// 构建查询部分
	var queryParts = map[string]interface{}{
		"bool": map[string]interface{}{
			"must": []map[string]interface{}{},
		},
	}

	// 处理 Match 条件
	if len(query.Match) > 0 {
		for field, value := range query.Match {
			matchQuery := map[string]interface{}{
				"match": map[string]interface{}{
					field: value,
				},
			}
			// 将 match 查询加入 must 中
			//queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), matchQuery)
			boolQuery, ok := queryParts["bool"].(map[string]interface{})
			if !ok {
				// handle error or initialize boolQuery
				boolQuery = map[string]interface{}{}
				queryParts["bool"] = boolQuery
			}

			mustQueries, ok := boolQuery["must"].([]map[string]interface{})
			if !ok {
				// handle error or initialize mustQueries
				mustQueries = []map[string]interface{}{}
				boolQuery["must"] = mustQueries
			}

			mustQueries = append(mustQueries, matchQuery)
			boolQuery["must"] = mustQueries
		}
	}

	// 添加 Term 查询条件
	if query.Term != nil {
		termQuery := map[string]interface{}{
			"term": query.Term,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), termQuery)
	}

	// 添加 Range 查询条件
	if query.Range != nil {
		rangeQuery := map[string]interface{}{
			"range": query.Range,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), rangeQuery)
	}

	// 添加 Bool 查询条件
	if query.Bool != nil {
		boolQuery := map[string]interface{}{
			"bool": query.Bool,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), boolQuery)
	}

	// 将查询条件转为 JSON
	queryJson, err := json.Marshal(queryParts)
	if err != nil {
		return fmt.Errorf("Error marshalling query condition: %s", err)
	}

	// 动态构建更新脚本
	var updateScriptParts []string
	for field, value := range updateFields {
		fmt.Println("field: ", field, "value: ", value)
		updateScriptParts = append(updateScriptParts, fmt.Sprintf("ctx._source.%s = params.%s;", field, field))
	}
	updateScript := strings.Join(updateScriptParts, " ")

	// 构建请求体
	bodyMap := map[string]interface{}{
		"script": map[string]interface{}{
			"source": updateScript,
			"lang":   "painless",
			"params": updateFields,
		},
		"query": json.RawMessage(queryJson),
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		return fmt.Errorf("Error marshalling request body: %s", err)
	}
	fmt.Println("---------------body: ", string(body))

	//// 将查询条件结构体转为 JSON
	//queryJson, err := json.Marshal(query)
	//if err != nil {
	//	return fmt.Errorf("Error marshalling query condition: %s", err)
	//}
	//
	//// 动态构建更新脚本
	//var updateScriptParts []string
	//for field, value := range updateFields {
	//	// 将每个字段更新拼接成 Painless 脚本
	//	updateScriptParts = append(updateScriptParts, fmt.Sprintf("ctx._source.%s = '%v';", field, value))
	//}
	//updateScript := strings.Join(updateScriptParts, " ")
	//
	//// 构建更新请求体
	//body := fmt.Sprintf(`{
	//	"script": {
	//		"source": "%s",
	//		"lang": "painless"
	//	},
	//	"query": %s
	//}`, updateScript, string(queryJson))
	//
	//fmt.Println("---------------body: ", string(body))

	// 执行 Update By Query 请求
	res, err := es.UpdateByQuery(
		[]string{index}, // 传递索引列表
		es.UpdateByQuery.WithBody(strings.NewReader(string(body))),
		es.UpdateByQuery.WithContext(context.Background()),
	)
	if err != nil {
		return fmt.Errorf("Error executing update by query: %s", err)
	}
	defer res.Body.Close()

	// 检查响应
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应结果
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		return fmt.Errorf("Error decoding response body: %s", err)
	}
	bbres, _ := json.Marshal(res)
	fmt.Println("--------------res: ", string(bbres), "-------------resMap: ", resMap)

	// 打印更新操作的结果
	if updated, ok := resMap["updated"]; ok {
		fmt.Printf("Successfully updated %v documents\n", updated)
	} else {
		return fmt.Errorf("Unexpected response format: %v", resMap)
	}

	return nil
}

// UpdateOperation 定义每个更新字段的操作类型和值
type UpdateOperation struct {
	Operation string      // 操作类型：add、subtract、multiply、divide 或 assign
	Value     interface{} // 更新值
}

// 构建动态更新脚本
func buildUpdateScript(operations map[string]UpdateOperation) (string, map[string]interface{}) {
	var scriptParts []string
	params := make(map[string]interface{})

	for field, operation := range operations {
		var script string
		switch operation.Operation {
		case "add", "+":
			script = fmt.Sprintf("ctx._source.%s += params.%s;", field, field)
		case "subtract", "-":
			script = fmt.Sprintf("ctx._source.%s -= params.%s;", field, field)
		case "multiply", "*":
			script = fmt.Sprintf("ctx._source.%s *= params.%s;", field, field)
		case "divide", "/":
			script = fmt.Sprintf("ctx._source.%s /= params.%s;", field, field)
		case "assign", "=":
			script = fmt.Sprintf("ctx._source.%s = params.%s;", field, field)
		default:
			continue // 忽略未知操作类型
		}
		scriptParts = append(scriptParts, script)
		params[field] = operation.Value
	}

	return strings.Join(scriptParts, " "), params
}

// UpdateDocumentsByQueryMulti 通用方法：查询并更新,支持普通赋值和四则运算
func (es *EsLowLevelClient) UpdateDocumentsByQueryMulti(index string, query QueryCondition, operations map[string]UpdateOperation) error {

	// 检查查询条件是否为空
	if query.Match == nil && query.Term == nil && query.Range == nil && query.Bool == nil {
		return errors.New("query condition is empty")
	}

	// 构建查询部分
	var queryParts = map[string]interface{}{
		"bool": map[string]interface{}{
			"must": []map[string]interface{}{},
		},
	}

	// 处理 Match 条件
	if len(query.Match) > 0 {
		for field, value := range query.Match {
			matchQuery := map[string]interface{}{
				"match": map[string]interface{}{
					field: value,
				},
			}
			// 将 match 查询加入 must 中
			//queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), matchQuery)
			boolQuery, ok := queryParts["bool"].(map[string]interface{})
			if !ok {
				// handle error or initialize boolQuery
				boolQuery = map[string]interface{}{}
				queryParts["bool"] = boolQuery
			}

			mustQueries, ok := boolQuery["must"].([]map[string]interface{})
			if !ok {
				// handle error or initialize mustQueries
				mustQueries = []map[string]interface{}{}
				boolQuery["must"] = mustQueries
			}

			mustQueries = append(mustQueries, matchQuery)
			boolQuery["must"] = mustQueries
		}
	}

	// 添加 Term 查询条件
	if query.Term != nil {
		termQuery := map[string]interface{}{
			"term": query.Term,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), termQuery)
	}

	// 添加 Range 查询条件
	if query.Range != nil {
		rangeQuery := map[string]interface{}{
			"range": query.Range,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), rangeQuery)
	}

	// 添加 Bool 查询条件
	if query.Bool != nil {
		boolQuery := map[string]interface{}{
			"bool": query.Bool,
		}
		queryParts["bool"].(map[string]interface{})["must"] = append(queryParts["bool"].(map[string]interface{})["must"].([]map[string]interface{}), boolQuery)
	}

	// 将查询条件转为 JSON
	queryJson, err := json.Marshal(queryParts)
	if err != nil {
		return fmt.Errorf("Error marshalling query condition: %s", err)
	}

	// 动态构建更新脚本
	// 构建动态更新脚本
	updateScript, params := buildUpdateScript(operations)

	// 构建请求体
	bodyMap := map[string]interface{}{
		"script": map[string]interface{}{
			"source": updateScript,
			"lang":   "painless",
			"params": params,
		},
		"query": json.RawMessage(queryJson),
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		return fmt.Errorf("Error marshalling request body: %s", err)
	}
	fmt.Println("---------------body: ", string(body))

	//// 将查询条件结构体转为 JSON
	//queryJson, err := json.Marshal(query)
	//if err != nil {
	//	return fmt.Errorf("Error marshalling query condition: %s", err)
	//}
	//
	//// 动态构建更新脚本
	//var updateScriptParts []string
	//for field, value := range updateFields {
	//	// 将每个字段更新拼接成 Painless 脚本
	//	updateScriptParts = append(updateScriptParts, fmt.Sprintf("ctx._source.%s = '%v';", field, value))
	//}
	//updateScript := strings.Join(updateScriptParts, " ")
	//
	//// 构建更新请求体
	//body := fmt.Sprintf(`{
	//	"script": {
	//		"source": "%s",
	//		"lang": "painless"
	//	},
	//	"query": %s
	//}`, updateScript, string(queryJson))
	//
	//fmt.Println("---------------body: ", string(body))

	// 执行 Update By Query 请求
	res, err := es.UpdateByQuery(
		[]string{index}, // 传递索引列表
		es.UpdateByQuery.WithBody(strings.NewReader(string(body))),
		es.UpdateByQuery.WithContext(context.Background()),
	)
	if err != nil {
		return fmt.Errorf("Error executing update by query: %s", err)
	}
	defer res.Body.Close()

	// 检查响应
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应结果
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		return fmt.Errorf("Error decoding response body: %s", err)
	}
	bbres, _ := json.Marshal(res)
	fmt.Println("--------------res: ", string(bbres), "-------------resMap: ", resMap)

	// 打印更新操作的结果
	if updated, ok := resMap["updated"]; ok {
		fmt.Printf("Successfully updated %v documents\n", updated)
	} else {
		return fmt.Errorf("Unexpected response format: %v", resMap)
	}

	return nil
}
