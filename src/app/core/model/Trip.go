package model

type Trip struct {
	Id            int64   `json:"tripId,omitempty"`
	OriginId      int64   `json:"originId,omitempty"`
	DestinationId int64   `json:"destinationId,omitempty"`
	Dates         string  `json:"dates,omitempty"`
	Price         float64 `json:"price,omitempty"`
}

type City struct {
	Id   int64  `json:"cityId,omitempty"`
	Name string `json:"city,omitempty"`
}
