package cf_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/juliencherry/envs/cf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("CLIManager", func() {

	var (
		cfAPI      string
		cfUsername string
		cfPassword string
		cfHome     string
		tempCFHome string
	)

	BeforeEach(func() {
		cfAPI = os.Getenv("CF_API")
		Expect(cfAPI).NotTo(BeEmpty(), "Must set CF_API environment variable")
		cfUsername = os.Getenv("CF_USERNAME")
		Expect(cfUsername).NotTo(BeEmpty(), "Must set CF_USERNAME environment variable")
		cfPassword = os.Getenv("CF_PASSWORD")
		Expect(cfPassword).NotTo(BeEmpty(), "Must set CF_PASSWORD environment variable")

		cfHome = os.Getenv("CF_HOME")
		var err error
		tempCFHome, err = ioutil.TempDir("", "envs-cf-home-")
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("CF_HOME", tempCFHome)
	})

	AfterEach(func() {
		os.Setenv("CF_HOME", cfHome)
		os.RemoveAll(tempCFHome)
	})

	It("runs `cf-login` with the given arguments", func() {
		cliManager := cf.CLIManager{}
		err := cliManager.Login(cfAPI, cfUsername, cfPassword, true)
		Expect(err).NotTo(HaveOccurred())

		command := exec.Command("cf", "target")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(Exit(0))
		Expect(session).To(Say(cfAPI))
		Expect(session).To(Say(cfUsername))
	})

	It("returns the output from `cf-login` if it errors", func() {
		cliManager := cf.CLIManager{}
		err := cliManager.Login("bad-api", "bad-username", "bad-password", true)
		Expect(err).To(MatchError(And(HavePrefix("failed to log in: "))))
	})
})
