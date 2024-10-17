package repositories

import "context"

// Uploader defines methods for an uploader repository
type Uploader interface {
	// Verify verifies that the file exists in our storage service
	Verify(ctx context.Context, url string) error

	// NOTE: Should be supported by most storage service
	//
	// PresignUpload presigns a url for the client to upload the file to our storage service
	PresignUpload(ctx context.Context, fileKey string) (string, error)

	// Delete deletes a file from our storage service
	Delete(ctx context.Context, url string) error
}
