<!-- ...existing code... -->

# go-iracing

## Overview

A simple Go client and types for the iRacing data API.

It Provides HTTP client, request/response types and example CLI usage to fetch members, tracks, cars, results and stats.

Features
- Lightweight HTTP client with authentication helpers: [`iracing.NewClient`](pkg/iracing/client.go) and [`iracing.Client`](pkg/iracing/client.go).
- Typed request/response structs under [`pkg/types`](pkg/types) (example: [`types.ResultsSeasonResultsReq`](pkg/types/results.go)).
- Example implementation in [cmd/main.go](cmd/main.go).

## Authentication

### iRacing legacy authentication

The API supports the legacy authentication which relies on the `` endpoint.

> This legacy authentication endpoint is deprecated, and can be disabled soon by iRacing (seen [announcement](https://forums.iracing.com/discussion/84226/legacy-authentication-removal-dec-9-2025/p1)). I encourage to move directly to the new method. However the new method requires to request credentials from iRacing staff, thus the legacy method can be used as a quick start, waiting for the credentials.

To use the legacy authentication:

```go
client, err := iracing.NewClient(
    iracing.WithAuthenticator(iracing.NewIrLegacyAuth(configEnv.Email, configEnv.Password)),
)
```

### iRacing OAuth authentication

The go-iracing uses the `password_limited` flow, as it is assumed that the API will be consumed from a server application without UI.

[See the full documentation here](https://oauth.iracing.com/oauth2/book/password_limited_flow.html)

Once you have received your client id and secret, use the authentication as follows:

```go
client, err := iracing.NewClient(
    iracing.WithAuthenticator(iracing.NewIrOAuth(
        configEnv.Email, configEnv.Password,
        configEnv.ClientID, configEnv.ClientSecret)),
)
```

## API coverage

The full iRacing API has been implemented. However the documentation from iRacing is quite weak, and despite the tests that I have tried to perform, it is totally possible that some unmarshalling fail in cases that I could not easily test. If it happens, don't hesitate to open an issue (or a PR) with all information that will help me to reproduce the case, such as:

* The API which has the issue
* The data which allows to request the API, typically: cust id, subsession id, ...

## License

This project is MIT licensed — see [LICENSE](LICENSE).