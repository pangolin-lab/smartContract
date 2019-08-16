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
    
    address payable private  _adminWallet;
    uint256 private _rate;
    ERC20 private _token;
    
    constructor(uint256 rate, address payable wallet, ERC20 token) public{
        require(rate > 0);
        require(wallet != address(0));
        require(address(token) != address(0));
    
        _rate = rate;
        _adminWallet = wallet;
        _token = token;
    }
    
    function () payable external{
        uint256 weiAmount = msg.value;
        require(weiAmount > 0); 
        
        uint256 tokenNo = weiAmount.mul(_rate); 
        require(remainingTokens() >= tokenNo);
        
        _token.transferFrom(_adminWallet, msg.sender, tokenNo); 
        emit TokensPurchased(msg.sender, weiAmount, tokenNo); 
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
    
    function changeRate(uint256 newRate) public onlyOwner{
        _rate = newRate;
    }
    
    function remainingTokens() public view returns (uint256) {
        return _token.allowance(_adminWallet, address(this));
    }
    
    function withDrawIncome() public onlyOwner{
        _adminWallet.transfer(address(this).balance);
    }
}