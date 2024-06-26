package webhooks

import (
	"reflect"
	"testing"

	"github.com/influxdata/telegraf/plugins/inputs/webhooks/artifactory"
	"github.com/influxdata/telegraf/plugins/inputs/webhooks/github"
	"github.com/influxdata/telegraf/plugins/inputs/webhooks/papertrail"
	"github.com/influxdata/telegraf/plugins/inputs/webhooks/particle"
	"github.com/influxdata/telegraf/plugins/inputs/webhooks/rollbar"
	"github.com/influxdata/telegraf/plugins/inputs/webhooks/zabbix"
)

func TestAvailableWebhooks(t *testing.T) {
	wb := NewWebhooks()
	expected := make([]Webhook, 0)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Github = &github.GithubWebhook{Path: "/github"}
	expected = append(expected, wb.Github)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Rollbar = &rollbar.RollbarWebhook{Path: "/rollbar"}
	expected = append(expected, wb.Rollbar)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Papertrail = &papertrail.PapertrailWebhook{Path: "/papertrail"}
	expected = append(expected, wb.Papertrail)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Particle = &particle.ParticleWebhook{Path: "/particle"}
	expected = append(expected, wb.Particle)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Artifactory = &artifactory.ArtifactoryWebhook{Path: "/artifactory"}
	expected = append(expected, wb.Artifactory)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}
	wb.Zabbix = &zabbix.ZabbixWebhook{Path: "/zabbix"}
	expected = append(expected, wb.Zabbix)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}
}
