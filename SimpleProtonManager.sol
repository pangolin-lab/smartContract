pragma solidity >=0.4.24;
import "./SafeMath.sol";
import "./Owned.sol";

contract SimpleProtonManager is owned{
    
    using SafeMath for uint256;
    
    mapping(bytes32=>address) public BoundedEthAddress;
    mapping(address=>bytes32[]) public AddressesUnderMyAccount; 
    mapping(address=>uint) public AddressesCounter; 
    uint256 public servicePrice = 1 ether;
    
    constructor() public{
        servicePrice = 1 ether;
    }
    
    function multiUnbind(bytes32[] memory addresses) public{
        uint available = 0;
        address admin = msg.sender;
        
        for (uint i = 0; i < addresses.length; i++){
             bytes32 addr = addresses[i];
            if (BoundedEthAddress[addr] != admin){
                continue;
            }
            
            available = available.add(1);
            delete BoundedEthAddress[addr];
            removeItem(addr);
        }
        
        AddressesCounter[admin] = AddressesCounter[admin].sub(available);
    }
    
    function multiBind(bytes32[] memory addresses) public{
        uint available = 0;
        address admin = msg.sender;
        
        for (uint i = 0; i < addresses.length; i++){
            bytes32 addr = addresses[i];
            if (BoundedEthAddress[addr] != address(0)){
                continue;
            }
            available = available.add(1);
            BoundedEthAddress[addr] = admin;
            AddressesUnderMyAccount[admin].push(addr);
        }
        AddressesCounter[admin] = AddressesCounter[admin].add(available);
    }
    
    function bind(bytes32 addr) public{
        require(BoundedEthAddress[addr] == address(0));
       
        BoundedEthAddress[addr] = msg.sender;
        AddressesUnderMyAccount[msg.sender].push(addr);
        AddressesCounter[msg.sender] = AddressesCounter[msg.sender].add(1);
    }
    
    function unbind(bytes32 addr) public{
        require(BoundedEthAddress[addr] == msg.sender);
        delete BoundedEthAddress[addr];
        AddressesCounter[msg.sender] = AddressesCounter[msg.sender].sub(1);
        removeItem(addr);
    }
    
    function removeItem(bytes32 addr) private{
        bytes32[] storage members = AddressesUnderMyAccount[msg.sender];
        for (uint i = 0; i < members.length; i++){
            if (addr == members[i]){
                delete members[i];
                return;
            }
        }
    }
    
    function check(bytes32 addr) public view returns(bool){
        address admin = BoundedEthAddress[addr];
        if (admin == address(0)){
            return false;
        }
        
        uint sumNeeded = servicePrice.mul(AddressesCounter[admin]);
        if (sumNeeded > admin.balance){
            return false;
        }
        
        return true;
    }
    
    function search(address ethAddr) public view returns(bytes32[]memory){
        return AddressesUnderMyAccount[ethAddr];
    }
    
     function changeServicePrice(uint256 price) public onlyOwner{
        servicePrice = price;
    }
}