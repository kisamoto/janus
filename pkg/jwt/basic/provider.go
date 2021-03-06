package basic

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kisamoto/janus/pkg/config"
	"github.com/kisamoto/janus/pkg/jwt/provider"
)

func init() {
	provider.Register("basic", &Provider{})
}

// Provider abstracts the authentication for github
type Provider struct {
	provider.Verifier
}

// Build acts like the constructor for a provider
func (gp *Provider) Build(config config.Credentials) provider.Provider {
	return &Provider{
		Verifier: provider.NewVerifierBasket(
			NewPasswordVerifier(userConfigToTeam(config.Basic.Users)),
		),
	}
}

// GetClaims returns a JWT Map Claim
func (gp *Provider) GetClaims(httpClient *http.Client) (jwt.MapClaims, error) {
	return jwt.MapClaims{}, nil
}

func userConfigToTeam(configUser map[string]string) []*user {
	users := []*user{}
	for u, p := range configUser {
		users = append(users, &user{
			Username: u,
			Password: p,
		})
	}
	return users
}
