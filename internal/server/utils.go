package server

import (
	"diffinlist/internal/auth"
	"diffinlist/internal/middleware"
	"net/http"
)

func getPageAuthInfo(r *http.Request) auth.AuthenticatedPage {
	user := middleware.UserFromContext(r.Context())
	pageData := auth.AuthenticatedPage{
		IsAuthorized: false,
		User: auth.UserUI{
			PathToPfp: "",
			Username:  "",
		},
	}

	if user != nil {
		pageData.IsAuthorized = true
		pageData.User = auth.UserUI{
			PathToPfp: user.PathToPfp,
			Username:  user.Username,
		}
	}

	return pageData
}
