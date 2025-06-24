package encryption

const (
	// Recommended argon2id parameters
	argon2idTime    uint32 = 1
	argon2idMemory  uint32 = 64 * 1024 // 64MB
	argon2idThreads uint8  = 4
	argon2idKeyLen  uint32 = 32

	memoryStr  = "memory"
	timeStr    = "time"
	threadsStr = "threads"

	randomSaltBytes = 16
)
