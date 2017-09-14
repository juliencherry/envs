package command_test

import (
	"io/ioutil"
	"os"

	"github.com/juliencherry/envs/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFAddTargetCommand", func() {
	var envStateFilepath string

	BeforeEach(func() {
		envStateFile, err := ioutil.TempFile("", "env-state")
		Expect(err).NotTo(HaveOccurred())
		envStateFilepath = envStateFile.Name()
		os.Setenv("ENV_STATE_FILEPATH", envStateFilepath)
	})

	AfterEach(func() {
		os.Remove(envStateFilepath)
		os.Unsetenv("ENV_STATE_FILEPATH")
	})

	It("allows multiple targets to be added", func() {
		someEnv := "some-env"
		msg, err := command.CFAddTargetCommand{}.Run([]string{someEnv})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		someOtherEnv := "some-other-env"
		msg, err = command.CFAddTargetCommand{}.Run([]string{someOtherEnv})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-other-env"`))

		contents, err := ioutil.ReadFile(envStateFilepath)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(contents)).To(Equal(someEnv + "\n" + someOtherEnv + "\n"))
	})

	It("does not add duplicate targets", func() {
		someEnv := "some-env"
		msg, err := command.CFAddTargetCommand{}.Run([]string{someEnv})
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(`Added target "some-env"`))

		_, err = command.CFAddTargetCommand{}.Run([]string{someEnv})
		Expect(err).To(MatchError(`target "some-env" already exists`))

		contents, err := ioutil.ReadFile(envStateFilepath)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(contents)).To(Equal(someEnv + "\n"))
	})

	Context("when no target is given", func() {
		It("indicates that a target must be specified", func() {
			_, err := command.CFAddTargetCommand{}.Run([]string{})
			Expect(err).To(MatchError("Missing required argument"))

			contents, err := ioutil.ReadFile(envStateFilepath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(contents)).To(BeEmpty())
		})
	})
})
