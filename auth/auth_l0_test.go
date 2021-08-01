package auth

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)
/*
#1 [writing table]
1. Define test struct: input ac and pw, should return expected
2. Create an instance of struct and build the test contents
3. Using loop to test all of the test contents
*/
type users struct {
	ac string
	pw string
	expected error
}

var addUser = []users{
	{"jane","1005",nil},
	{"test","test",nil},
}

func TestAuth_InputExistUsernamePassword_ReturnTrue(t *testing.T){
	for _, test := range addUser {
		if output := Auth(test.ac, test.pw); output != test.expected{
			t.Errorf("Input username %s and password %s return %q, not equal to expect %v", test.ac, test.pw, output, test.expected)
		}
	}
}
/*
#2 [if]
*/
func TestAuth_InputAuthorizedUsernamePassword_ReturnTrue(t *testing.T) {
	InputUsername := "jane"
	InputPassword := "1005"
	if Auth(InputUsername, InputPassword) != nil{
		t.Errorf("Input username %s and password %s to Auth, not return true.", InputUsername, InputPassword)
	}
}

func TestCheckPassword_InputInequalityPassword_ReturnError(t *testing.T){
	if CheckPassword("123", "1234") == nil {
		t.Error("Input password not equal")
	}
}

func TestAuth_InputUnauthorizedUsernamePassword_ReturnError(t *testing.T) {
	if err := Auth("123", "123"); err == nil{
		t.Error("Input username and password are unauthorized but return true")
	}}
/*
#3 [assert]
the Equal should input three parameters
*/
func TestCheckUserIsExist_InputExistUser_ReturnTrue(t *testing.T) {
	assert.Equal(t, true, CheckUserIsExist("jane"))
}
/*
#4 [read file]
*/
func TestReadFile_example(t *testing.T){
	data, err := ioutil.ReadFile("../../test/test.data")
	if err != nil{
		t.Fatal("Could not open file")
	}
	if string(data) != "Hello World"{
		t.Fatal("String contents do not match expected")
	}
}
/*
#5 [benchmark]
command: go test -bench BenchmarkAuth
*/
func BenchmarkAuth(b *testing.B) {
	for i := 0; i< b.N; i++{
		Auth("jane","1005")
	}
}

