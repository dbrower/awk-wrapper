# awk-wrapper

A quick utility to wrap the AWK command.
It converts input files from Mac line encodings into unix encodings.
It then passes it to AWK and then converts the output into dos line encoding.

The usage is

    ./awk-wrapper <awk command file> <input files>...

The output files are prefixed with `sierra-`, so example usage is

    $ ./awk-wrapper my.awk Period\ 1.txt period-2.txt
    Using awk file my.awk
    Processing Geriod 1.txt
          into sierra-Period 1.txt
    Processing period-2.txt
          into sierra-period-2.txt
