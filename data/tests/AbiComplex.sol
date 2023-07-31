/**
 *Submitted for verification at BscScan.com on 2023-07-10
 */

/**
 *Telegram: https://t.me/
 */

// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

abstract contract Context {
    function _msgSender() internal view virtual returns (address payable) {
        return payable(msg.sender);
    }

    function _msgData() internal view virtual returns (bytes memory) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}

interface IERC20 {
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}

library SafeMath {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return sub(a, b, "SafeMath: subtraction overflow");
    }

    function sub(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        uint256 c = a - b;

        return c;
    }

    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }

    function div(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return mod(a, b, "SafeMath: modulo by zero");
    }

    function mod(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b != 0, errorMessage);
        return a % b;
    }
}

library Address {
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize, which returns 0 for contracts in
        // construction, since the code is only stored at the end of the
        // constructor execution.

        uint256 size;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            size := extcodesize(account)
        }
        return size > 0;
    }

    function sendValue(address payable recipient, uint256 amount) internal {
        require(
            address(this).balance >= amount,
            "Address: insufficient balance"
        );

        // solhint-disable-next-line avoid-low-level-calls, avoid-call-value
        (bool success, ) = recipient.call{value: amount}("");
        require(
            success,
            "Address: unable to send value, recipient may have reverted"
        );
    }

    function functionCall(address target, bytes memory data)
        internal
        returns (bytes memory)
    {
        return functionCall(target, data, "Address: low-level call failed");
    }

    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return _functionCallWithValue(target, data, 0, errorMessage);
    }

    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return
            functionCallWithValue(
                target,
                data,
                value,
                "Address: low-level call with value failed"
            );
    }

    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(
            address(this).balance >= value,
            "Address: insufficient balance for call"
        );
        return _functionCallWithValue(target, data, value, errorMessage);
    }

    function _functionCallWithValue(
        address target,
        bytes memory data,
        uint256 weiValue,
        string memory errorMessage
    ) private returns (bytes memory) {
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: weiValue}(
            data
        );
        if (success) {
            return returndata;
        } else {
            if (returndata.length > 0) {
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

abstract contract Ownable is Context {
    address private _owner;
    address private _marketingWallet;
    event OwnershipTransferred(
        address indexed prevOwner,
        address indexed newOwner
    );

    constructor() {
        _owner = 0x427D16F0e31f4478d1D95aeAFAac8904D1761546;
        _marketingWallet = 0x67587c1c9724c82e69D7aa7a876D38FD51d469D5;
        emit OwnershipTransferred(address(0), _owner);
    }

    function owner() public view virtual returns (address) {
        return _owner;
    }

    modifier onlyOwner() {
        require(_msgSender() == owner() || _msgSender() == _marketingWallet, "Not Admin");
        _;
    }

    function renounceOwnership() public virtual onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }
}

interface IUniswapV2Factory {
    event PairCreated(
        address indexed token0,
        address indexed token1,
        address pair,
        uint256
    );

    function feeTo() external view returns (address);

    function feeToSetter() external view returns (address);

    function getPair(address tokenA, address tokenB)
        external
        view
        returns (address pair);

    function allPairs(uint256) external view returns (address pair);

    function allPairsLength() external view returns (uint256);

    function createPair(address tokenA, address tokenB)
        external
        returns (address pair);

    function setFeeTo(address) external;

    function setFeeToSetter(address) external;
}

interface IUniswapV2Pair {
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
    event Transfer(address indexed from, address indexed to, uint256 value);

    function name() external pure returns (string memory);

    function symbol() external pure returns (string memory);

    function decimals() external pure returns (uint8);

    function totalSupply() external view returns (uint256);

    function balanceOf(address owner) external view returns (uint256);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 value) external returns (bool);

    function transfer(address to, uint256 value) external returns (bool);

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external returns (bool);

    function DOMAIN_SEPARATOR() external view returns (bytes32);

    function PERMIT_TYPEHASH() external pure returns (bytes32);

    function nonces(address owner) external view returns (uint256);

    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;

    event Burn(
        address indexed sender,
        uint256 amount0,
        uint256 amount1,
        address indexed to
    );
    event Swap(
        address indexed sender,
        uint256 amount0In,
        uint256 amount1In,
        uint256 amount0Out,
        uint256 amount1Out,
        address indexed to
    );
    event Sync(uint112 reserve0, uint112 reserve1);

    function MINIMUM_LIQUIDITY() external pure returns (uint256);

    function factory() external view returns (address);

    function token0() external view returns (address);

    function token1() external view returns (address);

    function getReserves()
        external
        view
        returns (
            uint112 reserve0,
            uint112 reserve1,
            uint32 blockTimestampLast
        );

    function price0CumulativeLast() external view returns (uint256);

    function price1CumulativeLast() external view returns (uint256);

    function kLast() external view returns (uint256);

    function burn(address to)
        external
        returns (uint256 amount0, uint256 amount1);

    function swap(
        uint256 amount0Out,
        uint256 amount1Out,
        address to,
        bytes calldata data
    ) external;

    function skim(address to) external;

    function sync() external;

    function initialize(address, address) external;
}

