package auth

type Service interface {
    GenerateToken(userID, role string) (string, error)
    ValidateToken(tokenStr string) (*TokenData, error)

    HashPassword(password string) string
    CheckPassword(password, hash string) bool

    GenerateTempPassword() string
}

type TokenData struct {
    UserID string
    Role   string
}
