package bootstrap

import (
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kisamoto/janus/pkg/api"
	"github.com/pkg/errors"
)

// RegisterAPIContext registers godog suite context for handling API related steps
func RegisterAPIContext(s *godog.Suite, apiRepo api.Repository, ch chan<- api.ConfigurationMessage) {
	ctx := &apiContext{apiRepo: apiRepo, ch: ch}

	s.BeforeScenario(ctx.clearAPI)
}

type apiContext struct {
	apiRepo api.Repository
	ch      chan<- api.ConfigurationMessage
}

func (c *apiContext) clearAPI(arg interface{}) {
	data, err := c.apiRepo.FindAll()
	if err != nil {
		panic(errors.Wrap(err, "Failed to get all registered route specs"))
	}

	for _, definition := range data {
		c.ch <- api.ConfigurationMessage{
			Operation:     api.RemovedOperation,
			Configuration: definition,
		}
	}

	time.Sleep(time.Second)
}
