package reggiexpress
import (
  "fmt"
)
func main() {
  fg := NewFlowGraph();
  fg.Build("asdf");
  fmt.Print(fg.Process("asdf"))
}
