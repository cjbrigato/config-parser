// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *User) Init() {
	p.data = []types.User{}
}

func (p *User) GetParserName() string {
	return "user"
}

func (p *User) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *User) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *User) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.FetchError
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.User{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *User) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.InvalidData
	}
	switch newValue := data.(type) {
	case []types.User:
		p.data = newValue
	case *types.User:
		if index > -1 {
			p.data = append(p.data, types.User{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.User:
		if index > -1 {
			p.data = append(p.data, types.User{})
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

func (p *User) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.User:
		p.data = newValue
	case *types.User:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.IndexOutOfRange
		}
	case types.User:
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

func (p *User) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "user" {
		data, err := p.parse(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "User", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "User", Line: line}
}
