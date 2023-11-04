# API TestingTools

This repository demonstrates some of my work in writing API testing tools.

## How to Use

The API tests I've put here run against a 
small, public API, just enough to demonstrate a few tests.
Be sure that you have access to the Internet, and that API, by clicking the
link: [public cards API](https://deckofcardsapi.com/).
If you see a bright green page titled **Deck of Cards**, then you are
ready to try this out.

You can either just [read](#peruse-the-code) the code here or try it out in 
actual [use](#use-the-code).

### Peruse the Code

To read the code, go to the links in the table above, and click into the 
desired area (at present, there's only `Golang/Cards`). Click any file to read
the code.

To return to this page, click the "TestingTools" linked node in any slash ("/")
delimited range of nodes (a.k.a: breadcrumb).

You may also be interested in my [test plan](https://docs.google.com/spreadsheets/d/1TZbfLPdaYk1R5Fl56N6DukIsLcjtKRDLVovpysmGB6k/edit#gid=0), which is:
- A work in progress.
- Only for golang at the moment.

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

Plans for the future include Ruby, as well as Java, both using Cucumber.

### Golang

If you just want to see my code, you can remain in this GitHub repo, and click
on "Golang/Cards" in the file navigation table near the top of this page.

Click "into" any file listed there:
- `deck_test.go` contains the API functions, each with a corresponding 
"Test..." function to drive the API function and test for the responses. 
- `golangcards.go` contains Ginkgo and Gomega handlers, and tests for the 
responses' status code and and body.

To return to this page, click the "TestingTools" node/link in any slash ("/")
delimited range of nodes (a.k.a: a breadcrumb).

To use my code:

1. Ensure that you've installed Golang v. 1.20.x (or later), as I wrote
these tests with that version.
2. Ensure that you have Git or Git Desktop installed. 
3. Clone this repository:

    a. In a terminal window (Bash, KSH, DOS, etc.), do a `cd` to the parent 
       directory of your Github directories. It's something you will have
       chosen the name of, but it's often named `Git` or something similar.

    b. Copy the url of this page from the browser location bar.

    c. In the terminal, in your Git parent directory, start the command by
       typing `git clone ` then paste this repository's URL. You should
       end up with something like: 
        `git clone https://github.com/xylemase/TestingTools`

    d. Execute the command by hitting the `Enter` key.

    Expected result: You should see Git making a directory that is the local
    repository (repo). 
4. CD into your local repo of TestingTools, then descend into the `Golang` 
folder, then into the `Cards` subfolder. You can do it in one command:
    `cd TestingTools/Golang/Cards`.
5. Run `go test`. The first time you do this, you will see a messsage, or
series of messages from Go to initialize the directory and download the Ginkgo
and Gomega packages. The messages may not be perfectly clear, but in my  
experience, you should:

    a. Do a `cd` up to the top-level git-managed directory, "TestingTools": 
        `cd ../..`.

    b. Invoke `go mod init example.com` on the go test directory, like this:
        `go mod init example.com/Golang/Cards`

    c. Invoke `go mod tidy`. This will download `gomega` and `ginkgo`.

6. Re-enter the go test folder:

   `cd Golang/Cards`.
7. Run the test:

   `go test`
    
    Expected result: **`Pass`**.

If you like, you can learn more about the tests by running with higher 
verbosity on the test runner, using `go test -v`.

This verbosity setting pertains only to the golang test runner. It shows
you each test's start, end, and and the test-by-test status. Using this, you
can pinpoint where an error occurred, or give test context to other verbosity
or metatest data. If you include more than one `v`, the result is undefined. 
Doing so won't cause any great problem, but it will trigger a usage message, 
while failing the test.

There are also package-level flags that I've provided in the test package
`golangcards.go`. These are:
- (verbosity) `-v=1`
- `-meta_test` 

These flags pertain only to this particular test package. If you 
invoke either, or both of them,
you'll want the test-by-test context provided by the first -v, after `test`.

Without that context, you'll get a jumble of undifferentiated output. So, first
include the runner `-v` flag, and the runner flag `-args` to introduce the 
package-level flags. Then provide the package flags as desired. Examples:

    go test -v -args -v=1 

    go test -v -args -v=3

    go test -v -args -meta_test

    go test -v -args -v=1 -meta_test 
etc.
 
As you can see above, the package-level verbosity flag (the -v on the right)
takes an integer value. 
The higher the value, the more information is shared, up to a point. Currently
the level 3 is the maximum.

Including `-meta_test` will "test the test" by contriving failures in some 
selected test conditions. I "unsort" card orderings that should be sorted, and
sort some that should be unsorted. I also mess up the count of some cards in a
test condition, where they should appear exactly once in the deck. At the time
this is written, that's an exhaustive list of all that I do.

You'll **never** want to run with `-meta_test` in place during production use.
It's purely diagnostic, just to ensure that a desired error can be caught. 

### Python

1. Ensure that you are using Python 3.x. These instructions will not work with
Python 2.x.
2. (This step and further ones are to be determined).
