package models


type SalesInvoice struct {
	ClientID int              `json:"client_id"`
	Items    []SalesInvoiceItem `json:"items"`
	Customer Customer         `json:"customer"`
}


type SalesInvoiceItem struct {
	Product            Product `json:"product"`
	Quantity           float64 `json:"quantity"`
	SellingPrice       float64 `json:"selling_price"`
	TaxedSellingPrice float64 `json:"taxed_selling_price"`
}
