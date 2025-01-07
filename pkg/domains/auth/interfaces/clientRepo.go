package interfaces

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/infrastructures/models"
)

type IClientRepo interface {
	GetClientById(ctx context.Context, clientId string) (result *models.Client, err error)
	GetAllClients(ctx context.Context, page, pageSize *int32, search *string) (result []*dtos.PublicClientDto, total int64, err error)

	CreateClient(ctx context.Context, client *models.Client) (result *models.Client, err error)
	UpdateClient(ctx context.Context, client *models.Client) (result *models.Client, err error)
	UpdateStatus(ctx context.Context, clientId string, status bool) (err error)

	UpdateClientSecret(ctx context.Context, clientId, secret string) (err error)
}
