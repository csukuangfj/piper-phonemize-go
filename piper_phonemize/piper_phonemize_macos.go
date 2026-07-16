//go:build (darwin && amd64 && !ios) || (darwin && arm64 && !ios)
package piper_phonemize

// ============================================================
// Code Generated Automatically for macos platform, DO NOT EDIT MANUALLY!!
// ============================================================

import (
	pp "github.com/csukuangfj/piper-phonemize-go-macos"
)

// ============================================================
// Struct/Function Defines
// ============================================================

type PhonemizeResult = pp.PhonemizeResult
var GetVersionStr = pp.GetVersionStr
var _platformInitialize = pp.Initialize
var Phonemize = pp.Phonemize
var DeletePhonemizeResult = pp.DeletePhonemizeResult
