package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func retrieveRecord() (azidentity.AuthenticationRecord, error) {
	authRecordPath := os.Getenv("AUTH_CACHE_PATH")
	record := azidentity.AuthenticationRecord{}
	b, err := os.ReadFile(authRecordPath)
	if err == nil {
		err = json.Unmarshal(b, &record)
	}
	return record, err
}

func storeRecord(record azidentity.AuthenticationRecord) error {
	authRecordPath := os.Getenv("AUTH_CACHE_PATH")
	b, err := json.Marshal(record)
	if err == nil {
		err = os.WriteFile(authRecordPath, b, 0600)
	}
	return err
}

func InitializeUserSession() (*azidentity.InteractiveBrowserCredential, *msgraphsdk.GraphServiceClient, azidentity.Cache, error) {
	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	graphUserScopes := strings.Split(scopes, ",")

	record, err := retrieveRecord()
	if err != nil {
		log.Println(err)
	}

	c, err := cache.New(&cache.Options{Name: "TokenCache"})
	if err != nil {
		return nil, nil, azidentity.Cache{}, fmt.Errorf("Cache creation failed: %w", err)
	}

	credential, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		AuthenticationRecord: record,
		ClientID:             clientId,
		TenantID:             tenantId,
		Cache:                c,
	})
	if err != nil {
		return nil, nil, azidentity.Cache{}, fmt.Errorf("IteractiveBrowserCredential creation failed: %w", err)
	}

	if record == (azidentity.AuthenticationRecord{}) {
		// No stored record; call Authenticate to acquire one.
		// This will prompt the user to authenticate interactively.
		record, err = credential.Authenticate(context.Background(), &policy.TokenRequestOptions{Scopes: graphUserScopes})
		if err != nil {
			return nil, nil, azidentity.Cache{}, fmt.Errorf("Authentication failed: %w", err)
		}
		err = storeRecord(record)
		if err != nil {
			return nil, nil, azidentity.Cache{}, fmt.Errorf("Failed to store authentication record: %w", err)
		}
	}

	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, graphUserScopes)
	if err != nil {
		return nil, nil, azidentity.Cache{}, fmt.Errorf("Authentication provider creation failed: %w", err)
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return nil, nil, azidentity.Cache{}, fmt.Errorf("GraphRequestAdapter creation failed: %w", err)
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)

	return credential, client, c, nil
}
