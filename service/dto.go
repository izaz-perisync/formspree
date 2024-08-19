package service

import "time"

type Project struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Templates []int64 `json:"templates"`
	TimeZone  string  `json:"timeZone"`
	Domain    string  `json:"domain"`
	Time
}

type Form struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	LinkedEmail string `json:"linkedEmail"`
	Project     int64  `json:"project"`
	Key         string `json:"key"`
	Time
}

type Filter struct {
	Id              int64  `schema:"id"`
	SearchKey       string `schema:"searchKey"` // SearchKey represents the search key for filtering data.
	Role            int64  `schema:"role"`      // Role represents the role for filtering data.
	UserId          int64  `schema:"userId"`    // UserId represents the user ID for filtering data.
	SessionId       int64  `schema:"sessionId"` // SessionId represents the session ID for filtering data.
	NotificationId  int64  `schema:"Id"`        // NotificationId represents the notification ID for filtering data.
	Pagination             // Embedding Pagination for pagination settings.
	Active          bool   `schema:"active"`
	StoreSubmission bool   `schema:"storeSubmission"`
	Name            string `schema:"name"`
	ReCAPTCHA       bool   `schema:"reCAPTCHA"`
}

// Pagination represents the pagination settings.
type Pagination struct {
	Size int `schema:"size"` // Size represents the number of items per page.
	Page int `schema:"page"` // Page represents the page number.
}

type Time struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetPagination returns the pagination settings from the filter.
func (f *Filter) GetPagination() (size, offset int) {
	size = f.Pagination.Size
	if f.Pagination.Size == 0 {
		size = 10
	}

	if f.Pagination.Page > 1 {
		offset = (f.Pagination.Page - 1) * size
	}

	return
}
