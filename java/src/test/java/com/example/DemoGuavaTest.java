package com.example;

import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class DemoGuavaTest {

  @Test
  public void computeHello_withMultipleWords_returnsFormattedString() {
    String result = DemoGuava.computeHello("Java", "World");
    assertEquals("Hello, Java World!", result);
  }

  @Test
  public void computeHello_withSingleWord_returnsFormattedString() {
    String result = DemoGuava.computeHello("Bazel");
    assertEquals("Hello, Bazel!", result);
  }

  @Test
  public void computeHello_withNoWords_returnsHelloPrefixOnly() {
    String result = DemoGuava.computeHello();
    assertEquals("Hello, !", result);
  }
}
