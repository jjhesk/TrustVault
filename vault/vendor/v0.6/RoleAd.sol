// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./AccessControl.sol";
import "./Ownable.sol";

contract RoleAd is AccessControl, Ownable {
    bytes32 public constant whitelistAdmins = keccak256("_whitelistAdmins");

    function isOwner() public view returns (bool){
        return owner() == _msgSender();
    }

    //role end
    modifier onlyWhitelistAdmin() {
        require(hasRole(whitelistAdmins, _msgSender()) || isOwner(), "no auth");
        _;
    }

    function isWhitelistAdmin(address account) public view returns (bool) {
        return hasRole(whitelistAdmins, account) || isOwner();
    }

    function addWhitelistAdmin(address account) public onlyOwner {
        _grantRole(whitelistAdmins, account);
    }

    function removeWhitelistAdmin(address account) public onlyOwner {
        _revokeRole(whitelistAdmins, account);
    }


    modifier isHuman() {
        address q = _msgSender();
        uint codeLength;
        assembly {codeLength := extcodesize(q)}
        require(codeLength == 0, "humans only");
        require(tx.origin == _msgSender(), "humans only");
        _;
    }


}