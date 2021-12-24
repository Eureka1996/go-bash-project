package entities

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// 实现Interface接口的方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

// i,j为下标，Less方法就是决定使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	// 年龄从小到大排序
	return hs[i].Age > hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	tmp := hs[i]
	hs[i] = hs[j]
	hs[j] = tmp
}
