pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";

interface Token{
    function balanceOf(address account) external view returns (uint);
    function transfer(address to, uint256 value) external returns (bool);
    function getDeccimal() external pure returns (uint);
    function getAllowanceToken(address from, address to) external view returns (uint);
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
}

contract MicroPaySystem is owned{

    struct Client{
        bytes32 protonAddress;
        address ethAddress;
        uint refundAveragePrice;
    }
    struct MinerPool{
       string desc;
       address ethAddress;
       uint guaranteedToken;
    }
    struct Miner{
       bytes32 protonAddress;
       address ethAddress;
    }

    struct PayChannelItem{
        bytes32 protonAddr;
        address buyerEthAddr;
    }

    uint public SettlePriceForBuyer = 4;  //4M bit/ppnt;
    function ChangePriceForBuyer(uint newPrice) public onlyOwner{
        SettlePriceForBuyer = newPrice;
    }
    // uint public SettlePriceForMortgagor = 8; //8M bit/ppnt
    // function ChangePriceForMortgagor(uint newPrice) public onlyOwner{
    //     SettlePriceForMortgagor = newPrice;
    // }

    struct TopItem{
        address poolEthAddr;
        uint guaranteedNo;
    }

    using SafeMath for uint256;
    uint public constant MinMinerCostInToken = 1024000;
    uint public constant MinUserCostInToken = 100;
    uint public constant MinPoolCostInToken = 102400000;

    Token public PPNToken;
    uint public  TokenDecimals;
    TopItem[64] public TopPoolList;

    mapping(address=>MinerPool) AllMinerPools;
    mapping(bytes32=>PayChannelItem) NormalPayChannels;

    constructor(address ta) public{
        PPNToken = Token(ta);
        TokenDecimals = PPNToken.getDeccimal();
    }
    /*
    *
    * functions for user
    *
    */
    function BuyPacket(bytes32 targetAddress, uint tokenAmount, address poolEthAddr) public{
        require(tokenAmount > 0);
        require(PPNToken.balanceOf(msg.sender) > tokenAmount);

        MinerPool memory pool = AllMinerPools[poolEthAddr];
        require(pool.ethAddress != address(0));

        PayChannelItem storage item =  NormalPayChannels[targetAddress];

        if (item.buyerEthAddr == address(0)){
            _initNewMicroPayChannel(targetAddress, tokenAmount, pool, item);
        }else{
            _rechargeToPayChannel(targetAddress, tokenAmount, pool, item);
        }
    }

    function _initNewMicroPayChannel(bytes32 targetAddress, uint tokenAmount, MinerPool memory pool, PayChannelItem storage item) internal{

        item.buyerEthAddr = msg.sender;
    }

    function _rechargeToPayChannel(bytes32 targetAddress, uint tokenAmount, MinerPool memory pool, PayChannelItem storage item) internal{

    }

    /*
    *
    * functions for miner pool
    *
    */
    function RegAsMinerPool(uint guaranteeAmount, string memory desc) public {

        require(guaranteeAmount > MinPoolCostInToken);

        uint tokenNoWithDecimal = guaranteeAmount.mul(TokenDecimals);
        require(PPNToken.balanceOf(msg.sender) > tokenNoWithDecimal);

        MinerPool storage pool = AllMinerPools[msg.sender];
        require(pool.ethAddress == address(0));

        pool.ethAddress = msg.sender;
        pool.desc = desc;
        PPNToken.transfer(address(this), guaranteeAmount);
        pool.guaranteedToken = guaranteeAmount;

        if (guaranteeAmount > TopPoolList[63].guaranteedNo){
            _reSortTopList(msg.sender, guaranteeAmount);
        }
    }

    function ChangeDesc(string memory newDesc) public {
        MinerPool storage pool = AllMinerPools[msg.sender];
        pool.desc = newDesc;
    }

    function AddGuarantToken(uint tokenNo) public {

        uint noWithDecimal = tokenNo.mul(TokenDecimals);

        require(PPNToken.balanceOf(msg.sender) > noWithDecimal);
        MinerPool storage pool = AllMinerPools[msg.sender];

        pool.guaranteedToken = pool.guaranteedToken.add(noWithDecimal);

        if (pool.guaranteedToken > TopPoolList[63].guaranteedNo){
            _reSortTopList(msg.sender, pool.guaranteedToken);
        }
    }

    function RefundGuarantToken(uint tokenNo) public{

        uint noWithDecimal = tokenNo.mul(TokenDecimals);
        MinerPool storage pool = AllMinerPools[msg.sender];

        if (pool.guaranteedToken >= MinPoolCostInToken.add(noWithDecimal)){

        }
    }

    function _reSortTopList(address poolAddr, uint tokenNo) internal {
    }

    /*
    *
    * functions for miner
    *
    */
    function JoinMinerPool(address poolEthAddr) public {
    }
}