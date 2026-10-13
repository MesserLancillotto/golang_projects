package table;

import (
	"wave_function_collapse/cell"
)

type Table struct {
	CellWidth 	 int
	CellHeight 	 int 
	Shape 		 uint64
	Cells 		 []cell.Cell
	UpdateBuffer []Vector2
}

type Vector2 struct {
	row int
	col int
}

func (table *Table) SetCollapsedCell(row int, col int, value int) {
		table.Cells[row * int(table.Shape) + col].CollapseCell(value)
}

func CreateTable(
    possibleValuesCardinality int, 
    subcellWidth int, 
    subcellHeight int) *Table {
    var table *Table = new(Table)
    table.Shape = uint64(possibleValuesCardinality)
    num := int(table.Shape * table.Shape)
    table.Cells = make([]cell.Cell, int(num))
    table.CellWidth = subcellWidth
    table.CellHeight = subcellHeight
    for i := 0; i < num; i++ {
        table.Cells[i] = *cell.NewCell(uint64(possibleValuesCardinality))
    }
    return table
}

func (table *Table) updateRow(row int, col int) {
    var index int
    var shape int = int(table.Shape)
    for j := 0; j < shape; j++ {
        index = row * shape + j
        table.Cells[index].DeletePossibility(int(table.Cells[row * shape + col].FinalValue))
        table.triggerUpdate(row, j)   
    }
}

func (table *Table) updateColumn(row int, col int) {
    var index int
    var shape int = int(table.Shape)
    for i := 0; i < shape; i++ {
        index = i * shape + col
        table.Cells[index].DeletePossibility(int(table.Cells[row * shape + col].FinalValue))
        table.triggerUpdate(i, col)
    }
}

func (table *Table) updateCell(row int, col int) {
    var cell *cell.Cell
    shapeInt := int(table.Shape)
    startRow := (row / int(table.CellHeight)) * int(table.CellHeight)
    startCol := (col / int(table.CellWidth)) * int(table.CellWidth)
    endRow := startRow + int(table.CellHeight)
    endCol := startCol + int(table.CellWidth)
    sourceValue := int(table.Cells[row*shapeInt + col].FinalValue)
    for i := startRow; i < endRow; i++ {
        for j := startCol; j < endCol; j++ {
            cell = &table.Cells[i*shapeInt+j]
            cell.DeletePossibility(sourceValue)
            table.triggerUpdate(i, j)
        }
    }
}

func (table *Table) triggerUpdate(row int, col int) {
    var cell *cell.Cell = &table.Cells[row * int(table.Shape) + col]
    if cell.Checked || !cell.IsCollapsed() {
        return
    }
    cell.Checked = true
    table.updateRow(row, col)
    table.updateColumn(row, col)
    table.updateCell(row, col)
}

func (table *Table) WaveFunctionCollapse() {
	var shape int = int(table.Shape) 
	for i := 0; i < shape; i++ {
		for j := 0; j < shape; j++ {
			table.triggerUpdate(i, j)
		}
	}
}

