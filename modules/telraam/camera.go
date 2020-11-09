package telraam

import "fmt"

const (
	telraamSegmentPath = "https://telraam.net/en/location/"
)

// Story represents a story submission on HackerNews
type Camera struct {
	mac               int    `json:"mac"`
	segment_id        int    `json:"segment_id"`
	status            string `json:"active"`
	last_data_package string `json:"last_data_package"`
}

// CommentLink return the link to the HackerNews story comments page
func (camera *Camera) SegmentLink() string {
	return fmt.Sprintf("%s%d", telraamSegmentPath, camera.segment_id)
}
