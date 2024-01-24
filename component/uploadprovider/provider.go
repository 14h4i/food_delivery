package uploadprovider

import (
	"context"
	"food_delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	GetDomain() string
}
