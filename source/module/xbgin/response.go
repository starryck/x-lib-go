package xbgin

import (
	"github.com/starryck/x-lib-go/source/core/base/xbmodel"
	"github.com/starryck/x-lib-go/source/core/base/xbmtmsg"
)

type (
	MetaMessage = xbmtmsg.MetaMessage

	JSONResponseMeta     = xbmodel.JSONResponseMeta
	JSONResponsePageMeta = xbmodel.JSONResponsePageMeta
	JSONResponseBaseData = xbmodel.JSONResponseBaseData
	JSONResponsePageData = xbmodel.JSONResponsePageData
)

func NewJSONResponse(message *MetaMessage, data any, options *JSONResponseOptions) *JSONResponse {
	response := (&jsonResponseBuilder{message: message, data: data, options: options}).
		initialize().
		setCode().
		setMeta().
		setData().
		build()
	return response
}

type JSONResponse struct {
	Code int `json:"-"`
	Meta any `json:"meta"`
	Data any `json:"data"`
}

type jsonResponseBuilder struct {
	message  *MetaMessage
	data     any
	options  *JSONResponseOptions
	response *JSONResponse
}

type JSONResponseOptions struct {
	HTTPCode *int
	MetaArgs []any
	PageData *JSONResponsePageData
}

func (builder *jsonResponseBuilder) build() *JSONResponse {
	return builder.response
}

func (builder *jsonResponseBuilder) initialize() *jsonResponseBuilder {
	builder.response = &JSONResponse{}
	if builder.options == nil {
		builder.options = &JSONResponseOptions{}
	}
	return builder
}

func (builder *jsonResponseBuilder) setCode() *jsonResponseBuilder {
	code := builder.options.HTTPCode
	if code != nil {
		builder.response.Code = *code
	} else {
		builder.response.Code = builder.message.GetHTTPCode()
	}
	return builder
}

func (builder *jsonResponseBuilder) setMeta() *jsonResponseBuilder {
	if builder.options.PageData == nil {
		builder.response.Meta = builder.makeMeta()
	} else {
		builder.response.Meta = builder.makePageMeta()
	}
	return builder
}

func (builder *jsonResponseBuilder) setData() *jsonResponseBuilder {
	builder.response.Data = builder.data
	return builder
}

func (builder *jsonResponseBuilder) makeMeta() *JSONResponseMeta {
	meta := &JSONResponseMeta{
		JSONResponseBaseData: JSONResponseBaseData{
			Code:    builder.makeMetaCode(),
			Message: builder.makeMetaMessage(),
		},
	}
	return meta
}

func (builder *jsonResponseBuilder) makePageMeta() *JSONResponsePageMeta {
	data := *builder.options.PageData
	meta := &JSONResponsePageMeta{
		JSONResponseBaseData: JSONResponseBaseData{
			Code:    builder.makeMetaCode(),
			Message: builder.makeMetaMessage(),
		},
		JSONResponsePageData: JSONResponsePageData{
			PageIndex:   data.PageIndex,
			PageSize:    data.PageSize,
			PageConut:   data.PageConut,
			RecordCount: data.RecordCount,
		},
	}
	return meta
}

func (builder *jsonResponseBuilder) makeMetaCode() string {
	code := builder.message.GetOutCode()
	return code
}

func (builder *jsonResponseBuilder) makeMetaMessage() string {
	message := builder.message.GetOutText(builder.options.MetaArgs...)
	return message
}
