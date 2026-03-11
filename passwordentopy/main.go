package passwordentopy
Sure! Let’s create a beginner-friendly **“Password Entropy”** question in the style of **01-edu / z01 checkpoints**, and provide a **simple, working solution** without any complex libraries.

---

# **Question: Password Entropy**

**Instructions**

Write a function `PasswordEntropy` that takes **two integers**:

1. `length` — the length of the password
2. `charsetSize` — the number of possible characters for each position

The function should return a **string** containing the estimated password entropy in **bits**, followed by a newline (`\n`).

Use the formula:

[
\text{entropy} = L \cdot \log_2(N)
]

where:

* `L` = password length
* `N` = size of the character set

If the password length is 0 or charset size is 0, the entropy is 0.

**Function signature:**

```go
func PasswordEntropy(length int, charsetSize int) string
```

**Example:**

```go
fmt.Print(PasswordEntropy(6, 95)) // password length 6, 95 possible characters
fmt.Print(PasswordEntropy(0, 95))
fmt.Print(PasswordEntropy(10, 26))
```

**Expected output:**

```
39.42
0.00
47.03
```

---

# **Beginner-Friendly Solution**

```go
package main

import (
	"fmt"
	"math"
)

func PasswordEntropy(length int, charsetSize int) string {
	if length == 0 || charsetSize == 0 {
		return "0.00\n"
	}
	
	entropy := float64(length) * math.Log2(float64(charsetSize))
	return fmt.Sprintf("%.2f\n", entropy)
}

func main() {
	fmt.Print(PasswordEntropy(6, 95))  // 39.42
	fmt.Print(PasswordEntropy(0, 95))  // 0.00
	fmt.Print(PasswordEntropy(10, 26)) // 47.03
}
```

---

# **Explanation**

1. Convert `charsetSize` to `float64` to use `math.Log2`.
2. Multiply by `length` → gives **entropy in bits**.
3. Format the result to **two decimal places** for readability.
4. Handle **edge cases** (`length = 0` or `charsetSize = 0`) by returning `0.00`.

---

This solution is:

* **Beginner-friendly** ✅
* Works with **z01-style checkpoints** ✅
* **No advanced packages** besides `fmt` and `math` ✅

---

If you want, I can also create a **“bonus” version** that takes a **password string** instead of length and charset, and automatically calculates **entropy based on the types of characters used**.

Do you want me to do that?
