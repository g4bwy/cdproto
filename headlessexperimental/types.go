package headlessexperimental

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ScreenshotParams encoding options for a screenshot.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental#type-ScreenshotParams
type ScreenshotParams struct {
	Format           ScreenshotParamsFormat `json:"format,omitempty"`           // Image compression format (defaults to png).
	Quality          int64                  `json:"quality,omitempty"`          // Compression quality from range [0..100] (jpeg only).
	OptimizeForSpeed bool                   `json:"optimizeForSpeed,omitempty"` // Optimize image encoding for speed, not for resulting size (defaults to false)
}

// ScreenshotParamsFormat image compression format (defaults to png).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental#type-ScreenshotParams
type ScreenshotParamsFormat string

// String returns the ScreenshotParamsFormat as string value.
func (t ScreenshotParamsFormat) String() string {
	return string(t)
}

// ScreenshotParamsFormat values.
const (
	ScreenshotParamsFormatJpeg ScreenshotParamsFormat = "jpeg"
	ScreenshotParamsFormatPng  ScreenshotParamsFormat = "png"
	ScreenshotParamsFormatWebp ScreenshotParamsFormat = "webp"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ScreenshotParamsFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ScreenshotParamsFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ScreenshotParamsFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch ScreenshotParamsFormat(v) {
	case ScreenshotParamsFormatJpeg:
		*t = ScreenshotParamsFormatJpeg
	case ScreenshotParamsFormatPng:
		*t = ScreenshotParamsFormatPng
	case ScreenshotParamsFormatWebp:
		*t = ScreenshotParamsFormatWebp

	default:
		in.AddError(fmt.Errorf("unknown ScreenshotParamsFormat value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ScreenshotParamsFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
