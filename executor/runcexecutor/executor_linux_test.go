package runcexecutor

import (
	"context"
	"testing"

	"github.com/moby/buildkit/executor"
	"github.com/stretchr/testify/require"
)

func TestRuncRunInvalidContainerID(t *testing.T) {
	t.Parallel()

	// a client-provided container id must be rejected before it is joined
	// into the bundle path under the executor root
	w := &runcExecutor{root: t.TempDir()}
	_, err := w.Run(context.TODO(), "../../escape", executor.Mount{}, nil, executor.ProcessInfo{}, nil)
	require.ErrorContains(t, err, "invalid container id")
}
