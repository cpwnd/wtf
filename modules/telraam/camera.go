package telraam

import "fmt"

const (
	telraamSegmentPath = "https://telraam.net/en/location/"
)

type CameraResponse struct {
	Cameras []Camera `json:"camera"`
}

// Camera represents a camera instance on teelram-api
type Camera struct {
	Mac               int    `json:"mac"`
	Segment_id        int    `json:"segment_id"`
	Status            string `json:"active"`
	Last_data_package string `json:"last_data_package"`
}

// CommentLink return the link to the HackerNews story comments page
func (camera *Camera) SegmentLink() string {
	return fmt.Sprintf("%s%d", telraamSegmentPath, camera.Segment_id)
}
