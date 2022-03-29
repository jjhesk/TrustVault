pragma solidity ^0.6.12;

import "../vendor/v0.6/RoleIsland.sol";
import "../vendor/v0.6/ReentrancyGuard.sol";
import "./IVaultConfigProvider.sol";
import "../vendor/v0.6/SafeMath.sol";
import "../vendor/v0.6/SafeERC20.sol";

contract BaseVault is RoleIsland, ReentrancyGuard {
    using SafeMath for uint256;
    using SafeMath16 for uint16;
    using SafeMath8 for uint8;
    using SafeERC20 for IERC20;
    IVaultConfigProvider internal provider;

    function _updateProvider(address target) internal {
        provider = IVaultConfigProvider(target);
    }

    struct GameRound {
        uint256 dot;
        uint16 users;
        string hash;
    }

    // coin address
    mapping(address => uint256) public bank_vault;
    mapping(address => uint256) public pool_vault;
    // coin address => user address
    mapping(bytes32 => uint256) internal coin_user_vault;
    mapping(bytes32 => uint256) internal coin_user_rlock;
    // coin address and discount in percentage
    mapping(address => uint16) public pool_fee;
    // user and discount in percentage
    mapping(address => uint16) public pool_fee_discount;
    mapping(uint256 => GameRound) public round_vault;
    uint256 public game_round;

    event UsrDeposit(address coin, uint256 amount);
    event UsrWithdraw(address coin, uint256 amount);
    event DonateDeposit(address coin, uint256 amount);
    event WithdrawalGov(address coin, address beneficier, uint256 amount);
    event UpdatedWallet(address coin, address beneficier);
    event UpdatedBank(address coin, uint256 amount);
    event GameVault(string h, uint256 dot);

    modifier basicChecker(address coin, uint256 amount) {
        require(amount > 0, "BT:: invalid coin");
        require(Address.isContract(coin), "BT:: invalid coin contract");
        _;
    }

    function check(address coin, address user) external view returns (uint256[6] memory chr){
        bytes32 s = signature(user, coin);
        chr[0] = coin_user_vault[s];
        chr[1] = coin_user_rlock[s];
        chr[2] = bank_vault[coin];
        chr[3] = pool_vault[coin];
        chr[4] = pool_fee[coin];
        chr[5] = pool_fee_discount[user];
        return chr;
    }

    function signature(address user, address coin) public pure returns (bytes32){
        return keccak256(abi.encodePacked("_", coin, user));
    }

    function _updateBank(address coin, uint256 new_vault_balance) internal {
        uint256 ovault = bank_vault[coin];
        if (new_vault_balance == ovault) {
            return;
        }
        bank_vault[coin] = new_vault_balance;
        emit UpdatedBank(coin, new_vault_balance);
    }

    function _updateUser(address user, address coin, uint256 nw_vault) internal {
        bytes32 s = signature(user, coin);
        uint256 ovault = coin_user_vault[s];
        if (nw_vault == ovault) {
            return;
        }
        //coin_user_vault[coin][user] = nw_vault;
        if (nw_vault > ovault) {
            uint256 offset = nw_vault.sub(ovault);
            bank_vault[coin] = bank_vault[coin].sub(offset);
        } else {
            uint256 offset = ovault.sub(nw_vault);
            bank_vault[coin] = bank_vault[coin].add(offset);
        }
        coin_user_vault[s] = nw_vault;
        emit UpdatedWallet(coin, user);
    }

    function set_fee(address x, uint16 percent_1000) external nonReentrant onlyWhitelistAdmin {
        require(percent_1000 < 1000 && percent_1000 > 0, "BT:: wrong int");
        pool_fee[x] = percent_1000;
    }

    function set_round(uint256 n_round) external nonReentrant onlyWhitelistAdmin {
        require(n_round > 0, "BT:: wrong int");
        game_round = n_round;
    }

    function set_discount(address user, uint16 perc) external nonReentrant onlyWhitelistAdmin {
        require(perc < 100 && perc > 0, "BT:: wrong int");
        pool_fee_discount[user] = perc;
    }

    function with_fee(address x, address user, uint256 y) internal view returns (uint256 after_c, uint256 fee){
        fee = 0;
        if (pool_fee[x] > 0) {
            uint256 discount = 0;
            uint256 deduct = 0;
            if (pool_fee_discount[user] > 0) {
                discount = pool_fee_discount[user];
                fee = y.mul(uint256(pool_fee[x])).div(1000);
                deduct = fee.mul(discount).div(100);
                fee = fee.sub(deduct);
            } else {
                fee = y.mul(uint256(pool_fee[x])).div(1000);
            }
            after_c = y.sub(fee);
            return (after_c, fee);
        } else {
            after_c = y;
            return (after_c, 0);
        }
    }

    function _drawcoin_erc20(bytes32 ss, address user, address coin, uint256 amount) internal basicChecker (coin, amount) {
        uint256 vault = coin_user_vault[ss];
        require(vault >= amount, "BT:: vault is not enough");
        //require(lock_at + 4 < game_round, "BT:: it is currently in lockup");
        IERC20(coin).safeTransfer(user, amount);
        coin_user_vault[ss] = vault.sub(amount);
        pool_vault[coin] = pool_vault[coin].sub(amount);
        emit UsrWithdraw(coin, amount);
    }

}

