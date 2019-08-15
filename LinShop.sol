pragma solidity >=0.4.24;

import "./owned.sol";
import "./ERC20.sol";


contract LinShop is owned{ 
    using SafeMath for uint256;
    
    event TokensPurchased(
        address indexed purchaser, 
        uint256 value,
        uint256 amount
    );
    
    address private _adminWallet;
    uint256 private _rate;
    ERC20 private _token;
    
    constructor(uint256 rate, address wallet, ERC20 token) public{
        require(rate > 0);
        require(wallet != address(0));
        require(token != address(0));
    
        _rate = rate;
        _adminWallet = wallet;
        _token = token;
    }
    
    function () external payable {
        buyTokens();
    }
    
    function token() public view returns(ERC20) {
        return _token;
    }
    
    function adminWallet() public view returns(address) {
        return _adminWallet;
    }
    
    function rate() public view returns(uint256) {
        return _rate;
    } 
    
    function remainingTokens() public view returns (uint256) {
        return _token.allowance(_adminWallet, this);
    }
    
    function buyTokens() public payable {
        
        uint256 weiAmount = msg.value;
        require(weiAmount > 0);
        
        uint256 tokenNo = weiAmount.mul(_rate); 
        require(remainingTokens() >= tokenNo);
        
        _token.transferFrom(_adminWallet, msg.sender, tokenNo);
        
        emit TokensPurchased( msg.sender, weiAmount, tokenNo);
        
        _adminWallet.transfer(weiAmount);
    }
}