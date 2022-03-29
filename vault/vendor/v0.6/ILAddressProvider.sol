pragma solidity ^0.6.6;


interface ILAddressProvider {
    event USDTokenUpdated(address indexed newAddress);
    event FarmTokenUpdated(address indexed newAddress);
    event OracleUpdated(address indexed newAddress);
    event VRFUpdated(address indexed newAddress);
    event FarmUpdated(address indexed newAddress);
    event AlgoUpdated(address indexed newAddress);
    event SplitGroupUpdated(address indexed newAddress);
    event ReferralUpdated(address indexed newAddress);

    /**
    default functions
    **/

    event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy);
    event ProxyCreated(bytes32 id, address indexed newAddress);

    function setAddress(bytes32 id, address newAddress) external;

    function setAddressAsProxy(bytes32 id, address impl) external;

    function getAddress(bytes32 id) external view returns (address);

    function getUSDToken() external view returns (address);

    function getFarmToken() external view returns (address);

    function setUSDToken(address token) external;

    function setFarmToken(address token) external;

    function setOracle(address oracle) external;

    function getOracle() external view returns (address);

    function setReferralCola(address cola) external;

    function setFarmField(address farm) external;

    function getReferralCola() external view returns (address);

    function getFarmField() external view returns (address);

    function setVrf(address module) external;

    function setAlgoModel(address model) external;

    function getVrf() external view returns (address);

    function getAlgoModel() external view returns (address);

    function setGenesis(address person) external;

    function getGenesis() external view returns (address);

    function setSplitGroup(address group) external;

    function getSplitGroup() external view returns (address);
}