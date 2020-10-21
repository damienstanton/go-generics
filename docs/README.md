## Upshot from the draft design
- Functions can have an additional type parameter list that uses square brackets but otherwise looks like an ordinary parameter list: `func F[T any](p T) { ... }`.
- These type parameters can be used by the regular parameters and in the function body.
- Types can also have a type parameter list: `type M[T any] []T`.
- Each type parameter has a type constraint, just as each ordinary parameter has a type: `func F[T Constraint](p T) { ... }`.
- Type constraints are interface types.
- The new predeclared name any is a type constraint that permits any type.
- Interface types used as type constraints can have a list of predeclared types; only type arguments that match one of those types satisfy the constraint.
- Generic functions may only use operations permitted by the type constraint.
- Using a generic function or type requires passing type arguments.
- Type inference permits omitting the type arguments of a function call in common cases.

## Helpful Links

- [The wikipedia entry on generics (actually quite good)][3]
- [The full Go generics draft design document by Ian Lance Taylor and Robert Griesemer][1]
- [The `go2go` playground][2]

[1]: https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md
[2]: https://go2goplay.golang.org/
[3]: https://en.wikipedia.org/wiki/Generic_programming

© 2020 Damien Stanton

See LICENSE for details.