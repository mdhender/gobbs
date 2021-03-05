# Test Driven Development
Copied from https://en.wikipedia.org/wiki/Test-driven_development.

# Test-driven development cycle

The following sequence is based on Ken Beck's book Test-Driven Development by Example:

## Add a test
The adding of a new feature begins by writing a test that passes if (and only if) the feature's specifications are met.
The developer can discover these specifications by asking about use cases and user stories.
A key benefit of test-driven development is that it makes the developer focus on requirements before writing code.
This is in contrast with the usual practice, where unit tests are only written after code.

## Run all tests. The new test should fail for expected reasons
This shows that new code is actually needed for the desired feature.
It validates that the test harness is working correctly.
It rules out the possibility that the new test is flawed and will always pass.

## Write the simplest code that passes the new test
Inelegant or hard code is acceptable, as long as it passes the test.
The code will be honed anyway in Step 5. No code should be added beyond the tested functionality.

## All tests should now pass
If any fail, the new code must be revised until they pass.
This ensures the new code meets the test requirements and does not break existing features.

## Refactor as needed, using tests after each refactor to ensure that functionality is preserved
Code is refactored for readability and maintainability.
In particular, hard-coded test data should be removed.
Running the test suite after each refactor helps ensure that no existing functionality is broken.
Examples of refactoring:
* moving code to where it most logically belongs
* removing duplicate code
* making names self-documenting
* splitting methods into smaller pieces
* re-arranging inheritance hierarchies

## Repeat
The cycle above is repeated for each new piece of functionality.
Tests should be small and incremental, and commits made often.
That way, if new code fails some tests, the programmer can simply undo or revert rather than debug excessively.
When using external libraries, it is important not to write tests that are so small as to effectively test merely the library itself, unless there is some reason to believe that the library is buggy or not feature-rich enough to serve all the needs of the software under development.

# Development style
## Keep the unit small
For TDD, a unit is most commonly defined as a class, or a group of related functions often called a module.
Keeping units relatively small is claimed to provide critical benefits, including:

1. Reduced debugging effort:
When test failures are detected, having smaller units aids in tracking down errors.
1. Self-documenting tests:
Small test cases are easier to read and to understand.

# Best practices
## Test structure
Effective layout of a test case ensures all required actions are completed, improves the readability of the test case, and smooths the flow of execution.
Consistent structure helps in building a self-documenting test case.
A commonly applied structure for test cases has the following steps:

1. Setup:
Put the Unit Under Test (UUT) or the overall test system in the state needed to run the test.
1. Execution:
Trigger/drive the UUT to perform the target behavior and capture all output, such as return values and output parameters.
This step is usually very simple.
1. Validation:
Ensure the results of the test are correct.
These results may include explicit outputs captured during execution or state changes in the UUT.
1. Cleanup:
Restore the UUT or the overall test system to the pre-test state.
This restoration permits another test to execute immediately after this one.
In some cases in order to preserve the information for possible test failure analysis the cleanup should be starting the test just before the test's setup run.

## Practices to avoid, or "anti-patterns"
1. Having test cases depend on system state manipulated from previously executed test cases (i.e., you should always start a unit test from a known and pre-configured state).
1. Dependencies between test cases.
A test suite where test cases are dependent upon each other is brittle and complex.
Execution order should not be presumed.
Basic refactoring of the initial test cases or structure of the UUT causes a spiral of increasingly pervasive impacts in associated tests.
1. Interdependent tests.
Interdependent tests can cause cascading false negatives.
A failure in an early test case breaks a later test case even if no actual fault exists in the UUT, increasing defect analysis and debug efforts.
1. Testing precise execution behavior timing or performance.
1. Building "all-knowing oracles."
An oracle that inspects more than necessary is more expensive and brittle over time.
This very common error is dangerous because it causes a subtle but pervasive time sink across the complex project.
1. Testing implementation details.
1. Slow running tests.

# Sources
Beck, Kent (2002-11-08). Test-Driven Development by Example. Vaseem: Addison Wesley. ISBN 978-0-321-14653-3.