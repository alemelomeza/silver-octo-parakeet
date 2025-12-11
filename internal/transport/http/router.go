package httptransport

import (
    "net/http"
    "alemelomeza/silver-octo-parakeet/internal/service/auth"
)

type Router struct {
    mux       *http.ServeMux
    authSvc   auth.Service
    userHdl   *UserHandler
    taskHdl   *TaskHandler
}

func NewRouter(
    authSvc auth.Service,
    userHandler *UserHandler,
    taskHandler *TaskHandler,
) *Router {

    r := &Router{
        mux:     http.NewServeMux(),
        authSvc: authSvc,
        userHdl: userHandler,
        taskHdl: taskHandler,
    }

    r.registerRoutes()
    return r
}

func (r *Router) registerRoutes() {

    r.mux.Handle("/login", http.HandlerFunc(r.userHdl.Login))
    r.mux.Handle("/logout", http.HandlerFunc(r.userHdl.Logout))
    r.mux.Handle("/password/change", AuthMiddleware(r.authSvc, http.HandlerFunc(r.userHdl.ChangePassword)))

    adminOnly := func(h http.Handler) http.Handler {
        return AuthMiddleware(r.authSvc, RoleMiddleware("ADMIN")(h))
    }

    r.mux.Handle("/users", adminOnly(http.HandlerFunc(r.userHdl.Create)))
    r.mux.Handle("/users/list", adminOnly(http.HandlerFunc(r.userHdl.List)))
    r.mux.Handle("/users/update", adminOnly(http.HandlerFunc(r.userHdl.Update)))
    r.mux.Handle("/users/delete", adminOnly(http.HandlerFunc(r.userHdl.Delete)))

    r.mux.Handle("/tasks", adminOnly(http.HandlerFunc(r.taskHdl.Create)))
    r.mux.Handle("/tasks/update", adminOnly(http.HandlerFunc(r.taskHdl.Update)))
    r.mux.Handle("/tasks/delete", adminOnly(http.HandlerFunc(r.taskHdl.Delete)))

    execOnly := func(h http.Handler) http.Handler {
        return AuthMiddleware(r.authSvc, RoleMiddleware("EXECUTOR")(h))
    }

    r.mux.Handle("/tasks/my", execOnly(http.HandlerFunc(r.taskHdl.ListMy)))
    r.mux.Handle("/tasks/status", execOnly(http.HandlerFunc(r.taskHdl.UpdateStatus)))
    r.mux.Handle("/tasks/comment", execOnly(http.HandlerFunc(r.taskHdl.AddComment)))

    auditorOnly := func(h http.Handler) http.Handler {
        return AuthMiddleware(r.authSvc, RoleMiddleware("AUDITOR")(h))
    }
    r.mux.Handle("/tasks/all", auditorOnly(http.HandlerFunc(r.taskHdl.ListAll)))
}

func (r *Router) Handler() http.Handler {
    return r.mux
}
