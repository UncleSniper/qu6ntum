package core

type Definition interface {
	Location() *Location
	Define(*Engine, string) error
}
