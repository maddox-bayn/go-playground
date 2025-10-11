package strutils

import (
	"encoding/json"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func FilterExpensivePrice(data []byte, threhold float64) ([]Product, error) {
	var product []Product
	err := json.Unmarshal(data, &product)
	if err != nil {
		return nil, err
	}

	var result []Product
	for _, data := range product {
		if data.Price > threhold {
			result = append(result, data)
		}
	}
	return result, nil
}
