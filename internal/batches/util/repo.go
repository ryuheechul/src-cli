package util

import (
	"github.com/sourcegraph/sourcegraph/lib/batches/template"
	"github.com/sourcegraph/src-cli/internal/batches/graphql"
)

// GraphQLRepoToTemplatingRepo transforms a given *graphql.Repository into a
// template.TemplatingRepository.
func GraphQLRepoToTemplatingRepo(r *graphql.Repository) template.TemplatingRepository {
	return template.TemplatingRepository{
		ID:   r.ID,
		Name: r.Name,
		DefaultBranch: template.TemplatingBranch{
			Name:      r.DefaultBranch.Name,
			TargetOID: r.DefaultBranch.Target.OID,
		},
		FileMatches: r.FileMatches,
	}
}