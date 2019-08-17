pragma solidity >=0.4.24;

import "./owned.sol"; 
import "./safemath.sol";
import "./PPNToken.sol";

contract MinerManager is owned{ 
    
    using SafeMath for uint256;
    
    struct Miner{
       bytes32 protonAddress;
       address ethAddress;
    }
    
    PPNToken public token;
    uint public TokenDecimals;
    
    constructor(address ta) public{
        token = PPNToken(ta);
        TokenDecimals = token.getDeccimal();
    }
    
    
    function JoinMinerPool(address poolEthAddr) public {
    }
}