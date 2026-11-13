package python

import (
	"os"
	"os/exec"
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
	return content[:len(content)-2], err
}
