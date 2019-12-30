package org.mewx.kantv.command;

import picocli.CommandLine.*;

@Command(name = "test", description = "Testing command for Picocli.")
public class Test implements Runnable {

    @ParentCommand
    private KantvBase base;

    @Option(names = {"--test_option"}, description = "The test option as a String.")
    private String testOption = "";

    @Override
    public void run() {
        // Sample command: java -jar bazel-bin/kantv_deploy.jar --cookies=123 test --test_option=432
        System.out.println("Testing command.");
        if (!testOption.isEmpty()) {
            System.out.println("Command with option: " + testOption);
        } else {
            System.out.println("Use --help option for sub-commands to see usages.");
        }
        if (!base.cookies.isEmpty()) {
            System.out.println("Parent parameter: " + base.cookies);
        }
    }
}
