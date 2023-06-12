# recursive generic type reproduction

This repository showed how to reproduce the compiler unexpected behavior, or issue, bug when a struct that has a field to reference the outer struct that has a field to reference the initial struct by using generic in the way recursively: the compiler will complain `fatal error: all goroutines are asleep - deadlock!` or other behave unexpectedly in all different ways.

Well I could understand the limitations of the current implementation for generics and the "should avoid" with recursive types, but I still think this is a issue that the compiler or documentation should address when user attempts to do so. Especially it runs successfully without generating the needed assembly codes, and ran successfully when you change the package of `main_test.go` into `recursive_generic_type_issue_reproduction` without the `_test` suffix. Imaging, you, as being a Golang developer, might only find out what has been wrong after you released the library without the `_test` test packages test included and caused the compiler issue and complains for all your users.

Even though such unexpected behavior could be resolved by `GOEXPERIMENT=nounified` flag[^1] and reported rarely (or even nobody have mentioned so far), it's still a bit confusing for users to understand what's going on.

At least it jammed me for a while.

## Behaviors

- Compiler might complain the `fatal error: all goroutines are asleep - deadlock!` panics message when running the test by `go test -v ./...`.
- Compiler might be able to compile the needed assembly codes and run the test successfully when you write the test codes without the `_test` suffix (exported test).
- Compiler might complain about the `invalid recursive type` if you try to move the `T2[T any]` definition to another file.
- Compiler might not be able to generate the needed assembly codes for the generic recursive types and produces no error or warning messages silently.
- Compiler might complain about the `redecalred` if you try to reload the IDE somehow.

## How to reproduce the issue from this repository?

This is the most needed directory structure for reproducing the issue, each cases are well documented with the case, summary, and how-to-reproduce, you may navigate into the directory and run the command I provided in the `README.md` file to reproduce the issue by yourself on-demand. Or you could use the next section to click the link to the issue and case you want to reproduce by following the instructions.

```shell
.
├── minimum_repro
│   ├── deadlock_issue # issue 1: fatal error: all goroutines are asleep - deadlock!
│   │   ├── with_generics # at the case of using generics
│   │   ├── with_generics_with_pointer # at the case of using generics with pointer wrapped
│   │   └── with_generics_without_test_package # at the case of using generics without test package (no _test suffix)
│   ├── unstable_invalid_recursive_issue # issue 2: invalid recursive type
│   │   ├── seperated_files # at the case of seperated files
│   │   └── single_file # at the case of single file
│   └── unstable_redeclare_issue # issue 3: redeclared in this block
│       └── seperated_files # at the case of seperated files
└── original_scenario # the original scenario that I encountered the issue
```

### Issue 1: fatal error: all goroutines are asleep - deadlock!

- [when using generics with test package (`_test` suffix)](./minimum_repro/deadlock_issue/with_generics)
- [when using generics with pointer reference to the generic type](./minimum_repro/deadlock_issue/with_generics_with_pointer)
- [**disappeared** deadlock error when using generics **without** test package (`_test` suffix)](./minimum_repro/deadlock_issue/with_generics_without_test_package)

### Issue 2: invalid recursive type

- [when defining one of the type parameter of type `innerT` with union type](./minimum_repro/unstable_invalid_recursive_issue/seperated_files)
- [**disappeared** `invalid recursive type` error when defining one of the type parameter of type `innerT` with union type.](./minimum_repro/unstable_invalid_recursive_issue/single_file)

### Issue 3: redeclared in this block

- [when defining one of the type parameter of type `innerT` with union type](./minimum_repro/unstable_redeclare_issue/seperated_files)

[^1]: [cmd/compile: failed to compile some recursive generic type · Issue #54535 · golang/go](https://github.com/golang/go/issues/54535)
