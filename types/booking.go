package types

type Booking struct {
	ID                    string `json:"id"`
	UserName              string `json:"userName"`
	BookingStartTimestamp int64  `json:"bookingStartTimestamp"`
	BookingEndTimestamp   int64  `json:"bookingEndTimestamp"`
	NumberOfItems         int    `json:"numberOfItems"`
	SlotID                string `json:"slotId"`
}
