package parser

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"

	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

const (
	JSONConfigDefault string = `default`
	JSONConfigCustom  string = `custom`
)

type JSONParser interface {
	Marshal(orig interface{}) ([]byte, error)
	Unmarshal(blob []byte, dest interface{}) error
}

type jsonparser struct {
	log zerolog.Logger
	API jsoniter.API
	opt JSONOptions
}

type JSONOptions struct {
	Config                        string
	IndentionStep                 int
	MarshalFloatWith6Digits       bool
	EscapeHTML                    bool
	SortMapKeys                   bool
	UseNumber                     bool
	DisallowUnknownFields         bool
	TagKey                        string
	OnlyTaggedField               bool
	ValidateJSONRawMessage        bool
	ObjectFieldMustBeSimpleString bool
	CaseSensitive                 bool
}

func initJSONP(log zerolog.Logger, opt JSONOptions) JSONParser {
	var jsonAPI jsoniter.API

	switch opt.Config {
	case JSONConfigDefault:
		jsonAPI = jsoniter.ConfigDefault

	case JSONConfigCustom:
		jsonAPI = jsoniter.Config{
			IndentionStep:                 opt.IndentionStep,
			MarshalFloatWith6Digits:       opt.MarshalFloatWith6Digits,
			EscapeHTML:                    opt.EscapeHTML,
			SortMapKeys:                   opt.SortMapKeys,
			UseNumber:                     opt.UseNumber,
			DisallowUnknownFields:         opt.DisallowUnknownFields,
			TagKey:                        opt.TagKey,
			OnlyTaggedField:               opt.OnlyTaggedField,
			ValidateJsonRawMessage:        opt.ValidateJSONRawMessage,
			ObjectFieldMustBeSimpleString: opt.ObjectFieldMustBeSimpleString,
			CaseSensitive:                 opt.CaseSensitive,
		}.Froze()

	default:
		jsonAPI = jsoniter.ConfigCompatibleWithStandardLibrary
	}

	p := &jsonparser{
		log: log,
		API: jsonAPI,
		opt: opt,
	}

	return p
}

func (p *jsonparser) Marshal(orig interface{}) ([]byte, error) {
	stream := p.API.BorrowStream(nil)
	defer p.API.ReturnStream(stream)
	stream.WriteVal(orig)
	result := make([]byte, stream.Buffered())
	if stream.Error != nil {
		return nil, x.Wrap(stream.Error, "json_parser")
	}

	copy(result, stream.Buffer())

	return result, nil
}

func (p *jsonparser) Unmarshal(blob []byte, dest interface{}) error {
	iter := p.API.BorrowIterator(blob)
	defer p.API.ReturnIterator(iter)
	iter.ReadVal(dest)
	if iter.Error != nil {
		return x.Wrap(iter.Error, "json_parser")
	}

	return nil
}
