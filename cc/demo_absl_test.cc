#include "demo_absl.h"

#include "gtest/gtest.h"

namespace {

TEST(ComputeHelloTest, ReturnsCorrectGreeting) {
  EXPECT_EQ(compute_hello("World"), "Hello, World!");
}

TEST(ComputeHelloTest, HandlesEmptyString) {
  EXPECT_EQ(compute_hello(""), "Hello, !");
}

}  // namespace
