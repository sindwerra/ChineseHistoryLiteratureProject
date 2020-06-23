package main

import (
	"encoding/json"
	"fmt"
	_ "ginProject/docs"
	"ginProject/middleware"
	"ginProject/service"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 定义 Elasticsearch 索引和类型名字.
// 索引是具有不同类型的文档的集合，这个例子只定义了一个叫做 document 的类型。
const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

// Document 声明要建立索引的文档的主要结构。
type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

// elasticsearch 客户端
var (
	elasticClient *elastic.Client
)

type DocumentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DocumentResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type SearchResponse struct {
	Time      string             `json:"time"`
	Hits      string             `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}

func searchEndpoint(c *gin.Context) {
	// Parse request
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := SearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}
	// Transform search results before returning them
	docs := make([]DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc DocumentResponse
		json.Unmarshal(hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}

// @title Gin swagger
// @version 1.0
// @description swagger示例

// @contact.name sindwerra
// @contact.email sindwerra@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	gin.ForceConsoleColor()

	// 启动elasticsearch客户端
	var err error
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetSniff(false),
			)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	// 默认添加了Logger和Recovery中间件
	app := gin.Default()
	app.Use(middleware.Counter())

	index := app.Group("/")
	{
		index.GET("", service.IndexHandler)
	}

	user := app.Group("/user")
	login := user.Group("/login")
	{
		user.GET("", service.UserIndexHandler)
		user.GET("/register", service.UserRegisterHandler)
		login.GET("/", service.UserLoginHandler)
		login.GET("/v1", service.UserLoginV1Handler)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.POST("/documents", func(context *gin.Context) {
		var docs []DocumentRequest
		if err := context.BindJSON(&docs); err != nil {
			errorResponse(context, http.StatusBadRequest, "Malformed request body")
			return
		}

		bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
		for _, d := range docs {
			doc := Document{
				ID: shortid.MustGenerate(),
				Title: d.Title,
				CreatedAt: time.Now().UTC(),
				Content: d.Content,
			}
			bulk.Add(elastic.NewBulkIndexRequest().Id(doc.ID).Doc(doc))
		}
		if _, err := bulk.Do(context.Request.Context()); err != nil {
			log.Print(err)
			errorResponse(context, http.StatusInternalServerError, "Failed to create documents")
			return
		}
		context.Status(http.StatusOK)
	})
	app.GET("/search", searchEndpoint)

	if err = app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}