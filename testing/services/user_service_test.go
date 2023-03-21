package services_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/osalomon89/go-testing/mocks"
	"github.com/osalomon89/go-testing/services"

	"testing"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cart Suite")
}

var _ = Describe("Saving users with first and last name", func() {
	var ctrl *gomock.Controller
	var repositoryMock *mocks.MockUserRepository

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repositoryMock = mocks.NewMockUserRepository(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("When user has both names", func() {
		name := "Oscar"
		lastname := "Salomon"
		age := 18

		repositoryMock.EXPECT().
			SaveUser(name, lastname, age).
			Return(nil).
			Times(1)

		svc := services.NewUserService(repositoryMock)

		err := svc.CreateUser(name, lastname, age)
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("get user info", func() {
	var ctrl *gomock.Controller
	var repositoryMock *mocks.MockUserRepository

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repositoryMock = mocks.NewMockUserRepository(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	DescribeTable("pre validation",
		func(firstName, lastName string, age int, wantedError, repoError error, repoTimes int) {
			repositoryMock.EXPECT().
				SaveUser(firstName, lastName, age).
				Return(repoError).
				Times(repoTimes)

			svc := services.NewUserService(repositoryMock)

			err := svc.CreateUser(firstName, lastName, age)
			if wantedError != nil {
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(wantedError))
				return
			}
			Expect(err).NotTo(HaveOccurred())
		},
		Entry("When user has both names", "Oscar", "Salomon", 33, nil, nil, 1),
		Entry("When user has one name", "Oscar", "", 20, fmt.Errorf("name or lastname could not be empty"), nil, 0),
		Entry("When user has less than 18 years old", "Oscar", "Salomon", 15, fmt.Errorf("age cannot be less than 18"), nil, 0),
		Entry("When repository returns an error", "Oscar", "Salomon", 25, fmt.Errorf("error in repository: %w", errors.New("the repository error")), errors.New("the repository error"), 1),
	)
})
