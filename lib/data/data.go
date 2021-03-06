package data

import "sync"

// Response model the response to API call
type Response struct {
	Depth int         `json:"depth" description:"Depth of URL from seed website"`
	Title string      `json:"title" description:"Title of a site fetched by the crawler"`
	URL   string      `json:"url" description:"URL of a site fetched by the crawler"`
	Nodes []*Response `json:"nodes" description:"Children of a site fetched by the crawler"`
}

// Visited keeps track of visited sites to avoid loops
type Visited struct {
	sync.RWMutex
	M map[string]bool
}
