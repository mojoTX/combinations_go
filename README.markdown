
 # Mathematical Combinations / Permutations

With combinations, order does not matter. For instance, if you
pick 3 out of * 16 billiard balls, the number of possible
choices is 560.

  Formula:

        n!
  --------------
    r!(n - r)!

![n! / r!(n - r)! ](https://render.githubusercontent.com/render/math?math=\frac{n!}{r!(n%20-%20r)!})


```
$ go run combo.go 16 3
----------------------------------------------------------------

With combinations, order does not matter. For instance, if you
pick 3 out of * 16 billiard balls, the number of possible
choices is 560.

  Formula:

        n!
  --------------
    r!(n - r)!
----------------------------------------------------------------

For 16 elements, taken 3 at a time, ignoring order and without
repetition, you have 560 possible combinations.
```
