package middleware

import (
	"context"
	"diffinlist/internal/auth"
	db "diffinlist/internal/db/generated"
	"diffinlist/internal/utils"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type contextKey string

const userContextKey = contextKey("user")

// Helper to get user from context
func UserFromContext(ctx context.Context) *auth.User {
	user, ok := ctx.Value(userContextKey).(*auth.User)
	if !ok {
		return nil
	}
	return user
}

func getUserIDFromRequest(queries *db.Queries, r *http.Request) (pgtype.UUID, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		if err == http.ErrNoCookie {
			return pgtype.UUID{}, err
		}
		utils.LogError("error at COOKIE:", err)
		return pgtype.UUID{}, fmt.Errorf("error getting cookie: %w", err)
	}

	sesh, errSesh := queries.GetSessionValues(r.Context(), cookie.Value)
	if errSesh != nil {
		utils.LogError("error at session value from cookie:", err)
		return pgtype.UUID{}, errSesh
	}
	return sesh.UserID, nil
}

func getUserInfoFromSessionId(queries *db.Queries, r *http.Request) (*auth.User, error) {
	userId, err := getUserIDFromRequest(queries, r)
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, err
		}
		utils.LogError("error at userid from request:", err)
		return nil, fmt.Errorf("%w: %v", utils.ErrUnauthorized, err)
	}

	userDb, errUser := queries.GetUserInfo(r.Context(), userId)
	if errUser != nil {
		utils.LogError("error at user info:", errUser)
		return nil, fmt.Errorf("%w: %v", utils.ErrUnauthorized, errUser)
	}
	user := auth.User{
		ID:                userId.String(),
		PathToPfp:         userDb.PathToPfp,
		Username:          userDb.Username,
		DefaultPlaylistId: userDb.DefaultPlaylistID.String(),
	}
	return &user, nil

}

func HandlerWithUser(queries *db.Queries, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := getUserInfoFromSessionId(queries, r)
		if err != nil {
			utils.LogError("error at session value from cookie:", err)
			w.Header().Add("hx-redirect", "/login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithOptionalUser(queries *db.Queries, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := getUserInfoFromSessionId(queries, r)

		if err == nil && user != nil {
			// Create a new context with the user value
			ctx := context.WithValue(r.Context(), userContextKey, user)
			// Serve the next handler with the new context
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		next.ServeHTTP(w, r)
	})
}
