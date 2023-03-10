package types

const (
	ModuleName   = "coslms"
	StoreKey     = ModuleName
	QuerierRoute = ModuleName
)

var (
	AdminKey        = []byte{0x01}
	StudentKey      = []byte{0x02}
	ApplyLeavesKey  = []byte{0x03}
	AcceptLeavesKey = []byte{0x04}
	sequenceKey     = []byte{0x05}
)

func AdminStoreKey(address string) []byte {
	key := make([]byte, len(AdminKey)+len(address))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], address)
	return key
}

func StudentStoreKey(studentID string) []byte {
	key := make([]byte, len(StudentKey)+len(studentID))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentID)
	return key
}

func AcceptLeavesStoreKey(admin string, leaveID string) []byte {
	key := make([]byte, len(AcceptLeavesKey)+len(admin)+len(sequenceKey)+len(leaveID))
	copy(key, AcceptLeavesKey)
	copy(key[len(AcceptLeavesKey):], admin)
	copy(key, sequenceKey)
	copy(key[len(sequenceKey):], leaveID)
	return key
}
func ApplyLeavesStoreKey(admin string, leaveID string) []byte {
	key := make([]byte, len(ApplyLeavesKey)+len(admin)+len(sequenceKey)+len(leaveID))
	copy(key, ApplyLeavesKey)
	copy(key[len(ApplyLeavesKey):], admin)
	copy(key, sequenceKey)
	copy(key[len(sequenceKey):], leaveID)
	return key
}
