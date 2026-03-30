package com.example;

import com.google.common.base.Joiner;

class DemoGuava {
  // Takes a list of words and produces a greeting.
  public static String computeHello(String... words) {
    return "Hello, " + Joiner.on(" ").join(words) + "!";
  }
}
