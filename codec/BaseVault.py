from moody.contracttool import ContractTool


from codec.gen_py.tc20 import Tc20
from key import NETWORK, DEPLOYLIST, WALLETS, pri


class TokenBase(ContractTool):
    def defineToken(self, name: str) -> "TokenBase":
        if not self.hasContractName(name):
            print(f"Token {name} is not found..")
            exit(0)
        self.token: Tc20 = Tc20(self, self.getAddr(name))
        self.token.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineChainLinkToken(self, name: str) -> "TokenBase":
        if not self.hasContractName(name):
            print(f"Token {name} is not found..")
            exit(0)
        self.token: ChingToken = ChingToken(self, self.getAddr(name))
        self.token.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineLottoMinToken(self) -> "TokenBase":
        self.token: MintableVotingToken = MintableVotingToken(self, self.getAddr("LSL"))
        self.token.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineUSDTToken(self) -> "TokenBase":
        self.token: Tc20 = Tc20(self, self.getAddr("USDF"))
        self.token.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineVaultConfigProvider(self) -> "TokenBase":
        self.vault_config_provider: VaultConfigProvider = VaultConfigProvider(self, self.getAddr("VaultConfigProvider"))
        self.vault_config_provider.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineEnrockVault(self) -> "TokenBase":
        self.enrock: Enrockvault = Enrockvault(self, self.getAddr("Enrockvault"))
        self.enrock.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self

    def defineTVault(self) -> "TokenBase":
        self.tenclose: TestEnclose = TestEnclose(self, self.getAddr("TestEnclose"))
        self.tenclose.CallAutoConf(self).CallDebug(True).CallContractWait(self.waitSec)
        return self


class BsVault(TokenBase):

    def __init__(self, path_bs: str):
        # super().__init__(NETWORK, path_bs, OKTEST_DEPLOY, WALLETS)
        super().__init__(NETWORK, path_bs, DEPLOYLIST, WALLETS)
        self.withPOA()
        self.setWorkspace(path_bs).Auth(pri)
        self.connect(path_bs, False)
        self.OverrideGasConfig(6000000, 1059100000)
        self.OverrideChainConfig(10 ** 18, 6)

    def By(self, player: int) -> "BsVault":
        # switch account to the index A account
        self.AuthIndex(player)
        return self
