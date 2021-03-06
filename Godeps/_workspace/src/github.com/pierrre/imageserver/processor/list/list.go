// Package list provides a list of Image Processor
package list

import (
	"github.com/pierrre/imageserver"
	imageserver_processor "github.com/pierrre/imageserver/processor"
)

// ListProcessor represents a list of Image Processor
type ListProcessor []imageserver_processor.Processor

// Process processes the Image with the list of Image Processor
func (processor ListProcessor) Process(image *imageserver.Image, parameters imageserver.Parameters) (*imageserver.Image, error) {
	var err error
	for _, p := range processor {
		image, err = p.Process(image, parameters)
		if err != nil {
			return nil, err
		}
	}
	return image, nil
}
