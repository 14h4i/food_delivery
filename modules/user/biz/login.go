package userbiz

import (
	"context"
	"food_delivery/common"
	usermodel "food_delivery/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type LoginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenProvider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *LoginBiz {
	return &LoginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *LoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.UserLogin, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHasher := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHasher {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
