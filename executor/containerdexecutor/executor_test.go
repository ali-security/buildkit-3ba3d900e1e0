package containerdexecutor

import (
	"context"
	"testing"

	ctd "github.com/containerd/containerd/v2/client"
	"github.com/moby/buildkit/executor"
	gatewayapi "github.com/moby/buildkit/frontend/gateway/pb"
	"github.com/stretchr/testify/require"
)

func TestContainerdUnknownExitStatus(t *testing.T) {
	// There are assumptions in the containerd executor that the UnknownExitStatus
	// used in errdefs.ExitError matches the variable in the containerd package.
	if ctd.UnknownExitStatus != gatewayapi.UnknownExitStatus {
		t.Fatalf("containerd.UnknownExitStatus != errdefs.UnknownExitStatus")
	}
}

func TestContainerdRunInvalidContainerID(t *testing.T) {
	t.Parallel()

	// a client-provided container id must be rejected before it is used to
	// build any path or container state
	w := &containerdExecutor{}
	_, err := w.Run(context.TODO(), "../../escape", executor.Mount{}, nil, executor.ProcessInfo{}, nil)
	require.ErrorContains(t, err, "invalid container id")
}
