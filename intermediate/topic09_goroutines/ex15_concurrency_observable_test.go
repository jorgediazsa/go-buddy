package topic09_goroutines

import (
	"bytes"
	"strings"
	"testing"
)

func TestEX09_Ex15_RunInterleaved(t *testing.T) {
	var buf bytes.Buffer
	iterations := 100

	RunInterleaved(&buf, iterations)

	output := buf.String()
	if len(output) != iterations*2 {
		t.Errorf("Expected output length %d, got %d", iterations*2, len(output))
	}

	countA := strings.Count(output, "A")
	countB := strings.Count(output, "B")

	if countA != iterations || countB != iterations {
		t.Errorf("Expected %d A's and %d B's, got %d and %d", iterations, iterations, countA, countB)
	}

	// Heuristic: check if there's at least SOME interleaving
	// "AAA...BBB..." would have only one transition.
	transitions := 0
	for i := 0; i < len(output)-1; i++ {
		if output[i] != output[i+1] {
			transitions++
		}
	}

	if transitions < 2 {
		t.Logf("Warning: Very low interleaving detected (%d transitions). This can happen but might indicate lack of yielding.", transitions)
	}
}
