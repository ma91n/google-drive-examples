package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

func main() {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(1000).
		Fields("files(id, name)").
		Q(fmt.Sprintf("'%s' in parents", "＜フォルダID＞")). // 特定のフォルダ配下
		Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	for _, f := range r.Files {
		println(f.Name, f.Id)
	}

}
