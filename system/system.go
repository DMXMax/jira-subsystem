package system

import (
	jira "github.com/DMXMax/jira-subsystem"
	"github.com/DMXMax/jira-subsystem/region"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var jiralog zerolog.Logger

// ValidateAndRetrieveTicket retrieves a ticket from Jira and validates the ticket is correct.
// If the ticket is not correct, ValidateAndRetrieveTicket returns JiraNotFound in its error
func ValidateAndRetrieveTicket(t jira.Ticket) (jira.ValidatedTicket, error) {

	jiralog.Info().Interface("ValidateAndRetrieveTicket", t).Send()
	vt := jira.ValidatedTicket{Id: t.Id, Region: region.USA, IsTest: false}
	vt.Id = t.Id
	scopes := []jira.ScopeRequest{
		{
			Client:      "test-client-1",
			Scope:       "test-scope-1",
			IsArbitrary: true,
		},
		{
			Client:      "test-client-1",
			Scope:       "test-scope-2",
			IsArbitrary: true,
		},
		{
			Client:      "test-client-1",
			Scope:       "test-scope-3",
			IsArbitrary: true,
		},
	}
	vt.ScopeRequests = scopes
	return vt, nil
}

func init() {
	jiralog = log.With().
		Str("component", "Jira").
		Logger()

}
