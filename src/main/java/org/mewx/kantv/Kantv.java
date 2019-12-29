package org.mewx.kantv;

import picocli.CommandLine;
import picocli.CommandLine.*;

@CommandLine.Command(
        name = "kantv",
        description = "Kantv Downloader CLI",
        mixinStandardHelpOptions = true,
        version = "0.1.0"
)
public class Kantv implements Runnable {
    @Option(names = {"-v", "--verbose"},
            description = "Verbose mode. Helpful for troubleshooting. Multiple -v options increase the verbosity.")
    private boolean[] verbose = new boolean[0];

    public static void main(String[] args) {
        // By implementing Runnable or Callable, parsing, error handling and handling user
        // requests for usage help or version help can be done with one line of code.
        int exitCode = new CommandLine(new Kantv()).execute(args);
        System.exit(exitCode);
    }

    @Override
    public void run() {
        // do nothing
    }
}
