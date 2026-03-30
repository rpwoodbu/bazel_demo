import cowsay

def compute_hello(greeting: str) -> str:
   return cowsay.get_output_string("cow", "Hello, {}!".format(greeting))
