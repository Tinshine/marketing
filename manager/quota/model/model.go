package model

type Quota struct {
	// todo...
}

func (quota *Quota) TableName() string {
	return "quota"
}

func (q *Quota) Avaiable() bool {
	// todo...
	return false
}
