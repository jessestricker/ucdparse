package ucdparse

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type ParserResult map[string]RangeList

func Parse(r io.Reader) (ParserResult, error) {
	scan := bufio.NewScanner(r)
	result := make(ParserResult)
	for scan.Scan() {
		line := scan.Text()
		if posComment := strings.Index(line, "#"); posComment != -1 {
			line = line[:posComment] // remove trailing comment
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		posSemi := strings.Index(line, ";")
		if posSemi == -1 {
			continue
		}
		category := strings.TrimSpace(line[posSemi+1:])
		strRange := strings.TrimSpace(line[:posSemi])

		rge, err := parseRange(strRange)
		if err != nil {
			return nil, err
		}

		if rges, ok := result[category]; ok {
			result[category] = append(rges, rge)
		} else {
			result[category] = RangeList{rge}
		}
	}
	return result, nil
}

func parseRange(str string) (Range, error) {
	const sep = ".."
	posSep := strings.Index(str, sep)
	if posSep == -1 {
		v, err := strconv.ParseUint(str, 16, 32)
		if err != nil {
			return Range{}, err
		}
		return NewRangeSingle(uint32(v)), nil
	}
	v1, err := strconv.ParseUint(str[:posSep], 16, 32)
	if err != nil {
		return Range{}, err
	}
	v2, err := strconv.ParseUint(str[posSep+len(sep):], 16, 32)
	if err != nil {
		return Range{}, err
	}
	return NewRange(uint32(v1), uint32(v2)), nil
}
