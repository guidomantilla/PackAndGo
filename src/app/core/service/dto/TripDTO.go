package dto

type TripDTO struct {
	Id            int64   `json:"tripId,omitempty"`
	OriginId      int64   `json:"originId,omitempty"`
	Origin        string  `json:"origin,omitempty"`
	DestinationId int64   `json:"destinationId,omitempty"`
	Destination   string  `json:"destination,omitempty"`
	Dates         string  `json:"dates,omitempty"`
	Price         float64 `json:"price,omitempty"`
}
