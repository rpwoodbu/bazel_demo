#include "demo_absl.h"

#include <string>

#include "absl/strings/str_format.h"

std::string compute_hello(const std::string& name) {
  return absl::StrFormat("Hello, %s!", name);
}
