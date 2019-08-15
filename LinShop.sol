pragma solidity >=0.4.24;

import "./owned.sol";
import "./ERC20.sol";

interface PPNTToken {
    function balanceOf(address account) external view returns (uint);
    function getDeccimal() external pure returns (uint);
    function allowance(address _owner, address _spender) external view returns (uint256 remaining);
}

contract LinShop is owned{ 
    using SafeMath for uint256;
    
    event TokensPurchased(
        address indexed purchaser,
        address indexed beneficiary,
        uint256 value,
        uint256 amount
    );
    
    address private _wallet;
    uint256 private _rate;
    PPNTToken private _token;
    
    constructor(uint256 rate, address wallet, address token) public{
        require(rate > 0);
        require(wallet != address(0));
        require(token != address(0));
    
        _rate = rate;
        _wallet = wallet;
        _token = PPNTToken(token);
    }
    
    function () external payable {
        buyTokens(msg.sender);
    }
    
    function token() public view returns(PPNTToken) {
        return _token;
    }
    
    function wallet() public view returns(address) {
        return _wallet;
    }
    
    function rate() public view returns(uint256) {
        return _rate;
    } 
    
    function buyTokens(address beneficiary) public payable {
        uint256 weiAmount = msg.value;
    }
    
    
}