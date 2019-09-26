pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./PangolinToken.sol"; 

contract MicroPaySystem is owned{
    
    struct MinerPool{
        address mainAddr;
        address payer;
        uint    guaranteedNo;
        string shortName;
        string detailInfos;
        string seeds;
    }
    
    struct Channel{
        address mainAddr;
        uint remindTokens;
        uint remindPackets;
        uint256 expiration;
    }
    
    using SafeMath for uint256; 
    uint public PacketPrice = 16000000;  //(16M Bytes)/ppnt /-->/1M = 1000 KB = 1,000,000 Bytes;

    uint public MinPoolCostInToken;
    uint public MinMinerCostInToken;
    uint public TokenDecimals;  
    uint public Duration = 30 days;
    
    PangolinToken public token;
    uint32 public MinerPoolVersion;
    mapping(address =>uint32) public ChannelVersion;
    mapping(address=>MinerPool) public MinerPools;
    address[] public MinerPoolsAddresses;
    
    mapping(address=>mapping(address=>Channel)) public MicroPaymentChannels;
    mapping(address=>address[])public allSubPools;

    /********************************************************************************
    *                           Basic
    *********************************************************************************/
    
    function ChangeBandWithPrice(uint newPrice) public onlyOwner{
        PacketPrice = newPrice;
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
        
        MinPoolCostInToken = 51200 * TokenDecimals;
        MinMinerCostInToken = 512 * TokenDecimals;
    }

    /********************************************************************************
    *                           User
    *********************************************************************************/
    function BuyPacket(address user, uint tokenNo, address poolAddr) public{
        
        require(token.balanceOf(msg.sender) > tokenNo); 
        MinerPool memory pool = MinerPools[poolAddr];
        require(pool.mainAddr != address(0));
        
        token.transferFrom(msg.sender, address(this), tokenNo); 
        uint newPackets = tokenNo.div(TokenDecimals).mul(PacketPrice); 
        
        Channel storage ch = MicroPaymentChannels[user][pool.mainAddr];
        if (ch.mainAddr == address(0)){
            ch.mainAddr = user;
            allSubPools[user].push(pool.mainAddr);
        }
        
        ch.remindPackets += newPackets;
        ch.remindTokens += tokenNo;
        ch.expiration = now + Duration;
    }
    
    function TokenBalance(address userAddress) public view returns (uint, uint, uint){
        return (token.balanceOf(userAddress), userAddress.balance, token.allowance(userAddress, address(this)));
    }
    
    function AllMySubPools(address userAddress) public view returns (address[] memory){
        return allSubPools[userAddress];
    }
    
    /********************************************************************************
    *                           Pool
    *********************************************************************************/
    function RegAsMinerPool(uint gno, address mainAddr, string memory name, string memory desc, string memory seeds) public {
        require(gno > MinPoolCostInToken); 
        require(token.balanceOf(msg.sender) > gno); 
        
        MinerPool storage pool = MinerPools[mainAddr];
        require(pool.mainAddr == address(0));
        
        token.transferFrom(msg.sender, address(this), gno);
        
        MinerPoolsAddresses.push(mainAddr);
         
        pool.mainAddr = mainAddr;
        pool.payer = msg.sender;
        pool.guaranteedNo = gno;
        pool.shortName = name;
        pool.detailInfos = desc;
        pool.seeds = seeds;
        MinerPoolVersion += 1;
    }
    
    function GetPoolAddress() public view returns (address[] memory){
        return MinerPoolsAddresses;
    }
     
    function ChangeShortName(address mainAddr, string memory name) public{
       
        MinerPool storage pool = MinerPools[mainAddr]; 
        require(pool.mainAddr == msg.sender || pool.payer == msg.sender); 
        pool.shortName = name;
    }
    
    function ChangeDesc(address mainAddr, string memory desc) public{
        MinerPool storage pool = MinerPools[mainAddr]; 
        require(pool.mainAddr == msg.sender || pool.payer == msg.sender);
        pool.detailInfos = desc;
    }
    
    function ChangeSeed(address mainAddr, string memory seeds) public{
        MinerPool storage pool = MinerPools[mainAddr]; 
        require(pool.mainAddr == msg.sender || pool.payer == msg.sender);
        pool.seeds = seeds;
    }
    
    /********************************************************************************
    *                           Miner
    *********************************************************************************/
}