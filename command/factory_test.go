package command_test

import (
	"github.com/juliencherry/envs/command"
	"github.com/juliencherry/envs/command/commandfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {

	var (
		cfCLIWrapper = new(commandfakes.FakeCFCLIWrapper)
		stateManager = new(commandfakes.FakeStateManager)
	)

	It("builds a `cf-add-target`", func() {
		cmd, err := command.Build("cf-add-target", nil, stateManager)
		Expect(err).NotTo(HaveOccurred())
		cfAddTargetCommand, ok := cmd.(command.CFAddTargetCommand)
		Expect(ok).To(BeTrue())
		Expect(cfAddTargetCommand.StateManager).To(Equal(stateManager))
	})

	It("builds a `cf-target` command", func() {
		cmd, err := command.Build("cf-target", cfCLIWrapper, stateManager)
		Expect(err).NotTo(HaveOccurred())
		cfTargetCommand, ok := cmd.(command.CFTargetCommand)
		Expect(ok).To(BeTrue())
		Expect(cfTargetCommand.CFCLIWrapper).To(Equal(cfCLIWrapper))
		Expect(cfTargetCommand.StateManager).To(Equal(stateManager))
	})

	It("builds a `cf-targets` command", func() {
		cmd, err := command.Build("cf-targets", nil, stateManager)
		Expect(err).NotTo(HaveOccurred())
		cfTargetsCommand, ok := cmd.(command.CFTargetsCommand)
		Expect(ok).To(BeTrue())
		Expect(cfTargetsCommand.StateManager).To(Equal(stateManager))
	})

	It("errors if there is no command", func() {
		_, err := command.Build("some-bad-command", nil, stateManager)
		Expect(err).To(MatchError("cannot build command"))
	})
})
