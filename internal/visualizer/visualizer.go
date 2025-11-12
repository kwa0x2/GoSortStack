package visualizer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/kwa0x2/GoSortStack/internal/stack"
)

type Step struct {
	Operation string `json:"operation"`
	StackA    []int  `json:"stackA"`
	StackB    []int  `json:"stackB"`
}

type VisualizationData struct {
	Initial []int  `json:"initial"`
	Steps   []Step `json:"steps"`
}

var (
	visualizeMode bool
	steps         []Step
	initialStack  []int
)

func Enable() {
	visualizeMode = true
	steps = []Step{}
}

func Disable() {
	visualizeMode = false
}

func IsEnabled() bool {
	return visualizeMode
}

func stackToSlice(head *stack.Node) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func ShowStep(ops *stack.Operations, operation string) {
	if !visualizeMode {
		return
	}

	step := Step{
		Operation: operation,
		StackA:    stackToSlice(ops.StackA),
		StackB:    stackToSlice(ops.StackB),
	}

	steps = append(steps, step)
}

func ShowInitial(ops *stack.Operations) {
	if !visualizeMode {
		return
	}

	initialStack = stackToSlice(ops.StackA)
}

func ShowSummary() {
	if !visualizeMode {
		return
	}

	data := VisualizationData{
		Initial: initialStack,
		Steps:   steps,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error generating visualization data:", err)
		return
	}

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	fmt.Println(exeDir)
	templatePath := filepath.Join("../public/", "visualizer.html")
	htmlContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println("\n⚠ visualizer.html not found. Please create it in the same directory in public folder.")
		return
	}

	htmlWithData := string(htmlContent)

	embedScript := fmt.Sprintf("<script>const EMBEDDED_DATA = %s;</script>", string(jsonData))
	htmlWithData = embedScript + htmlWithData

	outputPath := filepath.Join("../public/", "visualization.html")
	err = os.WriteFile(outputPath, []byte(htmlWithData), 0644)
	if err != nil {
		fmt.Println("Error writing visualization file:", err)
		return
	}

	fmt.Printf("\n✓ Visualization file created: %s\n", outputPath)
	fmt.Printf("✓ Total operations: %d\n", len(steps))

	openBrowser(outputPath)
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // linux
		cmd = "xdg-open"
		args = []string{url}
	}

	err := exec.Command(cmd, args...).Start()
	if err == nil {
		fmt.Println("✓ Opening visualization in browser...")
	}
}

func GetOperationCount() int {
	return len(steps)
}
