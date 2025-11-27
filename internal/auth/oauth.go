package auth

import (
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func InitializeUserSession() (*azidentity.InteractiveBrowserCredential, *msgraphsdk.GraphServiceClient, error) {
	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	graphUserScopes := strings.Split(scopes, ",")

	// Create the device code credential
	credential, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		ClientID: clientId,
		TenantID: tenantId,
	})
	if err != nil {
		return nil, nil, err
	}

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, graphUserScopes)
	if err != nil {
		return nil, nil, err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return nil, nil, err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)

	return credential, client, nil
}
