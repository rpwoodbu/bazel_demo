#include <iostream>

#include "demo_absl.h"
#include "pugixml.hpp"

int main() {
  // Do a little XML parsing, just to exercise this dependency.
  pugi::xml_document doc;
  pugi::xml_parse_result result = doc.load_string("<greet>Bazel World</greet>");
  if (!result) {
    return 1;
  }

  const auto greeting = doc.child("greet").child_value();
  std::cout << compute_hello(greeting) << std::endl;
  return 0;
}
