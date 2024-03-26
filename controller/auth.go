package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	components "github.com/Mohammed785/goBlog/components/auth"
	"github.com/Mohammed785/goBlog/database/sqlc"
	"github.com/Mohammed785/goBlog/helpers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	queries   *sqlc.Queries
	storage   *session.Store
	validator helpers.Validator
}

func NewAuthController(queries *sqlc.Queries, storage *session.Store, validator helpers.Validator) *AuthController {
	return &AuthController{
		queries:   queries,
		storage:   storage,
		validator: validator,
	}
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	var credentials sqlc.CreateUserParams
	if err := c.Bind().Form(&credentials); err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "please validate your request body")
		return c.SendStatus(http.StatusInternalServerError)
	}
	if err := ac.validator.ValidateStruct(credentials); err != nil {
		return helpers.Render(c, components.AuthForm("/login", ac.validator.ParseValidationError(err)))
	}
	sess, err := ac.storage.Get(c)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "couldn't get session storage!! please try again.")
		return c.SendStatus(http.StatusInternalServerError)
	}
	user, err := ac.queries.GetUserByUsername(context.Background(), credentials.Username)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "user not found!!")
		return c.SendStatus(http.StatusUnauthorized)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "please check your password!!")
		return c.SendStatus(http.StatusUnauthorized)
	}
	sid := sess.ID()
	sess.Set("sid", sid)
	sess.Set("uid", user.Uid)
	sess.Set("isAdmin", user.IsAdmin.Bool)
	sess.Set("ip", c.Context().RemoteIP().String())
	sess.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
	sess.Set("ua", string(c.Request().Header.UserAgent()))

	err = sess.Save()
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "couldn't save your session!! please try again.")
		return c.SendStatus(http.StatusInternalServerError)
	}
	log.Println("isAdmin: ", user.IsAdmin.Bool)
	c.Append("HX-Location", "/")
	return c.SendStatus(http.StatusOK)
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	var credentials sqlc.CreateUserParams
	if err := c.Bind().Form(&credentials); err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	if err := ac.validator.ValidateStruct(credentials); err != nil {
		return helpers.Render(c, components.AuthForm("/register", ac.validator.ParseValidationError(err)))
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 10)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	err = ac.queries.CreateUser(context.Background(), sqlc.CreateUserParams{Username: credentials.Username, Password: string(hashedPassword)})
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	helpers.SendMsg(c, helpers.SUCCESS_MSG, "Account created successfully")
	c.Append("HX-Location", "/login")
	return c.SendStatus(http.StatusOK)
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	sess, err := ac.storage.Get(c)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "couldn't get session storage!! please try again.")
		return c.SendStatus(http.StatusInternalServerError)
	}
	err = sess.Destroy()
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "couldn't destroy session!! please try again.")
		return c.SendStatus(http.StatusInternalServerError)
	}
	c.Append("HX-Location", "/login")
	return c.SendStatus(http.StatusNoContent)
}
