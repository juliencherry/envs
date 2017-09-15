package command_test

import (
	"github.com/juliencherry/envs/command"
	"github.com/juliencherry/envs/command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFTargetsCommand", func() {
	var (
		cmd              command.CFTargetsCommand
		fakeStateManager *commandfakes.FakeStateManager
	)

	BeforeEach(func() {
		fakeStateManager = new(commandfakes.FakeStateManager)
		cmd = command.CFTargetsCommand{
			StateManager: fakeStateManager,
		}
	})

	It("returns the environment from the state file", func() {
		fakeStateManager.GetEnvsReturns([]string{
			"some-env",
			"some-other-env",
		}, nil)

		msg, err := cmd.Run(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal("some-env\nsome-other-env"))
		Expect(fakeStateManager.GetEnvsCallCount()).To(Equal(1))
	})

	Context("when no targets have been added", func() {
		It("indicates that no targets are available", func() {
			fakeStateManager.GetEnvsReturns([]string{}, nil)

			msg, err := cmd.Run(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(msg).To(Equal("No targets available"))
			Expect(fakeStateManager.GetEnvsCallCount()).To(Equal(1))
		})
	})
})
