package aoc2020

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Message struct {
	rules []Rules
}

type Rules struct {
	index                uint32
	subRules             []SubRules
	relation             []int32
	normalizedStringRule string
}

type SubRules struct {
	stringsRule []string
	numbersRule []uint32
}

type MessageInterface interface {
	BuildRuleTable(rline string)
	GetRules() []Rules
	GetRuleAt(index int) (Rules, error)
	TopDownMatchTraversal(t string, rule uint32) bool
}

const (
	OR = iota
)

func NewMessage() MessageInterface {
	msgCtx := &Message{rules: make([]Rules, 0)}
	return msgCtx
}

func _isDigit(buf string) bool {
	for i := 0; i < len(buf); i++ {
		if buf[i] < 0x30 || buf[i] > 0x39 {
			return false
		}
	}

	return true
}

func _isLowercaseString(buf string) bool {
	for i := 0; i < len(buf); i++ {
		if buf[i] < 97 || buf[i] > 122 {
			return false
		}
	}

	return true
}

func _isUppercaseString(buf string) bool {
	for i := 0; i < len(buf); i++ {
		if buf[i] < 65 || buf[i] > 90 {
			return false
		}
	}

	return true
}

func _isLogicalComparisonOperator(buf string) bool {
	switch buf {
	case "|":
		return true
	}

	return false
}

func _convertLogicalOperatorToAssociatedNumber(buf string) int32 {
	switch buf {
	case "|":
		return OR
	}

	return -1
}

func _parseRuleLine(rule string, m *Message) {
	reg := regexp.MustCompile(`([0-9]+|\||[a-zA-Z]+)`)
	matches := reg.FindAllString(rule, -1)
	index, _ := strconv.Atoi(matches[0])
	tmp := []uint32{}
	stmp := []string{}

	m.rules = append(
		m.rules,
		Rules{
			index:                uint32(index),
			subRules:             make([]SubRules, 0),
			relation:             make([]int32, 0),
			normalizedStringRule: "",
		},
	)

	rLen := len(m.rules)

	for i := 1; i < len(matches); i++ {
		if _isDigit(matches[i]) {
			val, _ := strconv.Atoi(matches[i])
			tmp = append(tmp, uint32(val))
			continue
		}

		if _isLowercaseString(matches[i]) || _isUppercaseString(matches[i]) {
			stmp = append(stmp, matches[i])
			continue
		}

		if _isLogicalComparisonOperator(matches[i]) {
			m.rules[rLen-1].relation = append(
				m.rules[rLen-1].relation,
				_convertLogicalOperatorToAssociatedNumber(matches[i]),
			)
		}

		m.rules[rLen-1].subRules = append(m.rules[rLen-1].subRules, SubRules{numbersRule: tmp, stringsRule: stmp})
		m.rules[rLen-1].normalizedStringRule = strings.Join(stmp, "")
		tmp = []uint32{}
		stmp = []string{}
	}

	if len(tmp) != 0 || len(stmp) != 0 {
		m.rules[rLen-1].subRules = append(m.rules[rLen-1].subRules, SubRules{numbersRule: tmp, stringsRule: stmp})
		m.rules[rLen-1].normalizedStringRule = strings.Join(stmp, "")
		tmp = []uint32{}
		stmp = []string{}
	}
}

func (m *Message) BuildRuleTable(rline string) {
	_parseRuleLine(rline, m)
}

func (m *Message) GetRules() []Rules {
	return m.rules
}

func (m *Message) GetRuleAt(index int) (Rules, error) {
	for i := 0; i < len(m.rules); i++ {
		if m.rules[i].index == uint32(index) {
			return m.rules[i], nil
		}
	}

	return Rules{}, fmt.Errorf("No rule with index [%v]\n", index)
}

func (m *Message) GetRuleRefAt(index int) (*Rules, error) {
	for i := 0; i < len(m.rules); i++ {
		if m.rules[i].index == uint32(index) {
			return &m.rules[i], nil
		}
	}

	return &Rules{}, fmt.Errorf("NO rule with index [%v]\n", index)
}

func _findTerminalRules(m *Message) []uint32 {
	termRules := make([]uint32, 0)

	for i := 0; i < len(m.rules); i++ {
		if len(m.rules[i].subRules) == 1 && len(m.rules[i].subRules[0].stringsRule) == 1 {
			termRules = append(termRules, m.rules[i].index)
		}
	}

	return termRules
}

func _contains(index uint32, rules []uint32) bool {
	for _, rule := range rules {
		if rule == index {
			return true
		}
	}

	return false
}

