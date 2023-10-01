package domain

import "time"

// initializing a data structure to keep the scraped data
type Article struct {
	ID           int
	SourceID     int
	SourceName   string
	URL          string
	Headline     string
	Summary      string
	ImageURL     string
	NormTags     []string
	ScrappedDate time.Time
}
