pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";

interface Token{
    function balanceOf(address account) external view returns (uint);
}

contract MicroPayChannel is owned{
    
    struct ProtonNode{
        bytes32 ProtonAddress;
        address EthAddress;
    }
    
    struct Client{
        ProtonNode baseInfo;
    }

    struct MinerPool{
       ProtonNode baseInfo;
    }

    struct Miner{
       ProtonNode baseInfo;
    } 
    
    using SafeMath for uint256;
    
    Token public TokenAddress;
    uint public BasicPricePerToken = 1 << 30;  //1G bit/ppnt;
    
    constructor(address ta) public{
        TokenAddress = Token(ta);
    }
    
    
}