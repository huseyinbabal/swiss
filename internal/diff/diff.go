package diff

import "strings"

// Lines computes a line-based diff using a standard LCS dynamic-programming
// algorithm. Unchanged lines are prefixed with "  ", deletions with "- ",
// and additions with "+ ".
func Lines(a, b string) (string, error) {
	aLines := strings.Split(a, "\n")
	bLines := strings.Split(b, "\n")

	n := len(aLines)
	m := len(bLines)

	// LCS length table.
	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if aLines[i] == bLines[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else if lcs[i+1][j] >= lcs[i][j+1] {
				lcs[i][j] = lcs[i+1][j]
			} else {
				lcs[i][j] = lcs[i][j+1]
			}
		}
	}

	var sb strings.Builder
	i, j := 0, 0
	for i < n && j < m {
		if aLines[i] == bLines[j] {
			sb.WriteString("  " + aLines[i] + "\n")
			i++
			j++
		} else if lcs[i+1][j] >= lcs[i][j+1] {
			sb.WriteString("- " + aLines[i] + "\n")
			i++
		} else {
			sb.WriteString("+ " + bLines[j] + "\n")
			j++
		}
	}
	for i < n {
		sb.WriteString("- " + aLines[i] + "\n")
		i++
	}
	for j < m {
		sb.WriteString("+ " + bLines[j] + "\n")
		j++
	}

	return strings.TrimRight(sb.String(), "\n"), nil
}
