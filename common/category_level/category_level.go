package common

type categoryLevel int

const (
	DEFAULT categoryLevel = iota
	LEVEL_ONE
	LEVEL_TWO
	LEVEL_THREE
)

func CategoryLevel(i int) categoryLevel {
	if i > int(LEVEL_THREE) || i < int(DEFAULT) {
		return DEFAULT
	}
	return categoryLevel(i)
}

func (i categoryLevel) Name() string {
	return [...]string{"ERROR", "一级分类", "二级分类", "三级分类"}[i]
}
