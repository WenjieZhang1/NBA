package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

const teamlistRe = `<a href="(https://nba.hupu.com/teams/[a-zA-Z0-9]+)"[^>]*>[^<]+</a>`

// ParseTeamlist to get team info
func ParseTeamlist(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(teamlistRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	visitedTeam := make(map[string]struct{})
	for _, match := range matches {
		team := strings.TrimPrefix(string(match[1]), "https://nba.hupu.com/teams/")
		if _, ok := visitedTeam[team]; !ok {
			visitedTeam[team] = struct{}{}
			result.Items = append(result.Items, "Team: "+team)
			result.Requests = append(result.Requests, engine.Request{
				URL:        string(match[1]),
				ParserFunc: ParseTeam})
		}
	}
	return result
}
