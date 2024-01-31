package appctx

import (
	"food_delivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider, secretKey string) *appContext {
	return &appContext{db: db, uploadprovider: uploadprovider, secretKey: secretKey}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.uploadprovider
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}
