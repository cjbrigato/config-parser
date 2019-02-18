// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *ErrorFile) GetParserName() string {
	return "errorfile"
}

func (p *ErrorFile) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *ErrorFile) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *ErrorFile) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.FetchError
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.ErrorFile{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *ErrorFile) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.InvalidData
	}
	switch newValue := data.(type) {
	case []types.ErrorFile:
		p.data = newValue
	case *types.ErrorFile:
		if index > -1 {
			p.data = append(p.data, types.ErrorFile{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.ErrorFile:
		if index > -1 {
			p.data = append(p.data, types.ErrorFile{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.InvalidData
	}
	return nil
}

func (p *ErrorFile) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.ErrorFile:
		p.data = newValue
	case *types.ErrorFile:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.IndexOutOfRange
		}
	case types.ErrorFile:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else if index == -1 {
			p.data = append(p.data, newValue)
		} else {
			return errors.IndexOutOfRange
		}
	default:
		return errors.InvalidData
	}
	return nil
}

func (p *ErrorFile) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "errorfile" {
		data, err := p.parse(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "ErrorFile", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ErrorFile", Line: line}
}
