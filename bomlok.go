package bomlok

type Bomlok interface {
	Bomlok_GetValue(string) any
	Bomlok_Fields() []string
}

//go:generate go run github.com/anton7r/bomlok/cmd/bomlok
