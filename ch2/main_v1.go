package main

func main() {
	calc_damage(1, 2, 3, 4)
}

// 変数名を省略しているため名前が何を意図しているのかコードを見るだけだとわからない
func calc_damage(p1, p2, d1, d2 int) {
	var d int = 0
	d = p1 + p2
	d = d - ((d1 + d2) / 2)
	if d < 0 {
		d = 0
	}
}
