package gotil

import (
	"fmt"
	"testing"
)

func TestFloat64ToString(t *testing.T) {

	floatStr := Float64ToString(0.0823, 2)
	fmt.Println(floatStr)
}
