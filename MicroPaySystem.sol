pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./PPNToken.sol"; 

contract MicroPaySystem is owned{
    
    struct MinerPool{
        address mainAddr;
        bytes32 viceAddr;
        uint guaranteedNo;
    }
    
    struct Channel{
        address sigAddr;
        uint remindTokens;
        uint remindPackets;
        uint256 expiration;
    }
    
    using SafeMath for uint256; 
    
    uint public PacketPrice = 16000000;  //(16M Bytes)/ppnt /-->/1M = 1000 KB = 1,000,000 Bytes;
    uint public MinUserCostInToken;
    uint public MinPoolCostInToken;
    uint public MinMinerCostInToken;
    uint public TokenDecimals;  
    uint public DurationInDays = 30 days;
    
    PPNToken public token;
    
    mapping(address=>MinerPool) public MinerPools;
    mapping(bytes32=>mapping(address=>Channel)) public MicroPaymentChannels;
    
    function ChangeBandWithPrice(uint newPrice) public onlyOwner{
        PacketPrice = newPrice;
    } 

    function ChangeMinUserCost(uint newCost) public onlyOwner{
        MinUserCostInToken = newCost;
    }
    
    function ChangeMinPoolCost(uint newCost) public onlyOwner{
        MinPoolCostInToken = newCost;
    }
    
    function ChangeMinMinerCost(uint newCost) public onlyOwner{
        MinMinerCostInToken = newCost;
    }
    
    function ChangeDuration(uint daysAfter) public onlyOwner{
        DurationInDays = daysAfter * 1 days;
    }
    
    constructor(address ta) public{
        token = PPNToken(ta);
        TokenDecimals = token.getDeccimal();
        
        MinUserCostInToken = 100 * TokenDecimals;
        MinPoolCostInToken = 102400000 * TokenDecimals;
        MinMinerCostInToken = 1024000 * TokenDecimals;
    }

    /********************************************************************************
    *                           User
    *********************************************************************************/
    function BuyPacket(bytes32 va, uint tokenNo, address poolAddr) public{
        
        require(tokenNo > MinUserCostInToken);
        require(token.balanceOf(msg.sender) > tokenNo);
        
        MinerPool memory pool = MinerPools[poolAddr];
        require(pool.mainAddr != address(0));
        
        token.transfer(address(this), tokenNo); 
        uint newPackets = tokenNo.div(TokenDecimals).mul(PacketPrice); 
        
        Channel storage ch = MicroPaymentChannels[va][pool.mainAddr];
        ch.remindPackets += newPackets;
        ch.remindTokens += tokenNo;
        ch.expiration = now + DurationInDays;
    }
    
    
    /********************************************************************************
    *                           Pool
    *********************************************************************************/
    function RegAsMinerPool(uint gno, bytes32 va) public {
        require(gno > MinPoolCostInToken); 
        require(token.balanceOf(msg.sender) > gno); 
        
         MinerPool storage pool = MinerPools[msg.sender];
         require(pool.mainAddr == address(0));
         
         token.transfer(address(this), gno);
         
         pool.mainAddr = msg.sender;
         pool.guaranteedNo = gno;
         pool.viceAddr = va;
    }
    
    /********************************************************************************
    *                           Miner
    *********************************************************************************/
}