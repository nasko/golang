package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankSquares, exists := cb[rank]

	if !exists {
		return 0
	}

	var countOfOccupied int
	for _, occupied := range rankSquares {
		if occupied {
			countOfOccupied++
		}
	}
	return countOfOccupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}

	var countOfOccupied int
	for _, rank := range cb {
		if rank[file-1] {
			countOfOccupied++
		}
	}
	return countOfOccupied
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var countOfSquares int
	for _, rank := range cb {
		for range rank {
			countOfSquares++
		}
	}
	return countOfSquares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var countOfOccupied int
	for rKey := range cb {
		countOfOccupied += CountInRank(cb, rKey)
	}
	return countOfOccupied
}
