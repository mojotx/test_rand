# Random Number Distribution Test

[![Generate code coverage badge](https://github.com/mojotx/test_rand/actions/workflows/coverage.yaml/badge.svg?branch=main)](https://github.com/mojotx/test_rand/actions/workflows/coverage.yaml)

## Summary

For 1000 iterations, calculate a random number between 0-99, and if the
number is less than 30, it is a challenger, otherwise a champion.

Take the challenger numbers, and calculate if how close they are to 30%
of the population.

Check the population standard deviation as a measure of closeness.

## Sample Output

```text
Mean: 29.944000
Standard Deviation (σ): 1.421142
Min: 26.200000
Max: 34.200000
```
