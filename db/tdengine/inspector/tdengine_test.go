package inspector_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"inspector"
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
		tc := inspector.TdengineConfig{
			Host:     "10.21.137.163:6041",
			Username: "root",
			Password: "taosdata",
			Data:     nil,
		}

		result,err := tc.Inspect(context.Background())

		Expect(result).To(Equal(true))
		Expect(err).To(Equal("Ping tdengine success!"))
	})

})

