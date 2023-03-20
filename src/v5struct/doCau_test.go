package v5struct

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		if got != want {
			t.Errorf("want %v,got %v", got, want)
		}
	})

}

/*
func TestArea(t *testing.T) {

		checkArea := func(t *testing.T, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got%.2f,want%.2f", got, want)
			}
		}

		t.Run("rectangle", func(t *testing.T) {
			rectangle := Rectangle{10.0, 10.0}
			checkArea(t, rectangle, 100.0)
		})

		t.Run("circle", func(t *testing.T) {
			circle := Circle{10.0}
			checkArea(t, circle, 314.1592653589793)

		})
}
*/
// 重构：表格驱动测试
func TestArea(t *testing.T) {
	// 声明结构体切片
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}

}
