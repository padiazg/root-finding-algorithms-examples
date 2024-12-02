# Numerical Root-Finding Algorithms

This repository contains demonstration code for an article published on [Numerical Root-Finding Part 2](https://padiazg.github.io/posts/numerical-root-finding-part2/). It includes implementations of three numerical root-finding algorithms: Bisection, Newton-Raphson, and Secant.

All three implementatins calculares the square-root of a fiven number, following the theory explained in the articles.

## Algorithms

### Bisection Method
The Bisection method is a simple and robust root-finding algorithm that repeatedly bisects an interval and then selects a subinterval in which a root must lie for further processing. It is guaranteed to converge if the function is continuous on the interval and the initial interval is chosen such that the function values at the endpoints have opposite signs.

**Example Usage:**
```go
bisection := NewBisection(threshold, maxiter, showDetails, summaryTable)
root := bisection.Calc(value)
```

### Newton-Raphson Method
The Newton-Raphson method is an iterative root-finding algorithm that uses the first derivative of the function to find successively better approximations to the roots (or zeroes) of a real-valued function. It is faster than the Bisection method but requires the derivative of the function.

**Example Usage:**
```go
newtonRaphson := NewNewtonRaphson(threshold, maxiter, showDetails, summaryTable)
root := newtonRaphson.Calc(value)
```

### Secant Method
The Secant method is similar to the Newton-Raphson method but does not require the derivative of the function. Instead, it uses a sequence of roots of secant lines to approximate the root. It is generally faster than the Bisection method and does not require the derivative, making it more broadly applicable than Newton-Raphson.

**Example Usage:**
```go
secant := NewSecant(threshold, maxiter, showDetails, summaryTable)
root := secant.Calc(value)
```

## Running the Code

To run the code with detailed output and a summary, use the following command:

```bash
go run main.go bisection.go newton-raphson.go secant.go --show-details --show-summary --values 2,5,19
```

To run the code without detailed output but with a summary, use:

```bash
go run main.go bisection.go newton-raphson.go secant.go --show-summary --values 2,5,19
```

## Flags

- `--show-details`: Displays detailed iteration information for each algorithm.
- `--show-summary`: Displays a summary table with the results of each algorithm.
- `--values`: Specifies a comma-separated list of values for which to find the roots.

## Running Benchmarks

To run the benchmarks for the algorithms, use the following command:

```bash
go test -bench=.
```

This will execute the benchmark tests defined in the test files, providing performance metrics for each algorithm.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Buy me a coffee
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/padiazgy)

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
