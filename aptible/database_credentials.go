package aptible

import (
	"github.com/aptible/go-deploy/client/operations"
)

type DatabaseCredential struct {
	ID            int64
	ConnectionURL string
	Type          string
}

func (c *Client) GetDatabaseCredentials(databaseID int64) ([]DatabaseCredential, error) {
	var databaseCredentials []DatabaseCredential

	params := operations.NewGetDatabasesDatabaseIDDatabaseCredentialsParams().WithDatabaseID(databaseID)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDDatabaseCredentials(params, c.Token)
	if err != nil {
		return databaseCredentials, err
	}

	for _, cred := range resp.Payload.Embedded.DatabaseCredentials {
		databaseCredential := DatabaseCredential{
			ID:            cred.ID,
			ConnectionURL: cred.ConnectionURL,
			Type:          cred.Type,
		}
		databaseCredentials = append(databaseCredentials, databaseCredential)
	}

	return databaseCredentials, nil
}
