package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtinme  int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtinme != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtinme)
	}

	aux := struct {
		ID        int64     `json:"id"`
		CreatedAt time.Time `json:"-"`
		Title     string    `json:"title"`
		Year      int32     `json:"year,omitempty"`
		Runtinme  string    `json:"runtime,omitempty"`
		Genres    []string  `json:"genres,omitempty"`
		Version   int32     `json:"version"`
	}{
		ID:       m.ID,
		Title:    m.Title,
		Year:     m.Year,
		Runtinme: runtime,
		Genres:   m.Genres,
		Version:  m.Version,
	}

	return json.Marshal(aux)
}
