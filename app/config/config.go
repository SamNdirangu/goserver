package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	errorHandlers "goserver/app/errors"

	hashing "github.com/thomasvvugt/fiber-hashing"
	argon_driver "github.com/thomasvvugt/fiber-hashing/driver/argon2id"
	bcrypt_driver "github.com/thomasvvugt/fiber-hashing/driver/bcrypt"

	"github.com/alexedwards/argon2id"
	"github.com/jameskeane/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/ace"
	"github.com/gofiber/template/amber"
	"github.com/gofiber/template/django"
	"github.com/gofiber/template/handlebars"
	"github.com/gofiber/template/html"
	"github.com/gofiber/template/jet"
	"github.com/gofiber/template/mustache"
	"github.com/gofiber/template/pug"
)

type Config struct {
	*viper.Viper

	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
}

func New() *Config {
	config := &Config{
		Viper: viper.New(),
	}

	// Set default configurations
	config.setDefaults()

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	config.SetErrorHandler(errorHandlers.DefaultErrorHandler)
	// TODO: Logger (Maybe a different zap object)
	// TODO: Add APP_KEY generation
	// TODO: Write changes to configuration file
	// Set Fiber configurations
	config.setFiberConfig()

	return config
}

func (config *Config) SetErrorHandler(errorHandler fiber.ErrorHandler) {
	config.errorHandler = errorHandler
}

