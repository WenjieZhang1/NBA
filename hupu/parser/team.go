package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const teamRe = `<a target="_blank" title="([^"]*)" href="(https://nba.hupu.com/players/[a-zA-Z0-9]+-[0-9]+.html)">`

// ParseTeam to get players
func ParseTeam(content []byte) engine.ParseResult {
	re := regexp.MustCompile(teamRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	visitedPlayer := make(map[string]struct{})
	for _, match := range matches {
		player := string(match[1])
		url := string(match[2])
		if _, ok := visitedPlayer[player]; !ok {
			visitedPlayer[player] = struct{}{}
			result.Requests = append(result.Requests, engine.Request{
				URL: url,
				ParserFunc: func(content []byte) engine.ParseResult {
					return ParseProfile(content, player, url)
				},
			})
		} else {
			log.Printf("Player already exist: %s", player)
		}
	}
	return result
}
