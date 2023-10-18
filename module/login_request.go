package module

type WXLoginByCodeType struct {
	Code string `json:"code" binding:"required"`
}
