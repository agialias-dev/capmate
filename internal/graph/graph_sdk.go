package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/agialias-dev/capmate/internal/auth"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

type UserSession struct {
	InteractiveBrowserCredential *azidentity.InteractiveBrowserCredential
	userClient                   *msgraphsdk.GraphServiceClient
	Token                        azidentity.Cache
}

func NewUserSession() *UserSession {
	us := &UserSession{}
	return us
}

func InitializeGraph(UserSession *UserSession) {
	credential, client, cache, err := auth.InitializeUserSession()
	if err != nil {
		log.Panicf("Error initializing Graph for user auth: %v\n", err)
	}
	UserSession.InteractiveBrowserCredential = credential
	UserSession.userClient = client
	UserSession.Token = cache
}

func GreetUser(UserSession *UserSession) {
	user, err := UserSession.GetUser()
	if err != nil {
		log.Panicf("Error getting user: %v\n", err)
	}

	fmt.Printf("Hello, %s!\n", *user.GetDisplayName())

	// For Work/school accounts, email is in Mail property
	// Personal accounts, email is in UserPrincipalName
	email := user.GetMail()
	if email == nil {
		email = user.GetUserPrincipalName()
	}

	fmt.Printf("Email: %s\n", *email)
	fmt.Println()
}

func GetAllCAPs(UserSession *UserSession) {
	err := UserSession.GetConditionalAccessPolicies()
	if err != nil {
		log.Panicf("Error making Graph call: %v", err)
	}
}

func (us *UserSession) GetUser() (models.Userable, error) {
	query := users.UserItemRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "mail", "userPrincipalName"},
	}

	return us.userClient.Me().Get(context.Background(),
		&users.UserItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &query,
		})
}

func (us *UserSession) GetConditionalAccessPolicies() error {
	result, err := us.userClient.Identity().ConditionalAccess().Policies().Get(context.Background(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Conditional Access Policies:")
	policies := result.GetValue()
	for _, policy := range policies {
		conditions := policy.GetConditions()
		grantControls := policy.GetGrantControls()
		sessionControls := policy.GetSessionControls()
		fmt.Printf("- Name:\t%s\n", *policy.GetDisplayName())
		fmt.Printf("- State:\t%s\n", *policy.GetState())
		fmt.Printf("- Conditions:\t%+v\n", conditions.GetApplications().GetIncludeApplications())
		fmt.Printf("- Grant Controls:\t%+v\n", grantControls)
		fmt.Printf("- Session Controls:\t%+v\n", sessionControls)
		fmt.Println()
	}
	return nil
}
