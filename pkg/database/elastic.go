package database

import (
	"fmt"

	"github.com/Ajulll22/belajar-microservice/pkg/security"
	"github.com/olivere/elastic/v7"
)

func ElasticConnect(ELASTIC_USER string, ELASTIC_PASS string, ELASTIC_PROTOCOL, ELASTIC_HOST string, ELASTIC_PORT string, ELASTIC_CLUSTER bool) *elastic.Client {
	elastic_username := ELASTIC_USER
	elastic_password := ELASTIC_PASS
	elastic_protocol := ELASTIC_PROTOCOL
	elastic_server := ELASTIC_HOST
	elastic_port := ELASTIC_PORT
	clear_password := security.Decrypt(elastic_password, "62277ecdae08d9e813ab17a4ec2db8c58db38e398617824a2ef035c64d3da4be")

	dsn := fmt.Sprintf("%s://%s:%s", elastic_protocol, elastic_server, elastic_port)

	client, err := elastic.NewClient(
		elastic.SetURL(dsn),
		elastic.SetBasicAuth(elastic_username, clear_password),
		elastic.SetSniff(ELASTIC_CLUSTER), // Disable sniffing jika tidak menggunakan cluster multi-node
	)
	if err != nil {
		panic("failed to connect elastic")
	}
	return client
}
