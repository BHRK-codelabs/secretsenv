package secretsenv

import (
	"context"
	"testing"
)

func TestResolverReadsFromEnv(t *testing.T) {
	t.Setenv("DEMO_UPSTREAM_API_KEY_UPSTREAM_API_KEY", "secret-value")

	resolver := New("demo_upstream_api_key")
	secret, err := resolver.Get(context.Background(), "upstream/api-key")
	if err != nil {
		t.Fatalf("get secret: %v", err)
	}
	if secret.Value != "secret-value" {
		t.Fatalf("unexpected secret value: %s", secret.Value)
	}
}

func TestResolverReturnsNotFound(t *testing.T) {
	resolver := New("")
	if _, err := resolver.Get(context.Background(), "upstream/api-key"); err == nil {
		t.Fatal("expected not found")
	}
}
