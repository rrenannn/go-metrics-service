package notification

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	db "go-metrics-service/db/sqlc"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNotification(ctx context.Context, typ, to string, body map[string]interface{}) (string, error) {
	id := uuid.New()

	payload := map[string]interface{}{
		"notification_id": id,
		"type":            typ,
		"to":              to,
		"body":            body,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	if err = s.repo.InsertNotification(ctx, db.InsertNotificationParams{
		NotificationID: id,
		Payload: pqtype.NullRawMessage{
			RawMessage: b,
			Valid:      true,
		},
		Status: "queued",
	}); err != nil {
		return "", err
	}

}
