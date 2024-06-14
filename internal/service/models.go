package service

type Error struct {
	Err     error
	Message string
	Status  int
}
