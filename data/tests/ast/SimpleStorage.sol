// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;


import "./MathLib.sol";

// This contract uses the MathLib library for safe math operations.
contract SimpleStorage {
    using MathLib for uint;
    uint storedData;

    // Use the add function from the MathLib to safely increment storedData.
    function increment(uint x) public {
        storedData = storedData.add(x);
    }

    // Use the sub function from the MathLib to safely decrement storedData.
    function decrement(uint x) public {
        storedData = storedData.sub(x);
    }

    // Get the value of storedData.
    function get() public view returns (uint) {
        return storedData;
    }
}
