package oci_test

import (
	. "github.com/go-skynet/LocalAI/pkg/oci" // Update with your module path
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OCI", func() {

	Context("when template is loaded successfully", func() {
		FIt("should evaluate the template correctly", func() {

			// https://registry.ollama.ai/v2/library/gemma/manifests/2b
			// "application/vnd.ollama.image.model"
			err := FetchImage("registry.ollama.ai/library/gemma", "sha256:c1864a5eb19305c40519da12cc543519e48a0697ecd30e15d5ac228644957d12", "/tmp/foo")
			Expect(err).NotTo(HaveOccurred())

		})
	})

})
