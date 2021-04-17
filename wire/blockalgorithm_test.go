package wire

import "testing"

func TestBlockAlgorithm(t *testing.T) {
	tests := []struct {
		name   string
		header BlockHeader
		want   Algorithm
	}{
		{
			"Scrypt + Version",
			BlockHeader{
				Version: int32(SCRYPT | (1 << 20)),
			},
			Algorithm(2048),
		},
		{
			"Groestl + Version",
			BlockHeader{
				Version: int32(GROESTL | (1 << 20)),
			},
			Algorithm(4096),
		},
		{
			"X17 + Version",
			BlockHeader{
				Version: int32(X17 | (1 << 20)),
			},
			Algorithm(6144),
		},
		{
			"Blake2s + Version",
			BlockHeader{
				Version: int32(BLAKE | (1 << 20)),
			},
			Algorithm(8192),
		},
		{
			"Lyra2Re + Version",
			BlockHeader{
				Version: int32(LYRA2RE | (1 << 20)),
			},
			Algorithm(20480),
		},
		{
			"Unkown",
			BlockHeader{
				Version: int32(LYRA2RE | (5 << 11)),
			},
			Algorithm(2048), // fallback to Scrypt
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.header.BlockAlgorithm(); got != tt.want {
				t.Errorf("BlockAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
