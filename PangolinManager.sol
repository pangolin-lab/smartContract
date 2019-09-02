pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./PangolinToken.sol"; 

contract MicroPaySystem is owned{
    
    struct MinerPool{
        address mainAddr;
        address payer;
        bytes32 subAddr;
        uint guaranteedNo;
        uint ID;
        uint8 poolType;
        string shortName;
        string detailInfos;
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
    uint public Duration = 30 days;
    
    PangolinToken public token;
    
    mapping(address=>MinerPool) public MinerPools;
    address[] public MinerPoolsAddresses;
    
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
        Duration = daysAfter * 1 days;
    }
    
    constructor(address ta) public{
        token = PangolinToken(ta);
        TokenDecimals = token.getDeccimal();
        
        MinUserCostInToken = 100 * TokenDecimals;
        MinPoolCostInToken = 51200 * TokenDecimals;
        MinMinerCostInToken = 512 * TokenDecimals;
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
        ch.expiration = now + Duration;
    }
    
    function TokenBalance(address userAddress) public view returns (uint, uint){
        return (token.balanceOf(userAddress), userAddress.balance);
    }
    
    /********************************************************************************
    *                           Pool
    *********************************************************************************/
    function RegAsMinerPool(uint gno, address mainAddr,  bytes32 subAddr, string memory name, string memory desc) public {
        require(gno > MinPoolCostInToken); 
        require(token.balanceOf(msg.sender) > gno); 
        
        MinerPool storage pool = MinerPools[mainAddr];
        require(pool.mainAddr == address(0));
        
        token.transferFrom(msg.sender, address(this), gno);
        
        pool.ID = MinerPoolsAddresses.length;
        MinerPoolsAddresses.push(mainAddr);
         
        pool.mainAddr = mainAddr;
        pool.payer = msg.sender;
        pool.subAddr = subAddr;
        pool.guaranteedNo = gno;
        pool.poolType = 0;
        pool.shortName = name;
        pool.detailInfos = desc;
    }
    
    function SetPoolType(address mainAddr, uint8 typ) public onlyOwner{
        MinerPool storage pool = MinerPools[mainAddr];
        require(pool.mainAddr != address(0));
        pool.poolType = typ;
    }
    
     function GetPoolSize() public view returns (uint){
         return MinerPoolsAddresses.length;
     }
     
     function GetPoolAddrees() public view returns (address[] memory){
        return MinerPoolsAddresses;
     }
    
    /********************************************************************************
    *                           Miner
    *********************************************************************************/
}