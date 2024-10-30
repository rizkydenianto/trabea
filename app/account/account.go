package account

import (
	"errors"
	"time"
	"trabea/app"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func passwordCheck(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(req ReqLogin) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": req.Email,
			"role":  req.Role,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	token, err := claims.SignedString([]byte(req.Password))
	return token, err
}

func GetLoginData(req ReqLogin) (*ResLogin, error) {
	conn, err := app.InitDB()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	rows, err := conn.Query(`
		SELECT password
		FROM users
		WHERE email = $1
	`, req.Email)
	if err != nil {
		println(err.Error())
		defer conn.Close()
		return nil, err
	}

	for rows.Next() {
		var password string
		if err := rows.Scan(&password); err != nil {
			println(err.Error())
			return nil, err
		}
		valid := passwordCheck(password, req.Password)
		if !valid {
			println(err.Error())
			err := errors.New("account not found")
			return nil, err
		}
	}

	token, err := createToken(req)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return &ResLogin{
		Token: token,
	}, nil
}
