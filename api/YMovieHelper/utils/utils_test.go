package utils_test

import (
	"strings"
	"testing"

	"github.com/itkmaingit/YMovieHelper/utils"
)

func TestModifyQuery(t *testing.T) {
	var testQuery strings.Builder
	testQuery.WriteString("UPDATE table SET ")
	testQuery.WriteString("name = ?, ")
	testQuery.WriteString("id = ?, ")

	result := utils.ModifyQuery(testQuery)
	expect := "UPDATE table SET name = ?, id = ? "

	if result != expect {
		t.Errorf("Expected %s, but got %s", expect, result)
	}
}
