package utils

import "testing"

func Test_IsPrime(t *testing.T) {
	//Arrange
	no := 37
	expectedResult := true

	//Act
	result := IsPrime(no)

	//Assert
	if expectedResult != result {
		t.Errorf("Expected %v but got %v\n", expectedResult, result)
	}
}

/* Data driven tests */
type testCase struct {
	name     string
	input    int
	expected bool
	actual   bool
}

var testCases = []testCase{
	testCase{name: "Testing_Prime:31", input: -1, expected: false},
	testCase{name: "Testing_Prime:31", input: 2, expected: true},
	testCase{name: "Testing_Prime:31", input: 31, expected: true},
	testCase{name: "Testing_Prime:61", input: 61, expected: true},
	testCase{name: "Testing_Prime:71", input: 71, expected: true},
	testCase{name: "Testing_Prime:73", input: 73, expected: true},
	testCase{name: "Testing_Prime:64", input: 64, expected: false},
	testCase{name: "Testing_Prime:83", input: 83, expected: true},
	testCase{name: "Testing_Prime:12", input: 12, expected: false},
}

func Test_IsPrime_Nos(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime(testCase.input)
			if testCase.expected != testCase.actual {
				t.Errorf("Expected %v but got %v\n", testCase.expected, testCase.actual)
			}
		})
	}
}
