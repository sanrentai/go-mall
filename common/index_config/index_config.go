package index_config

type indexConifg int

const (
	DEFAULT indexConifg = iota
	INDEX_SEARCH_HOTS
	INDEX_SEARCH_DOWN_HOTS
	INDEX_GOODS_HOT
	INDEX_GOODS_NEW
	INDEX_GOODS_RECOMMOND
)

func IndexConifg(i int) indexConifg {
	if i > int(INDEX_GOODS_RECOMMOND) || i < int(DEFAULT) {
		return DEFAULT
	}
	return indexConifg(i)
}

func (i indexConifg) Name() string {
	return [...]string{"DEFAULT", "INDEX_SEARCH_HOTS",
		"INDEX_SEARCH_DOWN_HOTS", "INDEX_GOODS_HOTS",
		"INDEX_GOODS_NEW", "INDEX_GOODS_RECOMMOND"}[i]
}