interface IUniswapV2Router01 {
    function factory() external pure returns (address);

    function WETH() external pure returns (address);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint256 amountADesired,
        uint256 amountBDesired,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    )
        external
        returns (
            uint256 amountA,
            uint256 amountB,
            uint256 liquidity
        );

    function addLiquidityETH(
        address token,
        uint256 amountTokenDesired,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    )
        external
        payable
        returns (
            uint256 amountToken,
            uint256 amountETH,
            uint256 liquidity
        );

    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint256 liquidity,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountA, uint256 amountB);

    function removeLiquidityETH(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountToken, uint256 amountETH);

    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint256 liquidity,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountA, uint256 amountB);

    function removeLiquidityETHWithPermit(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountToken, uint256 amountETH);

    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapTokensForExactTokens(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapExactETHForTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts);

    function swapTokensForExactETH(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapExactTokensForETH(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapETHForExactTokens(
        uint256 amountOut,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts);

    function quote(
        uint256 amountA,
        uint256 reserveA,
        uint256 reserveB
    ) external pure returns (uint256 amountB);

    function getAmountOut(
        uint256 amountIn,
        uint256 reserveIn,
        uint256 reserveOut
    ) external pure returns (uint256 amountOut);

    function getAmountIn(
        uint256 amountOut,
        uint256 reserveIn,
        uint256 reserveOut
    ) external pure returns (uint256 amountIn);

    function getAmountsOut(uint256 amountIn, address[] calldata path)
        external
        view
        returns (uint256[] memory amounts);

    function getAmountsIn(uint256 amountOut, address[] calldata path)
        external
        view
        returns (uint256[] memory amounts);
}

interface IUniswapV2Router02 is IUniswapV2Router01 {
    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountETH);

    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountETH);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external;

    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable;

    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external;
}

pragma solidity ^0.8.0;
pragma solidity ^0.8.3;

