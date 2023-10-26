package main

import (
	"fmt"
	"slices"

	"gopkg.in/yaml.v3"
)

// Photo is a struct for items in photos.yml file
type Photo struct {
	Path             string
	Width            int    `yaml:"width,omitempty"`
	Height           int    `yaml:"height,omitempty"`
	ThumbPath        string `yaml:"thumb,omitempty"`
	ThumbXOffset     int    `yaml:"thumb_x,omitempty"`
	ThumbYOffset     int    `yaml:"thumb_y,omitempty"`
	ThumbWidth       int    `yaml:"thumb_width,omitempty"`
	ThumbHeight      int    `yaml:"thumb_height,omitempty"`
	ThumbTotalWidth  int    `yaml:"thumb_total_width,omitempty"`
	ThumbTotalHeight int    `yaml:"thumb_total_height,omitempty"`
}

func (g *Generator) processPhotos(fileContent []byte) (interface{}, error) {
	var photos []Photo
	err := yaml.Unmarshal(fileContent, &photos)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling file: %v", err)
	}

	slices.Reverse(photos)

	return photos, nil
}
