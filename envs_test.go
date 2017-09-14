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
		pathToCLI, err = Build("github.com/juliencherry/envs", "-ldflags", "-X main.envStateFilepath="+envStateFilepath)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.Remove(envStateFilepath)
		CleanupBuildArtifacts()
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

	Describe("cf-targets", func() {
		Context("when no CF targets have been added", func() {
			It("outputs that no targets are available", func() {
				session := runCommand(pathToCLI, "cf-targets")
				Eventually(session).Should(Exit(0))
				Expect(session).To(Say("No targets available"))
			})
		})

		Context("when the envs file does not exist", func() {
			BeforeEach(func() {
				os.Remove(envStateFilepath)
			})

			It("outputs that no targets are available", func() {
				session := runCommand(pathToCLI, "cf-targets")
				Eventually(session).Should(Exit(0))
				Expect(session).To(Say("No targets available"))
			})
		})

		Context("when the command errors", func() {
			BeforeEach(func() {
				os.Remove(envStateFilepath)
				os.Mkdir(envStateFilepath, 0700)
			})

			AfterEach(func() {
				os.RemoveAll(envStateFilepath)
			})

			It("prints an error message", func() {
				session := runCommand(pathToCLI, "cf-targets")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say("Failed to list targets: "))
			})
		})
	})

	Describe("cf-add-target", func() {
		It("allows multiple targets to be added", func() {
			someEnv := "some-env"
			session := runCommand(pathToCLI, "cf-add-target", someEnv)
			Eventually(session).Should(Exit(0))

			someOtherEnv := "some-other-env"
			session = runCommand(pathToCLI, "cf-add-target", someOtherEnv)
			Eventually(session).Should(Exit(0))

			session = runCommand(pathToCLI, "cf-targets")
			Eventually(session).Should(Exit(0))
			Expect(session).To(Say(someEnv + "\n"))
			Expect(session).To(Say(someOtherEnv + "\n"))
		})

		It("does not add duplicate targets", func() {
			someEnv := "some-env"
			session := runCommand(pathToCLI, "cf-add-target", someEnv)
			Eventually(session).Should(Exit(0))

			session = runCommand(pathToCLI, "cf-add-target", someEnv)
			Eventually(session).Should(Exit(1))
			Expect(session).To(Say("target “" + someEnv + "” already exists"))

			session = runCommand(pathToCLI, "cf-targets")
			Eventually(session).Should(Exit(0))
			Expect(string(session.Out.Contents())).To(Equal(someEnv + "\n"))
		})

		Context("when no target is given", func() {
			It("indicates that a target must be specified", func() {
				session := runCommand(pathToCLI, "cf-add-target")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say("Must specify target to add"))
			})
		})

		Context("when the command errors", func() {
			BeforeEach(func() {
				os.Remove(envStateFilepath)
				os.Mkdir(envStateFilepath, 0700)
			})

			AfterEach(func() {
				os.RemoveAll(envStateFilepath)
			})

			It("prints an error message", func() {
				session := runCommand(pathToCLI, "cf-add-target", "some-env")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say("Failed to add the target: "))
			})
		})
	})
})

func runCommand(bin string, args ...string) *Session {
	command := exec.Command(bin, args...)
	session, err := Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
