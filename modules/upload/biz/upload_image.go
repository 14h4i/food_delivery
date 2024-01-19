package uploadbiz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

// type CreateImageStorage interface {
// 	CreateImage(context context.Context, data *common.Image) error
// }

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

type uploadBiz struct {
	provider UploadProvider
}

func NewUploadBiz(provider UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimesion(fileBytes)

	if err != nil {
		return nil, errors.New("file is not image")
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                              // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt) // 1203981028930.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", time.Now().UnixNano(), folder, fileName))

	if err != nil {
		return nil, errors.New("cannot save file")
	}

	img.Width = w
	img.Height = h
	// img.CloudName = "s3"
	img.Extension = fileExt

	return img, nil
}

func getImageDimesion(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
