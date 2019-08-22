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
    uint256 constant private _rateLevel1 = 300;
    uint256 constant private _rateLevel2 = 240;
    uint256 private level1 = 2.1e7 * (10 ** 18);
    uint256 public hasRaised = 0;
    PangolinToken private _token;
    
    constructor(address payable wallet, PangolinToken token) public{ 
        require(wallet != address(0));
        require(address(token) != address(0));  
        _adminWallet = wallet;
        _token = token;
        level1 = 2.1e7 * _token.getDeccimal();
    }
    
    function () payable external{
        uint256 weiAmount = msg.value;
        require(weiAmount > 0); 
        
        uint256 tokenNo = 0;
        if (hasRaised < level1){
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
    
    function token() public view returns(ERC20) {
        return _token;
    }
    
    function adminWallet() public view returns(address) {
        return _adminWallet;
    }
    
    function rate() public pure returns(uint256, uint256) {
        return (_rateLevel1, _rateLevel2);
    }
    
    function remainingTokens() public view returns (uint256) {
        return _token.allowance(_adminWallet, address(this));
    }
}