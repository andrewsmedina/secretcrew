package types

import (
	"secretcrew/modules/post/resolvers"

	"github.com/graphql-go/graphql"
)

var PostsType = &graphql.Field{
	Type: &graphql.List{
		OfType: graphql.String,
	},
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PostsResolver,
}