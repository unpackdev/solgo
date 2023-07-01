// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EnumContract {
    // Define an enum
    enum State { Waiting, Ready, Active }

    // Declare a state variable of type State
    State public state;

    // Initialize the state
    constructor() {
        state = State.Waiting;
    }

    // Function to check if state is Waiting
    function isWaiting() public view returns(bool) {
        return state == State.Waiting;
    }

    // Function to check if state is Ready
    function isReady() public view returns(bool) {
        return state == State.Ready;
    }

    // Function to check if state is Active
    function isActive() public view returns(bool) {
        return state == State.Active;
    }

    // Function to set state to Ready
    function makeReady() public {
        state = State.Ready;
    }

    // Function to set state to Active
    function makeActive() public {
        state = State.Active;
    }

    // Function to reset state to Waiting
    function reset() public {
        state = State.Waiting;
    }
}
