package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL = "https://developpement-skillkonnect.ngrok.app/api/v1"
)

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token   string      `json:"token"`
	User    UserProfile `json:"user"`
	Message string      `json:"message"`
}

// UserProfile represents user information
type UserProfile struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

// DashboardResponse represents the dashboard data
type DashboardResponse struct {
	User    UserProfile `json:"user"`
	Stats   interface{} `json:"stats"`
	Message string      `json:"message"`
}

// Worker represents a worker from the API
type Worker struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Profession string  `json:"profession"`
	Rating     float64 `json:"rating"`
	Reviews    int     `json:"reviews"`
	Distance   string  `json:"distance"`
	Price      string  `json:"price"`
	Available  bool    `json:"available"`
}

// WorkersResponse represents the workers list response
type WorkersResponse struct {
	Workers []Worker `json:"workers"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
}

// APIService handles all API calls
type APIService struct {
	baseURL string
	token   string
}

// NewAPIService creates a new API service
func NewAPIService() *APIService {
	return &APIService{
		baseURL: BaseURL,
	}
}

// SetToken sets the authorization token
func (api *APIService) SetToken(token string) {
	api.token = token
}

// Login authenticates a user
func (api *APIService) Login(email, password string) (*LoginResponse, error) {
	loginReq := LoginRequest{
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal login request: %w", err)
	}

	req, err := http.NewRequest("POST", api.baseURL+"/client/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed: %s", string(body))
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &loginResp, nil
}

// GetDashboard fetches dashboard data
func (api *APIService) GetDashboard() (*DashboardResponse, error) {
	// Validate token exists
	if api.token == "" {
		return nil, fmt.Errorf("no authentication token available")
	}

	req, err := http.NewRequest("GET", api.baseURL+"/client/dashboard", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+api.token)

	// Prevent following redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dashboard request failed: %s", string(body))
	}

	var dashResp DashboardResponse
	if err := json.Unmarshal(body, &dashResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &dashResp, nil
}

// GetWorkers fetches workers list with pagination
func (api *APIService) GetWorkers(page, limit int) (*WorkersResponse, error) {
	// Validate token exists
	if api.token == "" {
		return nil, fmt.Errorf("no authentication token available")
	}

	url := fmt.Sprintf("%s/client/workers?page=%d&limit=%d", api.baseURL, page, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+api.token)

	// Prevent following redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("workers request failed: %s", string(body))
	}

	var workersResp WorkersResponse
	if err := json.Unmarshal(body, &workersResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &workersResp, nil
}
