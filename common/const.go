package common

const (
	DBTypeCategory = 1
	DBTypePost     = 2
	DBTypeUser     = 3
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
