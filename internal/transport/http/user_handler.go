package httptransport

import (
	"alemelomeza/silver-octo-parakeet/internal/domain/user"
	useruc "alemelomeza/silver-octo-parakeet/internal/usecase/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
    LoginUC         *useruc.LoginUseCase
    CreateUC        *useruc.CreateUserUseCase
    UpdateUC        *useruc.UpdateUserUseCase
    DeleteUC        *useruc.DeleteUserUseCase
    ListUC          *useruc.ListUsersUseCase
    ChangePwdUC     *useruc.ChangePasswordUseCase
    LogoutUC        *useruc.LogoutUseCase
}

func NewUserHandler(
    login *useruc.LoginUseCase,
    create *useruc.CreateUserUseCase,
    update *useruc.UpdateUserUseCase,
    del *useruc.DeleteUserUseCase,
    list *useruc.ListUsersUseCase,
    change *useruc.ChangePasswordUseCase,
    logout *useruc.LogoutUseCase,
) *UserHandler {

    return &UserHandler{
        LoginUC:     login,
        CreateUC:    create,
        UpdateUC:    update,
        DeleteUC:    del,
        ListUC:      list,
        ChangePwdUC: change,
        LogoutUC:    logout,
    }
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&body)

    token, err := h.LoginUC.Execute(r.Context(), body.Username, body.Password)
    if err != nil {
        writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, map[string]string{
        "token": token,
    })
}

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Old string `json:"old_password"`
        New string `json:"new_password"`
    }
    json.NewDecoder(r.Body).Decode(&body)

    userID := r.Context().Value(CtxUserID).(string)

    err := h.ChangePwdUC.Execute(r.Context(), userID, body.Old, body.New)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, map[string]string{"status": "password updated"})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Username string `json:"username"`
        Role     string `json:"role"`
    }
    json.NewDecoder(r.Body).Decode(&body)

    role := r.Context().Value(CtxRole).(string)

    u, err := h.CreateUC.Execute(
        r.Context(),
        user.Role(role),
        body.Username,
        user.Role(body.Role),
    )

    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusCreated, u)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
    var body struct {
        UserID   string `json:"user_id"`
        Username string `json:"username"`
        Role     string `json:"role"`
    }
    json.NewDecoder(r.Body).Decode(&body)

    role := r.Context().Value(CtxRole).(string)

    u := &user.User{
        ID:       body.UserID,
        Username: body.Username,
        Role:     user.Role(body.Role),
    }

    err := h.UpdateUC.Execute(
        r.Context(),
        user.Role(role),
        u,
    )

    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, u)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
    var body struct {
        UserID string `json:"user_id"`
    }
    json.NewDecoder(r.Body).Decode(&body)

    role := r.Context().Value(CtxRole).(string)

    err := h.DeleteUC.Execute(r.Context(), user.Role(role), body.UserID)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, map[string]string{"status": "user deleted"})
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
    role := r.Context().Value(CtxRole).(string)

    users, err := h.ListUC.Execute(r.Context(), user.Role(role))
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, users)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")

    err := h.LogoutUC.Execute(r.Context(), token)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, map[string]string{"status": "logged out"})
}

