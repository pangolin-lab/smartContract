pragma solidity >=0.4.24;
import "owned.sol"

contract BlackBox is owned{
    
    address public MicroPaymentContract;
    
    function Upgrade(address newAddress) public onlyOwner{
        MicroPaymentContract = newAddress;
    }
}