package session

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

type Cookie playwright.Cookie

func (c Cookie) GetCookieAsString() string {
	return fmt.Sprintf("%s=%s", c.Name, c.Value)
}
