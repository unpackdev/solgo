// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./IERC20.sol";
import "./SafeMath.sol";

contract TokenSale {
    using SafeMath for uint256;

    IERC20 private token;
    address private owner;
    uint256 private tokenPrice;

    event TokensPurchased(address buyer, uint256 amount);

    constructor(address _tokenAddress, uint256 _tokenPrice) {
        token = IERC20(_tokenAddress);
        owner = msg.sender;
        tokenPrice = _tokenPrice;
    }

    function buyTokens(uint256 _amount) external {
        uint256 totalPrice = _amount.mul(tokenPrice);
        token.transferFrom(owner, msg.sender, _amount);
        emit TokensPurchased(msg.sender, _amount);
    }
}