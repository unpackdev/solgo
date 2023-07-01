// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// This contract will not work, it's not designed for it to work but to test the parser in
// many different ways, types, modifiers, etc.

contract MyStructs {
    struct NestedStruct {
        uint256 one;
        uint256 two;
        ClasicStruct myStruct;
    }

    struct ClasicStruct {
        uint256 one;
        uint256 two;
    }


    function nestedStructExample(ClasicStruct memory structOne, NestedStruct memory structTwo, uint Integer) public pure returns (NestedStruct memory structReturn) {}
}
