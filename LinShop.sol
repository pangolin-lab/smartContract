pragma solidity >=0.4.24;

import "./owned.sol";
import "./PangolinToken.sol";


contract LinShop is owned{ 
    using SafeMath for uint256;
    
    event TokensPurchased(
        address indexed purchaser, 
        uint256 value,
        uint256 amount
    );
    
    address payable private  _adminWallet;
    uint256 constant public _rateLevel1 = 300;
    uint256 public _rateLevel2 = 240;
    uint256 public TokenNoToSellInLevel1;
    uint256 public hasRaised = 0;
    PangolinToken private _token;
    
    constructor(address payable wallet, PangolinToken token) public{ 
        require(wallet != address(0));
        require(address(token) != address(0));  
        _adminWallet = wallet;
        _token = token;
        TokenNoToSellInLevel1 = 2.1e7 * _token.getDeccimal();
    }
    
    function () payable external{
        uint256 weiAmount = msg.value;
        require(weiAmount > 0); 
        
        uint256 tokenNo = 0;
        if (hasRaised < TokenNoToSellInLevel1){
            tokenNo = weiAmount.mul(_rateLevel1);
        }else{
            tokenNo = weiAmount.mul(_rateLevel2);
        } 
        require(remainingTokens() >= tokenNo);
        
        _token.transferFrom(_adminWallet, msg.sender, tokenNo); 
        emit TokensPurchased(msg.sender, weiAmount, tokenNo); 
        _adminWallet.transfer(weiAmount);
        hasRaised += tokenNo;
    }
    
    function SetLevel2Rate(uint newRate) public onlyOwner{
        _rateLevel2 = newRate;
    }
    
    function token() public view returns(ERC20) {
        return _token;
    }
    
    function adminWallet() public view returns(address) {
        return _adminWallet;
    }
    
    function remainingTokens() public view returns (uint256) {
        return _token.allowance(_adminWallet, address(this));
    }
}