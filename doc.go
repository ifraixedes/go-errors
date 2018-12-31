/*Package errors provides a simple API for creating and handling errors which
are identified by codes.

Rationale

Many times, a library, service, or any other kind of implementation needs to
deal with third party packages and it may need to identify errors for:

1. Perform some operation before returning it to the caller, for example doing
some kind of rollback operations.

2. Localize the error.

Meanwhile those 2 needs could be implemented with error values (defining them
as global variables of the package) or specific types, both of them aren't ideal
because:

1. Package variables values cannot be modified on each error instance, so they
cannot to be endowed with specific information which is useful for the
developers and operations teams for identifying the source of the problem.

2. Package variables aren't immutable, so their values, although unlikely, could
be erroneous overridden.

3. Specific types don't have the 2 previous problems, however when a considerable
list of different errors is needed, quite a few boilerplate is required for
creating them and the package documentation gets polluted.

4. Specific types must implement the same behavior over and over or use a base
type which must be embedded in all the types to have the same logic. This could
be fine in single package, but when you want to have the same mechanism on a
bunch of packages (think in medium/large implementations done by any company
whose code base is written in Go) is less than desirable to spread between
teams.

5. When using any of them and the errors are transmitted over the wire (this
problem arises when the dependency is a remote service), then the client must
identify those errors in order to reconstruct the error to use the same value
(when using package variables) or type (when using specific types) for allowing
the caller to be able to identify the error.

The 2 mentioned issues can be solved by errors which are identified by codes and
without the need of using package variables nor specific error types, just using
a minimal and simple public API exposed by this package. Nonetheless this
package doesn't attempt to fit to all the use cases, it fits to several of the
uses cases which were found in my experience, but they are not all. Hence,
before using it, assess  if it can bring the mentioned benefits to your
implementation (for example a minimal library may have enough just returning
standard errors or using a couple of package variables error values variables or
specific types).

Errors information

Errors are for users, but they must be useful for operations, too. In order to
achieve both, the error type of this package is endowed with several information.

1. A code and a static message. Both information is useful for users and
operations, because each code should be quite specific and self descriptive for
providing a synthesized information about the error which has happened. Having
a specific error code is also good to have each error properly documented, so it
can provide more detailed information about the error when its synthesized
information isn't enough.

2. An unique ID. Each error instance has an unique ID. For users, it's useful
because they could report it to the support team, a part of the code. Such ID
could be helpful to provide a customized feedback/response when needed and for
the operations team, could use the ID to correlate errors, when they are
registered/tracked in different operational systems or, for any reason, in the
same one several times.

3. Metadata. Despite that the code, and its associated message, should be
precise, the operations team needs more information about what happened when the
error has happened in some circumstances which aren't clear, for example the
input parameter values, variable values, etc. Developers should have a way to
provide such important context information when creating the errors and that's
what, in this package, is called metadata.

4. The call stack. Call stacks are ugly, but they provide the trace where the
error was originated and such information is very useful for the operations team
and maintainers, when the error has happened in unclear circumstances.

5. Original error. The most of the times, third party packages are used and,
obviously, those package don't probably use errors created by this package,
hence they don't have all the information, or a lest in the same way. Your
implementation has committed to return error with codes, but the original error
must be preserved because it may have additional information which may be useful
for the operations team and maintainers.

In summary, from the point of view of users, the error should be precise and
concise, without having any useless information and avoiding to leak important
information about the system; on the other hand, the operations team and
maintainers need much more information about the errors in order of being able
to understand the cause of the error.

Errors are standard errors

This package doesn't intentionally export the error type, because:

1. Functions' signatures, which return errors, should always return a standard
error.

2. Type assertion for finding out errors information is less than ideal in
comparison on having an API.


Hence, this package export some functions to get information about the error,
however it only intentionally allows to get a part of it, because the
information destined for operations is only thought to be exposed through
systems for such purpose, for example logging.


Printing the error

The error values returned by this package can be printed with more or less
information depending used verb and flags. Below you can see all the verbs and
flags and the information that each one prints, any other verb, it prints
nothing.

	fmt.Printf("%s", err) // Only code and message
	// Output
	// TestCode: an test code error has happened

	fmt.Printf("%v", err) // Code, message, ID and metadata
	// Output
	// TestCode: an test code error has happened
	//    id: a0d2dbe9-4aa9-47d2-9630-2cde8a3b1a0b
	//    metadata: [{"var1": a string},{"var2": 10}]

	fmt.Printf("%+v", err) // Code, message, ID, metadata and call stack.
	// Output
	// TestCode: an test code error has happened
	//    id: a0d2dbe9-4aa9-47d2-9630-2cde8a3b1a0b
	//    metadata: [{"var1": a string},{"var2": 10}]
	//    call stack:
	//    go.fraixed.es/errors.TestDerror_Format.func1
	//            /root/workspace/go-errors/derror_test.go:23
	//    go.fraixed.es/errors.TestDerror_Format.func2
	//            /root/workspace/go-errors/derror_test.go:31
	//    go.fraixed.es/errors.TestDerror_Format
	//            /root/workspace/go-errors/derror_test.go:35
	//    testing.tRunner
	//            /root/apps/go/src/testing/testing.go:827
	//    runtime.goexit
	//            /root/apps/go/src/runtime/asm_amd64.s:1333

	fmt.Printf("%+-v", err) // Code, message, ID, metadata and reduced call stack.
	// Output
	// TestCode: an test code error has happened
	//    id: 7ec98655-f9a7-4c38-ac3d-8b40d8c6d581
	//    metadata: [{"var1": a string},{"var2": 10}]
	//    call stack (compacted):
	//    go.fraixed.es/errors.TestDerror_Format.func1
	//    go.fraixed.es/errors.TestDerror_Format.func2
	//    go.fraixed.es/errors.TestDerror_Format
	//    testing.tRunner
	//    runtime.goexit


	fmt.Printf("%+v", err) // When err wraps another error, it's printed with this verb and flag.
	// Output
	// TestCode: an test code error has happened
	//    id: 7ec98655-f9a7-4c38-ac3d-8b40d8c6d581
	//    metadata: [{"var1": a string},{"var2": 10}]
	//    wrapped error: some external error
	//    call stack:
	//    go.fraixed.es/errors.TestDerror_Format.func1
	//            /root/workspace/go-errors/derror_test.go:24
	//    go.fraixed.es/errors.TestDerror_Format.func2
	//            /root/workspace/go-errors/derror_test.go:32
	//    go.fraixed.es/errors.TestDerror_Format
	//            /root/workspace/go-errors/derror_test.go:36
	//    testing.tRunner
	//            /root/apps/go/src/testing/testing.go:827
	//    runtime.goexit
	//            /root/apps/go/src/runtime/asm_amd64.s:1333

Therefore, developers should consider to use one another depending the needs and
the audience of those messages.
*/
package errors
