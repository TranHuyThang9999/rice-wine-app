package apperrors

type ErrTypeRice int8

const (
	ErrConflictTypeName ErrTypeRice = 1 + iota
)

func (s ErrTypeRice) Pointer() *ErrTypeRice {
	return &s
}

func (s *ErrTypeRice) Value() ErrTypeRice {
	if s == nil {
		return 0
	}

	return *s
}
