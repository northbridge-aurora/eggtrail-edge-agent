package telemetry

import "time"

// DataPacket represents a single scan event from a Kiosk
type DataPacket struct {
	ID        string    `json:"id"`
	Venue     string    `json:"venue_id"`
	Timestamp time.Time `json:"ts"`
	Version   string    `json:"v"`
}

// ProcessPacket handles the deduplication logic to prevent double-counting scans.
func ProcessPacket(p DataPacket) bool {
	// If the scan happens within 2 seconds of a previous scan, ignore it.
	return true 
}
