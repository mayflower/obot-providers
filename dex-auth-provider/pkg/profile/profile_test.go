package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchDexProfile(t *testing.T) {
	mockProfile := dexProfile{
		Sub:               "user-123",
		Email:             "user@example.com",
		EmailVerified:     true,
		Name:              "Test User",
		PreferredUsername: "testuser",
		Groups:            []string{"engineering", "admins"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedAuth := "Bearer mock_dex_token"
		if got := r.Header.Get("Authorization"); got != expectedAuth {
			http.Error(w, fmt.Sprintf("unexpected auth header: got %s", got), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockProfile)
	}))
	defer server.Close()

	profile, err := FetchDexProfile(context.Background(), "Bearer mock_dex_token", server.URL)
	if err != nil {
		t.Fatalf("FetchDexProfile returned error: %v", err)
	}

	if profile.Sub != mockProfile.Sub {
		t.Errorf("unexpected Sub: got %s, want %s", profile.Sub, mockProfile.Sub)
	}
	if profile.Email != mockProfile.Email {
		t.Errorf("unexpected Email: got %s, want %s", profile.Email, mockProfile.Email)
	}
	if profile.EmailVerified != mockProfile.EmailVerified {
		t.Errorf("unexpected EmailVerified: got %v, want %v", profile.EmailVerified, mockProfile.EmailVerified)
	}
	if profile.Name != mockProfile.Name {
		t.Errorf("unexpected Name: got %s, want %s", profile.Name, mockProfile.Name)
	}
	if profile.PreferredUsername != mockProfile.PreferredUsername {
		t.Errorf("unexpected PreferredUsername: got %s, want %s", profile.PreferredUsername, mockProfile.PreferredUsername)
	}
	if len(profile.Groups) != len(mockProfile.Groups) {
		t.Fatalf("unexpected Groups length: got %d, want %d", len(profile.Groups), len(mockProfile.Groups))
	}
	for i, group := range mockProfile.Groups {
		if profile.Groups[i] != group {
			t.Errorf("unexpected group at %d: got %s, want %s", i, profile.Groups[i], group)
		}
	}
}
