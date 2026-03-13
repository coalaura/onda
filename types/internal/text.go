package internal

import (
	"github.com/coalaura/onda/types"
	"github.com/coalaura/onda/types/internal/custom"
)

func init() {
	types.RegisterFallback(types.DetectFunc(custom.DetectText))
}
