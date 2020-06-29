# rando
A package to randomize the order of test execution for methods defined on a test
suite struct.

The go maintainers have [accepted a proposal](https://github.com/golang/go/issues/28592)
to randomize test order in go 1.16, but until that release is out, test order is static.
As such, this small package will allow you to randomize the order of the tests defined
on your test suite object.

I wrote something very similar to use internally at work, but it's reasonable
enough that other people might want to use it. Plus, it the syntax is sufficiently
lightweight that pulling out the library should be super easy with 1.16 is released.
