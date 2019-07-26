pragma solidity >=0.4.24;
import "./SafeMath.sol";
import "./Owned.sol";

interface token {
    function balanceOf(address account) external view returns (uint);
}

contract ProtonManager is owned{
    using SafeMath for uint256;
    
    uint256 constant decimals = 10 ** 18;
    uint256 public TokenNoPerProtonAccount = (1 << 10) * decimals;
    
    token public ProtonCoin;
    mapping(bytes32=>address) public ProtonUserMap;
    mapping(address=>uint) public EtherCounter;
    
    constructor(address protonTokenAddress) public{
        ProtonCoin = token(protonTokenAddress);
    }
    
    function multiUnbind(bytes32[] memory protonAddresses) public{
        uint no = protonAddresses.length;
        require(no > 1);
        
        uint actualNo = 0;
        for (uint i = 0; i < protonAddresses.length; i++){
            bytes32 protonAddr = protonAddresses[i];
            
            if (ProtonUserMap[protonAddr] != msg.sender){
                continue;
            }
            
            actualNo ++; 
            delete ProtonUserMap[protonAddr];
        }
        
        EtherCounter[msg.sender] = EtherCounter[msg.sender].sub(actualNo);
    }
    
    function multiBind(bytes32[] memory protonAddresses) public{
        uint no = protonAddresses.length;
        require(no > 1);
        
        uint protonBalance = ProtonCoin.balanceOf(msg.sender);
        uint totalNeed = TokenNoPerProtonAccount.mul(no + EtherCounter[msg.sender]);
        require(protonBalance >= totalNeed);
        
        uint actualNo = 0;
        for (uint i = 0; i < protonAddresses.length; i++){
            
            if (ProtonUserMap[protonAddresses[i]] != address(0)){
                continue;
            }
            
            actualNo ++; 
            ProtonUserMap[protonAddresses[i]] = msg.sender;
        }
        
        EtherCounter[msg.sender] = EtherCounter[msg.sender].add(actualNo);
    }
    
    function unbindProtonAddress(bytes32 protonAddress) public{
        address userAddr = ProtonUserMap[protonAddress];
        require(userAddr == msg.sender);
        
        delete ProtonUserMap[protonAddress];
        EtherCounter[msg.sender] = EtherCounter[msg.sender].sub(1);
    }
    
    function bindProtonAddress(bytes32 protonAddress) public{
        require(ProtonUserMap[protonAddress] == address(0));
        
        uint protonBalance = ProtonCoin.balanceOf(msg.sender);
        uint no = EtherCounter[msg.sender];
        uint totalNeed = TokenNoPerProtonAccount.mul(no + 1);
        require(protonBalance >= totalNeed);
        
        EtherCounter[msg.sender] = EtherCounter[msg.sender].add(1);
        ProtonUserMap[protonAddress] = msg.sender;
    }
    
    function checkProtonAddress(bytes32 protonAddress) public view returns(address, uint, uint){
        address userAddr = ProtonUserMap[protonAddress];
        if (userAddr == address(0)){
            return (address(0), 0, 0);
        }
        
        uint protonBalance = ProtonCoin.balanceOf(userAddr);
        return (userAddr, protonBalance, EtherCounter[userAddr]);
    }
    
    function checkBinder(address userAddr) public view returns(uint, uint, uint){ 
        uint ethBalance = userAddr.balance;
        uint protonBalance = ProtonCoin.balanceOf(userAddr);
        uint bindNo = EtherCounter[userAddr]; 
        return (ethBalance, protonBalance, bindNo);
    }
    
    function changeServicePrice(uint256 price) public onlyOwner{
        TokenNoPerProtonAccount = price.mul(decimals);
    }
}