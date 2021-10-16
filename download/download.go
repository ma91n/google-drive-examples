package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

func main() {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(1000).
		Fields("files(id, name, mimeType)").
		Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	for _, f := range r.Files {
		if f.MimeType == "application/vnd.google-apps.folder" {
			// フォルダの場合はスキップ
			continue
		}
		log.Println("Download:", f.Name)

		if err := download(srv, ctx, f.Name, f.Id); err != nil {
			log.Fatalf("Unable to download: %v", err)
		}
	}

}

func download(srv *drive.Service, ctx context.Context, name, id string) error {
	create, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer create.Close()

	resp, err := srv.Files.Get(id).Context(ctx).Download()
	if err != nil {
		return fmt.Errorf("get drive file: %w", err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(create, resp.Body); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
