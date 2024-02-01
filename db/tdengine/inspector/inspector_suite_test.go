package inspector_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInspector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inspector Suite")
}
