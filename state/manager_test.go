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

		Expect(manager.SaveEnv(someEnv)).To(Succeed())
		Expect(manager.SaveEnv(someOtherEnv)).To(Succeed())

		envs, err = manager.GetEnvs()
		Expect(err).NotTo(HaveOccurred())
		Expect(envs).To(Equal([]string{someEnv, someOtherEnv}))
	})

	It("does not allow the same environment to be added twice", func() {
		Expect(manager.SaveEnv(someEnv)).To(Succeed())
		Expect(manager.SaveEnv(someEnv)).To(MatchError("target already exists"))

		envs, err := manager.GetEnvs()
		Expect(err).NotTo(HaveOccurred())
		Expect(envs).To(Equal([]string{someEnv}))
	})

	Context("when the state file does not exist", func() {
		BeforeEach(func() {
			os.Remove(path)
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
				Expect(manager.SaveEnv(someEnv)).To(Succeed())
				Expect(path).To(BeAnExistingFile())
			})
		})
	})
})
