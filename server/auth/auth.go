package auth

import (
	"context"
	"fmt"
	"time"

	"example.com/model"

	"golang.org/x/crypto/bcrypt"
)

func Register(user *model.User) (int64, error) {
	var newID int64
	if len(user.Password) < 8 || len(user.Password) > 56 {
		return 0, fmt.Errorf("register %v: The password must contain between 8 and 56 characters", user.Name)
	}
	err := model.DB.QueryRow(context.Background(), "SELECT user_id FROM users WHERE email = $1", user.Email).Scan(&newID)
	if err == nil {
		return 0, fmt.Errorf("register %+v: Email already exists", model.User{Name: user.Name, Email: user.Email})
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("register %v: %v", user.Name, err)
	}
	now_time := time.Now().UTC()
	err = model.DB.QueryRow(context.Background(), `
	INSERT INTO users
		(name, password, email, create_time, update_time)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING user_id`, user.Name, hashedPwd, user.Email, now_time, now_time.UTC()).Scan(&newID)
	if err != nil {
		return 0, fmt.Errorf("register %v: %v", user.Name, err)
	}
	return newID, nil
}

func Login(user *model.User) (model.User, error) {
	var hashedPwd string
	err := model.DB.QueryRow(context.Background(), `
	SELECT 
		user_id, name, password, email, create_time, update_time
	FROM 
		users 
	WHERE 
		email = $1`, user.Email).Scan(&user.User_id, &user.Name, &hashedPwd, &user.Email, &user.Create_time, &user.Update_time)
	if err != nil {
		return model.User{}, fmt.Errorf("Login %+v: User doesn't exist", model.User{Name: user.Name, Email: user.Email})
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(user.Password))
	if err != nil {
		return model.User{}, fmt.Errorf("Login %+v: Wrong password", model.User{Name: user.Name, Email: user.Email})
	}
	return model.User{User_id: user.User_id, Name: user.Name, Email: user.Email, Create_time: user.Create_time, Update_time: user.Update_time}, nil
}
