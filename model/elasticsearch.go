package model

import "time"

// 定义 Elasticsearch 索引和类型名字.
// 索引是具有不同类型的文档的集合，这个例子只定义了一个叫做 document 的类型。
const (
	ElasticIndexName = "documents"
	ElasticTypeName  = "document"
)

// Document 声明要建立索引的文档的主要结构。
type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

// Document 请求的单个元素结构
type DocumentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 搜索返回的单个元素结构
type DocumentResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

// 搜索返回的响应结构
type SearchResponse struct {
	Time      string             `json:"time"`
	Hits      string             `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}
