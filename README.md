# Go Program Verification Survey

## Report

[Link to final report]()

## Slides & Video

[Link to final presentation slides]()
[Link to final presentation video]()

## Go verification tool benchmark set

This project is adapted from [benchmark-maker](https://github.com/nicolasdilley/benchmark-maker)

This is a tool which synthesize Go programs from three sets of code
elements and verify the resulting set with three different static checkers
Gomela, GCatch and Godel2. These programs are made from instantiating
contexts, programs with holes, with 3 different code elements:

1.  The declaration of a concurrency primitive (such as a channel, waitgroup or mutex)
2.  A code snippet that contains a bug related to concurrency primitive from 1.
3.  A potential bound.

### Installation

Mac Sillicon likely does not work, Linux system is prefered for reproducing the result.

To build the tool the Go ecosystem is required, also need to build and configure verification tools Godel2, GCatch, and Gomela.

To install Go, follow the instructions [here](https://go.dev/doc/install).

To install Godel2, follow instructions [here](https://github.com/jujuyuki/godel2).

To install Gomela, follow instructions [here](https://github.com/nicolasdilley/Gomela).

To install GCatch, follow instructions [here](https://github.com/system-pclub/GCatch/tree/master/GCatch).

If you already have everything installed, simply clone this repo.

`git clone https://github.com/nicolasdilley/benchmark-maker`

and build the tool.

`go get && go build`

### Usage

To synthesize programs, Benchmark Maker requires two folders which contain two
different code elements, the first containing the contexts and the second
which contains the code snippets. The repo contains example snippets and contexts you can direcly use.

#### Contexts

A context is a program which contains holes for the declaration of concurrency
primitives (CP) , code snippets (CS) and a potential bound (bound).

The contexts can declare what type of concurrency primitive they allow by
specifying as a comment declared a the top of the context:

1. `// type = ALL` If they allow all concurrency primitives (channel, mutex and waitgroups)
2. `// type = CH` If they only allow channels.

The value given to the bound in the program is also given as a comment as
follow:

`// bounds = 10 000 ,len(os.Args)` which will instantiate
programs with bounds of `10 000` and `len(os.Args)`for example.

A set of contexts can be found in the `contexts` folder.

#### Code snippets

The only two requirements for code snippets are to declared which concurrency
primitive they use and to use specific names for the variables of the
concurrency primitives. This is achieved by commenting at the top:

1. `// type = CH` for channels snippets. The channel is available via the variable `ch`.
2. `// type = MU` for sync.Mutex snippets. The mutex variable is `mu`.
3. `// type = RWMU` for sync.RWMutex snippets. The rwmutex variable is `mu`.
4. `// type = WG` for sync.Waitgroup snippets. The waitgroup variable is `wg`.

note: Godel2 and GCatch does not support Waitgroup yet...

#### Generating and verifying a set of synthetic programs

The tool expects as argument the folder containing the contexts followed by
the folder containing the code snippets.

To instantiate and verify the contexts in `contexts` with
the code snippets in `snippets`, run:

`./benchmark-maker contexts snippets`

The console will output the result of verifying each code snippets into all
contexts with each tool (given as a list `verifiers` in `verify.go`).
The output can also be found in `text.log`.

The output for each code snippets is formatted as follow. The first column
shows the name of the contexts and the bound used. The rest of the columns
show the results and the time taken by each tool in the order they were given
in the list `verifiers` in `verify.go`.

The verification results are:

1. `\cmark` for a bug was found.
2. `\xmark` for no bug reported.
3. `\crash` which reports that the tool crashed.
4. `\nosupport` which reports that the tool reported that it did not support the benchmark.
