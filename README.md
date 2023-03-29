# TestingTools

This repository demonstrates some of my work in writing testing tools.

## How to Use

The API tests I've put here run against a 
[public cards API](https://deckofcardsapi.com/). 
This is a small, public API, enough to demonstrate a few tests.
Be sure that you have access to the Internet, and this API, by clicking the
link. If you see a bright green page titled **Deck of Cards**, then you are
ready to try this out.

You can either just [read](#peruse_the_code) the code here or try it out in 
actual [use](#use_the_code).

### Peruse the Code

To read it, go to the links in the table above, and click into the desired 
area (at present, there's only `Golang/Cards`). Click any file to read the 
code.

You may also be interested in my [test plan](https://docs.google.com/spreadsheets/d/1TZbfLPdaYk1R5Fl56N6DukIsLcjtKRDLVovpysmGB6k/edit#gid=0), which is also only for golang at the moment.

### Use the Code

You can fork, or clone this repo locally, then descend into the appropriate
language folder. Run the tests as indicated under the language header in the 
[Languages](#languages) section.

## State of Progress

As of late March 2023, I've just started this repo. It will be a work in
progress for a little while as I add more tests, languages and testing 
protocols.

## Languages

I'm currently demonstrating only **[Golang](#golang)**, using:
- Golang's `testing` package.
- Ginkgo.
- Gomega.

I hope to eventually (soon) include **[Python](#python)** tests, using:
- Python's `requests` module.
- Pytest.
- Tavern. 

### Golang

If you just want to see my code, you can remain in this GitHub repo, and click

1. Ensure that Golang v. 1.20.x (or later) is installed, as I wrote
these tests with that version.
2. In your local repo of TestingTools, descend into the `Golang` folder, then 
into the `Cards` subfolder.
3. In the `Cards` folder, run `go test`. 
Expected result: **`Pass`**.

If you like, you can run with higher verbosity on `go test`: `go test -v`. This
verbosity setting pertains to the golang test runner. It will show you each 
test's status, so you can pinpoint where an error occurred. Including more
than one `v`, is undefined. Doing so will trigger a usage message, while 
failing the test.

There are also package-level flags (verbosity) -v flag, and another 
`-meta_test` flag that pertain
to this particular test package. If invoke either, or both package-level flags,
you'll want the test by test context. So, first include the runner flag, then
the runner flag `-args` to introduce the package-level flags. Then provide the
package flags as desired. E.g:

    go test -v -args -v=1 

    go test -v -args -v=3

    go test -v -args -meta_test

    go test -v -args -v=1 -meta_test 

As you can see above, the package-level verbosity flag takes an integer value. 
The higher the value, the more information is shared, up to a point. Currently
the level 3 is the maximum.

Including `-meta_test` will "test the test" by contriving failures in some 
selected test conditions. I "unsort" card orderings that should be sorted, and
sort some that should be unsorted. I also mess up the count of some cards in a
test condition where they should appear exactly once in the deck. At the time
this is written, that's an exhaustive list of all that I do.

You'll **never** want to run with `-meta_test` in place during production use.
It's purely diagnostic, to ensure that a desired error can be caught. 

### Python

1. Ensure that you are using Python 3.x. These instructions will not work with
Python 2.x.
2. (This step and further ones are to be determined).
