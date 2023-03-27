# TestingTools

This repository demonstrates my work in writing testing tools.

## How to Use

The API tests I've put here run against a [public cards API](https://deckofcardsapi.com/). 
This is a small, public API, enough to demonstrate a few tests.
Be sure that you have access to the Internet, and this API, by clicking the link. If you 
then see a bright green page titled **Deck of Cards**, then you are ready to try this out.

You can fork, or clone this repo locally, then descend into the appropriate language folder.
Run the tests as indicated under the language header in the [Languages](#languages) section.

## State of Progress

As of late March 2023, I've just started this repo. It will be a work in progress 
for a little while as I add more languages and testing protocols.

## Languages

I'm currently demonstrating only **[Golang](#golang)**, using:
- Golang's own Testing package.
- Ginkgo.
- Gomega.

I hope to eventually (soon) include **Python**, using:
- Pytest.
- Tavern. 

### Golang

1. Ensure that Golang v. 1.20.x (or later) is installed, as I wrote
these tests with that version.
2. In your local repo of TestingTools, descend into the `Golang` folder, then into the `Cards` subfolder.
3. In the `Cards` folder, run `go test`. 
Expected result: **`Pass`**.

If you like, you can run with higher verbosity on `go test`: `go test -v`, or `-vv` etc.
You can also edit and save deck_test.go to raise the internal verbosity setting there, to see more of what is happens
in the interaction with the Deck of Cards API:
1. Open the deck_test.go file in your favorite editor.
2. Find the line that says: `var verbosity = 0`.
3. Edit that line to raise the integer from 0 to some higher number: 1, 2, or 3. 
4. Save, then re-run the test. You'll see progressively more output with each higher verbosity.
5. Restore verbosity to zero, and save for production use. 
