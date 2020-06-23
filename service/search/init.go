package search

import (
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

// elasticsearch 客户端
var (
	elasticClient *elastic.Client
)

func InitElasticSearch() {
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
}
