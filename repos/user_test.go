package repos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/romanceresnak/go-grpc/types"
)

var _ = Describe("UsersRepo", func() {

	var (
		err  error
		user *types.User

		//user will be store correct without error
		setupData = func() {
			user, err = types.NewUser(&types.TempUser{
				FirstName:       "Roman",
				LastName:        "Doe",
				Email:           "roman.doe@yahoo.com",
				Password:        "1234",
				ConfirmPassword: "1234",
			})
			Ω(err).To(BeNil())
		}
	)

	BeforeEach(func() {
		clearDatabase()
		setupData()
	})

	//we expect error
	//Describe test
	Describe("Create", func() {
		Context("Failure", func() {
			It("should fail with a nil user", func() {
				err = gr.Users().Create(nil)
				Ω(err).NotTo(BeNil()) //we want error
				Ω(err.Error()).To(Equal("validator: (nil *types.User)"))
			})
		})
	})
})
