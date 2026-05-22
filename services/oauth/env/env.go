// Package env sets the environmemnt variables of the application
package env

import (
	"reflect"
	"time"

	"github.com/caarlos0/env/v6"
)

// Config :
var Config = struct {
	Env         string `env:"ENV,required"`
	ServiceName string `env:"JAEGER_SERVICE_NAME,required" envDefault:"oauth2-service"`
	MySQL       struct {
		Host     string `env:"WETIX_MYSQL_HOST,required"`
		Port     string `env:"WETIX_MYSQL_PORT,required"`
		Username string `env:"WETIX_MYSQL_USERNAME,required"`
		Password string `env:"WETIX_MYSQL_PASSWORD,required"`
		Database string `env:"WETIX_MYSQL_DATABASE,required"`
	}

	Secret              string `env:"WETIX_SECRET,required"`
	CookieDomain        string `env:"WETIX_COOKIE_DOMAIN,required"`
	SessionID           string `env:"WETIX_SESSION_ID,required"`
	WetixAPIOrigin      string `env:"WETIX_API_ORIGIN,required"`
	WebsiteOrigin       string `env:"WETIX_WEBSITE_ORIGIN,required"`
	OAuthOrigin         string `env:"WETIX_OAUTH_ORIGIN,required"`
	OpenAPIOrigin       string `env:"WETIX_OPEN_API_ORIGIN,required"`
	OpengraphOrigin     string `env:"WETIX_OPEN_GRAPH_ORIGIN,required"`
	InternalGraphOrigin string `env:"WETIX_INTERNAL_GRAPH_ORIGIN,required"`
}{}

// Init :
func Init() {
	if err := env.ParseWithFuncs(&Config,
		map[reflect.Type]env.ParserFunc{
			reflect.TypeOf(time.Location{}): func(val string) (interface{}, error) {
				loc, err := time.LoadLocation(val)
				if err != nil {
					return nil, err
				}
				return *loc, nil
			},
		}); err != nil {
		panic(err)
	}
}

// IsProduction :
func IsProduction() bool {
	return Config.Env == "production"
}

// IsDevelopment :
func IsDevelopment() bool {
	return Config.Env == "dev"
}

// IsLocal :
func IsLocal() bool {
	return Config.Env == "local"
}
