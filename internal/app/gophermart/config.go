package gophermart

import (
	"flag"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	runAddress           = "RUN_ADDRESS"
	databaseURI          = "DATABASE_URI"
	accuralSystemAddress = "ACCRUAL_SYSTEM_ADDRESS"
)

type config struct {
	RunAddress           string
	DatabaseURI          string
	AccuralSystemAddress string
}

func NewConfig() (*config, error) {
	c := &config{}
	c.configureViaFlags()
	c.configureViaEnvVars()

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *config) configureViaFlags() {
	flag.Func("a", "Server address", func(input string) error {
		c.RunAddress = input
		return nil
	})

	flag.Func("d", "Database URI", func(input string) error {
		c.DatabaseURI = input
		return nil
	})

	flag.Func("r", "Accural system address", func(input string) error {
		c.AccuralSystemAddress = input
		return nil
	})

	flag.Parse()
}

func (c *config) configureViaEnvVars() {
	if val := os.Getenv(runAddress); len(val) != 0 {
		c.RunAddress = val
	}

	if val := os.Getenv(databaseURI); len(val) != 0 {
		c.DatabaseURI = val
	}

	if val := os.Getenv(accuralSystemAddress); len(val) != 0 {
		c.AccuralSystemAddress = val
	}
}

func (c *config) validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.RunAddress, validation.Required),
		validation.Field(&c.DatabaseURI, validation.Required),
		validation.Field(&c.AccuralSystemAddress, validation.Required, is.URL),
	)
}
