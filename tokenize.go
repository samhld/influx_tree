package main

import (
	"fmt"
	"regexp"
	"strings"
)

type RuleTokenizer struct {
	// line string
	re *regexp.Regexp
}

type TokenizedRule struct {
	// words []Word
	words []Word
	ops   []Op
}

type Word struct {
	text    string
	index   int
	sibling *Word
}

type Op struct {
	text  string
	index int
}

func NewRuleTokenizer() *RuleTokenizer {
	return &RuleTokenizer{
		regexp.MustCompile(`(?P<words>[a-z A-Z _]+)|(?P<ops>[|>\s])`),
	}
}

func createBranches(tiers Tiers, fieldKeys []string) [][]string {
	fmt.Printf("tiers: %q\n", tiers)
	// var branches [][]string
	tree := make(Tree)
	tagValTiers := countTagValTiers(tiers)
	for t, tier := range tiers {
		if v, ok := tier.(*Measurement); ok {
			tree[t] = v.text
		}
	}

	// for f := range fieldKeys {
	// 	var branch []string
	// 	for t := 0; t < len(tiers); t++ {
	// 		tier := tiers[t]
	// 		switch v := tier.(type) {
	// 		case *Measurement:
	// 			branch = append(branch, v.text)
	// 		case *Key:
	// 			branch = append(branch, v.text)
	// 		case []*Value:
	// 			for _, val := range v {

	// 			}
	// 			branch = append(branch, v.text)
	// 		case *Field:
	// 			branch = append(branch, fieldKeys[f])
	// 		}
	// 	}
	// 	branches = append(branches, branch)

	}
	return branches
}

func makeNodes(kvsMap map[string][]string) []*Node {
	tagTier := 0
	var nodes []*Node
	var prevNodes []*Node
	for k, v := range kvsMap {
		kNode := &Node{k, nil, nil}
		if tagTier == 0 {
			top := &Node{"top", nil, nil}
			kNode.parent = top
		} else {
			kNode.parent = prevNodes[tagTier-1]
		}
		for _, vv := range v {
			vNode := &Node{vv, kNode, nil}
			kNode.children = append(kNode.children, vNode)
		}

		nodes = append(nodes, kNode)
	}
	return nodes
}

func countTagValTiers(tiers Tiers) int {
	count := 0
	for _, tier := range tiers {
		switch tier.(type) {
		case []*
		}
	}
	return count
}

func MapTokensToData(measAPI dataGetter, tokenizedRule *TokenizedRule) Tiers {
	tiers := make(Tiers)
	// tree := Tree{}
	tagTierTracker := 0 // increment each Tag tier; add to i to keep i aligned
	for i, word := range tokenizedRule.words {
		if word.text != "MEASUREMENT" && word.text != "FIELD" {
			tagTierCount := 0
			for tagKey, tagValues := range measAPI.keyValsMap {
				for _, val := range tagValues {
					newKey := &Key(tagKey, i+1, nil) 
					
					tiers[i+1] = append(tiers[i+1], &Key{tagKey, i+1, tiers.measurement, val})
					
				}
				
			
			}
		}
	// 	switch word.text {
	// 	case "MEASUREMENT":
	// 		tiers[i+tagTierTracker] = &Measurement{measAPI.getMeasurement(), i}
	// 	case "FIELD":
	// 		fields := measAPI.getFields()
	// 		// tiers[i] = []&Field{"FIELD", i}
	// 		tiers[i+tagTierTracker] = fields
	// 	default:
	// 		tiers[i+tagTierTracker] = &Key{word.text, i + tagTierTracker, nil, nil}
	// 		vals := measAPI.getTagKeyValues(word.text)
	// 		tagTierTracker++
	// 		var values []*Value
	// 		for _, val := range vals {
	// 			values = append(values, &Value{val, i + tagTierTracker, nil, nil})
	// 		}
	// 		tiers[i+tagTierTracker] = values
	// 	}
	// }

	return tiers
}

func (t *RuleTokenizer) Tokenize(rule string) *TokenizedRule {
	matches := t.re.FindAllStringSubmatch(rule, -1)
	tokenized := &TokenizedRule{}
	for i, match := range matches {
		if match[1] != "" { // 2nd position of match tuple represents a 'word' if not zero-value
			word := Word{match[1], i, nil}
			tokenized.words = append(tokenized.words, word)
		} else {
			op := Op{match[2], i}
			tokenized.ops = append(tokenized.ops, op)
		}
	}
	return tokenized
}

func (t *RuleTokenizer) FindSiblings(rule string) [][]string {
	detectedSibs := detectSiblingTokens(rule)
	var siblings [][]string
	for _, pipeSet := range detectedSibs {
		set := strings.Split(pipeSet[0], "|")
		siblings = append(siblings, set)
	}
	return siblings
}

func detectSiblingTokens(rule string) [][]string {
	var sibs [][]string
	re := regexp.MustCompile(`[^>|]+\|[^>|]+`)
	sibs = re.FindAllStringSubmatch(rule, -1)
	return sibs
}

func MeasIndex(line string) {
}

func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
}
