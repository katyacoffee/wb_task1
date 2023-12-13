package models

type DeliveryOrder struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     int64  `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

//TODO: payment, item models

type Order struct {
	DelOrder DeliveryOrder
	//TODO
}

type DbOrder struct {
	DelOrderID int64
	//TODO
}
