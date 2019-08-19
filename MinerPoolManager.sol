pragma solidity >=0.4.24;

import "./owned.sol"; 
import "./safemath.sol";
import "./PPNToken.sol";

contract MinerPoolManager is owned{
    using SafeMath for uint256;
    struct MinerPool{
       string desc;
       address ethAddress;
       uint guaranteedToken;
    }
    struct TopItem{
        address poolEthAddr;
        uint guaranteedNo;
    }

    uint public MinPoolCostInToken = 102400000;
    PPNToken public token;
    uint public TokenDecimals;
    mapping(address=>MinerPool) AllMinerPools;
    TopItem[64] public TopPoolList;
    
    constructor(address ta) public{
        token = PPNToken(ta);
        TokenDecimals = token.getDeccimal();
    }
    
    function RegAsMinerPool(uint guaranteeAmount, string memory desc) public {
        require(guaranteeAmount > MinPoolCostInToken);

        uint tokenNoWithDecimal = guaranteeAmount.mul(TokenDecimals);
        require(token.balanceOf(msg.sender) > tokenNoWithDecimal);

        MinerPool storage pool = AllMinerPools[msg.sender];
        require(pool.ethAddress == address(0));

        pool.ethAddress = msg.sender;
        pool.desc = desc;
        token.transfer(address(this), guaranteeAmount);
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
        require(token.balanceOf(msg.sender) > noWithDecimal);
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
}