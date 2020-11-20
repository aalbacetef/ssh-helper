package configfile

import "testing"

func TestHostEntry(t *testing.T) {
	validhost := HostEntry{
		Host:     "local-test",
		HostName: "local.test",
	}
	invalidhost := HostEntry{}

	t.Run("detects-valid-hosts", func(tt *testing.T) {
		if err := validhost.Valid(); err != nil {
			tt.Fatalf("valid host is reporting errors: %v", err)
		}
	})

	t.Run("detects-invalid-hosts", func(tt *testing.T) {
		if err := invalidhost.Valid(); err == nil {
			tt.Fatal("invalid host is not reporting errors")
		}
	})
}
