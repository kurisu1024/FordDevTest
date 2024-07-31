package dir_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"FordDevTest/dir"
	"github.com/stretchr/testify/suite"
)

const (
	dirRoot     = "root"
	dirSub      = "root/sub"
	fileOne     = "fileOne.txt"
	fileTwo     = "filetwo.txt"
	fileOneText = "This is file one text."
	fileTwoText = "This is "
)

func TestDirSuite(t *testing.T) {
	suite.Run(t, new(DirTestSuite))
}

type DirTestSuite struct {
	suite.Suite
	rootSize int64
	subSize  int64
}

// Create the dir and files required to setup the test.
func (s *DirTestSuite) SetupTest() {
	var err error
	var f *os.File
	defer func(file *os.File) {
		if file != nil {
			file.Close()
		}
	}(f)
	s.Require().NoError(os.Mkdir(dirRoot, os.ModeDir))
	s.Require().NoError(os.Mkdir(dirSub, os.ModeDir))

	f, err = os.Create(fmt.Sprintf("%s/%s", dirRoot, fileOne))
	s.Require().NoError(err)

	s.rootSize, err = io.Copy(f, strings.NewReader(fileOneText))
	s.Require().NoError(err)
	s.Require().NoError(f.Close())

	f, err = os.Create(fmt.Sprintf("%s/%s", dirSub, fileTwo))
	s.Require().NoError(err)

	s.subSize, err = io.Copy(f, strings.NewReader(fileTwoText))
	s.Require().NoError(err)
	s.Require().NoError(f.Close())
	f = nil

	s.rootSize += s.subSize

}

// Clean up the test env and delete the dir and files created.
func (s *DirTestSuite) TearDownTest() {
	s.Require().NoError(os.RemoveAll(dirRoot))
}

// TODO(cclark): More tests need to be added to cover more test cases. This
//
//	tests the simple straight forward test case.
func (s *DirTestSuite) TestDirSize() {
	got, err := dir.DirSize(dirRoot)
	s.Require().NoError(err)
	s.Require().Equal(s.rootSize, got)

	got, err = dir.DirSize(dirSub)
	s.Require().NoError(err)
	s.Require().Equal(s.subSize, got)
}
