package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/manicminer/hamilton/clients"

	"github.com/terraform-providers/terraform-provider-azuread/internal/common"
)

type Client struct {
	AadClient *graphrbac.UsersClient
	MsClient  *clients.UsersClient
}

func NewClient(o *common.ClientOptions) *Client {
	aadClient := graphrbac.NewUsersClientWithBaseURI(o.AadGraphEndpoint, o.TenantID)
	msClient := clients.NewUsersClient(o.TenantID)
	o.ConfigureClient(&msClient.BaseClient, &aadClient.Client)

	return &Client{
		AadClient: &aadClient,
		MsClient:  msClient,
	}
}
