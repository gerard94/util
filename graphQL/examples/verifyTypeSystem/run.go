package main

import (
	
	E	"github.com/gerard94/util/graphQL/exec"
	F	"path/filepath"
	R	"github.com/gerard94/util/resources"

)

const (
	
	verifyDir = "util/graphQL/examples/verifyTypeSystem"

)

var (
	
	verifyPath = R.FindDir(verifyDir)

)

func main () {
	E.VerifyTypeSystem(F.Join(verifyPath, "emptyDefinitions.txt"))
}

