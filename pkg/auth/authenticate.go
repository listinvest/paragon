package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kcarretto/paragon/ent"
)

const (
	SessionTokenLength = 256
)

type authKey string

var userContextKey authKey = "user"

// GetUser from the context, returns nil for non-user contexts.
func GetUser(ctx context.Context) *ent.User {
	if v := ctx.Value(userContextKey); v != nil {
		if usr, ok := v.(*ent.User); ok {
			return usr
		}
		panic(fmt.Errorf("Received non-user value for user context key: %v", v))
	}

	return nil
}

type MultiAuthenticator struct {
	ServiceAuth ServiceAuthenticator
	UserAuth    UserAuthenticator
}

func (auth MultiAuthenticator) Authenticate(w http.ResponseWriter, req *http.Request) (*http.Request, error) {
	req, err := auth.UserAuth.Authenticate(w, req)
	if err != nil {
		return nil, fmt.Errorf("failed user authentication: %w", err)
	}

	req, err = auth.ServiceAuth.Authenticate(w, req)
	if err != nil {
		return nil, fmt.Errorf("failed service authentication: %w", err)
	}

	return req, nil
}

// ServiceAuthenticator parses http requests for service headers and adds service context to the
// request where possible.
type ServiceAuthenticator struct {
	Graph *ent.Client
}

// Authenticate a request by wrapping it's context with the authenticated service identity. It will
// upsert new (unactivated) service identities if the public key is not already registered. Returns
// an error if invalid credentials are provided.
func (auth ServiceAuthenticator) Authenticate(w http.ResponseWriter, req *http.Request) (*http.Request, error) {
	// TODO: @Nick
	return req, nil
	// svcName := req.Header.Get(HeaderService)
	// pubKeyB64 := req.Header.Get(HeaderIdentity)
	// sig := req.Header.Get(HeaderSignature)
	// epoch := req.Header.Get(HeaderEpoch)

	// pubKey := ed25519.PublicKey([]byte(b64_decoded_pubkey))
}

// UserAuthenticator parses http requests for session cookies and adds user context to the request
// where possible.
type UserAuthenticator struct {
	Graph *ent.Client
}

// Authenticate a request by wrapping it's context with the logged in user. If no user is logged in,
// the original request is returned. Returns an error if it fails to find the logged in user or if
// invalid credentials are provided.
func (auth UserAuthenticator) Authenticate(w http.ResponseWriter, req *http.Request) (*http.Request, error) {
	// Get requested userID, unauthenticated otherwise
	userID, err := parseUserID(req)
	if err != nil {
		return req, nil
	}

	// Get session token, unauthenticated otherwise
	token, err := parseSessionToken(req)
	if err != nil {
		return req, nil
	}

	// Load requested user object, error if no matching user found
	user, err := auth.Graph.User.Get(req.Context(), userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}
	if user == nil {
		return nil, ErrNotAuthenticated
	}

	// Authenticate as requested user
	if !token.Equals(Secret(user.SessionToken)) {
		return nil, ErrNotAuthenticated
	}

	// User Successfully Authenticated
	return CreateUserSession(w, req, user), nil
}

// parseUserID from the userid cookie.
func parseUserID(req *http.Request) (int, error) {
	userCookie, err := req.Cookie(UserCookieName)
	if err != nil {
		return -1, fmt.Errorf("no user cookie set: %w", err)
	}

	userID, err := strconv.Atoi(userCookie.Value)
	if err != nil {
		return -1, fmt.Errorf("invalid user cookie: %w", err)
	}
	return userID, nil
}

// parseSessionToken from the session token cookie.
func parseSessionToken(req *http.Request) (Secret, error) {
	sessionCookie, err := req.Cookie(SessionCookieName)
	if err != nil {
		return "", fmt.Errorf("no session cookie set: %w", err)
	}

	sessionToken := Secret(sessionCookie.Value)
	if sessionToken == "" {
		return "", fmt.Errorf("invalid session cookie")
	}

	return sessionToken, nil
}