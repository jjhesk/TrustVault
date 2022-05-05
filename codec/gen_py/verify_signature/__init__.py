"""Generated wrapper for VerifySignature Solidity contract."""

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
# constructor for VerifySignature below.
try:
    # both mypy and pylint complain about what we're doing here, but this
    # works just fine, so their messages have been disabled here.
    from . import (  # type: ignore # pylint: disable=import-self
        VerifySignatureValidator,
    )
except ImportError:

    class VerifySignatureValidator(  # type: ignore
        Validator
    ):
        """No-op input validator."""

try:
    from .middleware import MIDDLEWARE  # type: ignore
except ImportError:
    pass





class GetEthSignedMessageHashMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the getEthSignedMessageHash method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("getEthSignedMessageHash")

    def validate_and_normalize_inputs(self, message_hash: Union[bytes, str])->any:
        """Validate the inputs to the getEthSignedMessageHash method."""
        self.validator.assert_valid(
            method_name='getEthSignedMessageHash',
            parameter_name='_messageHash',
            argument_value=message_hash,
        )
        return (message_hash)


    def block_call(self,message_hash: Union[bytes, str], debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method(message_hash)
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)

    def block_send(self, message_hash: Union[bytes, str],_valeth:int=0) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(message_hash)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': self.gas_limit,
                'gasPrice': self.gas_price_wei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if self.debug_method:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if self.auto_reciept is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.wait_for_transaction_receipt(txHash)
                    if self.debug_method:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                self._on_receipt_handle("get_eth_signed_message_hash", tx_receipt, txHash)

            if self.auto_reciept is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: get_eth_signed_message_hash")
            message = f"Error {er}: get_eth_signed_message_hash"
            self._on_fail("get_eth_signed_message_hash", message)
        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_eth_signed_message_hash: {message}")
            else:
                message = "Error Revert , Reason: Unknown"
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_eth_signed_message_hash. Reason: Unknown")

            self._on_fail("get_eth_signed_message_hash", message)

    def send_transaction(self, message_hash: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (message_hash) = self.validate_and_normalize_inputs(message_hash)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(message_hash).transact(tx_params.as_dict())

    def build_transaction(self, message_hash: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (message_hash) = self.validate_and_normalize_inputs(message_hash)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(message_hash).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, message_hash: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (message_hash) = self.validate_and_normalize_inputs(message_hash)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(message_hash).estimateGas(tx_params.as_dict())

class GetMessageHashMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the getMessageHash method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("getMessageHash")

    def validate_and_normalize_inputs(self, to: str, amount: int, message: str, nonce: int)->any:
        """Validate the inputs to the getMessageHash method."""
        self.validator.assert_valid(
            method_name='getMessageHash',
            parameter_name='_to',
            argument_value=to,
        )
        to = self.validate_and_checksum_address(to)
        self.validator.assert_valid(
            method_name='getMessageHash',
            parameter_name='_amount',
            argument_value=amount,
        )
        # safeguard against fractional inputs
        amount = int(amount)
        self.validator.assert_valid(
            method_name='getMessageHash',
            parameter_name='_message',
            argument_value=message,
        )
        self.validator.assert_valid(
            method_name='getMessageHash',
            parameter_name='_nonce',
            argument_value=nonce,
        )
        # safeguard against fractional inputs
        nonce = int(nonce)
        return (to, amount, message, nonce)


    def block_call(self,to: str, amount: int, message: str, nonce: int, debug:bool=False) -> Union[bytes, str]:
        _fn = self._underlying_method(to, amount, message, nonce)
        returned = _fn.call({
                'from': self._operate
            })
        return Union[bytes, str](returned)

    def block_send(self, to: str, amount: int, message: str, nonce: int,_valeth:int=0) -> Union[bytes, str]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(to, amount, message, nonce)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': self.gas_limit,
                'gasPrice': self.gas_price_wei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if self.debug_method:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if self.auto_reciept is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.wait_for_transaction_receipt(txHash)
                    if self.debug_method:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                self._on_receipt_handle("get_message_hash", tx_receipt, txHash)

            if self.auto_reciept is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: get_message_hash")
            message = f"Error {er}: get_message_hash"
            self._on_fail("get_message_hash", message)
        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_message_hash: {message}")
            else:
                message = "Error Revert , Reason: Unknown"
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, get_message_hash. Reason: Unknown")

            self._on_fail("get_message_hash", message)

    def send_transaction(self, to: str, amount: int, message: str, nonce: int, tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (to, amount, message, nonce) = self.validate_and_normalize_inputs(to, amount, message, nonce)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(to, amount, message, nonce).transact(tx_params.as_dict())

    def build_transaction(self, to: str, amount: int, message: str, nonce: int, tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (to, amount, message, nonce) = self.validate_and_normalize_inputs(to, amount, message, nonce)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(to, amount, message, nonce).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, to: str, amount: int, message: str, nonce: int, tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (to, amount, message, nonce) = self.validate_and_normalize_inputs(to, amount, message, nonce)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(to, amount, message, nonce).estimateGas(tx_params.as_dict())

class RecoverSignerMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the recoverSigner method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("recoverSigner")

    def validate_and_normalize_inputs(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str])->any:
        """Validate the inputs to the recoverSigner method."""
        self.validator.assert_valid(
            method_name='recoverSigner',
            parameter_name='_ethSignedMessageHash',
            argument_value=eth_signed_message_hash,
        )
        self.validator.assert_valid(
            method_name='recoverSigner',
            parameter_name='_signature',
            argument_value=signature,
        )
        return (eth_signed_message_hash, signature)


    def block_call(self,eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str], debug:bool=False) -> str:
        _fn = self._underlying_method(eth_signed_message_hash, signature)
        returned = _fn.call({
                'from': self._operate
            })
        return str(returned)

    def block_send(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str],_valeth:int=0) -> str:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(eth_signed_message_hash, signature)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': self.gas_limit,
                'gasPrice': self.gas_price_wei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if self.debug_method:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if self.auto_reciept is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.wait_for_transaction_receipt(txHash)
                    if self.debug_method:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                self._on_receipt_handle("recover_signer", tx_receipt, txHash)

            if self.auto_reciept is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: recover_signer")
            message = f"Error {er}: recover_signer"
            self._on_fail("recover_signer", message)
        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, recover_signer: {message}")
            else:
                message = "Error Revert , Reason: Unknown"
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, recover_signer. Reason: Unknown")

            self._on_fail("recover_signer", message)

    def send_transaction(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (eth_signed_message_hash, signature) = self.validate_and_normalize_inputs(eth_signed_message_hash, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(eth_signed_message_hash, signature).transact(tx_params.as_dict())

    def build_transaction(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (eth_signed_message_hash, signature) = self.validate_and_normalize_inputs(eth_signed_message_hash, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(eth_signed_message_hash, signature).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (eth_signed_message_hash, signature) = self.validate_and_normalize_inputs(eth_signed_message_hash, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(eth_signed_message_hash, signature).estimateGas(tx_params.as_dict())

class SplitSignatureMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the splitSignature method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("splitSignature")

    def validate_and_normalize_inputs(self, sig: Union[bytes, str])->any:
        """Validate the inputs to the splitSignature method."""
        self.validator.assert_valid(
            method_name='splitSignature',
            parameter_name='sig',
            argument_value=sig,
        )
        return (sig)


    def block_call(self,sig: Union[bytes, str], debug:bool=False) -> Tuple[Union[bytes, str], Union[bytes, str], int]:
        _fn = self._underlying_method(sig)
        returned = _fn.call({
                'from': self._operate
            })
        return (returned[0],returned[1],returned[2],)

    def block_send(self, sig: Union[bytes, str],_valeth:int=0) -> Tuple[Union[bytes, str], Union[bytes, str], int]:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(sig)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': self.gas_limit,
                'gasPrice': self.gas_price_wei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if self.debug_method:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if self.auto_reciept is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.wait_for_transaction_receipt(txHash)
                    if self.debug_method:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                self._on_receipt_handle("split_signature", tx_receipt, txHash)

            if self.auto_reciept is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: split_signature")
            message = f"Error {er}: split_signature"
            self._on_fail("split_signature", message)
        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, split_signature: {message}")
            else:
                message = "Error Revert , Reason: Unknown"
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, split_signature. Reason: Unknown")

            self._on_fail("split_signature", message)

    def send_transaction(self, sig: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (sig) = self.validate_and_normalize_inputs(sig)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(sig).transact(tx_params.as_dict())

    def build_transaction(self, sig: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (sig) = self.validate_and_normalize_inputs(sig)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(sig).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, sig: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (sig) = self.validate_and_normalize_inputs(sig)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(sig).estimateGas(tx_params.as_dict())

class VerifyMethod(ContractMethod): # pylint: disable=invalid-name
    """Various interfaces to the verify method."""

    def __init__(self, elib: MiliDoS, contract_address: str, contract_function: ContractFunction, validator: Validator=None):
        """Persist instance data."""
        super().__init__(elib, contract_address, validator)
        self._underlying_method = contract_function
        self.sign = validator.getSignature("verify")

    def validate_and_normalize_inputs(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str])->any:
        """Validate the inputs to the verify method."""
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='_signer',
            argument_value=signer,
        )
        signer = self.validate_and_checksum_address(signer)
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='_to',
            argument_value=to,
        )
        to = self.validate_and_checksum_address(to)
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='_amount',
            argument_value=amount,
        )
        # safeguard against fractional inputs
        amount = int(amount)
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='_message',
            argument_value=message,
        )
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='_nonce',
            argument_value=nonce,
        )
        # safeguard against fractional inputs
        nonce = int(nonce)
        self.validator.assert_valid(
            method_name='verify',
            parameter_name='signature',
            argument_value=signature,
        )
        return (signer, to, amount, message, nonce, signature)


    def block_call(self,signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str], debug:bool=False) -> bool:
        _fn = self._underlying_method(signer, to, amount, message, nonce, signature)
        returned = _fn.call({
                'from': self._operate
            })
        return bool(returned)

    def block_send(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str],_valeth:int=0) -> bool:
        """Execute underlying contract method via eth_call.

        :param tx_params: transaction parameters
        :returns: the return value of the underlying method.
        """
        _fn = self._underlying_method(signer, to, amount, message, nonce, signature)
        try:

            _t = _fn.buildTransaction({
                'from': self._operate,
                'gas': self.gas_limit,
                'gasPrice': self.gas_price_wei
            })
            _t['nonce'] = self._web3_eth.getTransactionCount(self._operate)

            if _valeth > 0:
                _t['value'] = _valeth

            if self.debug_method:
                print(f"======== Signing ✅ by {self._operate}")
                print(f"======== Transaction ✅ check")
                print(_t)

            if 'data' in _t:

                signed = self._web3_eth.account.sign_transaction(_t)
                txHash = self._web3_eth.sendRawTransaction(signed.rawTransaction)
                tx_receipt = None
                if self.auto_reciept is True:
                    print(f"======== awaiting Confirmation 🚸️ {self.sign}")
                    tx_receipt = self._web3_eth.wait_for_transaction_receipt(txHash)
                    if self.debug_method:
                        print("======== TX Result ✅")
                        print(tx_receipt)

                self._on_receipt_handle("verify", tx_receipt, txHash)

            if self.auto_reciept is False:
                time.sleep(self._wait)


        except ContractLogicError as er:
            print(f"{Bolors.FAIL}Error {er} {Bolors.RESET}: verify")
            message = f"Error {er}: verify"
            self._on_fail("verify", message)
        except ValueError as err:
            if "message" in err.args[0]:
                message = err.args[0]["message"]
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, verify: {message}")
            else:
                message = "Error Revert , Reason: Unknown"
                print(f"{Bolors.FAIL}Error Revert {Bolors.RESET}, verify. Reason: Unknown")

            self._on_fail("verify", message)

    def send_transaction(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> Union[HexBytes, bytes]:
        """Execute underlying contract method via eth_sendTransaction.

        :param tx_params: transaction parameters
        """
        (signer, to, amount, message, nonce, signature) = self.validate_and_normalize_inputs(signer, to, amount, message, nonce, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(signer, to, amount, message, nonce, signature).transact(tx_params.as_dict())

    def build_transaction(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> dict:
        """Construct calldata to be used as input to the method."""
        (signer, to, amount, message, nonce, signature) = self.validate_and_normalize_inputs(signer, to, amount, message, nonce, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(signer, to, amount, message, nonce, signature).buildTransaction(tx_params.as_dict())

    def estimate_gas(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str], tx_params: Optional[TxParams] = None) -> int:
        """Estimate gas consumption of method call."""
        (signer, to, amount, message, nonce, signature) = self.validate_and_normalize_inputs(signer, to, amount, message, nonce, signature)
        tx_params = super().normalize_tx_params(tx_params)
        return self._underlying_method(signer, to, amount, message, nonce, signature).estimateGas(tx_params.as_dict())

class SignatureGenerator(Signatures):
    """
        The signature is generated for this and it is installed.
    """
    def __init__(self, abi: any):
        super().__init__(abi)

    def get_eth_signed_message_hash(self) -> str:
        return self._function_signatures["getEthSignedMessageHash"]
    def get_message_hash(self) -> str:
        return self._function_signatures["getMessageHash"]
    def recover_signer(self) -> str:
        return self._function_signatures["recoverSigner"]
    def split_signature(self) -> str:
        return self._function_signatures["splitSignature"]
    def verify(self) -> str:
        return self._function_signatures["verify"]

# pylint: disable=too-many-public-methods,too-many-instance-attributes
class VerifySignature(ContractBase):
    """Wrapper class for VerifySignature Solidity contract.

    All method parameters of type `bytes`:code: should be encoded as UTF-8,
    which can be accomplished via `str.encode("utf_8")`:code:.
    """
    _fn_get_eth_signed_message_hash: GetEthSignedMessageHashMethod
    """Constructor-initialized instance of
    :class:`GetEthSignedMessageHashMethod`.
    """

    _fn_get_message_hash: GetMessageHashMethod
    """Constructor-initialized instance of
    :class:`GetMessageHashMethod`.
    """

    _fn_recover_signer: RecoverSignerMethod
    """Constructor-initialized instance of
    :class:`RecoverSignerMethod`.
    """

    _fn_split_signature: SplitSignatureMethod
    """Constructor-initialized instance of
    :class:`SplitSignatureMethod`.
    """

    _fn_verify: VerifyMethod
    """Constructor-initialized instance of
    :class:`VerifyMethod`.
    """

    SIGNATURES:SignatureGenerator = None

    def __init__(
        self,
        core_lib: MiliDoS,
        contract_address: str,
        validator: VerifySignatureValidator = None,
    ):
        """Get an instance of wrapper for smart contract.
        """
        # pylint: disable=too-many-statements
        super().__init__(contract_address, VerifySignature.abi())
        web3 = core_lib.w3

        if not validator:
            validator = VerifySignatureValidator(web3, contract_address)




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
        functions = self._web3_eth.contract(address=to_checksum_address(contract_address), abi=VerifySignature.abi()).functions
        self._signatures = SignatureGenerator(VerifySignature.abi())
        validator.bindSignatures(self._signatures)

        self._fn_get_eth_signed_message_hash = GetEthSignedMessageHashMethod(core_lib, contract_address, functions.getEthSignedMessageHash, validator)
        self._fn_get_message_hash = GetMessageHashMethod(core_lib, contract_address, functions.getMessageHash, validator)
        self._fn_recover_signer = RecoverSignerMethod(core_lib, contract_address, functions.recoverSigner, validator)
        self._fn_split_signature = SplitSignatureMethod(core_lib, contract_address, functions.splitSignature, validator)
        self._fn_verify = VerifyMethod(core_lib, contract_address, functions.verify, validator)


    
    
    
    def get_eth_signed_message_hash(self, message_hash: Union[bytes, str]) -> Union[bytes, str]:
        """
        Implementation of get_eth_signed_message_hash in contract VerifySignature
        Method of the function
    
        """
    
        self._fn_get_eth_signed_message_hash.callback_onfail = self._callback_onfail
        self._fn_get_eth_signed_message_hash.callback_onsuccess = self._callback_onsuccess
        self._fn_get_eth_signed_message_hash.auto_reciept = self.call_contract_enforce_tx_receipt
        self._fn_get_eth_signed_message_hash.gas_limit = self.call_contract_fee_amount
        self._fn_get_eth_signed_message_hash.gas_price_wei = self.call_contract_fee_price
        self._fn_get_eth_signed_message_hash.debug_method = self.call_contract_debug_flag
    
    
    
    
        return self._fn_get_eth_signed_message_hash.block_call(message_hash)
    
    
    
    
    
    def get_message_hash(self, to: str, amount: int, message: str, nonce: int) -> Union[bytes, str]:
        """
        Implementation of get_message_hash in contract VerifySignature
        Method of the function
    
        """
    
        self._fn_get_message_hash.callback_onfail = self._callback_onfail
        self._fn_get_message_hash.callback_onsuccess = self._callback_onsuccess
        self._fn_get_message_hash.auto_reciept = self.call_contract_enforce_tx_receipt
        self._fn_get_message_hash.gas_limit = self.call_contract_fee_amount
        self._fn_get_message_hash.gas_price_wei = self.call_contract_fee_price
        self._fn_get_message_hash.debug_method = self.call_contract_debug_flag
    
    
    
    
        return self._fn_get_message_hash.block_call(to, amount, message, nonce)
    
    
    
    
    
    def recover_signer(self, eth_signed_message_hash: Union[bytes, str], signature: Union[bytes, str]) -> str:
        """
        Implementation of recover_signer in contract VerifySignature
        Method of the function
    
        """
    
        self._fn_recover_signer.callback_onfail = self._callback_onfail
        self._fn_recover_signer.callback_onsuccess = self._callback_onsuccess
        self._fn_recover_signer.auto_reciept = self.call_contract_enforce_tx_receipt
        self._fn_recover_signer.gas_limit = self.call_contract_fee_amount
        self._fn_recover_signer.gas_price_wei = self.call_contract_fee_price
        self._fn_recover_signer.debug_method = self.call_contract_debug_flag
    
    
    
    
        return self._fn_recover_signer.block_call(eth_signed_message_hash, signature)
    
    
    
    
    
    def split_signature(self, sig: Union[bytes, str]) -> Tuple[Union[bytes, str], Union[bytes, str], int]:
        """
        Implementation of split_signature in contract VerifySignature
        Method of the function
    
        """
    
        self._fn_split_signature.callback_onfail = self._callback_onfail
        self._fn_split_signature.callback_onsuccess = self._callback_onsuccess
        self._fn_split_signature.auto_reciept = self.call_contract_enforce_tx_receipt
        self._fn_split_signature.gas_limit = self.call_contract_fee_amount
        self._fn_split_signature.gas_price_wei = self.call_contract_fee_price
        self._fn_split_signature.debug_method = self.call_contract_debug_flag
    
    
    
    
        return self._fn_split_signature.block_call(sig)
    
    
    
    
    
    def verify(self, signer: str, to: str, amount: int, message: str, nonce: int, signature: Union[bytes, str]) -> bool:
        """
        Implementation of verify in contract VerifySignature
        Method of the function
    
        """
    
        self._fn_verify.callback_onfail = self._callback_onfail
        self._fn_verify.callback_onsuccess = self._callback_onsuccess
        self._fn_verify.auto_reciept = self.call_contract_enforce_tx_receipt
        self._fn_verify.gas_limit = self.call_contract_fee_amount
        self._fn_verify.gas_price_wei = self.call_contract_fee_price
        self._fn_verify.debug_method = self.call_contract_debug_flag
    
    
    
    
        return self._fn_verify.block_call(signer, to, amount, message, nonce, signature)
    
    

    def CallContractWait(self, t_long:int)-> "VerifySignature":
        self._fn_get_eth_signed_message_hash.setWait(t_long)
        self._fn_get_message_hash.setWait(t_long)
        self._fn_recover_signer.setWait(t_long)
        self._fn_split_signature.setWait(t_long)
        self._fn_verify.setWait(t_long)
        return self


    @staticmethod
    def abi():
        """Return the ABI to the underlying contract."""
        return json.loads(
            '[{"inputs":[{"internalType":"bytes32","name":"_messageHash","type":"bytes32"}],"name":"getEthSignedMessageHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"},{"internalType":"string","name":"_message","type":"string"},{"internalType":"uint256","name":"_nonce","type":"uint256"}],"name":"getMessageHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"bytes32","name":"_ethSignedMessageHash","type":"bytes32"},{"internalType":"bytes","name":"_signature","type":"bytes"}],"name":"recoverSigner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"bytes","name":"sig","type":"bytes"}],"name":"splitSignature","outputs":[{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"},{"internalType":"uint8","name":"v","type":"uint8"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"_signer","type":"address"},{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"},{"internalType":"string","name":"_message","type":"string"},{"internalType":"uint256","name":"_nonce","type":"uint256"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"verify","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"pure","type":"function"}]'  # noqa: E501 (line-too-long)
        )

# pylint: disable=too-many-lines
