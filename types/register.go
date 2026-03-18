package types

// Register adds a standard custom detector.
func Register(d Detector) {}

// RegisterWeak adds a structural detector that is prone to false positives (e.g., LZMA).
func RegisterWeak(d Detector) {}

// RegisterFallback adds a detector of last resort (e.g., Plain Text).
func RegisterFallback(d Detector) {}

// RegisterSignature adds a standard magic byte signature.
func RegisterSignature(kind KindID, typ TypeID, offset int, magic []byte) {}

// RegisterMaskedSignature adds a magic byte signature with a bitmask.
func RegisterMaskedSignature(kind KindID, typ TypeID, offset int, magic []byte, mask []byte) {}

// RegisterWeakSignature adds a magic byte signature that is prone to false positives.
func RegisterWeakSignature(kind KindID, typ TypeID, offset int, magic []byte) {}
