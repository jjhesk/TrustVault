from moody.contracttool import ContractTool

from codec.gen_py.enrockvault import Enrockvault
from codec.gen_py.tc20 import Tc20
from codec.gen_py.test_enclose import TestEnclose
from codec.gen_py.vault_config_provider import VaultConfigProvider

from key import NETWORK


class TokenBase(ContractTool):
    def defineToken(self, name: str) -> "TokenBase":
        if not self.hasContractName(name):
            print(f"Token {name} is not found..")
            exit(0)
        self.token: Tc20 = Tc20(self, self.getAddr(name))
        self.token.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineTVault(self) -> "TokenBase":
        self.tenclose: TestEnclose = TestEnclose(self, self.getAddr("TestEnclose"))
        self.tenclose.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineVault(self) -> "TokenBase":
        self.envault: Enrockvault = Enrockvault(self, self.getAddr("Enrockvault"))
        self.envault.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineVaultConfigProvider(self) -> "TokenBase":
        self.vault_config_provider: VaultConfigProvider = VaultConfigProvider(self, self.getAddr("VaultConfigProvider"))
        self.vault_config_provider.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

class BsVault(TokenBase):

    def __init__(self, path_bs: str):
        # super().__init__(NETWORK, path_bs, OKTEST_DEPLOY, WALLETS)
        super().__init__(NETWORK, path_bs, {}, [])
        self.withPOA()
        self.setWorkspace(path_bs).Auth("")
        self.connect(path_bs, False)
        self.OverrideGasConfig(6000000, 1059100000)
        self.OverrideChainConfig(10 ** 18, 6)

    def By(self, player: int) -> "BsVault":
        # switch account to the index A account
        self.AuthIndex(player)
        return self
