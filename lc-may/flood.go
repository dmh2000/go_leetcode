package lcmay

import (
	"container/list"
)

type coord struct {
	row int
	col int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var oldColor int
	var queue *list.List
	var elem *list.Element
	var cell coord

	oldColor = image[sr][sc]

	// no change if newColor == oldColor, just return
	if oldColor == newColor {
		return image
	}

	// breadth first search with a queue
	queue = list.New()
	queue.PushBack(coord{sr, sc})
	for queue.Len() > 0 {
		// pop front of queue and get its value
		elem = queue.Front()
		cell = elem.Value.(coord)
		queue.Remove(elem)

		// paint the pixel
		image[cell.row][cell.col] = newColor

		// enqueue all the surrounding pixels that have original color

		// upper cell
		if cell.row > 0 {
			if image[cell.row-1][cell.col] == oldColor {
				queue.PushBack(coord{cell.row - 1, cell.col})
			}
		}
		// lower cell
		if cell.row < len(image)-1 {
			if image[cell.row+1][cell.col] == oldColor {
				queue.PushBack(coord{cell.row + 1, cell.col})
			}
		}
		// left cell
		if cell.col > 0 {
			if image[cell.row][cell.col-1] == oldColor {
				queue.PushBack(coord{cell.row, cell.col - 1})
			}
		}
		// right cell
		if cell.col < len(image[0])-1 {
			if image[cell.row][cell.col+1] == oldColor {
				queue.PushBack(coord{cell.row, cell.col + 1})
			}
		}
	}

	return image
}
