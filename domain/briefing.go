package domain

import "time"

// initializing a data structure to keep the scraped data
type Briefing struct {
	Topics []Topic
	Date   time.Time
}

// initializing a data structure to keep the scraped data
type Topic struct {
	Headline string
	Summary  string
	Sources  []int
}
