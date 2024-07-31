```
       ░▒▓█▓▒░  ░▒▓██████▓▒░   ░▒▓███████▓▒░ ░▒▓██████████████▓▒░  ░▒▓█▓▒░ ░▒▓███████▓▒░  ░▒▓████████▓▒░ 
       ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        
       ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        
       ░▒▓█▓▒░ ░▒▓████████▓▒░  ░▒▓██████▓▒░  ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓██████▓▒░   
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        
 ░▒▓██████▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓███████▓▒░  ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓████████▓▒░ 
                                                                                                         
                                                                                                         
```

Jasmine is an open source nearest point finder program with K-D Tree implementation in Go.

## What Are K-D Trees?
A K-D Tree is a binary tree in which each node represents a k-dimensional point. Every non-leaf node in the tree acts as a hyperplane, dividing the space into two partitions. This hyperplane is perpendicular to the chosen axis, which is associated with one of the K dimensions.

There are different strategies for choosing an axis when dividing, but the most common one would be to cycle through each of the K dimensions repeatedly and select a midpoint along it to divide the space. For instance, in the case of 2-dimensional points with x and y axes, we first split along the x-axis, then the y-axis, and then the x-axis again, continuing in this manner until all points are accounted for:

![K-D Tree](https://www.baeldung.com/wp-content/uploads/sites/4/2023/03/kdtree.png)
#### Introduction to K-D Trees. (2023). [baeldung.com](https://www.baeldung.com/cs/k-d-trees) (Accessed on: July 26, 2024)

## How Does a K-D Tree Work find nearest point?

![K-D Tree](https://upload.wikimedia.org/wikipedia/commons/3/36/Kdtree_animation.gif)
#### K-d tree. (2024). [wikimedia](https://en.wikipedia.org/wiki/K-d_tree) (Accessed on: July 26, 2024)


## Features

- [x] In-memory database
- [x] Backup and Restore file
- [x] gRpc
- [x] Rest API
- [x] Export to GeoJson
- [x] TTL (Time to Live)
- [x] OSRM Integration
- [ ] Valhalla Integration

## Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/mseptiaan/jasmine/internal/core
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkPostStore
BenchmarkPostStore-8        	 8055708	       132.5 ns/op
BenchmarkGetData
BenchmarkGetData-8          	 4027813	       280.1 ns/op
BenchmarkGetDataGeoJson
BenchmarkGetDataGeoJson-8   	 3060290	       407.7 ns/op
BenchmarkPostNearby
BenchmarkPostNearby-8       	 1236009	       918.5 ns/op
PASS
```



## Contributors
Thanks to all of the following who contributed to `JASMINE`:

<a href="https://github.com/mseptiaan/jasmine/graphs/contributors"><img src="https://contrib.rocks/image?repo=mseptiaan/jasmine&max=100&columns=16" /></a>