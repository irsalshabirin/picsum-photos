package mock

import (
	"fmt"

	"github.com/DMarby/picsum-photos/internal/database"
)

// Provider implements a mock image storage
type Provider struct {
}

// Get returns the image data for an image id
func (p *Provider) Get(id string) (i *database.Image, err error) {
	return nil, fmt.Errorf("get error")
}

// GetRandom returns a random image
func (p *Provider) GetRandom() (i *database.Image, err error) {
	return nil, fmt.Errorf("random error")
}

// GetRandomWithSeed returns a random image based on the given seed
func (p *Provider) GetRandomWithSeed(seed int64) (i *database.Image, err error) {
	return nil, fmt.Errorf("random error")
}

// ListAll returns a list of all the images
func (p *Provider) ListAll() ([]database.Image, error) {
	return nil, fmt.Errorf("list error")
}

// List returns a list of all the images with an offset/limit
func (p *Provider) List(offset, limit int) ([]database.Image, error) {
	return nil, fmt.Errorf("list error")
}

// Shutdown shuts down the database client
func (p *Provider) Shutdown() {}
