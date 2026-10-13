package cell;

import (
	"math/bits"
)

type Cell struct {
	PossibleValuesBitmap 		uint64
	FinalValue 					uint64
	possibleValuesCardinality 	uint64
	Checked 					bool
	collapsed 					bool
}

func NewCell(possibleValuesCardinality uint64) *Cell {
	var cell *Cell = new(Cell)
	var i uint64 = 0 
	for ; i < possibleValuesCardinality; i++ {
		cell.PossibleValuesBitmap |= uint64(1 << i)
	}
	cell.collapsed = false
	cell.Checked = false
	cell.FinalValue = 0
	cell.possibleValuesCardinality = possibleValuesCardinality
	return cell
}


func (cell *Cell) CollapseCell(value int) {
	cell.PossibleValuesBitmap = uint64(1 << value)
	cell.FinalValue = uint64(value)
	cell.Checked = false 
	cell.collapsed = true
}


func NewCollapsedCell(value uint64) *Cell {
	var cell *Cell = new(Cell)
	cell.collapsed = true
	cell.Checked = false
	cell.FinalValue = value
	return cell
}

func (cell *Cell) checkCollapse() bool {
	cell.collapsed = bits.OnesCount64(cell.PossibleValuesBitmap) == 1
	if !cell.collapsed {
		return false
	}
	for i := uint64(0); i < cell.possibleValuesCardinality; i++ {
		if uint64(1 << i) == cell.PossibleValuesBitmap {
			cell.FinalValue = i + 1
			break
		}
 	}
	return cell.collapsed
}

func (cell *Cell) DeletePossibility(value int) {
	if cell.collapsed {
		return
	}
	cell.PossibleValuesBitmap &^= uint64(1 << (value - 1))
	_ = cell.checkCollapse()
}

func (cell *Cell) IsCollapsed() bool {
	return cell.collapsed
}

func (cell *Cell) GetPossibleValuesBitmap() (uint64, bool) {
	return cell.PossibleValuesBitmap, cell.collapsed
}  

func (cell *Cell) GetValue() (uint64, bool) {
	if !cell.collapsed {
		return 0, false
	}
	return cell.FinalValue, true
}