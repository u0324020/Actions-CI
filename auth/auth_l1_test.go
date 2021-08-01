package auth

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)
/*
#1 [gomock]
Declare Controller -> Define mocked method and testing expect, return, times...
-> Call mocked function -> Finish the Controller
*/
func TestCertificate_InputAuthorizedUsernamePassword_ReturnNil(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAuth := NewMockCertificate(mockCtrl)
	mockAuth.EXPECT().Auth("jane","1005").Return(nil).Times(1)

	if res := mockAuth.Auth("jane", "1005"); res != nil{
		t.Error("Input username and password are authorized but return error")
	}
}

func TestCertificate_InputUnauthorizedUsernamePassword_ReturnError(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAuth := NewMockCertificate(mockCtrl)
	mockAuth.EXPECT().Auth("123","123").Return(errors.New("user is not exist")).Times(2)

	log.Println(mockAuth.Auth("123", "123"))
	if res := mockAuth.Auth("123", "123"); res == nil{
		t.Error("Input username and password are authorized but return error")
	}
}

/*
#2 [test suits]
SetupSuite-> (SetupTest->TestXxx->TearDownTest)*N times -> TearDownSuite
*/

type testCertificateL0 struct{
	suite.Suite
}

func TestCertificate_L0(t *testing.T){
	suite.Run(t, new(testCertificateL0))
}

func (suite * testCertificateL0) SetupSuite(){
	fmt.Println("# Not necessary")
	fmt.Println("In SetupSuite Session, it just do once at the beginning of the suite.")
}

func (suite * testCertificateL0) TearDownSuite(){
	fmt.Println("# Not necessary")
	fmt.Println("In TearDownSuite Session, it just do once at the ending of the suite.")
}

func (suite * testCertificateL0) SetupTest(){
	fmt.Println("# Not necessary")
	fmt.Println("In SetupTest Session, it do once at the beginning of every suite.")
}

func (suite * testCertificateL0) TearDownTest(){
	fmt.Println("# Not necessary")
	fmt.Println("In TearDownTest Session, it do once at the ending of the suite.")
}

func (suite * testCertificateL0) TestCertificate_WithoutArgument_ReturnErrorMessage(){
	//# Necessary
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	mockAuth := NewMockCertificate(mockCtrl)
	Input_AC := ""
	Input_PW := ""
	Want := errors.New("user is not exist")
	mockAuth.EXPECT().Auth(Input_AC,Input_PW).Return(Want).Times(1)
	result := mockAuth.Auth(Input_AC,Input_PW)
	suite.Equal(Want, result)
}

func (suite * testCertificateL0) TestCertificate_WithoutPassword_ReturnErrorMessage(){
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	mockAuth := NewMockCertificate(mockCtrl)
	Input_AC := "123"
	Input_PW := ""
	Want := errors.New("user is not exist")
	mockAuth.EXPECT().Auth(Input_AC,Input_PW).Return(Want).Times(1)
	result := mockAuth.Auth(Input_AC,Input_PW)
	suite.Equal(Want, result)
}

func (suite * testCertificateL0) TestCertificate_InputTestUser_ReturnNil(){
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	mockAuth := NewMockCertificate(mockCtrl)
	Input_AC := "test"
	Input_PW := "test"
	Want := errors.New("user is not exist")
	mockAuth.EXPECT().Auth(Input_AC,Input_PW).Return(Want).Times(1)
	result := mockAuth.Auth(Input_AC,Input_PW)
	suite.Equal(Want, result)
}