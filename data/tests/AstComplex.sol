// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library SafeMath {
    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }
}

interface IBaseContract {
    function baseFunction() external returns (uint);
}

contract BaseContract is IBaseContract {
    function baseFunction() public virtual override returns (uint) {
        return 1;
    }
}

contract TestContract is BaseContract {
    using SafeMath for uint;

    uint public stateVariable;
    uint private privateStateVariable;
    uint internal internalStateVariable;
    uint public constant constantStateVariable = 1;
    uint public immutable immutableStateVariable;
    address public owner;

    // More complex types
    uint[] public arrayStateVariable;
    mapping(address => uint) public mappingStateVariable;
    struct StructType {
        uint field1;
        address field2;
    }
    StructType public structStateVariable;

    event StateVariableChanged(uint oldValue, uint newValue);
    event OwnerChanged(address indexed oldOwner, address indexed newOwner);
    event CustomEvent(string message);
    
    enum CustomEnum { Option1, Option2, Option3 }
    
    error CustomError(string message);

    constructor(uint _stateVariable, uint _privateStateVariable) {
        stateVariable = _stateVariable;
        privateStateVariable = _privateStateVariable;
        owner = msg.sender;
        immutableStateVariable = block.timestamp;
    }

    function publicFunction(uint a, uint b) public returns (uint) {
        return a.add(b);
    }

    function privateFunction(uint a, uint b) private returns (uint) {
        uint result;
        assembly {
            result := sub(a, b)
        }
        return result;
    }

    function baseFunction() public override returns (uint) {
        return 2;
    }

    function functionWithModifier(uint a, uint b) public onlyOwner returns (uint) {
        return a.add(b);
    }

    function changeStateVariable(uint newValue) public {
        emit StateVariableChanged(stateVariable, newValue);
        stateVariable = newValue;
    }

    function changeOwner(address newOwner) public onlyOwner {
        emit OwnerChanged(owner, newOwner);
        owner = newOwner;
    }

    modifier onlyOwner {
        require(msg.sender == owner, "Not the owner");
        _;
    }
    
    fallback() external {
        revert("Fallback function called");
    }
    
    receive() external payable {
        emit CustomEvent("Receive function called");
    }
    
    function throwError() public pure {
        revert(CustomError("Error occurred"));
    }
    
    function getEnumValue(CustomEnum option) public pure returns (uint) {
        if (option == CustomEnum.Option1) {
            return 1;
        } else if (option == CustomEnum.Option2) {
            return 2;
        } else if (option == CustomEnum.Option3) {
            return 3;
        }
        
        return 0;
    }
}
