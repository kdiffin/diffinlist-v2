package auth

// user
type User struct {
	ID                string
	PathToPfp         string
	Username          string
	DefaultPlaylistId string
}

type Author struct {
	PathToPfp string
	Username  string
}

type UserUI struct {
	PathToPfp string
	Username  string
}

type AuthenticatedPage struct {
	IsAuthorized bool
	User         UserUI
}
