package state_test

import (
	"io/ioutil"
	"os"

	"github.com/juliencherry/envs/state"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {

	var (
		path         string
		manager      state.Manager
		someEnv      = "some-env"
		someOtherEnv = "some-other-env"
		someAPI      = "https://example.com/"
		someUsername = "some-username"
		somePassword = "some-password"
	)

	BeforeEach(func() {
		tempFile, err := ioutil.TempFile("", "state-manager-")
		Expect(err).NotTo(HaveOccurred())

		path = tempFile.Name()
		manager = state.Manager{Path: path}
	})

	AfterEach(func() {
		os.Remove(path)
	})

	It("saves and gets environments", func() {
		envs, err := manager.GetEnvs()
		Expect(err).NotTo(HaveOccurred())
		Expect(envs).To(BeEmpty())

		Expect(manager.SaveEnv(someEnv, someAPI, someUsername, somePassword, true)).To(Succeed())
		Expect(manager.SaveEnv(someOtherEnv, someAPI, someUsername, somePassword, true)).To(Succeed())

		envs, err = manager.GetEnvs()
		Expect(err).NotTo(HaveOccurred())
		Expect(envs).To(ConsistOf(someEnv, someOtherEnv))

		api, username, password, skipSSLValidation, err := manager.GetEnvDetails(someEnv)
		Expect(err).NotTo(HaveOccurred())
		Expect(api).To(Equal(someAPI))
		Expect(username).To(Equal(someUsername))
		Expect(password).To(Equal(somePassword))
		Expect(skipSSLValidation).To(Equal(true))
	})

	It("does not allow the same environment to be added twice", func() {
		Expect(manager.SaveEnv(someEnv, someAPI, someUsername, somePassword, false)).To(Succeed())
		Expect(manager.SaveEnv(someEnv, someAPI, someUsername, somePassword, false)).To(MatchError("target already exists"))

		envs, err := manager.GetEnvs()
		Expect(err).NotTo(HaveOccurred())
		Expect(envs).To(Equal([]string{someEnv}))
	})

	It("errors when trying to get an environment that has not been saved", func() {
		_, _, _, _, err := manager.GetEnvDetails(someEnv)
		Expect(err).To(MatchError("environment not found"))
	})

	Context("when the state file does not exist", func() {
		BeforeEach(func() {
			Expect(os.Remove(path)).To(Succeed())
		})

		Describe("#GetEnvs", func() {
			It("returns no environments", func() {
				envs, err := manager.GetEnvs()
				Expect(err).NotTo(HaveOccurred())
				Expect(envs).To(BeEmpty())
				Expect(path).NotTo(BeAnExistingFile())
			})
		})

		Describe("#SaveEnv", func() {
			It("creates the state file", func() {
				Expect(manager.SaveEnv(someEnv, someAPI, someUsername, somePassword, false)).To(Succeed())
				Expect(path).To(BeAnExistingFile())
			})
		})
	})
})
