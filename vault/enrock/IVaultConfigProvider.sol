pragma solidity ^0.6.12;

interface IVaultConfigProvider {
    function setGenesis(address person) external;
    function getGenesis() external view returns (address);
    event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy);
}