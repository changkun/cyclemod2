This module `cyclemod2` and a [brother module `cyclemod`](https://github.com/changkun/cyclemod) demonstrates
the following interesting relationship: 

Assume we have two modules: A and B.

- A@v1.0.0 depends B@v2.0.0
- B@v2.0.0 depends A@v0.2.0
- A@v0.2.0 depends B@v0.1.0
- B@v0.1.0 depends A@v0.1.0
- A@v0.1.0 has no dependency

That means, in A:

```
import "changkun.de/x/cyclemod2/v2"
```

and in B:

```
import "changkun.de/x/cyclemod"
```

Will this leads to the infamous "import cycle not allowed"?
