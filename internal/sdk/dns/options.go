package dns

import (
	common "github.com/certimate-go/plugin-provider-ctcccloud-smartdns-acmedns01challenger/internal/sdk/common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
