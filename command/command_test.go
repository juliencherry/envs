package command_test

import (
	"github.com/juliencherry/envs/command"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Command", func() {

	It("builds a `cf-add-target`", func() {
		cmd, err := command.Build("cf-add-target")
		Expect(err).NotTo(HaveOccurred())
		_, ok := cmd.(command.CFAddTargetCommand)
		Expect(ok).To(BeTrue())
	})

	It("builds a `cf-target` command", func() {
		cmd, err := command.Build("cf-target")
		Expect(err).NotTo(HaveOccurred())
		_, ok := cmd.(command.CFTargetCommand)
		Expect(ok).To(BeTrue())
	})

	It("builds a `cf-targets` command", func() {
		cmd, err := command.Build("cf-targets")
		Expect(err).NotTo(HaveOccurred())
		_, ok := cmd.(command.CFTargetsCommand)
		Expect(ok).To(BeTrue())
	})

	It("errors if there is no command", func() {
		_, err := command.Build("some-bad-command")
		Expect(err).To(MatchError("cannot build command"))
	})
})
