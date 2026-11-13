package python

import (
	"os"
	"os/exec"

	"github.com/wille/osutil"
)

// ExecutePython executes python
func ExecutePython(
	scriptPath string,
) ([]byte, error) {
	scriptData, err := os.ReadFile(
		scriptPath,
	)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(
		"python",
		"-c",
		string(scriptData),
	)
	content, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	osName := osutil.Name
	if osName == "Windows" {
		return content[:len(content)-2], nil
	}
	return content, nil
}
