package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

type Transational interface {
	Repository

	Begin(query query.Query, options ...Options) errors.Error
	End(query query.Query) errors.Error
}
