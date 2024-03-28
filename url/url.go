package url

import (
	"context"
	"crypto/rand"
	"encoding/base64"
)

type URL struct {
	ID  string
	URL string
}

type ShortenParams struct {
	URL string
}

//encore:api public method=POST path=/url
func Shorten(ctx context.Context, p *ShortenParams) (*URL, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	}
	return &URL{ID: id, URL: p.URL}, nil
}

func generateID() (string, error) {
	var data [6]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}
