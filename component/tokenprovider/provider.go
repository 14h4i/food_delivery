package tokenprovider

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		erors.New("token not found"),
		"token not found",
		"ErrNotFound",
	),
	
	ErrEncodingToken = common.NewCustomError(
		erors.New("error encoding token"),
		"error encoding token",
		"ErrEncodingToken",
	),

	ErrInvalidToken = common.NewCustomError(
		erors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken provided",
	),

	
)