package provider

import (
	"context"
	"os"

	"github.com/loft-sh/devpod/e2e/framework"
	"github.com/onsi/ginkgo/v2"
)

var _ = DevPodDescribe("devpod import workspace test suite", func() {
	ctx := context.Background()
	initialDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ginkgo.It("should import workspace", func() {
		tempDir, err := framework.CopyToTempDir("tests/provider/testdata")
		framework.ExpectNoError(err)
		ginkgo.DeferCleanup(framework.CleanupTempDir, initialDir, tempDir)

		f := framework.NewDefaultFramework(initialDir + "/bin")

		//todo: setup everything

		err = f.DevPodImportWorkspace(ctx, tempDir+"/workspace.yaml")
		framework.ExpectNoError(err)

		//todo: check that workspace is imported

	}
}