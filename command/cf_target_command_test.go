package command_test

import (
	"errors"

	"github.com/juliencherry/envs/command"
	"github.com/juliencherry/envs/command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFTargetCommand", func() {

	var (
		cmd          command.CFTargetCommand
		stateManager *commandfakes.FakeStateManager
		cf           *commandfakes.FakeCFCLIWrapper
	)

	BeforeEach(func() {
		stateManager = new(commandfakes.FakeStateManager)
		cf = new(commandfakes.FakeCFCLIWrapper)
		cmd = command.CFTargetCommand{
			CFCLIWrapper: cf,
			StateManager: stateManager,
		}
	})

	Context("when no environment name is given", func() {
		It("indicates that no environments are targeted", func() {
			msg, err := cmd.Run(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(msg).To(Equal("No environment targeted"))
		})
	})

	Context("when an environment name is given", func() {

		var (
			someEnv               string
			someAPI               string
			someUsername          string
			somePassword          string
			someSkipSSLValidation bool
		)

		BeforeEach(func() {
			someEnv = "some-env"
			someAPI = "some-api"
			someUsername = "some-username"
			somePassword = "some-password"
			someSkipSSLValidation = true
		})

		Context("when the state manager returns the environment", func() {

			BeforeEach(func() {
				stateManager.GetEnvDetailsReturns(someAPI, someUsername, somePassword, someSkipSSLValidation, nil)
			})

			It("targets the given environment", func() {
				msg, err := cmd.Run([]string{someEnv})
				Expect(err).NotTo(HaveOccurred())
				Expect(msg).To(Equal(`Targeted environment "some-env"`))

				Expect(stateManager.GetEnvDetailsCallCount()).To(Equal(1))
				Expect(stateManager.GetEnvDetailsArgsForCall(0)).To(Equal(someEnv))

				Expect(cf.LoginCallCount()).To(Equal(1))
				actualApi, actualUsername, actualPassword, actualSkipSSLValidation := cf.LoginArgsForCall(0)
				Expect(actualApi).To(Equal(someAPI))
				Expect(actualUsername).To(Equal(someUsername))
				Expect(actualPassword).To(Equal(somePassword))
				Expect(actualSkipSSLValidation).To(Equal(someSkipSSLValidation))
			})

			It("errors when the login command fails", func() {
				cf.LoginReturns(errors.New("some-error"))
				_, err := cmd.Run([]string{someEnv})
				Expect(err).To(MatchError(And(HavePrefix("cannot target environment: "))))
			})
		})

		Context("when the state manager returns no environment", func() {
			BeforeEach(func() {
				stateManager.GetEnvDetailsReturns("", "", "", false, errors.New("cannot find environment"))
			})

			It("errors", func() {
				_, err := cmd.Run([]string{someEnv})
				Expect(err).To(MatchError(And(HavePrefix("cannot target environment: "))))
				Expect(stateManager.GetEnvDetailsCallCount()).To(Equal(1))
				Expect(cf.LoginCallCount()).To(BeZero())
			})
		})
	})
})
