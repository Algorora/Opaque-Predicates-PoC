# Opaque-Predicates-PoC
This project is a proof-of-concept (PoC) demonstrating the use of opaque predicates in Go. Opaque predicates are conditions that are constructed to always evaluate to true or false, but their true nature isn't immediately obvious. This technique is used to obfuscate code and make static and dynamic analysis more challenging.

# Purpose
The purpose of this PoC is to illustrate how opaque predicates can be used to obscure the logic of a program. The provided Go code includes functions with complex conditions that always evaluate to true, which is used to demonstrate how such techniques can add complexity to code analysis.

# Code Description
OpaquePredicate1: A function that contains a condition designed to always be true.
OpaquePredicate2: Another function with a condition that also always evaluates to true.
PerformOperation: Uses the opaque predicates to determine the execution path. The function will always execute the block that multiplies the input by 2, demonstrating the use of opaque predicates to obscure the program flow.

# How It Works
The PerformOperation function will always return twice the input value because the opaque predicates used (OpaquePredicate1 and OpaquePredicate2) are constructed in a way that ensures their conditions are always true.

# Prerequisites
- Go 1.18 or later (or any version compatible with Go modules).
  
# Running the Code
Clone the Repository

'''sh
git clone [Algorora/Opaque-Predicates-PoC](https://github.com/Algorora/Opaque-Predicates-PoC.git)
cd opaque-predicates-poc
'''

-Run the Program

'''sh
go run main.go
'''

-Expected Output

'''sh
Result: 20
'''
This output demonstrates that the PerformOperation function executes the block that multiplies the input by 2.

# Example
If you run the program with the default value of x = 10, the output will be Result: 20, as the opaque predicates ensure that the code path performing the multiplication is always executed.

# Disclaimer
This code is provided for educational and demonstration purposes only. It is intended to illustrate the concept of opaque predicates and how they can be used to obfuscate code. The use of obfuscation techniques should comply with legal and ethical guidelines.
