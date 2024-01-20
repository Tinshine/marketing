package item

type Item struct {
	Id          int    `gorm:"id"`
	ItemName    string `gorm:"item_name"`
	Description string `gorm:"description"`
}

func (Item) TableName() string {
	return "item"
}
