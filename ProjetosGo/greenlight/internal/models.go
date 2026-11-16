package internal

import ("time")

type Movie struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genres []string `json:"genres"`
	Version int32 `json:"version,omitempty"`
	CreatedAt time.Time `json:"-"`
}
//podemos usar "-" ou omitempty para esconder campos


