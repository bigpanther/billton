package firebase

//go:generate mockgen -destination ../app/mock_firebase_test.go -package app github.com/bigpanther/billton/internal/firebase Firebase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"github.com/bigpanther/billton/internal/models"
	"google.golang.org/api/option"
)

type Firebase interface {
	SetClaims(c context.Context, u *models.User) error
	GetUser(c context.Context, token string) (*auth.UserRecord, error)
	SendAll(c context.Context, messages []*messaging.Message) error
	SubscribeToTopics(c context.Context, user *models.User, token string) error
	UnSubscribeToTopics(c context.Context, user *models.User, token string) error
}
type firebaseSdkClient struct {
	authClient      *auth.Client
	messagingClient *messaging.Client
}

var errMissingToken = errors.New("missing token")

// New returns an instance of Firebase
func New() (Firebase, error) {
	var credsJSONEncoded = os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON_ENCODED")
	credJSON, err := base64.StdEncoding.DecodeString(credsJSONEncoded)
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentialsJSON(credJSON)
	ctx := context.Background()
	client := &firebaseSdkClient{}
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		client = nil
		return nil, err
	}
	client.authClient, err = app.Auth(ctx)
	if err != nil {
		client = nil
		return nil, err
	}
	client.messagingClient, err = app.Messaging(ctx)
	if err != nil {
		client = nil
		return nil, err
	}
	return client, err
}

// SetClaims sets the custom claims for the user
func (client *firebaseSdkClient) SetClaims(c context.Context, u *models.User) error {
	claims := map[string]interface{}{
		"bpRole": u.Role,
	}
	err := client.authClient.SetCustomUserClaims(c, u.Username, claims)
	if err != nil {
		return err
	}
	return nil
}

// verifyIDToken return the auth token after verification
func (client *firebaseSdkClient) verifyIDToken(c context.Context, token string) (*auth.Token, error) {
	return client.authClient.VerifyIDToken(c, token)
}

// GetUser return the firebase user for the username
func (client *firebaseSdkClient) GetUser(c context.Context, token string) (*auth.UserRecord, error) {
	t, err := client.verifyIDToken(c, token)
	if err != nil {
		return nil, err
	}
	return client.authClient.GetUser(c, t.Subject)
}

// SendAll sends all messages to FCM topics
func (client *firebaseSdkClient) SendAll(c context.Context, messages []*messaging.Message) error {
	r, err := client.messagingClient.SendEach(c, messages)
	if err != nil {
		log.Println("error sending messages", err, r)
		return err
	}
	return nil
}

// SubscribeToTopics create subscription topics for a user
func (client *firebaseSdkClient) SubscribeToTopics(c context.Context, user *models.User, token string) error {
	if token == "" {
		return errMissingToken
	}

	var t = GetTopic(user)
	var topics = []string{t}
	for _, t := range topics {

		r, err := client.messagingClient.SubscribeToTopic(c, []string{token}, t)
		if err != nil {
			log.Println(user.ID, " role=", user.Role, " subscription failed to topic", t, r)
			return err
		}
		log.Println(user.ID, " role=", user.Role, " subscribed to topic", t, r)
	}
	return nil
}

// UnSubscribeToTopics removes subscription topics for a user
func (client *firebaseSdkClient) UnSubscribeToTopics(c context.Context, user *models.User, token string) error {
	if token == "" {
		return errMissingToken
	}
	topics := []string{GetConsumerTopic(user.ID.String())}
	if user.IsAdmin() {
		topics = append(topics, GetAdminTopic())
	}
	for _, t := range topics {
		r, err := client.messagingClient.UnsubscribeFromTopic(c, []string{token}, t)
		if err != nil {
			log.Println(user.ID, " role=", user.Role, " unsubscription failed to topic", t, r)
			// Don't fail here
			continue
		}
		log.Println(user.ID, " role=", user.Role, " unsubscribed to topic", t, r)
	}
	return nil
}

// GetTopic returns the topic name for the user
func GetTopic(user *models.User) string {

	if user.IsAdmin() {
		return GetAdminTopic()
	}
	if user.IsConsumer() {
		return GetConsumerTopic(user.ID.String())
	}
	return fmt.Sprintf("warrant_none_%s", user.ID)
}

// GetAdminTopic returns the topic for admin
func GetAdminTopic() string {
	return "warrant_admin"
}

// GetConsumerTopic returns the topic for consumer
func GetConsumerTopic(consumerID string) string {
	return fmt.Sprintf("warrant_consumer_%s", consumerID)
}
