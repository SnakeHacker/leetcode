package computeArea

/*
在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。

每个矩形由其左下顶点和右上顶点坐标表示，如图所示。

示例:

输入: -3, 0, 3, 4, 0, -1, 9, 2
输出: 45
说明: 假设矩形面积不会超出 int 的范围。
*/

/*
@Author: Mickey
*/

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum := (C-A)*(D-B) + (G-E)*(H-F)
	if A >= G || B >= H || C <= E || D <= F {
		return sum
	}

	var x0, y0, x1, y1 int

	if A < E {
		x0 = E
	} else {
		x0 = A
	}

	if B < F {
		y0 = F
	} else {
		y0 = B
	}

	if C < G {
		x1 = C
	} else {
		x1 = G
	}

	if D < H {
		y1 = D
	} else {
		y1 = H
	}

	return sum - (y1-y0)*(x1-x0)
}
