package normal

// 0000 必输
// 0000 0000 必输，因为8块石头，无论开局拿x块，都可以被对手通过拿4-x块达到4块必输
// 同理 0000 0000 0000 12块也是必输
func CanWinNim(n int) bool {
    return n % 4 != 0
}