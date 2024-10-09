package utils

// 分类名称映射表
var classificationMap = map[string]string{
	"A": "马克思主义、列宁主义、毛泽东思想",
	"B": "哲学",
	"C": "社会科学总论",
	"D": "政治、法律",
	"E": "军事",
	"F": "经济",
	"G": "文化、科学、教育、体育",
	"H": "语言、文字",
	"I": "文学",
	"J": "艺术",
	"K": "历史、地理",
	"N": "自然科学总论",
	"O": "数理科学和化学",
	"P": "天文学、地球科学",
	"Q": "生物科学",
	"R": "医药、卫生",
	"S": "农业科学",
	"T": "工业技术",
	"U": "交通运输",
	"V": "航空、航天",
	"X": "环境科学、劳动保护科学",
	"Z": "综合性图书",
}

// GetClassificationName 根据分类号获取分类名称
func GetClassificationName(classification string) string {
	if classification == "" {
		return ""
	}
	// 获取分类号的第一个字母
	firstLetter := classification[:1]
	// 在映射表中查找对应的分类名称
	name, ok := classificationMap[firstLetter]
	if !ok {
		return "未知分类"
	}
	return name
}
