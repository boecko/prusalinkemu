package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/prusa3d/prusa-connect/sdk/wirerpc"
)

// TestAPathThisModuleDoesNotServeIsRefusedByCode is the refusal check. A client
// branches on the code, so the test reads the code and not the sentence: a
// reworded sentence must not break a caller, and a changed code must.
func TestAPathThisModuleDoesNotServeIsRefusedByCode(t *testing.T) {
	recorder := httptest.NewRecorder()
	routes().ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/nothing-here", nil))

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("an unserved path answered %d, want %d", recorder.Code, http.StatusNotFound)
	}

	if got := recorder.Header().Get("Content-Type"); got != "application/problem+json" {
		t.Fatalf("the refusal is %s, want application/problem+json", got)
	}

	var problem struct {
		Code   string            `json:"code"`
		Title  string            `json:"title"`
		Detail string            `json:"detail"`
		Params map[string]string `json:"params"`
	}

	if err := json.Unmarshal(recorder.Body.Bytes(), &problem); err != nil {
		t.Fatalf("the refusal is not JSON: %v", err)
	}

	if want := RouteNotFound.Code().String(); problem.Code != want {
		t.Fatalf("the refusal carries %q, want %q", problem.Code, want)
	}

	// The English is kept beside the code, so a log line and an untranslated
	// screen still read.
	if problem.Title == "" || problem.Detail == "" {
		t.Fatalf("the refusal lost its English: title %q, detail %q", problem.Title, problem.Detail)
	}

	if problem.Params["path"] != "/nothing-here" {
		t.Fatalf("the refusal names the path %q, want /nothing-here", problem.Params["path"])
	}
}

// TestTheManifestPublishesTheCodesDefault holds module.yaml and this code to
// one default. An operator reads the manifest to learn what a setting does when
// they write nothing, so a manifest that publishes a different number than the
// code uses is documentation that lies.
func TestTheManifestPublishesTheCodesDefault(t *testing.T) {
	manifest, err := os.ReadFile("module.yaml")
	if err != nil {
		t.Fatalf("read module.yaml: %v", err)
	}

	if declared := `default: "` + strconv.Itoa(defaultPort) + `"`; !strings.Contains(string(manifest), declared) {
		t.Fatalf("module.yaml does not declare %s for the port setting", declared)
	}
}

func TestConfigureDefaultsThePort(t *testing.T) {
	saved := held.Load()
	held.Store(nil)
	t.Cleanup(func() { held.Store(saved) })

	for name, tc := range map[string]struct {
		config string
		want   int
	}{
		"no port written": {`{}`, defaultPort},
		"a port written":  {`{"port":8080}`, 8080},
	} {
		t.Run(name, func(t *testing.T) {
			params, err := json.Marshal(wirerpc.Settings{Config: json.RawMessage(tc.config)})
			if err != nil {
				t.Fatalf("render the settings: %v", err)
			}

			if _, err := methods()[wirerpc.MethodConfigure](context.Background(), wirerpc.Call{
				Method: wirerpc.MethodConfigure,
				Params: params,
			}); err != nil {
				t.Fatalf("configure refused the settings: %v", err)
			}

			if got := settings().Port; got != tc.want {
				t.Fatalf("the held port is %d, want %d", got, tc.want)
			}
		})
	}
}
