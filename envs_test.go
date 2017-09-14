package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Envs", func() {
	var (
		pathToCLI        string
		envStateFilepath string
	)

	BeforeEach(func() {
		envStateFile, err := ioutil.TempFile("", "env-state")
		Expect(err).NotTo(HaveOccurred())
		envStateFilepath = envStateFile.Name()
		pathToCLI, err = Build("github.com/juliencherry/envs")
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("ENV_STATE_FILEPATH", envStateFilepath)
	})

	AfterEach(func() {
		os.Unsetenv("ENV_STATE_FILEPATH")
		os.Remove(envStateFilepath)
		CleanupBuildArtifacts()
	})

	It("saves a targeted environment", func() {
		session := runCommand(pathToCLI, "cf-targets")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("No targets available"))

		session = runCommand(pathToCLI, "cf-add-target", "some-env")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say(`Added target "some-env"`))

		session = runCommand(pathToCLI, "cf-targets")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("some-env"))

		session = runCommand(pathToCLI, "cf-target")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("No environment targeted"))
	})

	Context("when an invalid command is used", func() {
		It("indicates that it was an invalid command", func() {
			session := runCommand(pathToCLI, "an-invalid-command")
			Eventually(session).Should(Exit(1))
			Expect(session).To(Say("Invalid command"))
		})
	})

	Context("when no command is used", func() {
		It("prints an error message", func() {
			session := runCommand(pathToCLI)
			Eventually(session).Should(Exit(1))
			Expect(session).To(Say("Invalid command"))
		})
	})
})

func runCommand(bin string, args ...string) *Session {
	command := exec.Command(bin, args...)
	session, err := Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
