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

		cfAPI      string
		cfUsername string
		cfPassword string
		cfHome     string
		tempCFHome string
	)

	BeforeEach(func() {
		envStateFile, err := ioutil.TempFile("", "env-state")
		Expect(err).NotTo(HaveOccurred())
		envStateFilepath = envStateFile.Name()
		pathToCLI, err = Build("github.com/juliencherry/envs", "-ldflags", "-X main.envStateFilepath="+envStateFilepath)
		Expect(err).NotTo(HaveOccurred())

		cfAPI = os.Getenv("CF_API")
		Expect(cfAPI).NotTo(BeEmpty(), "Must set CF_API environment variable")
		cfUsername = os.Getenv("CF_USERNAME")
		Expect(cfUsername).NotTo(BeEmpty(), "Must set CF_USERNAME environment variable")
		cfPassword = os.Getenv("CF_PASSWORD")
		Expect(cfPassword).NotTo(BeEmpty(), "Must set CF_PASSWORD environment variable")

		cfHome = os.Getenv("CF_HOME")
		tempCFHome, err = ioutil.TempDir("", "envs-cf-home-")
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("CF_HOME", tempCFHome)
	})

	AfterEach(func() {
		os.Remove(envStateFilepath)
		CleanupBuildArtifacts()
		os.Setenv("CF_HOME", cfHome)
		os.RemoveAll(tempCFHome)
	})

	It("saves a targeted environment", func() {
		session := runCommand(pathToCLI, "cf-targets")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("No targets available"))

		session = runCommand(pathToCLI, "cf-add-target",
			"--name", "some-env",
			"--api", cfAPI,
			"--username", cfUsername,
			"--password", cfPassword,
			"--skip-ssl-validation")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say(`Added target "some-env"`))

		session = runCommand(pathToCLI, "cf-targets")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("some-env"))

		session = runCommand(pathToCLI, "cf-target")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say("No environment targeted"))

		session = runCommand(pathToCLI, "cf-target", "some-env")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say(`Targeted environment "some-env"`))

		session = runCommand("cf", "target")
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say(cfAPI))
		Expect(session).To(Say(cfUsername))
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