func (config *Config) setDefaults() {
	// Set default App configuration
	config.SetDefault("APP_PORT", ":8080")
	config.SetDefault("APP_ENV", "local")

	// Set default database configuration
	config.SetDefault("DB_DRIVER", "mysql")
	config.SetDefault("DB_HOST", "localhost")
	config.SetDefault("DB_USERNAME", "fiber")
	config.SetDefault("DB_PASSWORD", "password")
	config.SetDefault("DB_PORT", 3306)
	config.SetDefault("DB_DATABASE", "boilerplate")

	// Set default hasher configuration
	config.SetDefault("HASHER_DRIVER", "argon2id")
	config.SetDefault("HASHER_MEMORY", 131072)
	config.SetDefault("HASHER_ITERATIONS", 4)
	config.SetDefault("HASHER_PARALLELISM", 4)
	config.SetDefault("HASHER_SALTLENGTH", 16)
	config.SetDefault("HASHER_KEYLENGTH", 32)
	config.SetDefault("HASHER_ROUNDS", bcrypt.DefaultRounds)

	// Set default Fiber configuration
	config.SetDefault("FIBER_PREFORK", false)
	config.SetDefault("FIBER_SERVERHEADER", "")
	config.SetDefault("FIBER_STRICTROUTING", false)
	config.SetDefault("FIBER_CASESENSITIVE", false)
	config.SetDefault("FIBER_IMMUTABLE", false)
	config.SetDefault("FIBER_UNESCAPEPATH", false)
	config.SetDefault("FIBER_ETAG", false)
	config.SetDefault("FIBER_BODYLIMIT", 4194304)
	config.SetDefault("FIBER_CONCURRENCY", 262144)
	config.SetDefault("FIBER_VIEWS", "html")
	config.SetDefault("FIBER_VIEWS_DIRECTORY", "resources/views")
	config.SetDefault("FIBER_VIEWS_RELOAD", false)
	config.SetDefault("FIBER_VIEWS_DEBUG", false)
	config.SetDefault("FIBER_VIEWS_LAYOUT", "embed")
	config.SetDefault("FIBER_VIEWS_DELIMS_L", "{{")
	config.SetDefault("FIBER_VIEWS_DELIMS_R", "}}")
	config.SetDefault("FIBER_READTIMEOUT", 0)
	config.SetDefault("FIBER_WRITETIMEOUT", 0)
	config.SetDefault("FIBER_IDLETIMEOUT", 0)
	config.SetDefault("FIBER_READBUFFERSIZE", 4096)
	config.SetDefault("FIBER_WRITEBUFFERSIZE", 4096)
	config.SetDefault("FIBER_COMPRESSEDFILESUFFIX", ".fiber.gz")
	config.SetDefault("FIBER_PROXYHEADER", "")
	config.SetDefault("FIBER_GETONLY", false)
	config.SetDefault("FIBER_DISABLEKEEPALIVE", false)
	config.SetDefault("FIBER_DISABLEDEFAULTDATE", false)
	config.SetDefault("FIBER_DISABLEDEFAULTCONTENTTYPE", false)
	config.SetDefault("FIBER_DISABLEHEADERNORMALIZING", false)
	config.SetDefault("FIBER_DISABLESTARTUPMESSAGE", false)
	config.SetDefault("FIBER_REDUCEMEMORYUSAGE", false)

	// Set default Custom Access Logger middleware configuration
	config.SetDefault("MW_ACCESS_LOGGER_ENABLED", true)
	config.SetDefault("MW_ACCESS_LOGGER_TYPE", "console")
	config.SetDefault("MW_ACCESS_LOGGER_FILENAME", "access.log")
	config.SetDefault("MW_ACCESS_LOGGER_MAXSIZE", 500)
	config.SetDefault("MW_ACCESS_LOGGER_MAXAGE", 28)
	config.SetDefault("MW_ACCESS_LOGGER_MAXBACKUPS", 3)
	config.SetDefault("MW_ACCESS_LOGGER_LOCALTIME", false)
	config.SetDefault("MW_ACCESS_LOGGER_COMPRESS", false)

	// Set default Force HTTPS middleware configuration
	config.SetDefault("MW_FORCE_HTTPS_ENABLED", false)

	// Set default Force trailing slash middleware configuration
	config.SetDefault("MW_FORCE_TRAILING_SLASH_ENABLED", false)

	// Set default HSTS middleware configuration
	config.SetDefault("MW_HSTS_ENABLED", false)
	config.SetDefault("MW_HSTS_MAXAGE", 31536000)
	config.SetDefault("MW_HSTS_INCLUDESUBDOMAINS", true)
	config.SetDefault("MW_HSTS_PRELOAD", false)

	// Set default Suppress WWW middleware configuration
	config.SetDefault("MW_SUPPRESS_WWW_ENABLED", true)

	// Set default Fiber Cache middleware configuration
	config.SetDefault("MW_FIBER_CACHE_ENABLED", false)
	config.SetDefault("MW_FIBER_CACHE_EXPIRATION", "1m")
	config.SetDefault("MW_FIBER_CACHE_CACHECONTROL", false)

	// Set default Fiber Compress middleware configuration
	config.SetDefault("MW_FIBER_COMPRESS_ENABLED", false)
	config.SetDefault("MW_FIBER_COMPRESS_LEVEL", 0)

	// Set default Fiber CORS middleware configuration
	config.SetDefault("MW_FIBER_CORS_ENABLED", false)
	config.SetDefault("MW_FIBER_CORS_ALLOWORIGINS", "*")
	config.SetDefault("MW_FIBER_CORS_ALLOWMETHODS", "GET,POST,HEAD,PUT,DELETE,PATCH")
	config.SetDefault("MW_FIBER_CORS_ALLOWHEADERS", "")
	config.SetDefault("MW_FIBER_CORS_ALLOWCREDENTIALS", false)
	config.SetDefault("MW_FIBER_CORS_EXPOSEHEADERS", "")
	config.SetDefault("MW_FIBER_CORS_MAXAGE", 0)

	// Set default Fiber CSRF middleware configuration
	config.SetDefault("MW_FIBER_CSRF_ENABLED", false)
	config.SetDefault("MW_FIBER_CSRF_TOKENLOOKUP", "header:X-CSRF-Token")
	config.SetDefault("MW_FIBER_CSRF_COOKIE_NAME", "_csrf")
	config.SetDefault("MW_FIBER_CSRF_COOKIE_SAMESITE", "Strict")
	config.SetDefault("MW_FIBER_CSRF_COOKIE_EXPIRES", "24h")
	config.SetDefault("MW_FIBER_CSRF_CONTEXTKEY", "csrf")

	// Set default Fiber ETag middleware configuration
	config.SetDefault("MW_FIBER_ETAG_ENABLED", false)
	config.SetDefault("MW_FIBER_ETAG_WEAK", false)

	// Set default Fiber Expvar middleware configuration
	config.SetDefault("MW_FIBER_EXPVAR_ENABLED", false)

	// Set default Fiber Favicon middleware configuration
	config.SetDefault("MW_FIBER_FAVICON_ENABLED", false)
	config.SetDefault("MW_FIBER_FAVICON_FILE", "")
	config.SetDefault("MW_FIBER_FAVICON_CACHECONTROL", "public, max-age=31536000")

	// Set default Fiber Limiter middleware configuration
	config.SetDefault("MW_FIBER_LIMITER_ENABLED", false)
	config.SetDefault("MW_FIBER_LIMITER_MAX", 5)
	config.SetDefault("MW_FIBER_LIMITER_DURATION", "1m")

	// Set default Fiber Monitor middleware configuration
	config.SetDefault("MW_FIBER_MONITOR_ENABLED", false)

	// Set default Fiber Pprof middleware configuration
	config.SetDefault("MW_FIBER_PPROF_ENABLED", false)

	// Set default Fiber Recover middleware configuration
	config.SetDefault("MW_FIBER_RECOVER_ENABLED", true)

	// Set default Fiber RequestID middleware configuration
	config.SetDefault("MW_FIBER_REQUESTID_ENABLED", false)
	config.SetDefault("MW_FIBER_REQUESTID_HEADER", "X-Request-ID")
	config.SetDefault("MW_FIBER_REQUESTID_CONTEXTKEY", "requestid")
}

