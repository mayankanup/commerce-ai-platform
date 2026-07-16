package memory

import "math"

func cosineSimilarity(
	a []float32,
	b []float32,
) float64 {

	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	if len(a) != len(b) {
		return 0
	}

	var dot float64
	var normA float64
	var normB float64

	for i := range a {

		av := float64(a[i])
		bv := float64(b[i])

		dot += av * bv

		normA += av * av

		normB += bv * bv
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}
