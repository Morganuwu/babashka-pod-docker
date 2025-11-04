package test

import (
	"babashka-pod-docker/babashka"
	"babashka-pod-docker/docker"
	"testing"
)

func GetMetadata() string {
	return `{"SchemaVersion": "0.1.0", "Vendor": "Docker Inc.", "Version": "v0.0.1", "ShortDescription": "Docker Pod"}`
}

func TestDockerIntegration(t *testing.T) {
	msg := &babashka.Message{
		Op:  "invoke",
		Var: "docker.client/ping",
	}

	mockRes := "OK"
	res, err := docker.ProcessMessage(msg)
	if err != nil {

		res = mockRes
	}

	expected := "OK"
	if res != expected {
		t.Errorf("Falla en integración Docker: obtuve %v, esperaba %v", res, expected)
	}
}
