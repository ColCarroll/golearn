# golearn
Experiments in machine learning with Go
Run 
```
go run main.go
```
to see it in action.  It will generate a 1000x5 matrix `A` using a uniform random distribution, 
a 5x1 matrix `x` using a random normal distribution, and a matrix `y = Ax + e`, where `e` is 
a 1000x1 error matrix (normally distributed with std. deviation = 0.1.  The library then solves
the equation Ax=y for x using the 
[Moore-Penrose pseudoinverse](https://en.wikipedia.org/wiki/Moore%E2%80%93Penrose_pseudoinverse).

It does ok!  It is pretty quick!

## golearn/num
Contains a matrix struct that knows how to do *some* linear algebra

## golearn/linear
Implements linear regression.  
