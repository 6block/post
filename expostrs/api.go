package expostrs

import (
	"github.com/spacemeshos/post/internal/postrs"
	"go.uber.org/zap"
)

type VerifyPosOptionsFunc = postrs.VerifyPosOptionsFunc

type ScryptParams = postrs.ScryptParams

type Provider = postrs.Provider

const ClassGPU = postrs.ClassGPU

const ClassCPU = postrs.ClassCPU

var ErrInvalidPos = postrs.ErrInvalidPos

func OpenCLProviders() ([]Provider, error) {
	return postrs.OpenCLProviders()
}

func NewScryptParams(n, r, p uint) ScryptParams {
	return postrs.NewScryptParams(n, r, p)
}

func ToFile(toFile uint32) VerifyPosOptionsFunc {
	return postrs.ToFile(toFile)
}

func VerifyPos(dataDir string, scryptParams ScryptParams, o ...VerifyPosOptionsFunc) error {
	return postrs.VerifyPos(dataDir, scryptParams, o...)
}

func WithFraction(fraction float64) VerifyPosOptionsFunc {
	return postrs.WithFraction(fraction)
}

func FromFile(fromFile uint32) VerifyPosOptionsFunc {
	return postrs.FromFile(fromFile)
}

func VerifyPosWithLogger(logger *zap.Logger) VerifyPosOptionsFunc {
	return postrs.VerifyPosWithLogger(logger)
}