contract Enrockvault is BaseVault {
    constructor(address _provider) public {
        _updateProvider(_provider);
        game_round = 0;
    }

    function updateProvider(address k) external onlyWhitelistAdmin {
        _updateProvider(k);
    }

    function deposit_erc20(address coin, uint256 amount) public basicChecker (coin, amount) {
        address usr = _msgSender();
        bytes32 _ss = signature(usr, coin);
        (uint256 f, uint256 fee) = with_fee(coin, usr, amount);
        IERC20(coin).safeTransferFrom(usr, address(this), f);
        if (fee > 0) {
            IERC20(coin).safeTransferFrom(usr, provider.getGenesis(), fee);
        }
        pool_vault[coin] = pool_vault[coin].add(f);
        coin_user_vault[_ss] = coin_user_vault[_ss].add(f);
        if (coin_user_rlock[_ss] != game_round) {
            coin_user_rlock[_ss] = game_round;
        }
        emit UsrDeposit(coin, f);
    }

    function donate_erc20(address coin, uint256 amount) public basicChecker (coin, amount) {
        IERC20(coin).safeTransferFrom(_msgSender(), address(this), amount);
        bank_vault[coin] = bank_vault[coin].add(amount);
        pool_vault[coin] = pool_vault[coin].add(amount);
        emit DonateDeposit(coin, amount);
    }

    function take_usr(address user, address coin, uint256 amount) external nonReentrant onlyGovernor {
        bytes32 _ss = signature(user, coin);
        uint256 lock_at = coin_user_rlock[_ss];
        require(game_round > lock_at, "BT:: it is currently in lockup");
        _drawcoin_erc20(_ss, user, coin, amount);
    }

    function up_single(address user, address coin, uint256 amount) external nonReentrant onlyGovernor {
        _updateUser(user, coin, amount);
    }

    function up_usr_batch(address[] calldata users, address[] calldata coins, uint256[] memory amounts) external nonReentrant onlyGovernor {
        require(users.length == coins.length, "BT:: wrong number");
        require(amounts.length == coins.length, "BT:: wrong number");
        require(users.length == amounts.length, "BT:: wrong number");
        uint8 _looper = uint8(users.length);
        for (uint8 i = 0; i < _looper; i++) {
            address user = users[i];
            address coin = coins[i];
            uint256 amount = amounts[i];
            _updateUser(user, coin, amount);
        }
    }

    function up_bank_vault(address[] calldata coins, uint256[] memory balances) external nonReentrant onlyGovernor {
        require(balances.length == coins.length, "BT:: wrong number");
        uint8 _looper = uint8(balances.length);
        for (uint8 i = 0; i < _looper; i++) {
            address coin = coins[i];
            uint256 amount = balances[i];
            _updateBank(coin, amount);
        }
    }

    function oracle_post_result(uint256 round, uint256 dot, string calldata hash, uint16 total_users) external nonReentrant onlyGovernor {
        round_vault[round].users = total_users;
        round_vault[round].dot = dot;
        round_vault[round].hash = hash;
        game_round = round;
        emit GameVault(hash, dot);
    }

}
