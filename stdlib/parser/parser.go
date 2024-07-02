package parser

import (
	"github.com/rs/zerolog"
)

type Parser interface {
	JSONParser() JSONParser
}

type parser struct {
	json JSONParser
	opt  Options
}

type Options struct {
	JSON JSONOptions
}

func Init(logger zerolog.Logger, opt Options) Parser {
	return &parser{
		json: initJSONP(logger, opt.JSON),
		opt:  opt,
	}
}

func (p *parser) JSONParser() JSONParser {
	return p.json
}
