package models

// AccessTime -...
type AccessTime struct {
	AccessTime int64  `json:"access_time"`
	URL        string `json:"url"`
}

// CounterStats - ...
type CounterStats struct {
	Counter int64  `json:"counter"`
	Handler string `json:"handler"`
}

// Admin Auth cred - ...
type AuthCred struct {
	User string
	Pass string
}
