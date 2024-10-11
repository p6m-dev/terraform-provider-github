package xpprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/integrations/terraform-provider-github/v6/github"
)

func GetProviderSchema(_ context.Context) (*schema.Provider, error) {
	return github.Provider(), nil
}
