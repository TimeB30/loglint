package linter

type Rule interface {
	Check(message *string) bool
	Name() string
}

type RuleLinter struct {
	rules []Rule
}

func NewRuleLinker(r []Rule) *RuleLinter {
	return &RuleLinter{rules: r}
}

func (rl *RuleLinter) Run(message *string) map[string]bool {
	results := make(map[string]bool)
	for _, rule := range rl.rules {
		results[rule.Name()] = rule.Check(message)
	}
	return results
}
