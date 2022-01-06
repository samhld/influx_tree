package main

import (
	"reflect"
	"testing"
)

func TestTokenizationRule(t *testing.T) {
	// line := "stats,region=us-west,app=cart count=12"
	ruleTokenizer := NewRuleTokenizer()
	t.Run("test parsing rule to map with indexed tokens", func(t *testing.T) {
		rule := "MEASUREMENT>region>app>FIELD"
		wantedTokenizedRule := &TokenizedRule{
			[]Word{
				{"MEASUREMENT", 0, nil},
				{"region", 2, nil},
				{"app", 4, nil},
				{"FIELD", 6, nil},
			},
			[]Op{
				{">", 1},
				{">", 3},
				{">", 5},
			},
		}
		gotTokenizedRule := ruleTokenizer.Tokenize(rule)
		assertEqual(t, gotTokenizedRule, wantedTokenizedRule)
	})
	t.Run("test making TokenizedRule with pipe op", func(t *testing.T) {
		rule := "MEASUREMENT>region>host|app>FIELD"
		wantedTokenizedRule := &TokenizedRule{
			[]Word{
				{"MEASUREMENT", 0, nil},
				{"region", 2, nil},
				{"host", 4, nil},
				{"app", 6, nil},
				{"FIELD", 8, nil},
			},
			[]Op{
				{">", 1},
				{">", 3},
				{"|", 5},
				{">", 7},
			},
		}
		gotTokenizedRule := ruleTokenizer.Tokenize(rule)
		assertEqual(t, gotTokenizedRule, wantedTokenizedRule)
	})
	t.Run("test siblings", func(t *testing.T) {
		rule := "MEASUREMENT>region>host|app>sib1|sib2>FIELD"
		ruleTokenizer := NewRuleTokenizer()
		want := [][]string{
			{"host", "app"},
			{"sib1", "sib2"},
		}
		got := ruleTokenizer.FindSiblings(rule)

		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot %q\nWant %q\n", got, want)
	}
}
