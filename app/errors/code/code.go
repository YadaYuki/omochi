package code

// common Error Code. ref: https://github.com/gilcrest/diy-go-api/blob/9dea2423ed084c14d251f4db014967eaa57f74be/domain/errs/errs.go

type Code string

const (
	NotExist       Code = "NOT_EXIST"
	InvalidRequest Code = "INVALID_REQUEST"
	InternalServer Code = "INTERNAL_SERVER"
)
