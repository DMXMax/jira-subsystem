package ticket

import "github.com/DMXMax/jira-subsystem/region"

type Ticket struct {
	Id string
}

type ScopeRequest struct {
	Client      string
	Scope       string
	IsArbitrary bool
}

type ValidatedTicket struct {
	Id            string
	IsTest        bool
	Region        region.ScopeRegion
	ScopeRequests []ScopeRequest
}
