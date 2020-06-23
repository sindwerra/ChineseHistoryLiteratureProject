package search

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/teris-io/shortid"
	"historyProject/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}

// Post documents to elasticsearch node
// @Summary 根据查询关键字搜索文档
// @Description 根据查询关键字搜索文档
// @ID 1
// @Tags Elastic
// @Accept  json
// @Produce  json
// @Router /elastic/search [get]
func Endpoint(c *gin.Context) {
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
		Index(model.ElasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := model.SearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}
	// Transform search results before returning them
	docs := make([]model.DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc model.DocumentResponse
		json.Unmarshal(hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}

// Post documents to elasticsearch node
// @Summary 提交elastic文档接口
// @Description 提交elastic文档接口
// @ID 2
// @Tags Elastic
// @Accept  json
// @Produce  json
// @Router /elastic/documents [post]
func PostDocument(context *gin.Context) {
	var docs []model.DocumentRequest
	if err := context.BindJSON(&docs); err != nil {
		errorResponse(context, http.StatusBadRequest, "Malformed request body")
		return
	}

	bulk := elasticClient.Bulk().Index(model.ElasticIndexName).Type(model.ElasticTypeName)
	for _, d := range docs {
		doc := model.Document{
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
}
