// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// This contract will not work, it's not designed for it to work but to test the parser in
// many different ways, types, modifiers, etc.

contract MyMappings {
    mapping(address=>uint) public simpleMapping;
    mapping(address=>mapping(address=>uint)) public doubleMapping;
    mapping(address=>mapping(address=>mapping(address=>uint))) public tripleMapping;
}
