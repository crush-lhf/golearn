package basics

//方式一 常规写法
func RectProps(length, height int) (int, int) {
	var area = length * height
	var per = length + height
	return area, per
}

//方式二 隐藏写法
func RectProps02(length, height int) (area, per int) {
	area = length * height
	per = length + height
	return
}


