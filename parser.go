package setupfiles

import (
	"math"
	"strings"
)

func parse(source string) ([]*file, error) {
	var files []*file
	for _, lines := range split(strings.Split(source, "\n")) {
		files = append(files, &file{
			lines[0], strings.Join(unindent(lines[1:]), "\n"),
		})
	}
	return files, nil
}

func split(lines []string) [][]string {
	xs := [][]string{}
	for _, line := range lines {
		if len(line) > 0 && line[0] != ' ' && line[0] != '\t' {
			xs = append(xs, []string{line})
		} else if len(xs) > 0 || len(line) > 0 {
			if len(xs) == 0 {
				xs = append(xs, []string{})
			}
			xs[len(xs)-1] = append(xs[len(xs)-1], line)
		}
	}
	return xs
}

func unindent(lines []string) []string {
	minCnt := minIndentCount(lines)
	for i, line := range lines {
		var cnt, k int
		for j, c := range line {
			if c == ' ' {
				cnt++
			} else if c == '\t' {
				cnt += 4
			}
			if cnt >= minCnt {
				k = j + 1
				break
			}
		}
		lines[i] = line[k:]
	}
	return lines
}

func minIndentCount(lines []string) int {
	cnt := math.MaxInt32
	for _, line := range lines {
		if newCnt := indentCount(line); 0 < newCnt && newCnt < cnt {
			cnt = newCnt
		}
	}
	if cnt == math.MaxInt32 {
		return 0
	}
	return cnt
}

func indentCount(line string) int {
	var cnt int
	for _, c := range line {
		if c == ' ' {
			cnt++
		} else if c == '\t' {
			cnt += 4
		} else {
			return cnt
		}
	}
	return cnt
}
