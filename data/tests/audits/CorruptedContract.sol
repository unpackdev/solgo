// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CorruptedContract {
    mapping(address => uint256) public balances;

    function deposit() external payable {
        require(msg.value > 0, "Deposit amount should be greater than 0");
        balances[msg.sender] += msg.value;
        bug();
    }

    function withdraw() external {
        uint256 amount = balances[msg.sender];
        require(amount > 0, "Insufficient balance");

        // This call can be exploited for reentrancy
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "Transfer failed");

        balances[msg.sender] = 0;
    }
}
