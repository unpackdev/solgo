// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Some additional comments that can be extracted

/** 
 * Multi line comments
 * are supported as well
*/

contract Dummy {
    uint256 public x;
    uint256 public y;

    constructor(uint256 _x, uint256 _y) {
        x = _x;
        y = _y;
    }
}
