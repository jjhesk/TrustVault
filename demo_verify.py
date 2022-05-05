# !/usr/bin/env python
# coding: utf-8
# from web3.auto import w3

from moody import Bolors

from key import ROOT
from codec.BaseVault import BsVault


def Prod() -> BsVault:
    # ========================== Of course
    j = BsVault(ROOT)
    return j


"""
from the request 
		Address  string `json:"wallet_address"`
		Signed   string `json:"signed_message"`
		Hash     string `json:"hash_message"`
		Original string `json:"original_message"`
		
"""

hash_message = ""
signed_message = ""

res = Prod().versign.recover_signer(
    eth_signed_message_hash=hash_message,
    signature=signed_message
)
# the message is now success or not
result = "success" if res is True else "failed"
print(f"The verification is {result}")
