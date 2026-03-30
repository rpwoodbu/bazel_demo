import unittest
import demo_cowsay

class TestDemoCowsay(unittest.TestCase):

    # Formatted message is present in the output.
    def test_compute_hello_contains_greeting(self):
        greeting = "Bazel User"
        output = demo_cowsay.compute_hello(greeting)

        self.assertIn("Hello, Bazel User!", output)

    # Output contains a cow.
    def test_compute_hello_contains_cow(self):
        greeting = "doesn't matter"
        output = demo_cowsay.compute_hello(greeting)
        
        # These are her ears.
        self.assertIn("^__^", output)

    # Works with the empty string.
    def test_compute_hello_empty_string(self):
        greeting = ""
        output = demo_cowsay.compute_hello(greeting)

        self.assertIn("Hello, !", output)


if __name__ == "__main__":
    unittest.main()
