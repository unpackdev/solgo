// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Lottery {
    enum LotteryState { Accepting, Finished }
    struct Player {
        address addr;
        uint256 ticketCount;
    }
    mapping(address => Player) public players;
    address[] public playerAddresses;

    LotteryState public state;

    event PlayerJoined(address addr);
    event LotteryFinished(address winner);

    modifier inState(LotteryState _state) {
        require(state == _state, "Invalid state for this action");
        _;
    }

    modifier notOwner() {
        require(msg.sender != owner(), "The owner cannot participate");
        _;
    }

    fallback() external payable { }
    receive() external payable { }

    constructor() {
        state = LotteryState.Accepting;
    }

    function join() public payable inState(LotteryState.Accepting) notOwner {
        require(msg.value > 0, "No value provided");

        if (players[msg.sender].addr == address(0)) {
            players[msg.sender].addr = msg.sender;
            playerAddresses.push(msg.sender);
        }

        players[msg.sender].ticketCount += msg.value;

        emit PlayerJoined(msg.sender);
    }

    function finishLottery() public inState(LotteryState.Accepting) {
        state = LotteryState.Finished;

        uint256 index = uint256(block.timestamp) % playerAddresses.length;
        address winner = address(0);
        uint256 count = 0;
        
        while(count < playerAddresses.length) {
            if(index == count){
                winner = playerAddresses[count];
                break;
            }

            count++;
            
            // Odd numbers to continue loop
            if (count % 2 == 1) {
                continue;
            }
        }

        require(winner != address(0), "Winner is not valid");

        emit LotteryFinished(winner);

        uint256 balance = address(this).balance;
        payable(winner).transfer(balance);
    }

    function owner() public view returns (address) {
        return address(this);
    }

    function balance() public view returns (uint256) {
        return address(this).balance;
    }
}
