package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/m1guelpf/chatgpt-telegram/src/session"
	"github.com/playwright-community/playwright-go"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	v *viper.Viper

	Cookies []session.Cookie
}

// LoadOrCreatePersistentConfig uses the default config directory for the current OS
// to load or create a config file named "chatgpt.json"
func LoadOrCreatePersistentConfig() (*Config, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Couldn't get user config dir: %v", err))
	}
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName("chatgpt")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := v.SafeWriteConfig(); err != nil {
				return nil, errors.New(fmt.Sprintf("Couldn't create config file: %v", err))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Couldn't read config file: %v", err))
		}
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing config: %v", err))
	}
	cfg.v = v

	return &cfg, nil
}

func (cfg *Config) LoadSessionFromCookie() error {
	filePath := cfg.v.GetString("COOKIE_FILE")
	if filePath == "" {
		return errors.New(fmt.Sprintf("cookie file not found. path: %s", filePath))
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var cookies []session.Cookie
	err = json.Unmarshal(file, &cookies)
	if err != nil {
		return err
	}

	cfg.Cookies = cookies

	return nil
}

func (cfg *Config) SetSession(results []*playwright.BrowserContextCookiesResult) {
	cfg.Cookies = []session.Cookie{}
	for _, r := range results {
		cfg.Cookies = append(cfg.Cookies, session.Cookie{
			Domain:   r.Domain,
			HttpOnly: r.HttpOnly,
			Expires:  r.Expires,
			Value:    r.Value,
			SameSite: string(r.SameSite),
			Secure:   r.Secure,
			Name:     r.Name,
			Path:     r.Path,
		})
	}
}

func (cfg *Config) GetCookiesAsString() string {
	var cookieStrings []string
	for _, cookie := range cfg.Cookies {
		cookieStrings = append(cookieStrings, cookie.GetCookieAsString())
	}

	return strings.Join(cookieStrings, ";")
}
