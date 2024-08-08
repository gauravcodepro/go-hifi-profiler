# go-hifi-profiler

- a type interface for the profiling of the pacbiohifi reads, making hapmers and selective filtering of the hapmers for bubble pop
- this takes a concurrency slice mapping for the filtering of the hapmers so that the reads with the bubble clog can be filtered out.
- A pacbiohifi go profiler that takes the pacbiohifi reads, makes the mers according to the length, filters the mers according to the critera.
- This is supported witht a desktop application to see that if your reads have high profiled genomic mers that could hinder the graph assembly.

Gaurav Sablok
