//go:build (!android && linux && arm64) || (!android && linux && amd64 && !musl) || (!android && linux && arm)
package piper_phonemize

// ============================================================
// Code Generated Automatically for linux platform, DO NOT EDIT MANUALLY!!
// ============================================================

import (
	pp "github.com/csukuangfj/piper-phonemize-go-linux"
)

// ============================================================
// Struct/Function Defines
// ============================================================

type PhonemizeResult = pp.PhonemizeResult
var GetVersionStr = pp.GetVersionStr
var _platformInitialize = pp.Initialize
var Phonemize = pp.Phonemize
var DeletePhonemizeResult = pp.DeletePhonemizeResult
