package types

// Register declares a standard custom detector for the AOT generator.
func Register(d Detector) {}

// RegisterWeak declares an AOT structural detector that is prone to false positives (e.g., LZMA).
func RegisterWeak(d Detector) {}

// RegisterFallback declares an AOT detector of last resort (e.g., Plain Text).
func RegisterFallback(d Detector) {}

// RegisterSignature declares a standard magic byte signature for the AOT generator.
func RegisterSignature(kind KindID, typ TypeID, offset int, magic []byte) {}

// RegisterMaskedSignature declares an AOT magic byte signature with a bitmask.
func RegisterMaskedSignature(kind KindID, typ TypeID, offset int, magic []byte, mask []byte) {}

// RegisterWeakMaskedSignature declares an AOT masked signature that is prone to false positives.
func RegisterWeakMaskedSignature(kind KindID, typ TypeID, offset int, magic []byte, mask []byte) {}

// RegisterWeakSignature declares an AOT magic byte signature that is prone to false positives.
func RegisterWeakSignature(kind KindID, typ TypeID, offset int, magic []byte) {}
