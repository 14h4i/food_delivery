package appctx

import (
	"food_delivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, uploadprovider: uploadprovider}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.uploadprovider
}
