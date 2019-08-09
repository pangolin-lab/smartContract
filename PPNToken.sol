pragma solidity >=0.4.21;

import "./ERC20.sol";

contract PPNToken is ERC20{
    
    string  public constant  name = "Proton Protocol Network Token";
    string  public constant  symbol = "PPNT";
    uint8   public constant  decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 1024e6 * (10 ** uint256(decimals));

    mapping (address => uint256) private _freezeToContract;
    address public FreezeTokenHolder;

    constructor() public{
        _mint(msg.sender, INITIAL_SUPPLY);
    }

    function getDeccimal() public pure returns (uint) {
        return 10 ** uint256(decimals);
    }

    function getAllowanceToken(address to) public view returns (uint){
        return _freezeToContract[to];
    }

    function transferFrom(address from, address to, uint256 value ) public returns (bool) {

        // require(value <= balanceOf(from));
        // require(value <= _allowed[from][msg.sender]);
        // require(to != address(0));

        // _balances[from] = _balances[from].sub(value);
        // _balances[to] = _balances[to].add(value);
        // _allowed[from][msg.sender] = _allowed[from][msg.sender].sub(value);
        // emit Transfer(from, to, value);
        return true;
    }
}