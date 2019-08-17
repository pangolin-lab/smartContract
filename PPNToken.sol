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
}