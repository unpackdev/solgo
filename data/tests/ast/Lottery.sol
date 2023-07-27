// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IDummyContract {
    function dummyFunction() external returns (bool);
}

contract Lottery {
    uint256 public constant DUMMY_CONSTANT = 12345;
    
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
    event ExternalCallSuccessful();
    event ExternalCallFailed(string reason);

    // Define custom errors
    error InvalidState();
    error OwnerCannotParticipate();
    error NoValueProvided();
    error InvalidWinner();
    error InvalidPlayerAddress();
    error OnlyOwnerCanCall();

    modifier inState(LotteryState _state) {
        if (state != _state) {
            revert InvalidState();
        }
        _;
    }

    modifier notOwner() {
        if (msg.sender == owner()) {
            revert OwnerCannotParticipate();
        }
        _;
    }

    fallback() external payable { }
    receive() external payable { }

    constructor() {
        state = LotteryState.Accepting;
    }

    function join() public payable inState(LotteryState.Accepting) notOwner {
        if (msg.value == 0) {
            revert NoValueProvided();
        }

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

        if (winner == address(0)) {
            revert InvalidWinner();
        }

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

    // New function using for loop
    function checkAllPlayers() public view returns (bool) {
        for (uint i = 0; i < playerAddresses.length; i++) {
            if (players[playerAddresses[i]].addr == address(0)) {
                revert InvalidPlayerAddress();
            }
        }
        return true;
    }

    // New function using revert
    function requireOwner() public view {
        if (msg.sender != owner()) {
            revert OnlyOwnerCanCall();
        }
    }

    function callExternalFunction(address externalContractAddress) public {
        IDummyContract dummyContract = IDummyContract(externalContractAddress);

        try dummyContract.dummyFunction() {
            emit ExternalCallSuccessful();
        } catch (bytes memory /*lowLevelData*/) {
            emit ExternalCallFailed("External contract failed");
        }
    }

    function integerToString(uint _i) internal pure 
      returns (string memory) {
      
      if (_i == 0) {
         return "0";
      }
      uint j = _i;
      uint len;
      
      while (j != 0) {
         len++;
         j /= 10;
      }
      bytes memory bstr = new bytes(len);
      uint k = len - 1;
      
      do {                   // do while loop	
         bstr[k--] = bytes1(uint8(48 + _i % 10));
         _i /= 10;
      }
      while (_i != 0);
      return string(bstr);
   }
}