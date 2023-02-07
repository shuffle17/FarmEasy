package service

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type NewFarmer struct {
	Id        uint   `db:"id" json:"id"`
	FirstName string `db:"fname" json:"fname"`
	LastName  string `db:"lname" json:"lname"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	Password  string `db:"password" json:"password"`
}

type NewMachine struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	BaseHourlyCharge uint   `json:"base_hourly_charge"`
	OwnerId          uint   `json:"owner_id"`
}
