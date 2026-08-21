package custom

import (
	"github.com/coalaura/wtf/types"
)

func DetectDeepAnalysis(b types.Buffer) *types.Metadata {
	if b.Len() < 64 {
		return nil
	}

	result := CalculateEntropy(b)
	if result == nil {
		return nil
	}

	// Don't analyze small files for encryption
	if b.Len() < 512 {
		return nil
	}

	// Check for common binary patterns that might be confused with encryption
	if isKnownBinaryPattern(b) {
		return nil
	}

	if result.IsEncrypted {
		return &types.Metadata{
			Kind:       types.KindHighEntropyData,
			Type:       types.TypeEncrypted,
			Confidence: types.ConfidenceLow,
		}
	}

	if result.IsHighEntropy {
		// Could be compressed, but we didn't match any compression signatures
		return &types.Metadata{
			Kind:       types.KindHighEntropyData,
			Type:       types.TypeHighEntropy,
			Confidence: types.ConfidenceLow,
		}
	}

	// Check for binary data that's not high entropy
	if result.ZeroBytes > 0 || result.PrintableBytes*100/min(b.Len(), 65536) < 50 {
		return &types.Metadata{
			Kind:       types.KindHighEntropyData,
			Type:       types.TypeBinary,
			Confidence: types.ConfidenceLow,
		}
	}

	return nil
}

func isKnownBinaryPattern(b types.Buffer) bool {
	// Skip if it looks like it might be a partial header of a known format
	// These are patterns that might appear at various offsets

	// Check for repeated patterns (like padding)
	if b.Len() >= 64 {
		// Check if first 32 bytes are all the same
		allSame := true
		first := b[0]

		for i := 1; i < 32 && i < b.Len(); i++ {
			if b[i] != first {
				allSame = false
				break
			}
		}

		if allSame {
			return true
		}
	}

	return false
}
