package minox_test

import (
	_ "github.com/joho/godotenv/autoload"
)

// func TestCustomerPost(t *testing.T) {
// 	clientID := os.Getenv("OAUTH_CLIENT_ID")
// 	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
// 	accessToken := os.Getenv("OAUTH_ACCESS_TOKEN")
// 	tokenURL := os.Getenv("OAUTH_TOKEN_URL")
// 	tenant, err := strconv.Atoi(os.Getenv("MINOX_TENANT"))
// 	administrationID := os.Getenv("MINOX_ADMINISTRATION")
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	oauthConfig := minox.NewOauth2Config()
// 	oauthConfig.ClientID = clientID
// 	oauthConfig.ClientSecret = clientSecret
//
// 	// set alternative token url
// 	if tokenURL != "" {
// 		oauthConfig.Endpoint.TokenURL = tokenURL
// 	}
//
// 	token := &oauth2.Token{
// 		AccessToken: accessToken,
// 	}
//
// 	// get http client with automatic oauth logic
// 	httpClient := oauthConfig.Client(oauth2.NoContext, token)
//
// 	client := minox.NewClient(httpClient, tenant)
// 	client.SetDebug(true)
// 	client.SetDisallowUnknownFields(true)
//
// 	req := client.NewCustomerPostRequest()
// 	req.PathParams().AdministrationID = administrationID
// 	req.QueryParams().JournalID = "MEMO"
//
// 	req.RequestBody() = minox.CustomerPost{
// 		Addresses: minox.Addresses{
// 			{
// 				AddressType:         "postadres",
// 				Name:                "Stadswandeling",
// 				ContactName:         "Kim ten Have",
// 				StreetnameAndNumber: "IJselstraat 37",
// 				PostalCode:          "1078 CB",
// 				City:                "Amsterdam",
// 				CountryCode:         "NL",
// 				EmailAddresses: []minox.EmailAddress{
// 					{
// 						EmailAddressType: "string",
// 						EmailAddress:     "destadswandeling020@gmail.com",
// 					},
// 				},
// 			},
// 		},
// 		Search:       []minox.CustomerSearch{},
// 		CustomFields: minox.CustomFields{},
// 		ExternalID:   "6918",
// 	}
//
// 	resp, err := req.Do()
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	log.Println(string(b))
// }
