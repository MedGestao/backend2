package model

type CellphoneUser struct {
	id     int
	userId int
	number string
}

func NewCellphoneUser(number string) CellphoneUser {
	return CellphoneUser{
		number: number,
	}
}

func (c CellphoneUser) GetId() int {
	return c.id
}

func (c CellphoneUser) GetUserId() int {
	return c.userId
}

func (c *CellphoneUser) SetUserId(userId int) {
	c.userId = userId
}

func (c CellphoneUser) GetNumber() string {
	return c.number
}

func (c *CellphoneUser) SetNumber(number string) {
	c.number = number
}
