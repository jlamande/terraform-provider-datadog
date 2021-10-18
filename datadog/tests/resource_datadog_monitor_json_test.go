package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatadogMonitorJSONImport(t *testing.T) {
	ctx, accProviders := testAccProviders(context.Background(), t)
	uniqueID := uniqueEntityName(ctx, t)
	accProvider := testAccProvider(t, accProviders)

	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: accProviders,
		// Use checkMonitorDestroy() from Monitor resource
		CheckDestroy: testAccCheckDatadogMonitorDestroy(accProvider),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogMonitorJSON(uniqueID),
			},
			{
				ResourceName:      "datadog_monitor_json.monitor_json",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDatadogMonitorJSON(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_monitor_json" "monitor_json" {
  monitor =<<-EOF
{
    "name": "%s",
    "type": "service check",
    "query": "\"ntp.in_sync\".over(\"*\").last(2).count_by_status()",
    "message": "Change the message triggers if any host's clock goes out of sync with the time given by NTP. The offset threshold is configured in the Agent's 'ntp.yaml' file.\n\nSee [Troubleshooting NTP Offset issues](https://docs.datadoghq.com/agent/troubleshooting/ntp for more details on cause and resolution.",
    "tags": [],
    "multi": true,
	"restricted_roles": null,
    "options": {
        "include_tags": true,
        "locked": false,
        "new_host_delay": 150,
        "notify_audit": false,
        "notify_no_data": false,
        "thresholds": {
            "warning": 1,
            "ok": 1,
            "critical": 1
        },
        "silenced": {}
    },
    "priority": null,
    "classification": "custom"
}
EOF
}`, uniq)

}
