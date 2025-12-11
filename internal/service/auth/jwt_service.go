package auth

import (
    "errors"
    "math/rand"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
    secretKey       string
    expirationHours int
}

type jwtClaims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func NewJWTService(secret string, expirationHours int) Service {
    return &jwtService{
        secretKey:       secret,
        expirationHours: expirationHours,
    }
}

func (s *jwtService) GenerateToken(userID, role string) (string, error) {
    claims := jwtClaims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.expirationHours) * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenStr string) (*TokenData, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(s.secretKey), nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(*jwtClaims)
    if !ok {
        return nil, errors.New("invalid claims")
    }

    return &TokenData{
        UserID: claims.UserID,
        Role:   claims.Role,
    }, nil
}


func (s *jwtService) HashPassword(password string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash)
}

func (s *jwtService) CheckPassword(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (s *jwtService) GenerateTempPassword() string {
    n := 10
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
