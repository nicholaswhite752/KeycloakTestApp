# KeycloakTestApp

Notes: 

In gocloak client.go the urls it creates for certs and token auth are for an old version. Had to modify functions
Should look into replacing or updating the client with the minimum needed functions with correct URLs

```
func (client *gocloak) getNewCerts(ctx context.Context, realm string) (*CertResponse, error) {
	const errMessage = "could not get newCerts"

	var result CertResponse
	resp, err := client.getRequest(ctx).
		SetResult(&result).
		Get("http://localhost:8080/realms/TestRealm/protocol/openid-connect/certs")

	if err := checkForError(resp, err, errMessage); err != nil {
		return nil, err
	}

	return &result, nil
}
```

```
// RetrospectToken calls the openid-connect introspect endpoint
func (client *gocloak) RetrospectToken(ctx context.Context, accessToken, clientID, clientSecret, realm string) (*RetrospecTokenResult, error) {
	const errMessage = "could not introspect requesting party token"

	var result RetrospecTokenResult
	resp, err := client.getRequestWithBasicAuth(ctx, clientID, clientSecret).
		SetFormData(map[string]string{
			"token_type_hint": "requesting_party_token",
			"token":           accessToken,
		}).
		SetResult(&result).
		Post("http://localhost:8080/realms/TestRealm/protocol/openid-connect/token/introspect")

	if err := checkForError(resp, err, errMessage); err != nil {
		return nil, err
	}

	return &result, nil
}
```