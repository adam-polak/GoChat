package controller

import (
	"github.com/labstack/echo/v4"
)

func (c *Controller) SignUp(context echo.Context) error {
	return nil
}

func (c *Controller) Login(context echo.Context) error {
	print("Logging in...")
	return context.Redirect(302, "/")
}

func (c *Controller) CorrectLoginKey(context echo.Context) error {
	return nil
}
