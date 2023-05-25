package models

// AccessTime -...
type AccessTime struct {
	AccessTime int64  `json:"access_time"`
	URL        string `json:"url"`
}

// CounterStats - ...
type CounterStats struct {
	Counter int64
	Handler string
}

// Admin Auth cred - ...
type AuthCred struct {
	User string
	Pass string
}
