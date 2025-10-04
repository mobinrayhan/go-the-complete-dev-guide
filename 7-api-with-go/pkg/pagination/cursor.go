package pagination

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

type Cursor struct {
	CreateAt time.Time `json:"createdAt"`
	Id       int       `json:"id"`
}

func EncodeCursor(c Cursor) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func DecodeCursor(s string) (Cursor, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, err
	}

	var c Cursor

	if err := json.Unmarshal(b, &c); err != nil {
		return Cursor{}, err
	}

	return c, nil
}

func (c Cursor) IsEmpty() bool {
	return c.CreateAt.IsZero() && c.Id == 0
}