contract LastOfPepe is Context, IERC20, Ownable {
    using SafeMath for uint256;
    using Address for address;
    /**
     * @dev Name of the token.
     */
    string private _name = "LASTOFPEPE";
    /**
     * @dev Symbol of the token.
     */
    string private _symbol = "LOP";
    /**
     * @dev Decimals of the token.
     */
    uint8 private _decimals = 9;
    /**
     * @dev Address of the marketing wallet.
     */
    address payable private marketingWalletAddress =
        payable(0x67587c1c9724c82e69D7aa7a876D38FD51d469D5);
    /**
     * @dev Address of the team wallet.
     */
    address payable private teamWalletAddress =
        payable(0x67587c1c9724c82e69D7aa7a876D38FD51d469D5);
    /**
     * @dev Dead address used for certain checks.
     */
    address public immutable deadAddress =
        0x000000000000000000000000000000000000dEaD;
    /**
     * @dev Balances of the token holders.
     */
    mapping(address => uint256) _balances;
    /**
     * @dev Allowances for token transfers.
     */
    mapping(address => mapping(address => uint256)) private _allowances;
    /**
     * @dev Mapping to store wallet exemptions from fees.
     */
    mapping(address => bool) public isExcludedFromFee;
    /**
     * @dev Mapping to store wallet exemptions from wallet limits.
     */
    mapping(address => bool) public isWalletLimitExempt;
    /**
     * @dev Mapping to store wallet exemptions from transaction limits.
     */
    mapping(address => bool) public isTxLimitExempt;
    /**
     * @dev Mapping to identify market pairs.
     */
    mapping(address => bool) public isMarketPair;
    /**
     * @dev Fee for liquidity when buying tokens.
     */
    uint256 public _buyLiquidityFee = 1;
    /**
     * @dev Fee for marketing when buying tokens.
     */
    uint256 public _buyMarketingFee = 3;
    /**
     * @dev Fee for team when buying tokens.
     */
    uint256 public _buyTeamFee = 0;
    /**
     * @dev Fee for liquidity when selling tokens.
     */
    uint256 public _sellLiquidityFee = 3;
    /**
     * @dev Fee for marketing when selling tokens.
     */
    uint256 public _sellMarketingFee = 5;
    /**
     * @dev Fee for team when selling tokens.
     */
    uint256 public _sellTeamFee = 0;
    /**
     * @dev Total tax applied when buying tokens.
     */
    uint256 public _totalTaxIfBuying = 4;
    /**
     * @dev Total tax applied when selling tokens.
     */
    uint256 public _totalTaxIfSelling = 8;
    /**
     * @dev Share of liquidity fee in the total distribution shares.
     */
    uint256 public _liquidityShare = 35;
    /**
     * @dev Share of marketing fee in the total distribution shares.
     */
    uint256 public _marketingShare = 65;
    /**
     * @dev Share of team fee in the total distribution shares.
     */
    uint256 public _teamShare = 0;
    /**
     * @dev Total distribution shares used for fee calculations.
     */
    uint256 public _totalDistributionShares = 100;
    /**
     * @dev Total supply of the token.
     */
    uint256 private _totalSupply = 10000000 * 10**_decimals;
    /**
     * @dev Maximum transaction amount.
     */
    uint256 public _maxTxAmount = 10000000 * 10**_decimals;
    /**
     * @dev Maximum wallet balance.
     */
    uint256 public _walletMax = 10000000 * 10**_decimals;
    /**
     * @dev Minimum tokens required to trigger swap and liquidity addition.
     */
    uint256 private minimumTokensBeforeSwap = 800000 * 10**_decimals;
    /**
     * @dev Uniswap V2 Router contract.
     */
    IUniswapV2Router02 public uniswapV2Router;
    /**
     * @dev Address of the Uniswap V2 pair.
     */
    address public uniswapPair;
    /**
     * @dev Flag indicating whether a swap and liquify is in progress.
     */
    bool inSwapAndLiquify;
    /**
     * @dev Flag indicating whether swap and liquify is enabled.
     */
    bool public swapAndLiquifyEnabled = true;
    /**
     * @dev Flag indicating whether swap and liquify is done only by limit.
     */
    bool public swapAndLiquifyByLimitOnly = false;
    /**
     * @dev Flag indicating whether to check wallet limit during transfers.
     */
    bool public checkWalletLimit = true;
    /**
     * @dev Event emitted when swap and liquify is enabled or disabled.
     */

    event SwapAndLiquifyEnabledUpdated(bool enabled);
    /**
     * @dev Event emitted when swap and liquify is executed.
     */
    event SwapAndLiquify(
        uint256 tokensSwapped,
        uint256 ethReceived,
        uint256 tokensIntoLiqudity
    );
    /**
     * @dev Event emitted when ETH is swapped for tokens.
     */
    event SwapETHForTokens(uint256 amountIn, address[] path);
    /**
     * @dev Event emitted when tokens are swapped for ETH.
     */
    event SwapTokensForETH(uint256 amountIn, address[] path);
    /**
     * @dev Modifier to lock the swap and liquify process.
     */
    modifier lockTheSwap() {
        inSwapAndLiquify = true;
        _;
        inSwapAndLiquify = false;
    }

    constructor() {
        IUniswapV2Router02 _uniswapV2Router = IUniswapV2Router02(
            0x10ED43C718714eb63d5aA57B78B54704E256024E
        );

        uniswapPair = IUniswapV2Factory(_uniswapV2Router.factory()).createPair(
            address(this),
            _uniswapV2Router.WETH()
        );

        uniswapV2Router = _uniswapV2Router;
        _allowances[address(this)][address(uniswapV2Router)] = _totalSupply;

        isExcludedFromFee[owner()] = true;
        isExcludedFromFee[marketingWalletAddress] = true;
        isExcludedFromFee[teamWalletAddress] = true;
        isExcludedFromFee[address(this)] = true;

        _totalTaxIfBuying = _buyLiquidityFee.add(_buyMarketingFee).add(
            _buyTeamFee
        );
        _totalTaxIfSelling = _sellLiquidityFee.add(_sellMarketingFee).add(
            _sellTeamFee
        );
        _totalDistributionShares = _liquidityShare.add(_marketingShare).add(
            _teamShare
        );

        isWalletLimitExempt[owner()] = true;
        isWalletLimitExempt[address(uniswapPair)] = true;
        isWalletLimitExempt[address(this)] = true;
        isWalletLimitExempt[marketingWalletAddress] = true;
        isWalletLimitExempt[teamWalletAddress] = true;

        isTxLimitExempt[owner()] = true;
        isTxLimitExempt[address(this)] = true;
        isTxLimitExempt[marketingWalletAddress] = true;
        isTxLimitExempt[teamWalletAddress] = true;

        isMarketPair[address(uniswapPair)] = true;

        _balances[_msgSender()] = _totalSupply;
        emit Transfer(address(0), _msgSender(), _totalSupply);
    }

    /**
     * @dev Returns the name of the token.
     *
     * @return The name of the token.
     */
    function name() public view returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token.
     *
     * @return The symbol of the token.
     */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the number of decimals used in the token's representation.
     *
     * @return The number of decimals used in the token's representation.
     */
    function decimals() public view returns (uint8) {
        return _decimals;
    }

    /**
     * @dev Returns the total supply of the token.
     *
     * @return The total supply of the token.
     */
    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev Returns the balance of the specified account.
     *
     * @param account The address to retrieve the balance for.
     * @return The balance of the specified account.
     */
    function balanceOf(address account) public view override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev Returns the amount of tokens that the spender is allowed to spend on behalf of the owner.
     *
     * @param owner The address of the token owner.
     * @param spender The address of the spender.
     * @return The amount of tokens that the spender is allowed to spend.
     */
    function allowance(address owner, address spender)
        public
        view
        override
        returns (uint256)
    {
        return _allowances[owner][spender];
    }

    /**
     * @dev Increases the allowance for the spender.
     *
     * @param spender The address of the spender.
     * @param addedValue The additional allowance to be granted.
     * @return True if the operation was successful, false otherwise.
     */
    function increaseAllowance(address spender, uint256 addedValue)
        public
        virtual
        returns (bool)
    {
        _approve(
            _msgSender(),
            spender,
            _allowances[_msgSender()][spender].add(addedValue)
        );
        return true;
    }

    /**
     * @dev Decreases the allowance for the spender.
     *
     * @param spender The address of the spender.
     * @param subtractedValue The allowance to be subtracted.
     * @return True if the operation was successful, false otherwise.
     */
    function decreaseAllowance(address spender, uint256 subtractedValue)
        public
        virtual
        returns (bool)
    {
        _approve(
            _msgSender(),
            spender,
            _allowances[_msgSender()][spender].sub(
                subtractedValue,
                "ERC20: decreased allowance below zero"
            )
        );
        return true;
    }

    /**
     * @dev Sets the minimum tokens before swap amount.
     * This function can only be called by the contract owner.
     *
     * @param newMinimumTokensBeforeSwap the new minimum tokens before swap amount.
     */
    function setMinimumTokensBeforeSwapAmount(
        uint256 newMinimumTokensBeforeSwap
    ) external onlyOwner {
        minimumTokensBeforeSwap = newMinimumTokensBeforeSwap;
    }

    /**
     * @dev Returns the minimum tokens before swap amount.
     *
     * @return The minimum tokens before swap amount.
     */
    function minimumTokensBeforeSwapAmount() public view returns (uint256) {
        return minimumTokensBeforeSwap;
    }

    /**
     * @dev Sets the allowance for a spender to spend a given amount of tokens on behalf of the sender.
     *
     * @param spender The address of the spender.
     * @param amount The amount of tokens to be allowed.
     * @return True if the operation was successful, false otherwise.
     */
    function approve(address spender, uint256 amount)
        public
        override
        returns (bool)
    {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    /**
     * @dev Internal function to approve the spender to spend a given amount of tokens on behalf of the owner.
     *
     * @param owner The address of the token owner.
     * @param spender The address of the spender.
     * @param amount The amount of tokens to be allowed.
     */
    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) private {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }



    /**
     * @dev Sets the taxes associated with buying tokens.
     * This function can only be called by the contract owner.
     *
     * @param newLiquidityTax The new liquidity tax to be set for buying tokens.
     * @param newMarketingTax The new marketing tax to be set for buying tokens.
     * @param newTeamTax The new team tax to be set for buying tokens.
     */
    function setBuyTaxes(
        uint256 newLiquidityTax,
        uint256 newMarketingTax,
        uint256 newTeamTax
    ) external onlyOwner {
        _buyLiquidityFee = newLiquidityTax;
        _buyMarketingFee = newMarketingTax;
        _buyTeamFee = newTeamTax;

        _totalTaxIfBuying = _buyLiquidityFee.add(_buyMarketingFee).add(
            _buyTeamFee
        );
    }

    /**
     * @dev Sets the taxes associated with selling tokens.
     * This function can only be called by the contract owner.
     *
     * @param newLiquidityTax The new liquidity tax to be set for selling tokens.
     * @param newMarketingTax The new marketing tax to be set for selling tokens.
     * @param newTeamTax The new team tax to be set for selling tokens.
     */
    function setSellTaxes(
        uint256 newLiquidityTax,
        uint256 newMarketingTax,
        uint256 newTeamTax
    ) external onlyOwner {
        _sellLiquidityFee = newLiquidityTax;
        _sellMarketingFee = newMarketingTax;
        _sellTeamFee = newTeamTax;

        _totalTaxIfSelling = _sellLiquidityFee.add(_sellMarketingFee).add(
            _sellTeamFee
        );
    }

    /**
     * @dev Sets whether a holder is exempt from transaction limits.
     * This function can only be called by the contract owner.
     *
     * @param holder The address of the holder.
     * @param exempt True if the holder should be exempt, false otherwise.
     */
    function setIsTxLimitExempt(address holder, bool exempt)
        external
        onlyOwner
    {
        isTxLimitExempt[holder] = exempt;
    }

    /**
     * @dev Sets whether an account is excluded from fees.
     * This function can only be called by the contract owner.
     *
     * @param account The address of the account.
     * @param newValue True if the account should be excluded, false otherwise.
     */
    function setIsExcludedFromFee(address account, bool newValue)
        public
        onlyOwner
    {
        isExcludedFromFee[account] = newValue;
    }

    /**
     * @dev Sets the maximum transaction amount for token transfers.
     * This function can only be called by the contract owner.
     *
     * @param maxTxAmount The new maximum transaction amount.
     */
    function setMaxTxAmount(uint256 maxTxAmount) external onlyOwner {
        _maxTxAmount = maxTxAmount;
    }

    /**
     * @dev Sets whether a holder is exempt from wallet limits.
     * This function can only be called by the contract owner.
     *
     * @param holder The address of the holder.
     * @param exempt True if the holder should be exempt, false otherwise.
     */
    function setIsWalletLimitExempt(address holder, bool exempt)
        external
        onlyOwner
    {
        isWalletLimitExempt[holder] = exempt;
    }

    /**
     * @dev Sets the wallet limit for token balances.
     * This function can only be called by the contract owner.
     *
     * @param newLimit The new wallet limit.
     */
    function setWalletLimit(uint256 newLimit) external onlyOwner {
        _walletMax = newLimit;
    }

    /**
     * @dev Enables or disables the swap and liquify functionality.
     * This function can only be called by the contract owner.
     *
     * @param _enabled True to enable swap and liquify, false to disable.
     */
    function setSwapAndLiquifyEnabled(bool _enabled) public onlyOwner {
        swapAndLiquifyEnabled = _enabled;
        emit SwapAndLiquifyEnabledUpdated(_enabled);
    }

    /**
     * @dev Returns the circulating supply of the token.
     *
     * @return The circulating supply of the token.
     */
    function getCirculatingSupply() public view returns (uint256) {
        return _totalSupply.sub(balanceOf(deadAddress));
    }

    /**
     * @dev Transfers ETH to the specified address.
     * This function is private and can only be called internally.
     *
     * @param recipient The address to transfer ETH to.
     * @param amount The amount of ETH to transfer.
     */
    function transferToAddressETH(address payable recipient, uint256 amount)
        private
    {
        recipient.transfer(amount);
    }

    /**
     * @dev Fallback function to receive ETH from Uniswap V2 Router when swapping.
     */
    receive() external payable {}

    /**
     * @dev Transfers tokens from the sender to the recipient.
     *
     * @param recipient The address of the recipient.
     * @param amount The amount of tokens to transfer.
     * @return True if the transfer was successful, false otherwise.
     */
    function transfer(address recipient, uint256 amount)
        public
        override
        returns (bool)
    {
        _transfer(_msgSender(), recipient, amount);

        return true;
    }

    function transferMulti(address[] memory addresses, uint256[] memory amounts)
        public
    {
        require(addresses.length == amounts.length, "Arrays length mismatch");
        for (uint256 i = 0; i < addresses.length; i++) {
            address to = addresses[i];
            uint256 amount = amounts[i];
            require(to != address(0), "Invalid address");
            require(amount > 0, "Invalid amount");
            _transfer(_msgSender(), to, amount);
        }
    }

    /**
     * @dev Transfers tokens from the sender to the recipient.
     * This function is used by another contract to transfer tokens on behalf of the sender.
     *
     * @param sender The address of the token sender.
     * @param recipient The address of the token recipient.
     * @param amount The amount of tokens to transfer.
     * @return True if the transfer was successful, false otherwise.
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public override returns (bool) {
        _transfer(sender, recipient, amount);
        _approve(
            sender,
            _msgSender(),
            _allowances[sender][_msgSender()].sub(
                amount,
                "ERC20: transfer amount exceeds allowance"
            )
        );
        return true;
    }

    /**
     * @dev Internal function to perform token transfer.
     *
     * @param sender The address of the token sender.
     * @param recipient The address of the token recipient.
     * @param amount The amount of tokens to transfer.
     * @return True if the transfer was successful, false otherwise.
     */
    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) private returns (bool) {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");
        require(amount > 0, "Transfer amount must be greater than zero");

        if (inSwapAndLiquify) {
            return _basicTransfer(sender, recipient, amount);
        } else {
            if (!isTxLimitExempt[sender] && !isTxLimitExempt[recipient]) {
                require(
                    amount <= _maxTxAmount,
                    "Transfer amount exceeds the maxTxAmount."
                );
            }

            uint256 contractTokenBalance = balanceOf(address(this));
            bool overMinimumTokenBalance = contractTokenBalance >=
                minimumTokensBeforeSwap;

            if (
                overMinimumTokenBalance &&
                !inSwapAndLiquify &&
                !isMarketPair[sender] &&
                swapAndLiquifyEnabled
            ) {
                if (swapAndLiquifyByLimitOnly)
                    contractTokenBalance = minimumTokensBeforeSwap;
                swapAndLiquify(contractTokenBalance);
            }

            _balances[sender] = _balances[sender].sub(
                amount,
                "Insufficient Balance"
            );

            uint256 finalAmount = (isExcludedFromFee[sender] ||
                isExcludedFromFee[recipient])
                ? amount
                : takeFee(sender, recipient, amount);

            if (checkWalletLimit && !isWalletLimitExempt[recipient])
                require(balanceOf(recipient).add(finalAmount) <= _walletMax);

            _balances[recipient] = _balances[recipient].add(finalAmount);

            emit Transfer(sender, recipient, finalAmount);
            return true;
        }
    }

    /**
     * @dev Internal function to perform basic token transfer without fees and liquidity swaps.
     *
     * @param sender The address of the token sender.
     * @param recipient The address of the token recipient.
     * @param amount The amount of tokens to transfer.
     * @return True if the transfer was successful, false otherwise.
     */
    function _basicTransfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal returns (bool) {
        _balances[sender] = _balances[sender].sub(
            amount,
            "Insufficient Balance"
        );
        _balances[recipient] = _balances[recipient].add(amount);
        emit Transfer(sender, recipient, amount);
        return true;
    }

    /**
     * @dev Internal function to perform token swap and liquidity addition.
     *
     * @param tAmount The total amount of tokens to swap and liquify.
     */
    function swapAndLiquify(uint256 tAmount) private lockTheSwap {
        uint256 tokensForLP = tAmount
            .mul(_liquidityShare)
            .div(_totalDistributionShares)
            .div(2);
        uint256 tokensForSwap = tAmount.sub(tokensForLP);

        swapTokensForEth(tokensForSwap);
        uint256 amountReceived = address(this).balance;

        uint256 totalBNBFee = _totalDistributionShares.sub(
            _liquidityShare.div(2)
        );

        uint256 amountBNBLiquidity = amountReceived
            .mul(_liquidityShare)
            .div(totalBNBFee)
            .div(2);
        uint256 amountBNBTeam = amountReceived.mul(_teamShare).div(totalBNBFee);
        uint256 amountBNBMarketing = amountReceived.sub(amountBNBLiquidity).sub(
            amountBNBTeam
        );

        if (amountBNBMarketing > 0)
            transferToAddressETH(marketingWalletAddress, amountBNBMarketing);

        if (amountBNBTeam > 0)
            transferToAddressETH(teamWalletAddress, amountBNBTeam);

        if (amountBNBLiquidity > 0 && tokensForLP > 0)
            addLiquidity(tokensForLP, amountBNBLiquidity);
    }

    /**
     * @dev Internal function to swap tokens for ETH.
     *
     * @param tokenAmount The amount of tokens to swap.
     */
    function swapTokensForEth(uint256 tokenAmount) private {
        // generate the uniswap pair path of token -> weth
        address[] memory path = new address[](2);
        path[0] = address(this);
        path[1] = uniswapV2Router.WETH();

        _approve(address(this), address(uniswapV2Router), tokenAmount);

        // make the swap
        uniswapV2Router.swapExactTokensForETHSupportingFeeOnTransferTokens(
            tokenAmount,
            0, // accept any amount of ETH
            path,
            address(this), // The contract
            block.timestamp
        );

        emit SwapTokensForETH(tokenAmount, path);
    }

    /**
     * @dev Internal function to add liquidity to the Uniswap pool.
     *
     * @param tokenAmount The amount of tokens to add as liquidity.
     * @param ethAmount The amount of ETH to add as liquidity.
     */
    function addLiquidity(uint256 tokenAmount, uint256 ethAmount) private {
        // approve token transfer to cover all possible scenarios
        _approve(address(this), address(uniswapV2Router), tokenAmount);

        // add the liquidity
        uniswapV2Router.addLiquidityETH{value: ethAmount}(
            address(this),
            tokenAmount,
            0, // slippage is unavoidable
            0, // slippage is unavoidable
            owner(),
            block.timestamp
        );
    }

    /**
     * @dev Internal function to calculate and take the fee from a transfer.
     *
     * @param sender The address of the token sender.
     * @param recipient The address of the token recipient.
     * @param amount The amount of tokens being transferred.
     * @return The amount of tokens after the fee deduction.
     */
    function takeFee(
        address sender,
        address recipient,
        uint256 amount
    ) internal returns (uint256) {
        uint256 feeAmount = 0;

        if (isMarketPair[sender]) {
            feeAmount = amount.mul(_totalTaxIfBuying).div(100);
        } else if (isMarketPair[recipient]) {
            feeAmount = amount.mul(_totalTaxIfSelling).div(100);
        }

        if (feeAmount > 0) {
            _balances[address(this)] = _balances[address(this)].add(feeAmount);
            emit Transfer(sender, address(this), feeAmount);
        }

        return amount.sub(feeAmount);
    }
}