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


def deployVerifySignature():
    # j = Prod()
    j.deploy("VerifySignature", [])


def deploy():
    # j = Prod()
    model_addr = j.getAddr("VaultConfigProvider")
    j.deploy("Enrockvault", [model_addr])
    j.defineEnrockVault()
    # node oracle
    j.enrock.add_governor("0xa255FF46Ab9aafC1d8bd138A9b6f9dDc64044Ea4")
    # testing use only
    j.enrock.add_governor("0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7")
    j.enrock.set_fee(j.getAddr("LSL"), 10)


def deployTV():
    # j = Prod()
    # j.deploy("TestEnclose", [])
    j.By(2)
    j.defineTVault()
    # node oracle
    # j.tenclose.add_governor("0xa255FF46Ab9aafC1d8bd138A9b6f9dDc64044Ea4")
    # testing use only
    # j.tenclose.add_governor("0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7")
    node = "0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7"

    coin = j.getAddr("LSL")
    j.defineToken("LSL")
    addr = j.getAddr("TestEnclose")
    stated = 10 ** 18 * 3
    # j.token.approve(addr, stated)
    # j.tenclose.deposit_erc20(coin, stated)
    v = j.tenclose.sig(user, coin)
    amount1 = j.tenclose.coin_user_vault(v)
    amount2 = j.tenclose.coin_user_vault(j.w3.toHex(v))
    amount3 = j.tenclose.coin_user_vault2(coin, user)
    print(f"Finally got {amount1} or {amount2} or {amount3}")


def deployvaultConfig():
    # j = Prod()
    j.deploy("VaultConfigProvider", [])
    j.defineVaultConfigProvider()
    j.vault_config_provider.set_genesis("0x6Ab612071217E1d7d80176A7bA25dFaD15f69bB0")


def deployEnrockAddGov():
    j.defineEnrockVault()
    # node oracle
    j.enrock.add_governor("0xa255FF46Ab9aafC1d8bd138A9b6f9dDc64044Ea4")
    # testing use only
    j.enrock.add_governor("0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7")


def deposit(coin_symbol: str, stated_amount: int):
    j.By(2).defineToken(coin_symbol).defineEnrockVault()
    address = j.getAddr("Enrockvault")
    j.token.approve(address, stated_amount)
    j.enrock.deposit_erc20(j.getAddr(coin_symbol), stated_amount)


def withdrawbysys():
    j.By(14).defineEnrockVault().defineToken("LSL")
    j.enrock.take_usr(user, j.getAddr("LSL"), 10 ** 18 * 18)


def CheckDeposit(coin_symbol: str):
    # j = Prod()
    j.defineToken(coin_symbol).defineEnrockVault()
    h = "0xa255FF46Ab9aafC1d8bd138A9b6f9dDc64044Ea4"
    ishegov = j.enrock.is_governor(h)
    print(f"is he the gov {ishegov} - {h}")
    # (a, c) = j.auth(2)
    bb = "0x55cF3C074B67c93AD13C78A483d42Bf07444FB2C"
    addr = j.getAddr(coin_symbol)
    res = j.enrock.check(addr, bb)
    r = j.enrock.game_round()
    print(f" the amount in the account is now: {r}, {Bolors.OK}{res}{Bolors.RESET}")
    # res = j.enrock.coin_vault_user(addr, bb)
    # print(f" the amount in the account in mapping coin_vault_user[{addr}][{bb}] is {res}")


def AdjustBack(coin_symbol: str, stated_amount: int):
    # j = Prod()
    j.defineEnrockVault()
    coin_addr = j.getAddr(coin_symbol)
    # usr = "0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7"
    # j.enrock.up_usr_batch([usr], [coin_addr], [stated_amount])

    j.enrock.up_single(user, coin_addr, stated_amount)


# deployVerifySignature()
# deployEnrockAddGov()
# deposit("LSL", 10 ** 18 * 1)


# deployvaultConfig()

# deployEnrockAddGov()
# AdjustBack("LSL", 10 ** 18 * 990)
"""deployvaultConfig()
deploy()
"""

user = "0x55cF3C074B67c93AD13C78A483d42Bf07444FB2C"

if __name__ == '__main__':
    j = Prod()
    # deploy()
    # deployTV()
    # deposit("LSL", 10 ** 18 * 13)
    """    
    deposit("LSL", 10 ** 18 * 2)
    deposit("LSL", 10 ** 18 * 3)
    deposit("LSL", 10 ** 18 * 3)
    deposit("LSL", 10 ** 18 * 3)
    deposit("LSL", 10 ** 18 * 3)
    AdjustBack("LSL", 10 ** 18 * 2)
    """

    # deposit("LSL", 10 ** 18 * 1100)
    # deposit("LSL", 10 ** 18 * 5000)
    """
    j.defineEnrockVault()
    addr = j.getAddr("LSL")
    print(addr)
    who = "0xF4adebDEde39865bC76C6C1Af7Ceb74daeb524a7"
    j.enrock.set_round(3)
    g = j.enrock.signature(who, addr)
    decoded = j.w3.toHex(g)
    print(f"Nw is up {decoded}")
    """
    # AdjustBack("LSL", 10 ** 18 * 2)
    # withdrawbysys()
    CheckDeposit("LSL")
    # deployEnrockAddGov()
