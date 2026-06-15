package main

import (
	"testing"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
)

func TestNewDexLegacyOptionsKeepsEmailClaim(t *testing.T) {
	legacyOpts := newDexLegacyOptions(Options{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		IssuerURL:    "https://dex.example.com",
		Scopes:       "openid email profile",
	})

	if got := legacyOpts.LegacyProvider.OIDCEmailClaim; got != options.OIDCEmailClaim {
		t.Fatalf("unexpected OIDC email claim: got %q, want %q", got, options.OIDCEmailClaim)
	}
	if got := legacyOpts.LegacyProvider.UserIDClaim; got != options.OIDCEmailClaim {
		t.Fatalf("unexpected deprecated user ID claim: got %q, want %q", got, options.OIDCEmailClaim)
	}

	oauthProxyOpts, err := legacyOpts.ToOptions()
	if err != nil {
		t.Fatalf("ToOptions returned error: %v", err)
	}
	if len(oauthProxyOpts.Providers) != 1 {
		t.Fatalf("unexpected provider count: got %d, want 1", len(oauthProxyOpts.Providers))
	}

	oidcConfig := oauthProxyOpts.Providers[0].OIDCConfig
	if got := oidcConfig.EmailClaim; got != options.OIDCEmailClaim {
		t.Fatalf("unexpected converted OIDC email claim: got %q, want %q", got, options.OIDCEmailClaim)
	}
	if got := oidcConfig.UserIDClaim; got != options.OIDCEmailClaim {
		t.Fatalf("unexpected converted deprecated user ID claim: got %q, want %q", got, options.OIDCEmailClaim)
	}
}
