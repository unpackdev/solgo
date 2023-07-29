// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

// This is a library for performing safe math operations.
library MathLib {
    // Safely add two numbers.
    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        require(c >= a, "Addition overflow");

        return c;
    }

    // Safely subtract two numbers.
    function sub(uint a, uint b) internal pure returns (uint) {
        require(b <= a, "Subtraction underflow");
        uint c = a - b;

        return c;
    }

    // Safely multiply two numbers.
    function mul(uint a, uint b) internal pure returns (uint) {
        if (a == 0) {
            return 0;
        }

        uint c = a * b;
        require(c / a == b, "Multiplication overflow");

        return c;
    }

    // Safely divide two numbers.
    function div(uint a, uint b) internal pure returns (uint) {
        require(b > 0, "Division by zero");
        uint c = a / b;

        return c;
    }
}
