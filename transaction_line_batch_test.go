package minox_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	minox "github.com/omniboost/go-minox"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
)

func TestTransactionLineBatchPut(t *testing.T) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	accessToken := os.Getenv("OAUTH_ACCESS_TOKEN")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")
	tenant, err := strconv.Atoi(os.Getenv("MINOX_TENANT"))
	if err != nil {
		t.Error(err)
	}

	oauthConfig := minox.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		AccessToken: accessToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client := minox.NewClient(httpClient, tenant)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewTransactionLineBatchPutRequest()
	req.PathParams().AdministrationID = "99998"
	req.PathParams().BatchID = uuid.NewV4()

	line := minox.TransactionLinePut{}
	line.Account.ID = "1234"
	line.Period.FiscalYear = 2006
	line.Journal.ID = "MEMO"
	req.RequestBody().TransactionLineBatchPut = append(req.RequestBody().TransactionLineBatchPut, line)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
