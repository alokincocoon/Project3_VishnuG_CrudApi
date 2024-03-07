package dbhandler

import (
	"context"
	"log"
	"os"

	"github.com/edgedb/edgedb-go"
)

func ConnectDb() (*edgedb.Client, error) {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{
		Password:    edgedb.NewOptionalStr(os.Getenv("EDGEDB_SERVER_PASSWORD")),
		Host:        os.Getenv("EDGEDB_HOST"),
		TLSSecurity: os.Getenv("EDGEDB_CLIENT_TLS_SECURITY"),
	})
	if err != nil {
		log.Println("==========Error in db connection===== ", err)
		log.Fatal(err)
	}
	return client, err
}
