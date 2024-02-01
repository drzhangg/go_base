package test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("checking tdengine inspector", func() {
	//tc := TdengineConfig{
	//	Host:     "10.21.137.163:6041",
	//	Username: "root",
	//	Password: "taosdata",
	//	Data:     nil,
	//}
	//
	//BeforeEach(func() {
	//	By("当测试不通过时，打印 【BeforeEach】")
	//})

	It("check result", func() {
		tc := TdengineConfig{
			Host:     "10.21.137.163:6041",
			Username: "root",
			Password: "taosdata",
			Data:     nil,
		}

		Expect(tc.Inspect(context.Background())).To(Equal(true))
	})

	//Describe("1", func() {
	//	Context("2", func() {
	//		It("3", func() {
	//			Expect(tc.Inspect(context.Background())).To(Equal(true))
	//		})
	//	})
	//})
})
