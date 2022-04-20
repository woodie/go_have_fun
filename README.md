# ternary

A ternary module for go.

## Usage

```go
import (
	"netpress.com/fun"
)

recount := fun.ternary[int](c > 5, 0, c+5)
color := fun.ternary[string](brightness > 0, "cyan", "white")
```

# Advantage

One line is easier to read than five (in some cases).

```diff
-	if self.d == 0 {
-		self.d = 3
-	} else {
-		self.d = self.d - 1
-	}
+	self.d = fun.ternary[int](self.d == 0, 3, self.d-1)
}
```
