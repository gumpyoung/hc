package pair

type Container interface {
    SetByte(key, value byte)
    SetBytes(key byte, value []byte)
    SetString(key byte, value string)
    
    GetByte(key byte) byte
    GetBytes(key byte) []byte
    GetString(key byte) string
}