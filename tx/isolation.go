package tx

type IsolationLevel byte

const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)

type IsolationLevelOption interface {
	WithLevel(level IsolationLevel) IsolationLevelOption
	GetLevel() IsolationLevel

	WithReadOnly(readOnly bool) IsolationLevelOption
	IsReadOnly() bool
}

type isolationLevel struct {
	level      IsolationLevel
	isReadOnly bool
}

func NewIsolationLevel() IsolationLevelOption {
	return new(isolationLevel)
}

func (m *isolationLevel) WithLevel(level IsolationLevel) IsolationLevelOption {
	m.level = level
	return m
}

func (m *isolationLevel) GetLevel() IsolationLevel {
	return m.level
}

func (m *isolationLevel) WithReadOnly(readOnly bool) IsolationLevelOption {
	m.isReadOnly = readOnly
	return m
}

func (m *isolationLevel) IsReadOnly() bool {
	return m.isReadOnly
}
