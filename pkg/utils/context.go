package utils

import (
	"context"
	"errors"
	"strconv"
)

func GetUserIdFromContext(ctx context.Context) (int, error) {
	userIdAny := ctx.Value("userId")
	switch v := userIdAny.(type) {
	case string:
		if val, err := strconv.Atoi(v); err != nil {
			return 0, err
		} else {
			return val, nil
		}
	case int:
		return v, nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	default:
		return 0, errors.New("unknown type of user id")
	}
}
