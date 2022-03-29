pragma solidity ^0.6.12;

import "./IVaultConfigProvider.sol";
import "../vendor/v0.6/RoleAd.sol";

contract VaultConfigProvider is RoleAd, IVaultConfigProvider {
    mapping(bytes32 => address) internal _addresses;
    /**
        * @dev Returns an address by id
        * @return The address
        */
    function getAddress(bytes32 id) public view returns (address) {
        return _addresses[id];
    }

    /**
    * @dev Sets an address for an id replacing the address saved in the addresses map
    * IMPORTANT Use this function carefully, as it will do a hard replacement
    * @param id The id
    * @param newAddress The address to set
    */
    function setAddress(bytes32 id, address newAddress) internal {
        _addresses[id] = newAddress;
        emit AddressSet(id, newAddress, false);
    }



    constructor() public {}

    bytes32 private constant GENESIS = 'GENESIS';


    function setGenesis(address person) external override {
        setAddress(GENESIS, person);
    }

    function getGenesis() external override view returns (address){
        return getAddress(GENESIS);
    }

}
