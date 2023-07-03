// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";

contract MyToken is Initializable, ERC20Upgradeable, AccessControlUpgradeable, PausableUpgradeable {
    using SafeERC20Upgradeable for IERC20Upgradeable;
    
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    uint256 public subscriptionAmount;
    mapping(address => uint256) public subscriptionBalance;
    mapping(address => bool) public isSubscribed;

    // Rewards
    mapping(address => uint256) public rewards;
    address[] public rewardedUsers;

    function initialize(string memory name, string memory symbol, uint256 initialSupply, uint256 _subscriptionAmount) public initializer {
        __ERC20_init(name, symbol);
        __AccessControl_init();
        __Pausable_init();

        subscriptionAmount = _subscriptionAmount;

        _mint(msg.sender, initialSupply);
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(MINTER_ROLE, msg.sender);
    }

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }


    // Allows users to purchase subscription using the token
    function purchaseSubscription() external whenNotPaused {
        require(subscriptionBalance[msg.sender] < subscriptionAmount, "Already subscribed");

        transfer(address(this), subscriptionAmount);
        subscriptionBalance[msg.sender] += subscriptionAmount;
        isSubscribed[msg.sender] = true;

        emit SubscriptionPurchased(msg.sender, subscriptionAmount);
    }
    
    // Allows users to cancel subscription and receive a refund
    function cancelSubscription() external whenNotPaused {
        require(isSubscribed[msg.sender], "Not subscribed");
        uint256 refundAmount = subscriptionBalance[msg.sender];
        require(refundAmount > 0, "No subscription balance to refund");
        
        transfer(msg.sender, refundAmount);
        subscriptionBalance[msg.sender] = 0;
        isSubscribed[msg.sender] = false;
        
        emit SubscriptionCanceled(msg.sender, refundAmount);
    }
    
    // Gets the subscription status of a user
    function getSubscriptionStatus(address user) external view returns (bool) {
        return isSubscribed[user];
    }
    
    // Allows the contract owner to reward users with tokens
    function reward(address user, uint256 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        transfer(user, amount);
        rewards[user] += amount;
        
        // If the user hasn't been rewarded before, add them to the list of rewarded users
        if (rewards[user] == amount) {
            rewardedUsers.push(user);
        }
        
        emit UserRewarded(user, amount);
    }
    
    // Gets the total amount of rewards for a user
    function getRewards(address user) external view returns (uint256) {
        return rewards[user];
    }
    
    // Gets the list of all rewarded users
    function getRewardedUsers() external view returns (address[] memory) {
        return rewardedUsers;
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);

        if (isSubscribed[from] && subscriptionBalance[from] >= amount) {
            subscriptionBalance[from] -= amount;
            subscriptionBalance[to] += amount;
        }
    }

    // Allows the contract owner to update the subscription amount
    function updateSubscriptionAmount(uint256 newAmount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        subscriptionAmount = newAmount;
    }

    // Allows minters to mint new tokens
    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    // Events
    event SubscriptionPurchased(address indexed user, uint256 amount);
    event SubscriptionCanceled(address indexed user, uint256 amount);
    event UserRewarded(address indexed user, uint256 amount);
}
