package command_test

import (
	"io/ioutil"
	"os"

	"github.com/juliencherry/envs/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFTargetsCommand", func() {
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

	It("returns the environment from the state file", func() {
		environments := "some-env\nsome-other-env"
		Expect(ioutil.WriteFile(envStateFilepath, []byte(environments), 0600)).To(Succeed())
		msg, err := command.CFTargetsCommand{}.Run(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal(environments))
	})

	Context("when no targets have been added", func() {
		It("indicates that no targets are available", func() {
			msg, err := command.CFTargetsCommand{}.Run(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(msg).To(Equal("No targets available"))
		})
	})

	Context("when the envs file does not exist", func() {
		BeforeEach(func() {
			os.Remove(envStateFilepath)
		})

		It("outputs that no targets are available", func() {
			msg, err := command.CFTargetsCommand{}.Run(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(msg).To(Equal("No targets available"))
		})
	})
})
