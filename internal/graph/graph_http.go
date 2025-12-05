package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	objects "github.com/agialias-dev/capmate/internal/json"
)

func GetHTTPCAPs(us *UserSession) error {
	URL := os.Getenv("GRAPH_BASE_URL") + "identity/conditionalAccess/policies"
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	graphScopes := strings.Split(scopes, ",")
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return err
	}

	cred, err := us.InteractiveBrowserCredential.GetToken(ctx, policy.TokenRequestOptions{Scopes: graphScopes})
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+cred.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var policies []objects.ConditionalAccessPolicy
	err = json.Unmarshal(body, &policies)
	if err != nil {
		return err
	}
	for _, policy := range policies {
		fmt.Println("Name:", policy.DisplayName)
	}
	return nil
}
