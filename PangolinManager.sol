pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";


interface PPNTToken {
    function balanceOf(address account) external view returns (uint);
    function getDeccimal() external pure returns (uint);
}


contract PangolinManager is owned{
    using SafeMath for uint256;
    uint constant public TokenNoForOneUser = 100;
    uint TokenNoPerUser;
    PPNTToken public token;
    
    mapping(bytes32=>address) public PangolinUserRecord;
    mapping(address=>uint) public EtherCounter;
    
    constructor (address tokenAddr) public{
        token = PPNTToken(tokenAddr);
        TokenNoPerUser = token.getDeccimal().mul(TokenNoForOneUser);
    }
    
    
    function unbind(bytes32 addr) public{
        address userAddr = PangolinUserRecord[addr];
        require(userAddr == msg.sender);
        
        delete PangolinUserRecord[addr];
        EtherCounter[msg.sender] = EtherCounter[msg.sender].sub(1);
    }
    
    function bind(bytes32 addr) public{
        require(PangolinUserRecord[addr] == address(0));
        
        uint totalNeed = TokenNoPerUser.mul(EtherCounter[msg.sender]);
        require(token.balanceOf(msg.sender) >= totalNeed);
        
        EtherCounter[msg.sender] = EtherCounter[msg.sender].add(1);
        PangolinUserRecord[addr] = msg.sender;
    }
    
    function check(bytes32 addr) public view returns(address, uint, uint){
        address userAddr = PangolinUserRecord[addr];
        if (userAddr == address(0)){
            return (address(0), 0, 0);
        }
        
        uint linBalance = token.balanceOf(userAddr);
        return (userAddr, linBalance, EtherCounter[userAddr]);
    }
    
    function bindingInfo(address userAddr) public view returns(uint, uint, uint){ 
        uint ethBalance = userAddr.balance;
        uint linBalance = token.balanceOf(userAddr);
        uint bindNo = EtherCounter[userAddr]; 
        return (ethBalance, linBalance, bindNo);
    }
    
    function changeServicePrice(uint256 price) public onlyOwner{
        TokenNoPerUser = price.mul(token.getDeccimal());
    }
}