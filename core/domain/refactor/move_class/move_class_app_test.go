package move_class

import (
	"fmt"
	. "github.com/onsi/gomega"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMoveClassApp(t *testing.T) {
	g := NewGomegaWithT(t)

	config := filepath.FromSlash("../../../../_fixtures/refactor/move.config")
	path := filepath.FromSlash("../../../../_fixtures/refactor")

	absPath, _ := filepath.Abs(path)
	app := NewMoveClassApp(config, filepath.FromSlash(absPath))
	app.Analysis()
	app.Refactoring()

	cmd := exec.Command("pwd")
	output, _ := cmd.CombinedOutput()
	fmt.Println("////////////////////////")
	fmt.Println(string(output))

	// todo: fix in CI
	stat, _ := os.Stat(filepath.FromSlash(absPath + "/move/b/ImportForB.java"))
	g.Expect(stat.Name()).To(Equal("ImportForB.java"))

	g.Expect(true).To(Equal(true))
}