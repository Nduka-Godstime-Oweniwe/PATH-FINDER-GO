Title: PATH-FINDER

Description: PATH-FINDER is a terminal-based Go application that determines whether a valid path exists between a Start point (`S`) and an End point (`E`) on a 2D board.
The application uses recursion and backtracking to explore different possible paths while avoiding blocked cells marked as `X`.


Users can
* Input custom boards
* View shortest paths
* View longest paths
* View all valid paths
* Replay the game


Features
* Dynamic board size
* Row-by-row board input
* Full board input mode
* Input validation
* Recursive path solving
* Backtracking algorithm
* Multiple valid path detection
* Shortest path visualization
* Longest path visualization
* Beautiful arrow path display
* Replay system

Technologies Used
* Go (Golang)
* Recursion
* Backtracking
* Arrays & Slices
* 2D Boards
* Terminal UI

Project Structure
```text
path-finder/
│
├── go.mod
├── main.go
├── README.md
│
└── functions/
    │
    ├── createboard.go
    ├── printboard.go
    ├── solve.go
    ├── userchoice.go
    └── userinput.go
```

Shared Data Types
Board Type
```go
[][]string
```

Main Functions

UserInput():

Handles everything related to collecting and validating user input.

The user can:
* input all rows at once
* input rows one after another

Returns:
    ([]string, bool)

Example:

[]string{
	"S..X",
	".X..",
	"...E",
}
CreateBoard()

Converts the user input into a 2D board.

Input:

[]string

Returns:

[][]string
Solve()

This is the main recursive function used for solving the board.

The function:

finds possible paths
avoids walls (X)
prevents invalid movement
uses backtracking when a path fails

Returns:

([][]string, bool)

The recursion and backtracking logic was one of the hardest parts of the project because we had to avoid infinite loops while still checking all possible paths.

PrintBoard()

Prints the solved board using arrows to show the path taken.

Input:

[][]string
UserChoice()

Allows the user to choose between:

shortest path
longest path
all paths
go back

How The Solver Works
1. The program finds the position of `S`
2. It checks possible directions:
   * Right
   * Left
   * Up
   * Down
3. The algorithm recursively explores valid cells
4. Invalid cells are ignored:
   * `X`
   * out of bounds
   * visited paths
5. If a path fails, backtracking occurs
6. Valid paths are stored and displayed


Backtracking:
The project uses recursion with backtracking.

When the solver reaches a dead end:
- the current move is undone
- the algorithm returns to the previous position
- another path is attempted

This continues until:
- `E` is found
- or all paths fail


Example:
Input:
```text
S..X
.X..
...E
```


Possible Output:
```text
S  ➡️  ➡️  X

.  X  ⬇️  .

.  .  ➡️  E
```


User Menu:
```text
1. Shortest Path
2. Longest Path
3. All Path
4. Go Back
```


Input Rules:
* Only these characters are allowed:
```text
S E X .
```
* Only one `S`
* Only one `E`
* All rows must have equal length


Team Contributions:

| Member | Responsibility |
|--------|----------------|
| Isreal | UserInput & Validation |
| Joshua | CreateBoard |
| Nduka | Solve Logic & Backtracking |
| Samuel | PrintBoard |
| Abayomi | Main Flow Integration |
| Esther | Documentation & Testing |


Challenges Faced
* 
* 


Lessons Learned
Through this project we learned:
* Recursion
* Backtracking
* Team collaboration
* Debugging

Authors
* Isreal
* Joshua
* Nduka
* Samuel
* Abayomi
* Esther 
