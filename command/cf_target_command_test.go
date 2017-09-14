package command_test

import (
	"github.com/juliencherry/envs/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFTargetCommand", func() {

	It("indicates that no environments are targeted", func() {
		cmd := command.CFTargetCommand{}
		msg, err := cmd.Run(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal("No environment targeted"))
	})
})
