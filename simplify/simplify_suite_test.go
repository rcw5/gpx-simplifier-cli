package simplify_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSimplify(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Simplify Suite")
}
