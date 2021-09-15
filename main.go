package reggiexpress
import (
  "fmt"
)
func main() {
  fg := newFlowGraph();
  fmt.Print(fg.process("asdf"))
}
