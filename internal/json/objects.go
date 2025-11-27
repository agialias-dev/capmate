package objects

import "time"

type ConditionalAccessPolicy struct {
	OdataContext       string    `json:"@odata.context"`
	MicrosoftGraphTips string    `json:"@microsoft.graph.tips"`
	ID                 string    `json:"id"`
	TemplateID         any       `json:"templateId"`
	DisplayName        string    `json:"displayName"`
	CreatedDateTime    time.Time `json:"createdDateTime"`
	ModifiedDateTime   time.Time `json:"modifiedDateTime"`
	State              string    `json:"state"`
	Conditions         struct {
		UserRiskLevels             []string `json:"userRiskLevels"`
		SignInRiskLevels           []any    `json:"signInRiskLevels"`
		ClientAppTypes             []string `json:"clientAppTypes"`
		ServicePrincipalRiskLevels []any    `json:"servicePrincipalRiskLevels"`
		InsiderRiskLevels          any      `json:"insiderRiskLevels"`
		Platforms                  any      `json:"platforms"`
		Locations                  any      `json:"locations"`
		Devices                    any      `json:"devices"`
		ClientApplications         any      `json:"clientApplications"`
		Applications               struct {
			IncludeApplications                         []string `json:"includeApplications"`
			ExcludeApplications                         []any    `json:"excludeApplications"`
			IncludeUserActions                          []any    `json:"includeUserActions"`
			IncludeAuthenticationContextClassReferences []any    `json:"includeAuthenticationContextClassReferences"`
			ApplicationFilter                           any      `json:"applicationFilter"`
		} `json:"applications"`
		Users struct {
			IncludeUsers                 []string `json:"includeUsers"`
			ExcludeUsers                 []any    `json:"excludeUsers"`
			IncludeGroups                []any    `json:"includeGroups"`
			ExcludeGroups                []string `json:"excludeGroups"`
			IncludeRoles                 []any    `json:"includeRoles"`
			ExcludeRoles                 []any    `json:"excludeRoles"`
			IncludeGuestsOrExternalUsers any      `json:"includeGuestsOrExternalUsers"`
			ExcludeGuestsOrExternalUsers any      `json:"excludeGuestsOrExternalUsers"`
		} `json:"users"`
	} `json:"conditions"`
	GrantControls struct {
		Operator                           string   `json:"operator"`
		BuiltInControls                    []string `json:"builtInControls"`
		CustomAuthenticationFactors        []any    `json:"customAuthenticationFactors"`
		TermsOfUse                         []any    `json:"termsOfUse"`
		AuthenticationStrengthOdataContext string   `json:"authenticationStrength@odata.context"`
		AuthenticationStrength             struct {
			ID                                    string    `json:"id"`
			CreatedDateTime                       time.Time `json:"createdDateTime"`
			ModifiedDateTime                      time.Time `json:"modifiedDateTime"`
			DisplayName                           string    `json:"displayName"`
			Description                           string    `json:"description"`
			PolicyType                            string    `json:"policyType"`
			RequirementsSatisfied                 string    `json:"requirementsSatisfied"`
			AllowedCombinations                   []string  `json:"allowedCombinations"`
			CombinationConfigurationsOdataContext string    `json:"combinationConfigurations@odata.context"`
			CombinationConfigurations             []any     `json:"combinationConfigurations"`
		} `json:"authenticationStrength"`
	} `json:"grantControls"`
	SessionControls struct {
		DisableResilienceDefaults       any `json:"disableResilienceDefaults"`
		ApplicationEnforcedRestrictions any `json:"applicationEnforcedRestrictions"`
		CloudAppSecurity                any `json:"cloudAppSecurity"`
		PersistentBrowser               any `json:"persistentBrowser"`
		SignInFrequency                 struct {
			Value              any    `json:"value"`
			Type               any    `json:"type"`
			AuthenticationType string `json:"authenticationType"`
			FrequencyInterval  string `json:"frequencyInterval"`
			IsEnabled          bool   `json:"isEnabled"`
		} `json:"signInFrequency"`
	} `json:"sessionControls"`
}
