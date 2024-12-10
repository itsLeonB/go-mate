package util

import (
	"context"

	"github.com/google/uuid"
	"github.com/rotisserie/eris"
)

func GetUUIDFromContext(ctx context.Context, key string) (uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Value(key).(string))
	if err != nil {
		return uuid.Nil, eris.Wrap(err, "error while parsing uuid")
	}

	return id, nil
}
