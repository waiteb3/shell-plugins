package datadog

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func Executable_dogshell() schema.Executable {
	return schema.Executable{
		Runs:      []string{"dog"},
		Name:      "Dogshell",
		DocsURL:   sdk.URL("https://docs.datadoghq.com/developers/guide/dogshell-quickly-use-datadog-s-api-from-terminal-shell/"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Credentials: []schema.CredentialType{
			APIKey(),
		},
	}
}