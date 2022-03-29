"""Generated wrapper for TestEnclose Solidity contract."""

# pylint: disable=too-many-arguments

import json
from typing import (  # pylint: disable=unused-import
    Any,
    List,
    Optional,
    Tuple,
    Union,
)
import time
from eth_utils import to_checksum_address
from mypy_extensions import TypedDict  # pylint: disable=unused-import
from hexbytes import HexBytes
from web3 import Web3
from web3.contract import ContractFunction
from web3.datastructures import AttributeDict
from web3.providers.base import BaseProvider
from web3.exceptions import ContractLogicError
from moody.m.bases import ContractMethod, Validator, ContractBase, Signatures
from moody.m.tx_params import TxParams
from moody.libeb import MiliDoS
from moody import Bolors

# Try to import a custom validator class definition; if there isn't one,
# declare one that we can instantiate for the default argument to the
# constructor for TestEnclose below.
try:
    # both mypy and pylint complain about what we're doing here, but this
    # works just fine, so their messages have been disabled here.
    from . import (  # type: ignore # pylint: disable=import-self
        TestEncloseValidator,
    )
except ImportError:

    class TestEncloseValidator(  # type: ignore
        Validator
    ):
        """No-op input validator."""

try:
    from .middleware import MIDDLEWARE  # type: ignore
except ImportError:
    pass





class DefaultAdminRoleMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the DEFAULT_ADMIN_ROLE method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("DEFAULT_ADMIN_ROLE")



    def block_call(self, debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: default_admin_role")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, default_admin_role: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, default_admin_role. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class AddGovernorMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the addGovernor method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("addGovernor")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the addGovernor method."""
        self.validator.assert_valid(
            method_name='addGovernor',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: add_governor")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, add_governor: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, add_governor. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class AddWhitelistAdminMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the addWhitelistAdmin method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("addWhitelistAdmin")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the addWhitelistAdmin method."""
        self.validator.assert_valid(
            method_name='addWhitelistAdmin',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: add_whitelist_admin")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, add_whitelist_admin: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, add_whitelist_admin. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class CoinUserVaultMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the coin_user_vault method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("coin_user_vault")

    def validate_and_normalize_inputs(self, index_0: Union[bytes, str])->any:
        """Validate the inputs to the coin_user_vault method."""
        self.validator.assert_valid(
            method_name='coin_user_vault',
            parameter_name='index_0',
            argument_value=index_0,
        )
        return (index_0)



    def block_call(self,index_0: Union[bytes, str], debug:bool=False) -> int:
        _fn = self._underlying_method(index_0)
        returned = _fn.call({
                'from': self._operate
            })
        return int(returned)
    def block_send(self, index_0: Union[bytes, str],_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> int:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(index_0)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: coin_user_vault")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, coin_user_vault: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, coin_user_vault. Reason: Unknown")


    def send_transaction(self, index_0: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (index_0) = self.validate_and_normalize_inputs(index_0)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0).transact(tx_params.as_dict())

    def build_transaction(self, index_0: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (index_0) = self.validate_and_normalize_inputs(index_0)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, index_0: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (index_0) = self.validate_and_normalize_inputs(index_0)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0).estimateGas(tx_params.as_dict())

class CoinUserVault2Method(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the coin_user_vault2 method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("coin_user_vault2")

    def validate_and_normalize_inputs(self, index_0: str, index_1: str)->any:
        """Validate the inputs to the coin_user_vault2 method."""
        self.validator.assert_valid(
            method_name='coin_user_vault2',
            parameter_name='index_0',
            argument_value=index_0,
        )
        index_0 = self.validate_and_checksum_address(index_0)
        self.validator.assert_valid(
            method_name='coin_user_vault2',
            parameter_name='index_1',
            argument_value=index_1,
        )
        index_1 = self.validate_and_checksum_address(index_1)
        return (index_0, index_1)



    def block_call(self,index_0: str, index_1: str, debug:bool=False) -> int:
        _fn = self._underlying_method(index_0, index_1)
        returned = _fn.call({
                'from': self._operate
            })
        return int(returned)
    def block_send(self, index_0: str, index_1: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> int:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(index_0, index_1)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: coin_user_vault2")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, coin_user_vault2: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, coin_user_vault2. Reason: Unknown")


    def send_transaction(self, index_0: str, index_1: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (index_0, index_1) = self.validate_and_normalize_inputs(index_0, index_1)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0, index_1).transact(tx_params.as_dict())

    def build_transaction(self, index_0: str, index_1: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (index_0, index_1) = self.validate_and_normalize_inputs(index_0, index_1)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0, index_1).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, index_0: str, index_1: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (index_0, index_1) = self.validate_and_normalize_inputs(index_0, index_1)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(index_0, index_1).estimateGas(tx_params.as_dict())

class DepositErc20Method(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the deposit_erc20 method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("deposit_erc20")

    def validate_and_normalize_inputs(self, coin: str, amount: int)->any:
        """Validate the inputs to the deposit_erc20 method."""
        self.validator.assert_valid(
            method_name='deposit_erc20',
            parameter_name='coin',
            argument_value=coin,
        )
        coin = self.validate_and_checksum_address(coin)
        self.validator.assert_valid(
            method_name='deposit_erc20',
            parameter_name='amount',
            argument_value=amount,
        )
        # safeguard against fractional inputs
        amount = int(amount)
        return (coin, amount)



    def block_send(self, coin: str, amount: int,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(coin, amount)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: deposit_erc20")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, deposit_erc20: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, deposit_erc20. Reason: Unknown")


    def send_transaction(self, coin: str, amount: int, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (coin, amount) = self.validate_and_normalize_inputs(coin, amount)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(coin, amount).transact(tx_params.as_dict())

    def build_transaction(self, coin: str, amount: int, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (coin, amount) = self.validate_and_normalize_inputs(coin, amount)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(coin, amount).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, coin: str, amount: int, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (coin, amount) = self.validate_and_normalize_inputs(coin, amount)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(coin, amount).estimateGas(tx_params.as_dict())

class GetRoleAdminMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the getRoleAdmin method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("getRoleAdmin")

    def validate_and_normalize_inputs(self, role: Union[bytes, str])->any:
        """Validate the inputs to the getRoleAdmin method."""
        self.validator.assert_valid(
            method_name='getRoleAdmin',
            parameter_name='role',
            argument_value=role,
        )
        return (role)



    def block_call(self,role: Union[bytes, str], debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method(role)
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)
    def block_send(self, role: Union[bytes, str],_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: get_role_admin")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_admin: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_admin. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).estimateGas(tx_params.as_dict())

class GetRoleMemberMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the getRoleMember method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("getRoleMember")

    def validate_and_normalize_inputs(self, role: Union[bytes, str], index: int)->any:
        """Validate the inputs to the getRoleMember method."""
        self.validator.assert_valid(
            method_name='getRoleMember',
            parameter_name='role',
            argument_value=role,
        )
        self.validator.assert_valid(
            method_name='getRoleMember',
            parameter_name='index',
            argument_value=index,
        )
        # safeguard against fractional inputs
        index = int(index)
        return (role, index)



    def block_call(self,role: Union[bytes, str], index: int, debug:bool=False) -> str:
        _fn = self._underlying_method(role, index)
        returned = _fn.call({
                'from': self._operate
            })
        return str(returned)
    def block_send(self, role: Union[bytes, str], index: int,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> str:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role, index)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: get_role_member")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_member: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_member. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], index: int, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role, index) = self.validate_and_normalize_inputs(role, index)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, index).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], index: int, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role, index) = self.validate_and_normalize_inputs(role, index)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, index).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], index: int, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role, index) = self.validate_and_normalize_inputs(role, index)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, index).estimateGas(tx_params.as_dict())

class GetRoleMemberCountMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the getRoleMemberCount method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("getRoleMemberCount")

    def validate_and_normalize_inputs(self, role: Union[bytes, str])->any:
        """Validate the inputs to the getRoleMemberCount method."""
        self.validator.assert_valid(
            method_name='getRoleMemberCount',
            parameter_name='role',
            argument_value=role,
        )
        return (role)



    def block_call(self,role: Union[bytes, str], debug:bool=False) -> int:
        _fn = self._underlying_method(role)
        returned = _fn.call({
                'from': self._operate
            })
        return int(returned)
    def block_send(self, role: Union[bytes, str],_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> int:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: get_role_member_count")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_member_count: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_role_member_count. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role) = self.validate_and_normalize_inputs(role)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role).estimateGas(tx_params.as_dict())

class GovernorMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the governor method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("governor")



    def block_call(self, debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: governor")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, governor: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, governor. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class GrantRoleMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the grantRole method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("grantRole")

    def validate_and_normalize_inputs(self, role: Union[bytes, str], account: str)->any:
        """Validate the inputs to the grantRole method."""
        self.validator.assert_valid(
            method_name='grantRole',
            parameter_name='role',
            argument_value=role,
        )
        self.validator.assert_valid(
            method_name='grantRole',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (role, account)



    def block_send(self, role: Union[bytes, str], account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role, account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: grant_role")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, grant_role: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, grant_role. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).estimateGas(tx_params.as_dict())

class HasRoleMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the hasRole method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("hasRole")

    def validate_and_normalize_inputs(self, role: Union[bytes, str], account: str)->any:
        """Validate the inputs to the hasRole method."""
        self.validator.assert_valid(
            method_name='hasRole',
            parameter_name='role',
            argument_value=role,
        )
        self.validator.assert_valid(
            method_name='hasRole',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (role, account)



    def block_call(self,role: Union[bytes, str], account: str, debug:bool=False) -> bool:
        _fn = self._underlying_method(role, account)
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, role: Union[bytes, str], account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role, account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: has_role")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, has_role: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, has_role. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).estimateGas(tx_params.as_dict())

class IsGovernorMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the isGovernor method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("isGovernor")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the isGovernor method."""
        self.validator.assert_valid(
            method_name='isGovernor',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_call(self,account: str, debug:bool=False) -> bool:
        _fn = self._underlying_method(account)
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: is_governor")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_governor: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_governor. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class IsLockedMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the isLocked method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("isLocked")



    def block_call(self, debug:bool=False) -> bool:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: is_locked")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_locked: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_locked. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class IsOwnerMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the isOwner method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("isOwner")



    def block_call(self, debug:bool=False) -> bool:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: is_owner")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_owner: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_owner. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class IsPausedMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the isPaused method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("isPaused")



    def block_call(self, debug:bool=False) -> bool:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: is_paused")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_paused: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_paused. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class IsWhitelistAdminMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the isWhitelistAdmin method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("isWhitelistAdmin")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the isWhitelistAdmin method."""
        self.validator.assert_valid(
            method_name='isWhitelistAdmin',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_call(self,account: str, debug:bool=False) -> bool:
        _fn = self._underlying_method(account)
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)
    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: is_whitelist_admin")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_whitelist_admin: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, is_whitelist_admin. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class LockMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the lock method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("lock")



    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: lock")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, lock: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, lock. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class OwnerMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the owner method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("owner")



    def block_call(self, debug:bool=False) -> str:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return str(returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> str:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: owner")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, owner: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, owner. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class PauseScMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the pauseSc method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("pauseSc")



    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: pause_sc")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, pause_sc: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, pause_sc. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class RemoveGovernorMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the removeGovernor method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("removeGovernor")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the removeGovernor method."""
        self.validator.assert_valid(
            method_name='removeGovernor',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: remove_governor")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, remove_governor: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, remove_governor. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class RemoveWhitelistAdminMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the removeWhitelistAdmin method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("removeWhitelistAdmin")

    def validate_and_normalize_inputs(self, account: str)->any:
        """Validate the inputs to the removeWhitelistAdmin method."""
        self.validator.assert_valid(
            method_name='removeWhitelistAdmin',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (account)



    def block_send(self, account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: remove_whitelist_admin")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, remove_whitelist_admin: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, remove_whitelist_admin. Reason: Unknown")


    def send_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).transact(tx_params.as_dict())

    def build_transaction(self, account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (account) = self.validate_and_normalize_inputs(account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(account).estimateGas(tx_params.as_dict())

class RenounceOwnershipMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the renounceOwnership method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("renounceOwnership")



    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: renounce_ownership")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, renounce_ownership: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, renounce_ownership. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class RenounceRoleMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the renounceRole method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("renounceRole")

    def validate_and_normalize_inputs(self, role: Union[bytes, str], account: str)->any:
        """Validate the inputs to the renounceRole method."""
        self.validator.assert_valid(
            method_name='renounceRole',
            parameter_name='role',
            argument_value=role,
        )
        self.validator.assert_valid(
            method_name='renounceRole',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (role, account)



    def block_send(self, role: Union[bytes, str], account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role, account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: renounce_role")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, renounce_role: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, renounce_role. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).estimateGas(tx_params.as_dict())

class RevokeRoleMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the revokeRole method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("revokeRole")

    def validate_and_normalize_inputs(self, role: Union[bytes, str], account: str)->any:
        """Validate the inputs to the revokeRole method."""
        self.validator.assert_valid(
            method_name='revokeRole',
            parameter_name='role',
            argument_value=role,
        )
        self.validator.assert_valid(
            method_name='revokeRole',
            parameter_name='account',
            argument_value=account,
        )
        account = self.validate_and_checksum_address(account)
        return (role, account)



    def block_send(self, role: Union[bytes, str], account: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(role, account)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: revoke_role")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, revoke_role: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, revoke_role. Reason: Unknown")


    def send_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).transact(tx_params.as_dict())

    def build_transaction(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, role: Union[bytes, str], account: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (role, account) = self.validate_and_normalize_inputs(role, account)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(role, account).estimateGas(tx_params.as_dict())

class SigMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the sig method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("sig")

    def validate_and_normalize_inputs(self, user: str, coin: str)->any:
        """Validate the inputs to the sig method."""
        self.validator.assert_valid(
            method_name='sig',
            parameter_name='user',
            argument_value=user,
        )
        user = self.validate_and_checksum_address(user)
        self.validator.assert_valid(
            method_name='sig',
            parameter_name='coin',
            argument_value=coin,
        )
        coin = self.validate_and_checksum_address(coin)
        return (user, coin)


    def block_call(self,user: str, coin: str, debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method(user, coin)
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)

    def block_send(self, user: str, coin: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(user, coin)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: sig")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, sig: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, sig. Reason: Unknown")


    def send_transaction(self, user: str, coin: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (user, coin) = self.validate_and_normalize_inputs(user, coin)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(user, coin).transact(tx_params.as_dict())

    def build_transaction(self, user: str, coin: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (user, coin) = self.validate_and_normalize_inputs(user, coin)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(user, coin).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, user: str, coin: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (user, coin) = self.validate_and_normalize_inputs(user, coin)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(user, coin).estimateGas(tx_params.as_dict())

class TransferOwnershipMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the transferOwnership method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("transferOwnership")

    def validate_and_normalize_inputs(self, new_owner: str)->any:
        """Validate the inputs to the transferOwnership method."""
        self.validator.assert_valid(
            method_name='transferOwnership',
            parameter_name='newOwner',
            argument_value=new_owner,
        )
        new_owner = self.validate_and_checksum_address(new_owner)
        return (new_owner)



    def block_send(self, new_owner: str,_gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(new_owner)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: transfer_ownership")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, transfer_ownership: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, transfer_ownership. Reason: Unknown")


    def send_transaction(self, new_owner: str, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (new_owner) = self.validate_and_normalize_inputs(new_owner)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(new_owner).transact(tx_params.as_dict())

    def build_transaction(self, new_owner: str, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (new_owner) = self.validate_and_normalize_inputs(new_owner)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(new_owner).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, new_owner: str, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (new_owner) = self.validate_and_normalize_inputs(new_owner)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(new_owner).estimateGas(tx_params.as_dict())

class UnlockMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the unlock method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("unlock")



    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: unlock")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, unlock: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, unlock. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class UnpauseScMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the unpauseSc method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("unpauseSc")



    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> None:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: unpause_sc")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, unpause_sc: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, unpause_sc. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class WhitelistAdminsMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the whitelistAdmins method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("whitelistAdmins")



    def block_call(self, debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method()
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)
    def block_send(self, _gaswei:int,_pricewei:int,_valeth:int=0,_debugtx: bool = False,_receipList: bool = False) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method()
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': _gaswei,
                'gasPrice': _pricewei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if _debugtx:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if _receipList is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.waitForTransactionReceipt(txHash)
                    if _debugtx:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                print(f"======== TX blockHash ✅")
                if tx_receipt is not None:
                    print(f"{Bolors.OK}{tx_receipt.blockHash.hex()}{Bolors.RESET}")
                else:
                    print(f"{Bolors.WARNING}{txHash.hex()}{Bolors.RESET} - broadcast hash")

            if _receipList is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: whitelist_admins")

        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, whitelist_admins: {message}")
            else:
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, whitelist_admins. Reason: Unknown")


    def send_transaction(self, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().transact(tx_params.as_dict())

    def build_transaction(self, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().buildTransaction(tx_params.as_dict())

    def estimate_gas(self, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method().estimateGas(tx_params.as_dict())

class SignatureGenerator(Signatures):
    """
        The signature is generated for this and it is installed.
    """
    def __init__(self, abi: any):
        super().__init__(abi)

    def default_admin_role(self) -> str:
        return self._function_signatures["DEFAULT_ADMIN_ROLE"]
    def add_governor(self) -> str:
        return self._function_signatures["addGovernor"]
    def add_whitelist_admin(self) -> str:
        return self._function_signatures["addWhitelistAdmin"]
    def coin_user_vault(self) -> str:
        return self._function_signatures["coin_user_vault"]
    def coin_user_vault2(self) -> str:
        return self._function_signatures["coin_user_vault2"]
    def deposit_erc20(self) -> str:
        return self._function_signatures["deposit_erc20"]
    def get_role_admin(self) -> str:
        return self._function_signatures["getRoleAdmin"]
    def get_role_member(self) -> str:
        return self._function_signatures["getRoleMember"]
    def get_role_member_count(self) -> str:
        return self._function_signatures["getRoleMemberCount"]
    def governor(self) -> str:
        return self._function_signatures["governor"]
    def grant_role(self) -> str:
        return self._function_signatures["grantRole"]
    def has_role(self) -> str:
        return self._function_signatures["hasRole"]
    def is_governor(self) -> str:
        return self._function_signatures["isGovernor"]
    def is_locked(self) -> str:
        return self._function_signatures["isLocked"]
    def is_owner(self) -> str:
        return self._function_signatures["isOwner"]
    def is_paused(self) -> str:
        return self._function_signatures["isPaused"]
    def is_whitelist_admin(self) -> str:
        return self._function_signatures["isWhitelistAdmin"]
    def lock(self) -> str:
        return self._function_signatures["lock"]
    def owner(self) -> str:
        return self._function_signatures["owner"]
    def pause_sc(self) -> str:
        return self._function_signatures["pauseSc"]
    def remove_governor(self) -> str:
        return self._function_signatures["removeGovernor"]
    def remove_whitelist_admin(self) -> str:
        return self._function_signatures["removeWhitelistAdmin"]
    def renounce_ownership(self) -> str:
        return self._function_signatures["renounceOwnership"]
    def renounce_role(self) -> str:
        return self._function_signatures["renounceRole"]
    def revoke_role(self) -> str:
        return self._function_signatures["revokeRole"]
    def sig(self) -> str:
        return self._function_signatures["sig"]
    def transfer_ownership(self) -> str:
        return self._function_signatures["transferOwnership"]
    def unlock(self) -> str:
        return self._function_signatures["unlock"]
    def unpause_sc(self) -> str:
        return self._function_signatures["unpauseSc"]
    def whitelist_admins(self) -> str:
        return self._function_signatures["whitelistAdmins"]

# pylint: disable=too-many-public-methods,too-many-instance-attributes
class TestEnclose(ContractBase):
    """Wrapper class for TestEnclose Solidity contract."""
    _fn_default_admin_role: DefaultAdminRoleMethod
    """Constructor-initialized instance of
    :class:`DefaultAdminRoleMethod`.
    """

    _fn_add_governor: AddGovernorMethod
    """Constructor-initialized instance of
    :class:`AddGovernorMethod`.
    """

    _fn_add_whitelist_admin: AddWhitelistAdminMethod
    """Constructor-initialized instance of
    :class:`AddWhitelistAdminMethod`.
    """

    _fn_coin_user_vault: CoinUserVaultMethod
    """Constructor-initialized instance of
    :class:`CoinUserVaultMethod`.
    """

    _fn_coin_user_vault2: CoinUserVault2Method
    """Constructor-initialized instance of
    :class:`CoinUserVault2Method`.
    """

    _fn_deposit_erc20: DepositErc20Method
    """Constructor-initialized instance of
    :class:`DepositErc20Method`.
    """

    _fn_get_role_admin: GetRoleAdminMethod
    """Constructor-initialized instance of
    :class:`GetRoleAdminMethod`.
    """

    _fn_get_role_member: GetRoleMemberMethod
    """Constructor-initialized instance of
    :class:`GetRoleMemberMethod`.
    """

    _fn_get_role_member_count: GetRoleMemberCountMethod
    """Constructor-initialized instance of
    :class:`GetRoleMemberCountMethod`.
    """

    _fn_governor: GovernorMethod
    """Constructor-initialized instance of
    :class:`GovernorMethod`.
    """

    _fn_grant_role: GrantRoleMethod
    """Constructor-initialized instance of
    :class:`GrantRoleMethod`.
    """

    _fn_has_role: HasRoleMethod
    """Constructor-initialized instance of
    :class:`HasRoleMethod`.
    """

    _fn_is_governor: IsGovernorMethod
    """Constructor-initialized instance of
    :class:`IsGovernorMethod`.
    """

    _fn_is_locked: IsLockedMethod
    """Constructor-initialized instance of
    :class:`IsLockedMethod`.
    """

    _fn_is_owner: IsOwnerMethod
    """Constructor-initialized instance of
    :class:`IsOwnerMethod`.
    """

    _fn_is_paused: IsPausedMethod
    """Constructor-initialized instance of
    :class:`IsPausedMethod`.
    """

    _fn_is_whitelist_admin: IsWhitelistAdminMethod
    """Constructor-initialized instance of
    :class:`IsWhitelistAdminMethod`.
    """

    _fn_lock: LockMethod
    """Constructor-initialized instance of
    :class:`LockMethod`.
    """

    _fn_owner: OwnerMethod
    """Constructor-initialized instance of
    :class:`OwnerMethod`.
    """

    _fn_pause_sc: PauseScMethod
    """Constructor-initialized instance of
    :class:`PauseScMethod`.
    """

    _fn_remove_governor: RemoveGovernorMethod
    """Constructor-initialized instance of
    :class:`RemoveGovernorMethod`.
    """

    _fn_remove_whitelist_admin: RemoveWhitelistAdminMethod
    """Constructor-initialized instance of
    :class:`RemoveWhitelistAdminMethod`.
    """

    _fn_renounce_ownership: RenounceOwnershipMethod
    """Constructor-initialized instance of
    :class:`RenounceOwnershipMethod`.
    """

    _fn_renounce_role: RenounceRoleMethod
    """Constructor-initialized instance of
    :class:`RenounceRoleMethod`.
    """

    _fn_revoke_role: RevokeRoleMethod
    """Constructor-initialized instance of
    :class:`RevokeRoleMethod`.
    """

    _fn_sig: SigMethod
    """Constructor-initialized instance of
    :class:`SigMethod`.
    """

    _fn_transfer_ownership: TransferOwnershipMethod
    """Constructor-initialized instance of
    :class:`TransferOwnershipMethod`.
    """

    _fn_unlock: UnlockMethod
    """Constructor-initialized instance of
    :class:`UnlockMethod`.
    """

    _fn_unpause_sc: UnpauseScMethod
    """Constructor-initialized instance of
    :class:`UnpauseScMethod`.
    """

    _fn_whitelist_admins: WhitelistAdminsMethod
    """Constructor-initialized instance of
    :class:`WhitelistAdminsMethod`.
    """

    SIGNATURES:SignatureGenerator = None

    def __init__(
        self,
        core_lib: MiliDoS,
        contract_address: str,
        validator: TestEncloseValidator = None,
    ):
        """Get an instance of wrapper for smart contract.
        """
        # pylint: disable=too-many-statements
        super().__init__()
        self.contract_address = contract_address
        web3 = core_lib.w3

        if not validator:
            validator = TestEncloseValidator(web3, contract_address)




        # if any middleware was imported, inject it
        try:
            MIDDLEWARE
        except NameError:
            pass
        else:
            try:
                for middleware in MIDDLEWARE:
                    web3.middleware_onion.inject(
                         middleware['function'], layer=middleware['layer'],
                    )
            except ValueError as value_error:
                if value_error.args == ("You can't add the same un-named instance twice",):
                    pass

        self._web3_eth = web3.eth
        functions = self._web3_eth.contract(address=to_checksum_address(contract_address), abi=TestEnclose.abi()).functions
        signed = SignatureGenerator(TestEnclose.abi())
        validator.bindSignatures(signed)
        self.SIGNATURES = signed
        self._fn_default_admin_role = DefaultAdminRoleMethod(core_lib, contract_address, functions.DEFAULT_ADMIN_ROLE, validator)
        self._fn_add_governor = AddGovernorMethod(core_lib, contract_address, functions.addGovernor, validator)
        self._fn_add_whitelist_admin = AddWhitelistAdminMethod(core_lib, contract_address, functions.addWhitelistAdmin, validator)
        self._fn_coin_user_vault = CoinUserVaultMethod(core_lib, contract_address, functions.coin_user_vault, validator)
        self._fn_coin_user_vault2 = CoinUserVault2Method(core_lib, contract_address, functions.coin_user_vault2, validator)
        self._fn_deposit_erc20 = DepositErc20Method(core_lib, contract_address, functions.deposit_erc20, validator)
        self._fn_get_role_admin = GetRoleAdminMethod(core_lib, contract_address, functions.getRoleAdmin, validator)
        self._fn_get_role_member = GetRoleMemberMethod(core_lib, contract_address, functions.getRoleMember, validator)
        self._fn_get_role_member_count = GetRoleMemberCountMethod(core_lib, contract_address, functions.getRoleMemberCount, validator)
        self._fn_governor = GovernorMethod(core_lib, contract_address, functions.governor, validator)
        self._fn_grant_role = GrantRoleMethod(core_lib, contract_address, functions.grantRole, validator)
        self._fn_has_role = HasRoleMethod(core_lib, contract_address, functions.hasRole, validator)
        self._fn_is_governor = IsGovernorMethod(core_lib, contract_address, functions.isGovernor, validator)
        self._fn_is_locked = IsLockedMethod(core_lib, contract_address, functions.isLocked, validator)
        self._fn_is_owner = IsOwnerMethod(core_lib, contract_address, functions.isOwner, validator)
        self._fn_is_paused = IsPausedMethod(core_lib, contract_address, functions.isPaused, validator)
        self._fn_is_whitelist_admin = IsWhitelistAdminMethod(core_lib, contract_address, functions.isWhitelistAdmin, validator)
        self._fn_lock = LockMethod(core_lib, contract_address, functions.lock, validator)
        self._fn_owner = OwnerMethod(core_lib, contract_address, functions.owner, validator)
        self._fn_pause_sc = PauseScMethod(core_lib, contract_address, functions.pauseSc, validator)
        self._fn_remove_governor = RemoveGovernorMethod(core_lib, contract_address, functions.removeGovernor, validator)
        self._fn_remove_whitelist_admin = RemoveWhitelistAdminMethod(core_lib, contract_address, functions.removeWhitelistAdmin, validator)
        self._fn_renounce_ownership = RenounceOwnershipMethod(core_lib, contract_address, functions.renounceOwnership, validator)
        self._fn_renounce_role = RenounceRoleMethod(core_lib, contract_address, functions.renounceRole, validator)
        self._fn_revoke_role = RevokeRoleMethod(core_lib, contract_address, functions.revokeRole, validator)
        self._fn_sig = SigMethod(core_lib, contract_address, functions.sig, validator)
        self._fn_transfer_ownership = TransferOwnershipMethod(core_lib, contract_address, functions.transferOwnership, validator)
        self._fn_unlock = UnlockMethod(core_lib, contract_address, functions.unlock, validator)
        self._fn_unpause_sc = UnpauseScMethod(core_lib, contract_address, functions.unpauseSc, validator)
        self._fn_whitelist_admins = WhitelistAdminsMethod(core_lib, contract_address, functions.whitelistAdmins, validator)

    
    
    def event_ownership_transferred(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event ownership_transferred in contract TestEnclose
        Get log entry for OwnershipTransferred event.
                :param tx_hash: hash of transaction emitting OwnershipTransferred event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.OwnershipTransferred().processReceipt(tx_receipt)
    
    
    def event_role_granted(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event role_granted in contract TestEnclose
        Get log entry for RoleGranted event.
                :param tx_hash: hash of transaction emitting RoleGranted event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.RoleGranted().processReceipt(tx_receipt)
    
    
    def event_role_revoked(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event role_revoked in contract TestEnclose
        Get log entry for RoleRevoked event.
                :param tx_hash: hash of transaction emitting RoleRevoked event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.RoleRevoked().processReceipt(tx_receipt)
    
    
    def event_usr_deposit(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event usr_deposit in contract TestEnclose
        Get log entry for UsrDeposit event.
                :param tx_hash: hash of transaction emitting UsrDeposit event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.UsrDeposit().processReceipt(tx_receipt)
    
    
    def event_contract_paused(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event contract_paused in contract TestEnclose
        Get log entry for contractPaused event.
                :param tx_hash: hash of transaction emitting contractPaused event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.contractPaused().processReceipt(tx_receipt)
    
    
    def event_traillock(
            self, tx_hash: Union[HexBytes, bytes]
    ) -> Tuple[AttributeDict]:
        """
        Implementation of event traillock in contract TestEnclose
        Get log entry for traillock event.
                :param tx_hash: hash of transaction emitting traillock event
        """
        tx_receipt = self._web3_eth.getTransactionReceipt(tx_hash)
        return self._web3_eth.contract(address=to_checksum_address(self.contract_address), abi=TestEnclose.abi()).events.traillock().processReceipt(tx_receipt)

    
    
    
    def default_admin_role(self) -> Union[bytes, str]:
        """
        Implementation of default_admin_role in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_default_admin_role.block_call()
    
    
    
    def add_governor(self, account: str) -> None:
        """
        Implementation of add_governor in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_add_governor.block_send(account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def add_whitelist_admin(self, account: str) -> None:
        """
        Implementation of add_whitelist_admin in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_add_whitelist_admin.block_send(account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def coin_user_vault(self, index_0: Union[bytes, str]) -> int:
        """
        Implementation of coin_user_vault in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_coin_user_vault.block_call(index_0)
    
    
    
    def coin_user_vault2(self, index_0: str, index_1: str) -> int:
        """
        Implementation of coin_user_vault2 in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_coin_user_vault2.block_call(index_0, index_1)
    
    
    
    def deposit_erc20(self, coin: str, amount: int) -> None:
        """
        Implementation of deposit_erc20 in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_deposit_erc20.block_send(coin, amount, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def get_role_admin(self, role: Union[bytes, str]) -> Union[bytes, str]:
        """
        Implementation of get_role_admin in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_get_role_admin.block_call(role)
    
    
    
    def get_role_member(self, role: Union[bytes, str], index: int) -> str:
        """
        Implementation of get_role_member in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_get_role_member.block_call(role, index)
    
    
    
    def get_role_member_count(self, role: Union[bytes, str]) -> int:
        """
        Implementation of get_role_member_count in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_get_role_member_count.block_call(role)
    
    
    
    def governor(self) -> Union[bytes, str]:
        """
        Implementation of governor in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_governor.block_call()
    
    
    
    def grant_role(self, role: Union[bytes, str], account: str) -> None:
        """
        Implementation of grant_role in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_grant_role.block_send(role, account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def has_role(self, role: Union[bytes, str], account: str) -> bool:
        """
        Implementation of has_role in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_has_role.block_call(role, account)
    
    
    
    def is_governor(self, account: str) -> bool:
        """
        Implementation of is_governor in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_is_governor.block_call(account)
    
    
    
    def is_locked(self) -> bool:
        """
        Implementation of is_locked in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_is_locked.block_call()
    
    
    
    def is_owner(self) -> bool:
        """
        Implementation of is_owner in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_is_owner.block_call()
    
    
    
    def is_paused(self) -> bool:
        """
        Implementation of is_paused in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_is_paused.block_call()
    
    
    
    def is_whitelist_admin(self, account: str) -> bool:
        """
        Implementation of is_whitelist_admin in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_is_whitelist_admin.block_call(account)
    
    
    
    def lock(self) -> None:
        """
        Implementation of lock in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_lock.block_send(self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def owner(self) -> str:
        """
        Implementation of owner in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_owner.block_call()
    
    
    
    def pause_sc(self) -> None:
        """
        Implementation of pause_sc in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_pause_sc.block_send(self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def remove_governor(self, account: str) -> None:
        """
        Implementation of remove_governor in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_remove_governor.block_send(account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def remove_whitelist_admin(self, account: str) -> None:
        """
        Implementation of remove_whitelist_admin in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_remove_whitelist_admin.block_send(account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def renounce_ownership(self) -> None:
        """
        Implementation of renounce_ownership in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_renounce_ownership.block_send(self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def renounce_role(self, role: Union[bytes, str], account: str) -> None:
        """
        Implementation of renounce_role in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_renounce_role.block_send(role, account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def revoke_role(self, role: Union[bytes, str], account: str) -> None:
        """
        Implementation of revoke_role in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_revoke_role.block_send(role, account, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def sig(self, user: str, coin: str) -> Union[bytes, str]:
        """
        Implementation of sig in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
        return self._fn_sig.block_call(user, coin)
    
    
    
    
    
    def transfer_ownership(self, new_owner: str) -> None:
        """
        Implementation of transfer_ownership in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_transfer_ownership.block_send(new_owner, self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def unlock(self) -> None:
        """
        Implementation of unlock in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_unlock.block_send(self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def unpause_sc(self) -> None:
        """
        Implementation of unpause_sc in contract TestEnclose
        Method of the function
    
    
    
        """
    
        return self._fn_unpause_sc.block_send(self.call_contract_fee_amount,self.call_contract_fee_price,0,self.call_contract_debug_flag, self.call_contract_enforce_tx_receipt)
    
    
    
    
    
    
    
    def whitelist_admins(self) -> Union[bytes, str]:
        """
        Implementation of whitelist_admins in contract TestEnclose
        Method of the function
    
    
    
        """
    
    
    
    
    
        return self._fn_whitelist_admins.block_call()

    def CallContractWait(self, t_long:int)-> "TestEnclose":
        self._fn_default_admin_role.setWait(t_long)
        self._fn_add_governor.setWait(t_long)
        self._fn_add_whitelist_admin.setWait(t_long)
        self._fn_coin_user_vault.setWait(t_long)
        self._fn_coin_user_vault2.setWait(t_long)
        self._fn_deposit_erc20.setWait(t_long)
        self._fn_get_role_admin.setWait(t_long)
        self._fn_get_role_member.setWait(t_long)
        self._fn_get_role_member_count.setWait(t_long)
        self._fn_governor.setWait(t_long)
        self._fn_grant_role.setWait(t_long)
        self._fn_has_role.setWait(t_long)
        self._fn_is_governor.setWait(t_long)
        self._fn_is_locked.setWait(t_long)
        self._fn_is_owner.setWait(t_long)
        self._fn_is_paused.setWait(t_long)
        self._fn_is_whitelist_admin.setWait(t_long)
        self._fn_lock.setWait(t_long)
        self._fn_owner.setWait(t_long)
        self._fn_pause_sc.setWait(t_long)
        self._fn_remove_governor.setWait(t_long)
        self._fn_remove_whitelist_admin.setWait(t_long)
        self._fn_renounce_ownership.setWait(t_long)
        self._fn_renounce_role.setWait(t_long)
        self._fn_revoke_role.setWait(t_long)
        self._fn_sig.setWait(t_long)
        self._fn_transfer_ownership.setWait(t_long)
        self._fn_unlock.setWait(t_long)
        self._fn_unpause_sc.setWait(t_long)
        self._fn_whitelist_admins.setWait(t_long)
        return self


    @staticmethod
    def abi():
        """Return the ABI to the underlying contract."""
        return json.loads(
            '[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"role","type":"bytes32"},{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"sender","type":"address"}],"name":"RoleGranted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"role","type":"bytes32"},{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"sender","type":"address"}],"name":"RoleRevoked","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"a","type":"address"},{"indexed":false,"internalType":"uint256","name":"b","type":"uint256"}],"name":"UsrDeposit","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"value","type":"bool"}],"name":"contractPaused","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"value","type":"uint8"}],"name":"traillock","type":"event"},{"inputs":[],"name":"DEFAULT_ADMIN_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"addGovernor","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"addWhitelistAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"index_0","type":"bytes32"}],"name":"coin_user_vault","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"index_0","type":"address"},{"internalType":"address","name":"index_1","type":"address"}],"name":"coin_user_vault2","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"coin","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"deposit_erc20","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"}],"name":"getRoleAdmin","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getRoleMember","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"}],"name":"getRoleMemberCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"governor","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"grantRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"hasRole","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"isGovernor","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isLocked","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isOwner","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isPaused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"isWhitelistAdmin","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"lock","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pauseSc","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"removeGovernor","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"removeWhitelistAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"renounceRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"revokeRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"},{"internalType":"address","name":"coin","type":"address"}],"name":"sig","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unlock","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpauseSc","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"whitelistAdmins","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"}]'  # noqa: E501 (line-too-long)
        )

# pylint: disable=too-many-lines
