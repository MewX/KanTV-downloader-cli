package org.mewx.kantv.command;

import picocli.CommandLine.*;

@Command(
        name = "kantv",
        description = "Kantv Downloader CLI",
        mixinStandardHelpOptions = true,
        version = "0.1.0",
        subcommands = {
                Test.class
        }
)
public class KantvBase implements Runnable {

    @Option(names = {"-v", "--verbose"},
            description = "Verbose mode. Helpful for troubleshooting. Multiple -v options increase the verbosity.")
    boolean[] verbose = new boolean[0];

    @Option(names = {"--cookies"}, description = "Cookies used in every network request.")
    String cookies = "";

    @Override
    public void run() {
        System.out.println("Use --help option to see usages.");
    }
}
