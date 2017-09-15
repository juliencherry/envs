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

		someEnv  = "some-env"
		api      = "https://example.com/"
		username = "some-username"
		password = "some-password"
	)

	BeforeEach(func() {
		fakeStateManager = new(commandfakes.FakeStateManager)
		cmd = command.CFAddTargetCommand{
			StateManager: fakeStateManager,
		}
	})

	It("saves the given environment", func() {
		msg, err := cmd.Run([]string{
			"--name", someEnv,
			"--api", api,
			"--username", username,
			"--password", password,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		Expect(fakeStateManager.SaveEnvCallCount()).To(Equal(1))
		actualEnv, actualAPI, actualUsername, actualPassword, actualSkipSSLValidation := fakeStateManager.SaveEnvArgsForCall(0)
		Expect(actualEnv).To(Equal(someEnv))
		Expect(actualAPI).To(Equal(api))
		Expect(actualUsername).To(Equal(username))
		Expect(actualPassword).To(Equal(password))
		Expect(actualSkipSSLValidation).To(BeFalse())
	})

	It("allows flags in any order", func() {
		msg, err := cmd.Run([]string{
			"--password", password,
			"--username", username,
			"--api", api,
			"--name", someEnv,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		Expect(fakeStateManager.SaveEnvCallCount()).To(Equal(1))
		actualEnv, actualAPI, actualUsername, actualPassword, actualSkipSSLValidation := fakeStateManager.SaveEnvArgsForCall(0)
		Expect(actualEnv).To(Equal(someEnv))
		Expect(actualAPI).To(Equal(api))
		Expect(actualUsername).To(Equal(username))
		Expect(actualPassword).To(Equal(password))
		Expect(actualSkipSSLValidation).To(BeFalse())
	})

	It("requires the `name` flag", func() {
		_, err := cmd.Run([]string{
			"--api", api,
			"--username", username,
			"--password", password,
		})
		Expect(err).To(MatchError("`name` flag required"))
		Expect(fakeStateManager.SaveEnvCallCount()).To(BeZero())
	})

	It("requires the `api` flag", func() {
		_, err := cmd.Run([]string{
			"--name", someEnv,
			"--username", username,
			"--password", password,
		})
		Expect(err).To(MatchError("`api` flag required"))
		Expect(fakeStateManager.SaveEnvCallCount()).To(BeZero())
	})

	It("requires the `username` flag", func() {
		_, err := cmd.Run([]string{
			"--name", someEnv,
			"--api", api,
			"--password", password,
		})
		Expect(err).To(MatchError("`username` flag required"))
		Expect(fakeStateManager.SaveEnvCallCount()).To(BeZero())
	})

	It("requires the `password` flag", func() {
		_, err := cmd.Run([]string{
			"--name", someEnv,
			"--api", api,
			"--username", username,
		})
		Expect(err).To(MatchError("`password` flag required"))
		Expect(fakeStateManager.SaveEnvCallCount()).To(BeZero())
	})

	It("accepts the `--skip-ssl-validation` flag", func() {
		msg, err := cmd.Run([]string{
			"--name", someEnv,
			"--api", api,
			"--username", username,
			"--password", password,
			"--skip-ssl-validation",
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		Expect(fakeStateManager.SaveEnvCallCount()).To(Equal(1))
		_, _, _, _, actualSkipSSLValidation := fakeStateManager.SaveEnvArgsForCall(0)
		Expect(actualSkipSSLValidation).To(BeTrue())
	})

	It("accepts short-form flags", func() {
		msg, err := cmd.Run([]string{
			"-n", someEnv,
			"-a", api,
			"-u", username,
			"-p", password,
			"-s",
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		Expect(fakeStateManager.SaveEnvCallCount()).To(Equal(1))
		actualEnv, actualAPI, actualUsername, actualPassword, actualSkipSSLValidation := fakeStateManager.SaveEnvArgsForCall(0)
		Expect(actualEnv).To(Equal(someEnv))
		Expect(actualAPI).To(Equal(api))
		Expect(actualUsername).To(Equal(username))
		Expect(actualPassword).To(Equal(password))
		Expect(actualSkipSSLValidation).To(BeTrue())
	})

	It("errors when the flags cannot be parsed", func() {
		_, err := cmd.Run([]string{
			"--invalid-flag",
		})
		Expect(err).To(MatchError(And(HavePrefix("cannot parse flags: "))))
		Expect(fakeStateManager.SaveEnvCallCount()).To(BeZero())
	})
})
