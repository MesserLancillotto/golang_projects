package main

import (
	"fmt"
	"wave_function_collapse/table"
)

func main() {
	easy()
	medium()

}

func PrintTable(t *table.Table) {
	shape := int(t.Shape)
	for i := 0; i < shape; i++ {
		for j := 0; j < shape; j++ {
			cell := t.Cells[i*shape+j]
			val, ok := cell.GetValue()
			if ok {
				fmt.Printf("[%d]\t", val)
			} else {
				bitmap, _ := cell.GetPossibleValuesBitmap()
				var possible []int
				for b := uint64(0); b < t.Shape; b++ {
					if (bitmap & (1 << b)) != 0 {
						possible = append(possible, int(b+1))
					}
				}
				fmt.Printf("%v\t", possible)
			}
		}
		fmt.Println()
	}
}

func easy() {
		var t *table.Table = table.CreateTable(9, 3, 3)

	fmt.Println("Easy")

	t.SetCollapsedCell(0, 0, 5)
	t.SetCollapsedCell(0, 1, 3)
	t.SetCollapsedCell(0, 4, 7)
	t.SetCollapsedCell(1, 0, 6)
	t.SetCollapsedCell(1, 3, 1)
	t.SetCollapsedCell(1, 4, 9)
	t.SetCollapsedCell(1, 5, 5)
	t.SetCollapsedCell(2, 1, 9)
	t.SetCollapsedCell(2, 2, 8)
	t.SetCollapsedCell(2, 7, 6)

	t.SetCollapsedCell(3, 0, 8)
	t.SetCollapsedCell(3, 4, 6)
	t.SetCollapsedCell(3, 8, 3)
	t.SetCollapsedCell(4, 0, 4)
	t.SetCollapsedCell(4, 3, 8)
	t.SetCollapsedCell(4, 5, 3)
	t.SetCollapsedCell(4, 8, 1)
	t.SetCollapsedCell(5, 0, 7)
	t.SetCollapsedCell(5, 4, 2)
	t.SetCollapsedCell(5, 8, 6)

	t.SetCollapsedCell(6, 1, 6)
	t.SetCollapsedCell(6, 6, 2)
	t.SetCollapsedCell(6, 7, 8)
	t.SetCollapsedCell(7, 3, 4)
	t.SetCollapsedCell(7, 4, 1)
	t.SetCollapsedCell(7, 5, 9)
	t.SetCollapsedCell(7, 8, 5)
	t.SetCollapsedCell(8, 4, 8)
	t.SetCollapsedCell(8, 7, 7)
	t.SetCollapsedCell(8, 8, 9)

	t.WaveFunctionCollapse()

	PrintTable(t)
}

func medium() {
	fmt.Println("Medium")

	var t *table.Table = table.CreateTable(9, 3, 3)

	t.SetCollapsedCell(0, 2, 3)
	t.SetCollapsedCell(0, 3, 2)
	t.SetCollapsedCell(0, 5, 6)
	t.SetCollapsedCell(0, 6, 8)
	t.SetCollapsedCell(1, 0, 1)
	t.SetCollapsedCell(1, 1, 7)
	t.SetCollapsedCell(1, 7, 5)
	t.SetCollapsedCell(2, 3, 7)
	t.SetCollapsedCell(2, 4, 8)
	t.SetCollapsedCell(2, 5, 3)
	t.SetCollapsedCell(3, 1, 3)
	t.SetCollapsedCell(3, 4, 1)
	t.SetCollapsedCell(3, 7, 9)
	t.SetCollapsedCell(4, 0, 8)
	t.SetCollapsedCell(4, 3, 5)
	t.SetCollapsedCell(4, 5, 2)
	t.SetCollapsedCell(4, 8, 4)
	t.SetCollapsedCell(5, 1, 2)
	t.SetCollapsedCell(5, 4, 9)
	t.SetCollapsedCell(5, 7, 6)
	t.SetCollapsedCell(6, 3, 3)
	t.SetCollapsedCell(6, 4, 6)
	t.SetCollapsedCell(6, 5, 5)
	t.SetCollapsedCell(7, 1, 5)
	t.SetCollapsedCell(7, 7, 2)
	t.SetCollapsedCell(7, 8, 8)
	t.SetCollapsedCell(8, 2, 9)
	t.SetCollapsedCell(8, 3, 1)
	t.SetCollapsedCell(8, 5, 4)
	t.SetCollapsedCell(8, 6, 3)


	t.WaveFunctionCollapse()

	PrintTable(t)
}