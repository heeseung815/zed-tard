package storage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type StorageConfig struct {
	AccountName   string
	AccountKey    string
	ContainerName string
}

var containerClient azblob.ContainerClient

func ConfigureClient(cfg *StorageConfig) {
	accountName := cfg.AccountName
	accountKey := cfg.AccountKey
	if len(accountName) == 0 || len(accountKey) == 0 {
		// TODO: add to warning log
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
	ctx := context.Background()

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		// TODO: add to error log
	}

	serviceClient, err := azblob.NewServiceClientWithSharedKey(url, credential, nil)
	if err != nil {
		// TODO: add to error log
	}

	containerName := cfg.ContainerName
	containerClient = serviceClient.NewContainerClient(containerName)
	_, err = containerClient.Create(ctx, nil)
	if err != nil {
		// TODO: add to info log (already exists)
	}

}

func UploadBlockBlob() bool {
	return true
}

func DownloadBlockBlob(blobName string) string {
	blobClient := containerClient.NewBlockBlobClient(blobName)

	get, err := blobClient.Download(context.Background(), nil)
	if err != nil {
		// TODO : add to error log
	}

	downloadedData := &bytes.Buffer{}
	reader := get.Body(&azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(reader)
	if err != nil {
		// TODO : add to error log
	}
	err = reader.Close()
	if err != nil {
		// TODO : add to error log
	}

	return downloadedData.String()
}

/*
	[Clients]

	Three different clients are provided to interact with the various components of the Blob Service:

	1. ServiceClient
		* Get and set account settings
		* Query, create, and delete containers within the account
	2. ContainerClient
		* Get and set container access settings, properties, and metadata
		* Create, delete and query blobs within the container
		* ContainerLeaseClient to support container lease management
	3. BlobClient
		* AppendBlobClient, BlokBlobClient, and PageBlobClient
		* Get and set blob properties
		* Perform CRUD operations on a given blob
		* BlobLeaseClient to support blob lease management

*/
