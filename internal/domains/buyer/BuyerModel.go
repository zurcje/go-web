package buyer

type Buyer struct {
	Id         int64  `json:"id"`
	CardNumber int64  `json: "cardNumber"`
	FirstName  string `json:"first_name"`
	LasttName  string `json:"last_name"`
}
