package api

import (
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) Handler {
	return Handler{
		dbStore: dbStore,
	}
}

// RegisterUser The handler our lambda will route to whe we need register our new users
func (api Handler) RegisterUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request payload",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
				Body:       "Invalid user credentials",
				StatusCode: http.StatusBadRequest,
			},
			fmt.Errorf("username or password cannot be empty")
	}

	// does the user with this username already exist?
	isUserExist, err := api.dbStore.DoesUserExist(registerUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
				Body:       "Error checking user existence",
				StatusCode: http.StatusInternalServerError,
			},
			fmt.Errorf("error while checking if user exists: %w", err)
	}

	if isUserExist {
		return events.APIGatewayProxyResponse{
				Body:       "User already exists",
				StatusCode: http.StatusConflict,
			},
			fmt.Errorf("a user with this username already exists: %v", err)
	}

	user, err := types.NewUser(registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
				Body:       "Error creating user",
				StatusCode: http.StatusInternalServerError,
			},
			fmt.Errorf("error while creating user: %w", err)
	}

	// now we know a user does not exist
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
				Body:       "Error inserting user",
				StatusCode: http.StatusInternalServerError,
			},
			fmt.Errorf("error while inserting user into db: %w", err)
	}
	return events.APIGatewayProxyResponse{
		Body:       "User created successfully",
		StatusCode: http.StatusCreated,
	}, nil
}

func (api Handler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var loginUser LoginUser
	err := json.Unmarshal([]byte(request.Body), &loginUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
				Body:       "Invalid request payload",
				StatusCode: http.StatusBadRequest,
			},
			fmt.Errorf("error while unmarshalling body: %w", err)
	}

	user, err := api.dbStore.GetUser(loginUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
				Body:       "Internal server error",
				StatusCode: http.StatusInternalServerError,
			},
			fmt.Errorf("error while getting user from db: %w", err)
	}

	if !types.ValidatePassword(user.PasswordHash, loginUser.Password) {
		return events.APIGatewayProxyResponse{
				Body:       "Invalid user credentials",
				StatusCode: http.StatusBadRequest,
			},
			fmt.Errorf("invalid username or password")
	}

	accessToken := types.CreateToken(user)
	successMsg := fmt.Sprintf(`{"access_token": "%s"}`, accessToken)

	return events.APIGatewayProxyResponse{
		Body:       successMsg,
		StatusCode: http.StatusOK,
	}, nil

}
