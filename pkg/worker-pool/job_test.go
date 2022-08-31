package workerpool

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobExecuteReturnSuccess(t *testing.T) {
	expectedArgs := []int{1, 2, 3}
	job := Job{
		Id: 1,
		Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
			require.IsType(t, expectedArgs, args)
			require.Equal(t, expectedArgs, args)
			return 1, nil
		},
		Args: expectedArgs,
	}

	result := job.execute(context.Background())
	require.Equal(t, 1, result.JobId)
	require.Nil(t, result.Error)
	require.IsType(t, 1, result.Value)
	require.Equal(t, 1, result.Value)
}

func TestJobExecuteReturnError(t *testing.T) {
	expectedArgs := []int{1, 2, 3}
	job := Job{
		Id: 1,
		Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
			require.IsType(t, expectedArgs, args)
			require.Equal(t, expectedArgs, args)
			return nil, errors.New("Something went wrong")
		},
		Args: expectedArgs,
	}

	result := job.execute(context.Background())
	require.Equal(t, 1, result.JobId)
	require.Nil(t, result.Value)
	require.Error(t, result.Error)
}
