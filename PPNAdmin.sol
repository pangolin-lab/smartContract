pragma solidity >=0.4.24;

import "./owned.sol"; 

contract PPNAdmin is owned{ 
    address public MinerPoolManager;
    address public MinerManager;
    address public PayChannel;
    address public TokenAddress;
    
    constructor(address initToken) public{
        TokenAddress = initToken;
    }
    
    function ChangeMicroPay(address pc) public onlyOwner{
        PayChannel = pc;
    }
    
    function ChangeMinerManager(address mm) public onlyOwner{
        MinerPoolManager = mm;
    }
    
    function ChangeMinerPoolManager(address mpm) public onlyOwner{
        MinerPoolManager = mpm;
    }
}