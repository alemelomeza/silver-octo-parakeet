package httptransport

import (
    "context"
    "encoding/json"
    "net/http"
    "strings"

    "alemelomeza/silver-octo-parakeet/internal/service/auth"
)

type ctxKey string

const (
    CtxUserID ctxKey = "user_id"
    CtxRole   ctxKey = "role"
)

func AuthMiddleware(authSvc auth.Service, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing token"})
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")

        data, err := authSvc.ValidateToken(token)
        if err != nil {
            writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
            return
        }

        ctx := context.WithValue(r.Context(), CtxUserID, data.UserID)
        ctx = context.WithValue(ctx, CtxRole, data.Role)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func RoleMiddleware(roles ...string) func(http.Handler) http.Handler {

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            role := r.Context().Value(CtxRole)
            if role == nil {
                writeJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
                return
            }

            for _, allowed := range roles {
                if role.(string) == allowed {
                    next.ServeHTTP(w, r)
                    return
                }
            }

            writeJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
        })
    }
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}
