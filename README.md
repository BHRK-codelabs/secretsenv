# secretsenv

`secretsenv` is an environment-backed adapter for `secretskit`.

It is useful for local development or runtimes where secrets are injected as environment variables.

Example:

```powershell
$env:DEMO_UPSTREAM_API_KEY_UPSTREAM_API_KEY="demo-api-key"
```

```go
resolver := secretsenv.New("demo_upstream_api_key")
secret, err := resolver.Get(ctx, "upstream/api-key")
```
