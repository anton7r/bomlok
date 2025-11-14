package bomlok

type Bomlok interface {
	Bomlok_GetValue(string) any
	Bomlok_Fields() []string
}