func _multiContains(x []uint32, y []uint32) bool {
	if len(x) == 0 {
		return false
	}

	for _, _x := range x {
		if !_contains(_x, y) {
			return false
		}
	}

	return true
}

func _combine(a string, b string) string {
	aSplitted := strings.Split(a, "|")
	bSplitted := strings.Split(b, "|")
	result := make([]string, 0)

	var buffer bytes.Buffer

	for _, _a := range aSplitted {
		for _, _b := range bSplitted {
			buffer.WriteString(_a)
			buffer.WriteString(_b)
			result = append(result, buffer.String())
			buffer.Reset()
		}
	}

	return strings.Join(result, "|")
}

func _filterRules__part1(m *Message) []uint32 {
	termRules := _findTerminalRules(m)
	normRules := make([]uint32, 0)
	isPure := true

	for _, rule := range m.rules {
		for _, subRule := range rule.subRules {
			if !_multiContains(subRule.numbersRule, termRules) {
				isPure = false
				break
			}
		}

		if isPure && len(rule.normalizedStringRule) == 0 {
			normRules = append(normRules, rule.index)
		}

		isPure = true
	}

	return normRules
}

func _filterRules__part2(m *Message) []uint32 {
	termRules := _findTerminalRules(m)
	normRules := make([]uint32, 0)
	isPure := true

	for _, idx := range _filterRules__part1(m) {
		termRules = append(termRules, idx)
	}

	for _, rule := range m.rules {
		for _, subRule := range rule.subRules {
			if !_multiContains(subRule.numbersRule, termRules) {
				isPure = false
				break
			}
		}

		if isPure && len(rule.normalizedStringRule) == 0 {
			normRules = append(normRules, rule.index)
		}

		isPure = true
	}

	return normRules
}

func _filterRules__part3(m *Message) []uint32 {
	termRules := _findTerminalRules(m)
	normRules := make([]uint32, 0)

	// append filtered index from _filterRules__part1
	// function call.
	for _, idx := range _filterRules__part1(m) {
		termRules = append(termRules, idx)
	}

	// append filtered index from _filterRules__part2
	// function call.
	for _, idx := range _filterRules__part2(m) {
		termRules = append(termRules, idx)
	}

	for _, rule := range m.rules {
		if !_contains(rule.index, termRules) {
			normRules = append(normRules, rule.index)
		}
	}

	return normRules
}

func _normalizeRules__step1(m *Message) {
	var buffer bytes.Buffer

	for _, pr := range _filterRules__part1(m) {
		rule, _ := m.GetRuleRefAt(int(pr))

		for i, _ := range rule.subRules {
			buffer.WriteString(rule.normalizedStringRule)

			if i > 0 && len(rule.relation) > 0 {
				if rule.relation[i-1] == OR {
					buffer.WriteString("|")
				}
			}

			for _, index := range rule.subRules[i].numbersRule {
				tRule, _ := m.GetRuleAt(int(index))
				buffer.WriteString(strings.Join(tRule.subRules[0].stringsRule, ""))
			}

			rule.normalizedStringRule = buffer.String()
			buffer.Reset()
		}

		//fmt.Printf("%v: %v\n", rule.index, rule.normalizedStringRule)
	}
}

func _normalizeRules__step2(m *Message) {
	var buffer bytes.Buffer

	for _, pr := range _filterRules__part2(m) {
		rule, _ := m.GetRuleRefAt(int(pr))

		var tmp Rules

		for i, _ := range rule.subRules {
			buffer.WriteString(rule.normalizedStringRule)

			if i > 0 && len(rule.relation) > 0 {
				if rule.relation[i-1] == OR {
					buffer.WriteString("|")
				}
			}

			for x, index := range rule.subRules[i].numbersRule {
				if x == 0 {
					tmp, _ = m.GetRuleAt(int(index))
					continue
				}

				tRule, _ := m.GetRuleAt(int(index))
				buffer.WriteString(_combine(tmp.normalizedStringRule, tRule.normalizedStringRule))
			}

			rule.normalizedStringRule = buffer.String()
			buffer.Reset()
		}

		//fmt.Printf("%v: %v\n", rule.index, rule.normalizedStringRule)
	}
}

func _normalizeRules__step3(m *Message) {
}

func _normalizeRules(m *Message) {
	_normalizeRules__step1(m)
	_normalizeRules__step2(m)
	_normalizeRules__step3(m)
}

func (m *Message) TopDownMatchTraversal(t string, rule uint32) bool {
	_normalizeRules(m)
	return true
}
