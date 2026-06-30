package main

import (
	"os"
	"testing"

	"github.com/metacubex/mihomo/component/oix"
)

func TestApplyOIXOverridesPropagatesFlagsToRuntime(t *testing.T) {
	oldToken := oixToken
	oldProviderName := oixProviderName
	oldEnvToken, hadEnvToken := os.LookupEnv("OIX_TOKEN")
	oldEnvProvider, hadEnvProvider := os.LookupEnv("OIX_PROVIDER_NAME")
	oix.SetProviderName("")
	t.Cleanup(func() {
		oixToken = oldToken
		oixProviderName = oldProviderName
		oix.SetProviderName("")
		if hadEnvToken {
			os.Setenv("OIX_TOKEN", oldEnvToken)
		} else {
			os.Unsetenv("OIX_TOKEN")
		}
		if hadEnvProvider {
			os.Setenv("OIX_PROVIDER_NAME", oldEnvProvider)
		} else {
			os.Unsetenv("OIX_PROVIDER_NAME")
		}
	})

	os.Unsetenv("OIX_TOKEN")
	os.Unsetenv("OIX_PROVIDER_NAME")
	oixToken = "cli-token"
	oixProviderName = "cliProvider"

	applyOIXOverrides()

	if got := os.Getenv("OIX_TOKEN"); got != "cli-token" {
		t.Fatalf("expected OIX_TOKEN from CLI flag, got %q", got)
	}
	if got := os.Getenv("OIX_PROVIDER_NAME"); got != "cliProvider" {
		t.Fatalf("expected OIX_PROVIDER_NAME from CLI flag, got %q", got)
	}
	if got := oix.ProviderFile(); got != "cliProvider" {
		t.Fatalf("expected provider file from CLI flag, got %q", got)
	}
}
