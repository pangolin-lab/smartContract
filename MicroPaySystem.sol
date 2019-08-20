pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./PPNToken.sol"; 

contract MicroPaySystem is owned{
    
    struct TopItem{
        address poolEthAddr;
        uint guaranteedNo;
    }
    
    struct MinerPool{
        address mainAddr;
        bytes32 viceAddr;
    }
    
     
    struct Miner{
       bytes32 protonAddress;
       address ethAddress;
    }
    
    using SafeMath for uint256; 
    
    uint public BandWithPerToken = 4;  //4M bytes/ppnt;
    uint public MinUserCostInToken = 100;
    uint public MinPoolCostInToken = 102400000;
    uint public MinMinerCostInToken = 1024000;
    
    PPNToken public token;
    uint public  TokenDecimals; 
    mapping(address=>MinerPool) MinerPools;
    
    function ChangeBandWithPrice(uint newPrice) public onlyOwner{
        BandWithPerToken = newPrice;
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
    
    constructor(address ta) public{
        token = PPNToken(ta);
        TokenDecimals = token.getDeccimal();
    }

    /********************************************************************************
    *                           User
    *********************************************************************************/
    function BuyPacket(uint8 typ, bytes32 targetAddress, uint tokenAmount, address poolEthAddr) public{
        require(tokenAmount > MinUserCostInToken);
        require(token.balanceOf(msg.sender) > tokenAmount);
        
        MinerPool memory pool = MinerPools[poolEthAddr];
        require(pool.mainAddr != address(0));
        

        // PayChannelItem storage item =  NormalPayChannels[targetAddress];

        // if (item.buyerEthAddr == address(0)){
        //     _initNewMicroPayChannel(targetAddress, tokenAmount, pool, item);
        // }else{
        //     _rechargeToPayChannel(targetAddress, tokenAmount, pool, item);
        // }
    }
    
    
    /********************************************************************************
    *                           Pool
    *********************************************************************************/
    function RegAsMinerPool(uint guaranteeAmount, string memory desc) public {
        require(guaranteeAmount > MinPoolCostInToken);

        uint tokenNoWithDecimal = guaranteeAmount.mul(TokenDecimals);
        require(token.balanceOf(msg.sender) > tokenNoWithDecimal); 
    }

    function _reSortTopList(address poolAddr, uint tokenNo) internal {
    }
    
    /********************************************************************************
    *                           Miner
    *********************************************************************************/
}