func (config *Config) getFiberViewsEngine() fiber.Views {
	var viewsEngine fiber.Views
	switch strings.ToLower(config.GetString("FIBER_VIEWS")) {
	case "ace":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".ace")
		}
		engine := ace.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine
	case "amber":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".amber")
		}
		engine := amber.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine
	case "django":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".django")
		}
		engine := django.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT"))
		viewsEngine = engine

	case "handlebars":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".hbs")
		}
		engine := handlebars.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine

	case "jet":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".jet")
		}
		engine := jet.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine

	case "mustache":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".mustache")
		}
		engine := mustache.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine

	case "pug":
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".pug")
		}
		engine := pug.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine

	// Use the official html/template package by default
	default:
		// Set file extension dynamically to FIBER_VIEWS
		if config.GetString("FIBER_VIEWS_EXTENSION") == "" {
			config.Set("FIBER_VIEWS_EXTENSION", ".html")
		}
		engine := html.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
		engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
			Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
			Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
			Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
		viewsEngine = engine

	}
	return viewsEngine
}

func (config *Config) setFiberConfig() {
	config.fiber = &fiber.Config{
		Prefork:                   config.GetBool("FIBER_PREFORK"),
		ServerHeader:              config.GetString("FIBER_SERVERHEADER"),
		StrictRouting:             config.GetBool("FIBER_STRICTROUTING"),
		CaseSensitive:             config.GetBool("FIBER_CASESENSITIVE"),
		Immutable:                 config.GetBool("FIBER_IMMUTABLE"),
		UnescapePath:              config.GetBool("FIBER_UNESCAPEPATH"),
		ETag:                      config.GetBool("FIBER_ETAG"),
		BodyLimit:                 config.GetInt("FIBER_BODYLIMIT"),
		Concurrency:               config.GetInt("FIBER_CONCURRENCY"),
		Views:                     config.getFiberViewsEngine(),
		ReadTimeout:               config.GetDuration("FIBER_READTIMEOUT"),
		WriteTimeout:              config.GetDuration("FIBER_WRITETIMEOUT"),
		IdleTimeout:               config.GetDuration("FIBER_IDLETIMEOUT"),
		ReadBufferSize:            config.GetInt("FIBER_READBUFFERSIZE"),
		WriteBufferSize:           config.GetInt("FIBER_WRITEBUFFERSIZE"),
		CompressedFileSuffix:      config.GetString("FIBER_COMPRESSEDFILESUFFIX"),
		ProxyHeader:               config.GetString("FIBER_PROXYHEADER"),
		GETOnly:                   config.GetBool("FIBER_GETONLY"),
		ErrorHandler:              config.errorHandler,
		DisableKeepalive:          config.GetBool("FIBER_DISABLEKEEPALIVE"),
		DisableDefaultDate:        config.GetBool("FIBER_DISABLEDEFAULTDATE"),
		DisableDefaultContentType: config.GetBool("FIBER_DISABLEDEFAULTCONTENTTYPE"),
		DisableHeaderNormalizing:  config.GetBool("FIBER_DISABLEHEADERNORMALIZING"),
		DisableStartupMessage:     config.GetBool("FIBER_DISABLESTARTUPMESSAGE"),
		ReduceMemoryUsage:         config.GetBool("FIBER_REDUCEMEMORYUSAGE"),
	}
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}

func (config *Config) GetHasherConfig() hashing.Config {
	if strings.ToLower(config.GetString("HASHER_DRIVER")) == "bcrypt" {
		return hashing.Config{
			Driver: bcrypt_driver.New(bcrypt_driver.Config{
				Complexity: config.GetInt("HASHER_ROUNDS"),
			})}
	} else {
		return hashing.Config{
			Driver: argon_driver.New(argon_driver.Config{
				Params: &argon2id.Params{
					Memory:      config.GetUint32("HASHER_MEMORY"),
					Iterations:  config.GetUint32("HASHER_ITERATIONS"),
					Parallelism: uint8(config.GetInt("HASHER_PARALLELISM")),
					SaltLength:  config.GetUint32("HASHER_SALTLENGTH"),
					KeyLength:   config.GetUint32("HASHER_KEYLENGTH"),
				}})}
	}
}
