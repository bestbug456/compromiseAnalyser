package compromiseAnalyser

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os/exec"
)

// CheckEvairomentVariable retrive all
// the envarioment variable and create
// hash from them.
// If something went wrong an error is return.
func CheckEvairomentVariable() ([]byte, error) {
	cmd := exec.Command("printenv")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	data := sha256.New().Sum(out.Bytes())
	return data, nil
}

// CheckPackageInstalled retrive all
// the pakcage install on the docker
// container and create an hash from
// them. NOTE: this command works
// only on alpine .
// If something went wrong an error is return.
func CheckPackageInstalled(system string) ([]byte, error) {
	var cmd *exec.Cmd
	switch system {
	case "alpine":
		cmd = exec.Command("apk info")
	// case "osx":
	// 	cmd = exec.Command("pkgutil --pkgs")
	default:
		return nil, fmt.Errorf("system not supported.")
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	data := sha256.New().Sum(out.Bytes())
	return data, nil
}

// CheckListOfUsers retrive the list
// of users and create an hash from them.
// NOTE: this command works only on alpine.
// If something went wrong an error is return.
func CheckListOfUsers(system string) ([]byte, error) {
	var cmd *exec.Cmd
	switch system {
	case "alpine":
		cmd = exec.Command("cut -d: -f1 /etc/passwd")
	// case "osx":
	// 	cmd = exec.Command("dscl . list /Users | grep -v '_'")
	default:
		return nil, fmt.Errorf("system not supported.")
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	fmt.Println(out.String())
	data := sha256.New().Sum(out.Bytes())
	return data, nil
}

// MakeAllCheck make all the check previusly
// show, concat all the result and create a
// single hash. If something went wrong an
// error is return.
func MakeAllCheck(system string) ([]byte, error) {
	data, err := CheckEvairomentVariable()
	if err != nil {
		return nil, err
	}
	finalString := hex.EncodeToString(data)

	data, err = CheckPackageInstalled(system)
	if err != nil {
		return nil, err
	}
	finalString += hex.EncodeToString(data)
	data, err = CheckListOfUsers(system)
	if err != nil {
		return nil, err
	}
	finalString += hex.EncodeToString(data)
	return sha256.New().Sum([]byte(finalString)), nil
}
