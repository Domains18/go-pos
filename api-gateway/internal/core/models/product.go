package models


type Product struct {
	SKU            string      `json:"sku"`
	Name           string      `json:"name"`
	Price          float64     `json:"price"`
	TaxRate        float64     `json:"tax_rate"`
	TaxedPrice     float64     `json:"taxed_price"`
	Image          string      `json:"image"`
	ProductCategory ProductGroup `json:"product_category"`
	SalePrice      float64     `json:"sale_price"`
	TaxedSalePrice float64     `json:"taxed_sale_price"`
	Quantity       float64     `json:"quantity"`
	Available      bool        `json:"available"`
}

type ProductGroup struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Parent   int    `json:"parent"`
	IsGroup  bool   `json:"is_group"`
	Disabled bool   `json:"disabled"`
}