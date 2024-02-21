package user

type SdkUser struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (u *SdkUser) GetId() string {
	return u.Id
}
