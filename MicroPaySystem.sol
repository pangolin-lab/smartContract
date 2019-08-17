pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./PPNToken.sol";

contract MicroPaySystem is owned{

    struct Client{
        bytes32 protonAddress;
        address ethAddress;
        uint refundAveragePrice;
    }
    

    struct PayChannelItem{
        bytes32 protonAddr;
        address buyerEthAddr;
    }

    uint public SettlePriceForBuyer = 4;  //4M bit/ppnt;
    function ChangePriceForBuyer(uint newPrice) public onlyOwner{
        SettlePriceForBuyer = newPrice;
    }

    using SafeMath for uint256;
    uint public MinMinerCostInToken = 1024000;
    uint public MinUserCostInToken = 100;

    PPNToken public token;
    uint public  TokenDecimals;

    mapping(bytes32=>PayChannelItem) NormalPayChannels;

    constructor(address ta) public{
        token = PPNToken(ta);
        TokenDecimals = token.getDeccimal();
    }
    /*
    *
    * functions for user
    *
    */
    function BuyPacket(uint8 typ, bytes32 targetAddress, uint tokenAmount, address poolEthAddr) public{
        require(tokenAmount > MinUserCostInToken);
        require(token.balanceOf(msg.sender) > tokenAmount);

        // MinerPool memory pool = AllMinerPools[poolEthAddr];
        // require(pool.ethAddress != address(0));

        // PayChannelItem storage item =  NormalPayChannels[targetAddress];

        // if (item.buyerEthAddr == address(0)){
        //     _initNewMicroPayChannel(targetAddress, tokenAmount, pool, item);
        // }else{
        //     _rechargeToPayChannel(targetAddress, tokenAmount, pool, item);
        // }
    }
}