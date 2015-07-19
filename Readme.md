# tisemu 

`tisemu` is a [TIS-100](http://www.zachtronics.com/tis-100/) emulator. TIS-100
is a rad open-ended programming game where you solve puzzles by programming a
multi-node machine in a kind of assembly language. You should probably go play
the game instead of messing around here.

`tisemu` supports the `T21` execution node, the `T30` stack memory node, and a
terminal based visualization mode (using termbox). I think there are plans to
add more node types to the game. When that happens I'll probably add them here.

## Machine Maps

In the game, each puzzle can have a different array of nodes and node types.
When using `tisemu`, the machine can be describe in a simple map file. The map
file format is as follows:

```
<COLS>
<ROWS>
<DISPLAY>
<NODETYPE>
...
```

`COLS` is the number of columns in the node array. `ROWS` is the number of rows
in the node array. `DISPLAY` describes the display capabilities. This value can
either be `F` for no display, or `T <COLS> <ROWS>` for a display with a
provided gemetry. `NODETYPE` describes a type for each node in the array,
starting from the top left proceeding to the bottom right. There should be
`COLS * ROWS` lines. The supported values are currently `T21` and `T30`. If no
map file is provided, a default map of 4x3 (12 total) T21 nodes with no display
will be used. See the examples directory for some example map files. 

## Code Files

Code for all nodes in a machine lives in one file. It's best to play the game
and understand the code first. `tisemu` tries to keep the same format as the
game's save files, but differs slightly in node numbering. The game does not
appear to maintain numbering for "bugged" nodes and `T30` nodes, but `tisemu`
does. If you want to plug in save files from the game you might need to
renumber the nodes. I might fix this in the future.

## Running Code

Input and output to the machine is given on the command line in the format
`-in=<node>,<file>` or `-out=<node>,<file>`. For example, `-in=1,in.a` will
open file `in.a` and write its lines as input to node `1`.

### Examples

```
./tisemu signal-divider.tis -map=standard.map -in=1,in.a -in=2,in.b -out=9,out.q -out=10,out.r
```

```
./tisemu sequence-indexer.tis -map=memory.map -in=0,in.v -in=2,in.x -out=10,/dev/stdout
```

```
./tisemu display.tis -map=display.map
```

