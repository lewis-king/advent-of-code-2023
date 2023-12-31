package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day19/input/input.txt")
	workflows, parts := parseWorkflows(lines)
	fmt.Println(workflows)
	fmt.Println(parts)
	runningTotal := 0
	for _, part := range parts {
		derivedKey := "in"
		for derivedKey != "A" && derivedKey != "R" || derivedKey == "in" {
			rules := strings.Split(workflows[derivedKey], ",")
			noMatch := rules[len(rules)-1]
			for i, rule := range rules {
				if i == len(rules)-1 {
					derivedKey = noMatch
					break
				}
				variable := strings.Split(rule, "")[0]
				operator := strings.Split(rule, "")[1]
				unsanitisedValue := strings.Replace(strings.Split(rule, ":")[0], variable, "", 1)
				sanitisedValue := strings.Replace(strings.Replace(unsanitisedValue, ">", "", 1), "<", "", 1)
				value, _ := strconv.Atoi(sanitisedValue)
				if operator == ">" {
					if part[variable] > value {
						derivedKey = strings.Split(rule, ":")[1]
						break
					}
				} else if operator == "<" {
					if part[variable] < value {
						derivedKey = strings.Split(rule, ":")[1]
						break
					}
				} else if operator == ">" {
					if part[variable] > value {
						derivedKey = strings.Split(rule, ":")[1]
						break
					}
				}
			}
		}
		if derivedKey == "A" {
			for _, value := range part {
				runningTotal += value
			}
			fmt.Println(runningTotal)
		}
	}
}

func parseWorkflows(lines []string) (map[string]string, []map[string]int) {
	workflows := make(map[string]string)
	parts := make([]map[string]int, 0)

	seenDivider := false
	for _, line := range lines {
		if line == "" {
			seenDivider = true
			continue
		}
		if seenDivider {
			partsLine := strings.Replace(strings.Replace(line, "{", "", 1), "}", "", 1)
			individualParts := strings.Split(partsLine, ",")
			partAndValues := make(map[string]int)
			for _, part := range individualParts {
				value, _ := strconv.Atoi(strings.Split(part, "=")[1])
				partAndValues[strings.Split(part, "=")[0]] = value
			}
			parts = append(parts, partAndValues)
		} else {
			key := strings.Split(line, "{")[0]
			value := strings.Replace(strings.Split(line, "{")[1], "}", "", 1)
			workflows[key] = value
		}
	}
	return workflows, parts
}
