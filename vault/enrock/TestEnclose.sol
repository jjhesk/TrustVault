pragma solidity ^0.6.12;

import {RoleIsland} from "../vendor/v0.6/RoleIsland.sol";
import {ReentrancyGuard} from "../vendor/v0.6/ReentrancyGuard.sol";
import {SafeMath, SafeMath16, SafeMath8} from "../vendor/v0.6/SafeMath.sol";
import {SafeERC20} from "../vendor/v0.6/SafeERC20.sol";
import {IERC20} from "../vendor/v0.6/IERC20.sol";
import {Address} from "../vendor/v0.6/Address.sol";

contract TestEnclose is RoleIsland {
    using SafeMath for uint256;
    using SafeMath16 for uint16;
    using SafeMath8 for uint8;
    using SafeERC20 for IERC20;

    event UsrDeposit(address a, uint256 b);

    mapping(bytes32 => uint256) public coin_user_vault;
    mapping(address => mapping(address => uint256)) public coin_user_vault2;

    modifier basicChecker(address coin, uint256 amount) {
        require(amount > 0, "BT:: invalid coin");
        require(Address.isContract(coin), "BT:: invalid coin contract");
        _;
    }

    function sig(address user, address coin) public pure returns (bytes32){
        return keccak256(abi.encodePacked("_", coin, user));
    }

    function deposit_erc20(address coin, uint256 amount) public basicChecker (coin, amount) {
        address usr = msg.sender;
        bytes32 _ss = sig(usr, coin);
        IERC20(coin).safeTransferFrom(usr, address(this), amount);
        coin_user_vault[_ss] = coin_user_vault[_ss].add(amount);

        uint256 original = coin_user_vault2[coin][usr];
        coin_user_vault2[coin][usr] = original.add(amount);

        emit UsrDeposit(coin, amount);
    }


}
