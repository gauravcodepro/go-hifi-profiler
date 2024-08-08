# go-hifi-profiler

- Profiling of the pacbiohifi reads, making mers and selective filtering of the mers for bubble pop. Some bubbles get clogged dur to composition.
- This allows to filter the same. 
- A pacbiohifi go profiler that takes the pacbiohifi reads, makes the mers according to the length, filters the mers according to the critera.
- This is supported witht a desktop application to see that if your reads have high profiled genomic mers that could hinder the graph assembly.

```
go build main.go
 # provide the path to the pacbiohifi reads file
```
- implementing a embed and also the cli so that it will take the merqury results and will run the concurrent primitives on the same.
- writing a byte stitcher that will stich the reads after adaptive filtering. 

Gaurav Sablok
