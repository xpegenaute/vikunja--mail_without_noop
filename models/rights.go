package models

// Rights defines rights methods
type Rights interface {
	IsAdmin(*User) bool
	CanWrite(*User) bool
	CanRead(*User) bool
	CanDelete(*User, int64) bool
	CanUpdate(*User, int64) bool
	CanCreate(*User, int64) bool
}
