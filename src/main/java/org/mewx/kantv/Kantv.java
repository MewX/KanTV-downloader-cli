package org.mewx.kantv;

import org.mewx.kantv.command.KantvBase;
import picocli.CommandLine;

public class Kantv {
    public static void main(String[] args) {
        // By implementing Runnable or Callable, parsing, error handling and handling user
        // requests for usage help or version help can be done with one line of code.
        int exitCode = new CommandLine(new KantvBase()).execute(args);
        System.exit(exitCode);
    }
}
