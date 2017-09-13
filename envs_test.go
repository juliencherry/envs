package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Envs", func() {

	var pathToCLI string

	BeforeSuite(func() {
		var err error
		pathToCLI, err = Build("github.com/juliencherry/envs")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when no CF targets have been added", func() {
		It("outputs that no targets are available", func() {
			command := exec.Command(pathToCLI, "cf-targets")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(Exit(0))
			Expect(session).To(Say("No targets available"))
		})
	})
})
