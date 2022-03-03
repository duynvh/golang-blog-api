package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"golang-blog-api/common"
	"golang-blog-api/component/uploadprovider"
	"golang-blog-api/modules/upload/uploadmodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type ImageStore interface {
	Create(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore ImageStore
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore ImageStore) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	// if err := biz.imgStore.Create(ctx, img); err != nil {
	// 	return nil, uploadmodel.ErrCannotSaveFile(err)
	// }

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
