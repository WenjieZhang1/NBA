package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var positionRe = regexp.MustCompile(`<p>位置：([^<]+)</p>`)
var heightRe = regexp.MustCompile(`<p>身高：([^<]+)</p>`)
var weightRe = regexp.MustCompile(`<p>体重：([^<]+)</p>`)

// ParseProfile to get player info
func ParseProfile(content []byte, name string) engine.ParseResult {
	playerinfo := model.PlayerInfo{
		Name: name,
	}
	match := positionRe.FindSubmatch(content)
	if match != nil {
		playerinfo.Position = string(match[1])
	}
	match = heightRe.FindSubmatch(content)
	if match != nil {
		playerinfo.Height = string(match[1])
	}
	match = weightRe.FindSubmatch(content)
	if match != nil {
		playerinfo.Weight = string(match[1])
	}

	result := engine.ParseResult{
		Items: []interface{}{playerinfo},
	}

	return result
}
