package main

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"log"
)

func main() {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	var paging string

	for {
		r, err := srv.Files.List().PageSize(1000).
			Fields("nextPageToken, files(id, name, parents)").
			PageToken(paging).
			Context(ctx).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		for _, f := range r.Files {
			fmt.Printf("%s %s %+v\n", f.Name, f.Id, f.Parents)
		}

		paging = r.NextPageToken
		if len(paging) == 0 {
			break
		}
	}

}

