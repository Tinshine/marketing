package model

type Rule struct {
	AppId  uint   `json:"app_id"`
	RuleId string `json:"rule_id"`
	ActId  string `json:"act_id"`
	// todo...
}

func (r *Rule) TableName() string {
	return "rule"
}
