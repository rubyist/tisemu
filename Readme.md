Just playing around emulating TIS-100 in Go.

Example:
```
./tisemu signal-divider.tis -map=standard.map -in=1,in.a -in=2,in.b -out=9,out.q -out=10,out.r
```

```
./tisemu sequence-indexer.tis -map=memory.map -in=0,in.v -in=2,in.x -out=10,/dev/stdout
```

```
./tisemu display.tis -map=display.map
```

