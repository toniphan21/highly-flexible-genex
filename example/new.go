package example

//go:generate go run ../main.go

type Repository interface {
}

type Mailer interface {
}

type Dispatcher interface {
}

type Status int

const (
	StatusInactive Status = iota // inactive
	StatusActive
)
