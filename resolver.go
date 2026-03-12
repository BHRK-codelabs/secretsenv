package secretsenv

import (
	"context"
	"os"
	"strings"

	"github.com/BHRK-codelabs/secretskit"
)

type Resolver struct {
	prefix string
}

func New(prefix string) *Resolver {
	return &Resolver{prefix: strings.TrimSpace(prefix)}
}

func (r *Resolver) Get(_ context.Context, name string) (secretskit.Secret, error) {
	key := r.envKey(name)
	value, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(value) == "" {
		return secretskit.Secret{}, secretskit.ErrSecretNotFound
	}
	return secretskit.Secret{Name: name, Value: value}, nil
}

func (r *Resolver) GetMany(ctx context.Context, names []string) (map[string]secretskit.Secret, error) {
	result := make(map[string]secretskit.Secret, len(names))
	for _, name := range names {
		secret, err := r.Get(ctx, name)
		if err != nil {
			return nil, err
		}
		result[name] = secret
	}
	return result, nil
}

func (r *Resolver) envKey(name string) string {
	parts := []string{}
	if r.prefix != "" {
		parts = append(parts, r.prefix)
	}
	parts = append(parts, name)

	raw := strings.Join(parts, "_")
	replacer := strings.NewReplacer("/", "_", "-", "_", ".", "_", " ", "_")
	return strings.ToUpper(replacer.Replace(raw))
}
