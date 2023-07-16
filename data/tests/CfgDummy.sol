// SPDX-License-Identifier: MIT
pragma solidity ^0.8.5;

// Some additional comments that can be extracted

/** 
 * Multi line comments
 * are supported as well
*/

contract Dummy {
    uint256 public x;
    uint256 public y;

    constructor(uint256 _x, uint256 _y) {
        setDefaults();
        increment(10);
        uint256 z = 0;
        x = _x;
        y = _y;
    }

    function setDefaults() private {
        x = 1;
        y = 2;
    }

    function add() public view returns (uint256) {
        return x + y;
    }

    function increment(uint256 i) public {
        x += i;
        y += i;
    }
}
