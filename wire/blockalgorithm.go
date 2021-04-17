package wire

// AlgorithmBitMask to extract the algo id from the actual version number
const AlgorithmBitMask int32 = (15 << 11)

// OldAlgorithmBitMask to extract the algo id from the actual version number
const OldAlgorithmBitMask int32 = (10 << 11)

// Algorithm represents all available algorithm signatures within a block version
type Algorithm int32

// represents all algorithms possible with the current core node
const (
	SCRYPT  Algorithm = (1 << 11)
	GROESTL Algorithm = (2 << 11)
	X17     Algorithm = (3 << 11)
	BLAKE   Algorithm = (4 << 11)
	LYRA2RE Algorithm = (10 << 11)
)

// BlockAlgorithm returns the block algorithm that has been within the blockheader
func (h *BlockHeader) BlockAlgorithm() Algorithm {
	var extractedAlgorithm Algorithm = Algorithm(h.Version & AlgorithmBitMask)

	switch extractedAlgorithm {
	case SCRYPT:
		return SCRYPT
	case GROESTL:
		return GROESTL
	case X17:
		return X17
	case BLAKE:
		return BLAKE
	case LYRA2RE:
		return LYRA2RE
	default:
		// Everything that possibly moves out of range will identified as scrypt.
		return SCRYPT
	}
}

// GetAlgorithmString returns the block algorithm string that has been within the blockheader
func (h *BlockHeader) GetAlgorithmString() string {
	var extractedAlgorithm Algorithm = Algorithm(h.Version & AlgorithmBitMask)

	switch extractedAlgorithm {
	case SCRYPT:
		return "scrypt"
	case GROESTL:
		return "Groestl"
	case X17:
		return "x17"
	case BLAKE:
		return "blake2s"
	case LYRA2RE:
		return "lyra2re"
	default:
		// Everything that possibly moves out of range will identified as scrypt.
		return "scrypt"
	}
}
