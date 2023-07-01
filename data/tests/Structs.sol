// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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

    struct StructWithArray {
        uint256 one;
        uint256[] array;
    }

    struct StructWithMapping {
        uint256 one;
        mapping(address => uint256) map;
    }

    struct StructWithNestedArray {
        uint256 one;
        ClasicStruct[] structArray;
    }

    struct StructWithNestedMapping {
        uint256 one;
        mapping(address => ClasicStruct) structMap;
    }

    function nestedStructExample(ClasicStruct memory structOne, NestedStruct memory structTwo, uint Integer) public pure returns (NestedStruct memory structReturn) {}

    function structWithArrayExample(StructWithArray memory structOne) public pure {}

    function structWithNestedArrayExample(StructWithNestedArray memory structOne) public pure {}
}
