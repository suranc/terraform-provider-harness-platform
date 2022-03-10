package sweep_test

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness-platform/internal/sweep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	client := nextgen.NewAPIClient(nextgen.DefaultConfiguration())
	sweep.SweeperClient = client
	resource.TestMain(m)
}
