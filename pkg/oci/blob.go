package oci

import (
	"context"
	"fmt"
	"io"
	"os"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/registry/remote"
)

func FetchImage(r, reference, dst string) error {
	// 0. Create a file store
	fs, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fs.Close()

	// 1. Connect to a remote repository
	ctx := context.Background()
	repo, err := remote.NewRepository(r)
	if err != nil {
		return fmt.Errorf("failed to create repository: %v", err)
	}
	// Note: The below code can be omitted if authentication is not required
	//	repo.Client = &auth.Client{
	//		Client: retry.DefaultClient,
	//		Cache:  auth.NewCache(),
	//	Credential: auth.StaticCredential(reg, auth.Credential{
	//		Username: "username",
	//		Password: "password",
	//	}),
	//}

	// 2. Copy from the remote repository to the file store

	desc, reader, err := oras.Fetch(ctx, repo, reference, oras.DefaultFetchOptions)
	if err != nil {
		return fmt.Errorf("failed to fetch image: %v", err)
	}

	// 3. Write the file to the file store
	_, err = io.Copy(fs, reader)
	if err != nil {
		return err
	}

	fmt.Println("manifest descriptor:", desc)

	return nil
}
