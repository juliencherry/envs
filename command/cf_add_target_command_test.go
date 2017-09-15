package command_test

import (
	"github.com/juliencherry/envs/command"
	"github.com/juliencherry/envs/command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFAddTargetCommand", func() {
	var (
		cmd              command.CFAddTargetCommand
		fakeStateManager *commandfakes.FakeStateManager
	)

	BeforeEach(func() {
		fakeStateManager = new(commandfakes.FakeStateManager)
		cmd = command.CFAddTargetCommand{
			StateManager: fakeStateManager,
		}
	})

	It("saves an environment", func() {
		someEnv := "some-env"
		msg, err := cmd.Run([]string{someEnv})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		Expect(fakeStateManager.SaveEnvCallCount()).To(Equal(1))
		Expect(fakeStateManager.SaveEnvArgsForCall(0)).To(Equal(someEnv))
	})

	Context("when no target is given", func() {
		It("indicates that a target must be specified", func() {
			_, err := cmd.Run([]string{})
			Expect(err).To(MatchError("Missing required argument"))
		})
	})
})
