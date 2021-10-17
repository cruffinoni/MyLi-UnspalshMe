package args

import (
	"os"
	"testing"
)

func checkProgContent(t *testing.T, test ProgArg, ref ProgArg) {
	if test.Query != ref.Query {
		t.Fatalf("invalid query; wanted %v, got %v\n", ref.Query, test.Query)
	}
	if test.Page != ref.Page {
		t.Fatalf("invalid number of page; wanted %v, got %v\n", ref.Page, test.Page)
	}
}

func TestArgParsing(t *testing.T) {
	var fakeProgArgs = []string{
		projectName,
		"dog",
		"18",
	}
	os.Args = fakeProgArgs
	prog, err := New()
	if err != nil {
		t.Fatalf("err is not nil: %v\n", err)
	}
	checkProgContent(t, prog, ProgArg{
		Page:  18,
		Query: "dog",
	})
}

func TestArgOptionalParsing(t *testing.T) {
	var fakeProgArgs = []string{
		projectName,
		"cat",
	}
	os.Args = fakeProgArgs
	prog, err := New()
	if err != nil {
		t.Fatalf("err is not nil: %v\n", err)
	}
	checkProgContent(t, prog, ProgArg{
		Page:  1,
		Query: "cat",
	})
}

func TestInvalidArgNumber(t *testing.T) {
	var fakeProgArgs = []string{
		projectName,
	}
	os.Args = fakeProgArgs
	_, err := New()
	if err == nil {
		t.Fatal("err is nil and should not be\n")
	}
}

func TestParsingError(t *testing.T) {
	var fakeProgArgs = []string{
		projectName,
		"dog",
		"0123456789789456321",
	}
	os.Args = fakeProgArgs
	_, err := New()
	if err == nil {
		t.Fatal("err is nil and should not be\n")
	}
}

func TestHelpMsgArg(t *testing.T) {
	var fakeProgArgs = []string{
		projectName,
		"-h",
	}
	os.Args = fakeProgArgs
	_, err := New()
	if err != HelpMessage {
		t.Fatalf("err should equal to '%v'\n", HelpMessage)
	}
}
