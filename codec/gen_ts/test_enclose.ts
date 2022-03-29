/* DO NOT EDIT THE BELOW FILE AS THIS IS LIKELY WILL BE GENERATED AGAIN AND REWRITE OVER IT */

// tslint:disable:no-consecutive-blank-lines ordered-imports align trailing-comma enum-naming
// tslint:disable:whitespace no-unbound-method no-trailing-whitespace
// tslint:disable:no-unused-variable

import * as ethers from "ethers"
// eslint-disable-next-line import/named
import {
  assert,
  schemas,
  // eslint-disable-next-line import/named
  SubscriptionManager,
  // eslint-disable-next-line import/named
  BaseContract,
  // eslint-disable-next-line import/named
  EventCallback,
  // eslint-disable-next-line import/named
  IndexedFilterValues,
  // eslint-disable-next-line import/named
  BlockRange,
  // eslint-disable-next-line import/named
  DecodedLogArgs,
  // eslint-disable-next-line import/named
  LogWithDecodedArgs,
  // eslint-disable-next-line import/named
  MethodAbi
} from "vue-blocklink"

import {
  BatchRequest,
  Extension,
  Log,
  PromiEvent,
  provider,
  Providers,
  RLPEncodedTransaction,
  Transaction,
  TransactionConfig,
  TransactionReceipt,
  Common,
  hardfork,
  chain,
  BlockNumber,
  LogsOptions,
  PastLogsOptions
} from "web3-core"

import { Utils, AbiItem } from "web3-utils"
import Web3 from "web3"
import BN from "BN.js"

// tslint:enable:no-unused-variable
export interface ContractInterface {
    DEFAULT_ADMIN_ROLE():Promise<string>
    addGovernor(account: string,):Promise<void>
    addWhitelistAdmin(account: string,):Promise<void>
    coin_user_vault(index_0: string,):Promise<BN>
    coin_user_vault2(index_0: string,index_1: string,):Promise<BN>
    deposit_erc20(coin: string,amount: BN,):Promise<void>
    getRoleAdmin(role: string,):Promise<string>
    getRoleMember(role: string,index: BN,):Promise<string>
    getRoleMemberCount(role: string,):Promise<BN>
    governor():Promise<string>
    grantRole(role: string,account: string,):Promise<void>
    hasRole(role: string,account: string,):Promise<boolean>
    isGovernor(account: string,):Promise<boolean>
    isLocked():Promise<boolean>
    isOwner():Promise<boolean>
    isPaused():Promise<boolean>
    isWhitelistAdmin(account: string,):Promise<boolean>
    lock():Promise<void>
    owner():Promise<string>
    pauseSc():Promise<void>
    removeGovernor(account: string,):Promise<void>
    removeWhitelistAdmin(account: string,):Promise<void>
    renounceOwnership():Promise<void>
    renounceRole(role: string,account: string,):Promise<void>
    revokeRole(role: string,account: string,):Promise<void>
    sig(user: string,coin: string,):Promise<string>
    transferOwnership(newOwner: string,):Promise<void>
    unlock():Promise<void>
    unpauseSc():Promise<void>
    whitelistAdmins():Promise<string>
}





export enum TestEncloseEvents {
    OwnershipTransferred = 'OwnershipTransferred',
    RoleGranted = 'RoleGranted',
    RoleRevoked = 'RoleRevoked',
    UsrDeposit = 'UsrDeposit',
    contractPaused = 'contractPaused',
    traillock = 'traillock',
}

export interface TestEncloseOwnershipTransferredEventArgs extends DecodedLogArgs {
    previousOwner: string;
    newOwner: string;
}

export interface TestEncloseRoleGrantedEventArgs extends DecodedLogArgs {
    role: string;
    account: string;
    sender: string;
}

export interface TestEncloseRoleRevokedEventArgs extends DecodedLogArgs {
    role: string;
    account: string;
    sender: string;
}

export interface TestEncloseUsrDepositEventArgs extends DecodedLogArgs {
    a: string;
    b: BN;
}

export interface TestEnclosecontractPausedEventArgs extends DecodedLogArgs {
    value: boolean;
}

export interface TestEnclosetraillockEventArgs extends DecodedLogArgs {
    value: BN;
}


export type TestEncloseEventArgs =
    | TestEncloseOwnershipTransferredEventArgs
    | TestEncloseRoleGrantedEventArgs
    | TestEncloseRoleRevokedEventArgs
    | TestEncloseUsrDepositEventArgs
    | TestEnclosecontractPausedEventArgs
    | TestEnclosetraillockEventArgs;




/* istanbul ignore next */
// tslint:disable:array-type
// tslint:disable:no-parameter-reassignment
// tslint:disable-next-line:class-name
export class TestEncloseContract extends BaseContract implements ContractInterface{
    /**
     * @ignore
     */
public static deployedBytecode = '60806040526001805461ffff60a01b1916905534801561001e57600080fd5b50600061002961007c565b600180546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610080565b3390565b61165b8061008f6000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806391d1485411610104578063ca15c873116100a2578063e43581b811610071578063e43581b8146104c5578063eecdac88146104eb578063f2fde38b14610511578063f83d08ba14610537576101da565b8063ca15c87314610448578063d16544f014610465578063d547741f14610491578063dfc877d3146104bd576101da565b8063a69df4b5116100de578063a69df4b5146103f5578063a83b18ec146103fd578063b187bd261461041a578063bb5f747b14610422576101da565b806391d14854146103b9578063a217fddf146103e5578063a4e2d634146103ed576101da565b80636897e9741161017c5780637ff7a6e51161014b5780637ff7a6e5146103285780638da5cb5b146103565780638f32d59b1461037a5780639010d07c14610396576101da565b80636897e974146102cc578063715018a6146102f25780637362d9c8146102fa5780637ef581cc14610320576101da565b8063248a9ca3116101b8578063248a9ca3146102315780632f2ff15d1461024e57806336568abe1461027a5780633c4a25d0146102a6576101da565b80630c340a24146101df578063171ea1f8146101f95780631975ebaf14610227575b600080fd5b6101e761053f565b60408051918252519081900360200190f35b6101e76004803603604081101561020f57600080fd5b506001600160a01b0381358116916020013516610551565b61022f61056e565b005b6101e76004803603602081101561024757600080fd5b50356105eb565b61022f6004803603604081101561026457600080fd5b50803590602001356001600160a01b0316610600565b61022f6004803603604081101561029057600080fd5b50803590602001356001600160a01b0316610667565b61022f600480360360208110156102bc57600080fd5b50356001600160a01b03166106c8565b61022f600480360360208110156102e257600080fd5b50356001600160a01b031661073b565b61022f6107ab565b61022f6004803603602081101561031057600080fd5b50356001600160a01b031661084d565b6101e76108bd565b6101e76004803603604081101561033e57600080fd5b506001600160a01b03813581169160200135166108cf565b61035e610922565b604080516001600160a01b039092168252519081900360200190f35b610382610931565b604080519115158252519081900360200190f35b61035e600480360360408110156103ac57600080fd5b508035906020013561095b565b610382600480360360408110156103cf57600080fd5b50803590602001356001600160a01b031661097a565b6101e7610992565b610382610997565b61022f6109a9565b6101e76004803603602081101561041357600080fd5b5035610a16565b610382610a28565b6103826004803603602081101561043857600080fd5b50356001600160a01b0316610a3a565b6101e76004803603602081101561045e57600080fd5b5035610a62565b61022f6004803603604081101561047b57600080fd5b506001600160a01b038135169060200135610a79565b61022f600480360360408110156104a757600080fd5b50803590602001356001600160a01b0316610c03565b61022f610c5c565b610382600480360360208110156104db57600080fd5b50356001600160a01b0316610cce565b61022f6004803603602081101561050157600080fd5b50356001600160a01b0316610ce8565b61022f6004803603602081101561052757600080fd5b50356001600160a01b0316610d58565b61022f610e51565b6000805160206115ad83398151915281565b600360209081526000928352604080842090915290825290205481565b61058d6000805160206115ad833981519152610588610ebc565b61097a565b8061059b575061059b610931565b6105d6576040805162461bcd60e51b81526020600482015260076024820152660dcde40c2eae8d60cb1b604482015290519081900360640190fd5b6001805460ff60a81b1916600160a81b179055565b60009081526020819052604090206002015490565b60008281526020819052604090206002015461061e90610588610ebc565b6106595760405162461bcd60e51b815260040180806020018281038252602f8152602001806114e8602f913960400191505060405180910390fd5b6106638282610ec0565b5050565b61066f610ebc565b6001600160a01b0316816001600160a01b0316146106be5760405162461bcd60e51b815260040180806020018281038252602f8152602001806115f7602f913960400191505060405180910390fd5b6106638282610f29565b6106d0610ebc565b6001546001600160a01b03908116911614610720576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b6107386000805160206115ad83398151915282610ec0565b50565b610743610ebc565b6001546001600160a01b03908116911614610793576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b61073860008051602061151783398151915282610f29565b6107b3610ebc565b6001546001600160a01b03908116911614610803576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b6001546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600180546001600160a01b0319169055565b610855610ebc565b6001546001600160a01b039081169116146108a5576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b61073860008051602061151783398151915282610ec0565b60008051602061151783398151915281565b60408051605f60f81b6020808301919091526bffffffffffffffffffffffff19606085811b8216602185015286901b16603583015282516029818403018152604990920190925280519101205b92915050565b6001546001600160a01b031690565b600061093b610ebc565b6001600160a01b031661094c610922565b6001600160a01b031614905090565b60008281526020819052604081206109739083610f92565b9392505050565b60008281526020819052604081206109739083610f9e565b600081565b60018054600160a01b900460ff161490565b6109c36000805160206115ad833981519152610588610ebc565b806109d157506109d1610931565b610a0c576040805162461bcd60e51b81526020600482015260076024820152660dcde40c2eae8d60cb1b604482015290519081900360640190fd5b610a14610fb3565b565b60026020526000908152604090205481565b60018054600160a81b900460ff161490565b6000610a546000805160206115178339815191528361097a565b8061091c575061091c610931565b600081815260208190526040812061091c90611000565b818160008111610ac4576040805162461bcd60e51b8152602060048201526011602482015270212a1d1d1034b73b30b634b21031b7b4b760791b604482015290519081900360640190fd5b610acd8261100b565b610b1e576040805162461bcd60e51b815260206004820152601a60248201527f42543a3a20696e76616c696420636f696e20636f6e7472616374000000000000604482015290519081900360640190fd5b336000610b2b82876108cf565b9050610b426001600160a01b038716833088611047565b600081815260026020526040902054610b5b90866110a7565b6000828152600260209081526040808320939093556001600160a01b03808a16835260038252838320908616835290522054610b9781876110a7565b6001600160a01b03808916600081815260036020908152604080832094891683529381529083902093909355815190815291820188905280517f089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d899281900390910190a150505050505050565b600082815260208190526040902060020154610c2190610588610ebc565b6106be5760405162461bcd60e51b815260040180806020018281038252603081526020018061155d6030913960400191505060405180910390fd5b610c766000805160206115ad833981519152610588610ebc565b80610c845750610c84610931565b610cbf576040805162461bcd60e51b81526020600482015260076024820152660dcde40c2eae8d60cb1b604482015290519081900360640190fd5b6001805460ff60a81b19169055565b600061091c6000805160206115ad8339815191528361097a565b610cf0610ebc565b6001546001600160a01b03908116911614610d40576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b6107386000805160206115ad83398151915282610f29565b610d60610ebc565b6001546001600160a01b03908116911614610db0576040805162461bcd60e51b8152602060048201819052602482015260008051602061158d833981519152604482015290519081900360640190fd5b6001600160a01b038116610df55760405162461bcd60e51b81526004018080602001828103825260268152602001806115376026913960400191505060405180910390fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b610e6b6000805160206115ad833981519152610588610ebc565b80610e795750610e79610931565b610eb4576040805162461bcd60e51b81526020600482015260076024820152660dcde40c2eae8d60cb1b604482015290519081900360640190fd5b610a14611101565b3390565b6000828152602081905260409020610ed89082611153565b1561066357610ee5610ebc565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610f419082611168565b1561066357610f4e610ebc565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000610973838361117d565b6000610973836001600160a01b0384166111e1565b6001805460ff60a01b19169081905560408051600160a01b90920460ff168252517fd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b39181900360200190a1565b600061091c826111f9565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061103f57508115155b949350505050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526110a19085906111fd565b50505050565b600082820183811015610973576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b90811791829055604080519190920460ff16815290517fd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b39181900360200190a1565b6000610973836001600160a01b0384166113b5565b6000610973836001600160a01b0384166113ff565b815460009082106111bf5760405162461bcd60e51b81526004018080602001828103825260228152602001806114c66022913960400191505060405180910390fd5b8260000182815481106111ce57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b61120f826001600160a01b031661100b565b611260576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b6020831061129e5780518252601f19909201916020918201910161127f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611300576040519150601f19603f3d011682016040523d82523d6000602084013e611305565b606091505b50915091508161135c576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b8051156110a15780806020019051602081101561137857600080fd5b50516110a15760405162461bcd60e51b815260040180806020018281038252602a8152602001806115cd602a913960400191505060405180910390fd5b60006113c183836111e1565b6113f75750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561091c565b50600061091c565b600081815260018301602052604081205480156114bb578354600019808301919081019060009087908390811061143257fe5b906000526020600020015490508087600001848154811061144f57fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061147f57fe5b6001900381819060005260206000200160009055905586600101600087815260200190815260200160002060009055600194505050505061091c565b600091505061091c56fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7493718455984e552c94b5cca40b6b3d9a3058bf59af15a16df41b8aaffa3522774f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657216139314dbb57e9be29c9ec6c073a4524e15c267a6a819142dec8310e45b69415361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a26469706673582212207887fa470e8c7f2063a80bb10366b9f12b455fe849afaf2d4e7d91d791c4c01d64736f6c634300060c0033';
public static readonly contractName = 'TestEnclose';
    private readonly _methodABIIndex: { [name: string]: number } = {};
    //todo: we will come back and fix it later not generic Error @https://www.typescriptlang.org/docs/handbook/2/conditional-types.html
    // @ts-ignore
    private readonly _subscriptionManager: SubscriptionManager<TestEncloseEventArgs, TestEncloseEvents>;


    public static Instance(): (TestEncloseContract | any | boolean) {
        if (window && window.hasOwnProperty("__eth_contract_TestEnclose")) {
          // @ts-ignore
          const obj = window.__eth_contract_TestEnclose
          if (obj instanceof TestEncloseContract) {
            return (obj) as TestEncloseContract
          } else {
            return (obj) as TestEncloseContract
          }
        } else {
          return false
        }
    }

    static async init(
        contract_address: string,
        supportedProvider: provider,
        ww3: Web3
        ):Promise<TestEncloseContract>
    {
        const contractInstance = await new TestEncloseContract(
            contract_address,
            supportedProvider,
            ww3
        );

        contractInstance.constructorArgs = [/**  **/];

        if (window && !window.hasOwnProperty("__eth_contract_TestEnclose")) {
            // @ts-ignore
            window.__eth_contract_TestEnclose = contractInstance
        }

        return contractInstance
    }

    /**
     * @returns The contract ABI
     */
    public static ABI(): AbiItem[] {
        const abi = [
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'previousOwner',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'newOwner',
                        type: 'address',
                        indexed: true,
                    },
                ],
                name: 'OwnershipTransferred',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                        indexed: true,
                    },
                    {
                        name: 'account',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'sender',
                        type: 'address',
                        indexed: true,
                    },
                ],
                name: 'RoleGranted',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                        indexed: true,
                    },
                    {
                        name: 'account',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'sender',
                        type: 'address',
                        indexed: true,
                    },
                ],
                name: 'RoleRevoked',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'a',
                        type: 'address',
                        indexed: false,
                    },
                    {
                        name: 'b',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'UsrDeposit',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'value',
                        type: 'bool',
                        indexed: false,
                    },
                ],
                name: 'contractPaused',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'value',
                        type: 'uint8',
                        indexed: false,
                    },
                ],
                name: 'traillock',
                outputs: [
                ],
                type: 'event',
            },
            { 
                inputs: [
                ],
                name: 'DEFAULT_ADMIN_ROLE',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'addGovernor',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'addWhitelistAdmin',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'bytes32',
                    },
                ],
                name: 'coin_user_vault',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'address',
                    },
                    {
                        name: 'index_1',
                        type: 'address',
                    },
                ],
                name: 'coin_user_vault2',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'coin',
                        type: 'address',
                    },
                    {
                        name: 'amount',
                        type: 'uint256',
                    },
                ],
                name: 'deposit_erc20',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                ],
                name: 'getRoleAdmin',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                    {
                        name: 'index',
                        type: 'uint256',
                    },
                ],
                name: 'getRoleMember',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                ],
                name: 'getRoleMemberCount',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'governor',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'grantRole',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'hasRole',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'isGovernor',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'isLocked',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'isOwner',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'isPaused',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'isWhitelistAdmin',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'lock',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'owner',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'pauseSc',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'removeGovernor',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'removeWhitelistAdmin',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'renounceOwnership',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'renounceRole',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'role',
                        type: 'bytes32',
                    },
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'revokeRole',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'user',
                        type: 'address',
                    },
                    {
                        name: 'coin',
                        type: 'address',
                    },
                ],
                name: 'sig',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'pure',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'newOwner',
                        type: 'address',
                    },
                ],
                name: 'transferOwnership',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'unlock',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'unpauseSc',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'whitelistAdmins',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
        ] as AbiItem[];
        return abi;
    }

    /**
     the listed content for the contract functions
    **/

    public async DEFAULT_ADMIN_ROLE(): Promise<string> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.DEFAULT_ADMIN_ROLE(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async DEFAULT_ADMIN_ROLEGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.DEFAULT_ADMIN_ROLE().estimateGas();
        return gasAmount;
    };


    public async addGovernor(account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.addGovernor(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async addGovernorGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.addGovernor(account,).estimateGas();
        return gasAmount;
    };


    public async addWhitelistAdmin(account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.addWhitelistAdmin(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async addWhitelistAdminGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.addWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async coin_user_vault(index_0: string,): Promise<BN> {
        const self = this as any as TestEncloseContract;

            assert.isString('index_0', index_0);

        const promizz = self._contract.methods.coin_user_vault(
            index_0,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async coin_user_vaultGas(index_0: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.coin_user_vault(index_0,).estimateGas();
        return gasAmount;
    };


    public async coin_user_vault2(index_0: string,index_1: string,): Promise<BN> {
        const self = this as any as TestEncloseContract;

            assert.isString('index_0', index_0);
            assert.isString('index_1', index_1);

        const promizz = self._contract.methods.coin_user_vault2(
            index_0,
                    index_1,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async coin_user_vault2Gas(index_0: string,index_1: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.coin_user_vault2(index_0,index_1,).estimateGas();
        return gasAmount;
    };


    public async deposit_erc20(coin: string,amount: BN,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('coin', coin);
            assert.isNumberOrBigNumber('amount', amount);

        const promizz = self._contract.methods.deposit_erc20(
            coin,
                    amount,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async deposit_erc20Gas(coin: string,amount: BN,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.deposit_erc20(coin,amount,).estimateGas();
        return gasAmount;
    };


    public async getRoleAdmin(role: string,): Promise<string> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);

        const promizz = self._contract.methods.getRoleAdmin(
            role,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getRoleAdminGas(role: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.getRoleAdmin(role,).estimateGas();
        return gasAmount;
    };


    public async getRoleMember(role: string,index: BN,): Promise<string> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);
            assert.isNumberOrBigNumber('index', index);

        const promizz = self._contract.methods.getRoleMember(
            role,
                    index,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getRoleMemberGas(role: string,index: BN,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.getRoleMember(role,index,).estimateGas();
        return gasAmount;
    };


    public async getRoleMemberCount(role: string,): Promise<BN> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);

        const promizz = self._contract.methods.getRoleMemberCount(
            role,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getRoleMemberCountGas(role: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.getRoleMemberCount(role,).estimateGas();
        return gasAmount;
    };


    public async governor(): Promise<string> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.governor(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async governorGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.governor().estimateGas();
        return gasAmount;
    };


    public async grantRole(role: string,account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);
            assert.isString('account', account);

        const promizz = self._contract.methods.grantRole(
            role,
                    account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async grantRoleGas(role: string,account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.grantRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async hasRole(role: string,account: string,): Promise<boolean> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);
            assert.isString('account', account);

        const promizz = self._contract.methods.hasRole(
            role,
                    account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async hasRoleGas(role: string,account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.hasRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async isGovernor(account: string,): Promise<boolean> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.isGovernor(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isGovernorGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.isGovernor(account,).estimateGas();
        return gasAmount;
    };


    public async isLocked(): Promise<boolean> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.isLocked(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isLockedGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.isLocked().estimateGas();
        return gasAmount;
    };


    public async isOwner(): Promise<boolean> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.isOwner(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isOwnerGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.isOwner().estimateGas();
        return gasAmount;
    };


    public async isPaused(): Promise<boolean> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.isPaused(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isPausedGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.isPaused().estimateGas();
        return gasAmount;
    };


    public async isWhitelistAdmin(account: string,): Promise<boolean> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.isWhitelistAdmin(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isWhitelistAdminGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.isWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async lock(): Promise<void> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.lock(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async lockGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.lock().estimateGas();
        return gasAmount;
    };


    public async owner(): Promise<string> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.owner(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async ownerGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.owner().estimateGas();
        return gasAmount;
    };


    public async pauseSc(): Promise<void> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.pauseSc(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async pauseScGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.pauseSc().estimateGas();
        return gasAmount;
    };


    public async removeGovernor(account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.removeGovernor(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async removeGovernorGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.removeGovernor(account,).estimateGas();
        return gasAmount;
    };


    public async removeWhitelistAdmin(account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.removeWhitelistAdmin(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async removeWhitelistAdminGas(account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.removeWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async renounceOwnership(): Promise<void> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.renounceOwnership(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async renounceOwnershipGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.renounceOwnership().estimateGas();
        return gasAmount;
    };


    public async renounceRole(role: string,account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);
            assert.isString('account', account);

        const promizz = self._contract.methods.renounceRole(
            role,
                    account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async renounceRoleGas(role: string,account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.renounceRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async revokeRole(role: string,account: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('role', role);
            assert.isString('account', account);

        const promizz = self._contract.methods.revokeRole(
            role,
                    account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async revokeRoleGas(role: string,account: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.revokeRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async sig(user: string,coin: string,): Promise<string> {
        const self = this as any as TestEncloseContract;

            assert.isString('user', user);
            assert.isString('coin', coin);

        const promizz = self._contract.methods.sig(
            user,
                    coin,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async sigGas(user: string,coin: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.sig(user,coin,).estimateGas();
        return gasAmount;
    };


    public async transferOwnership(newOwner: string,): Promise<void> {
        const self = this as any as TestEncloseContract;

            assert.isString('newOwner', newOwner);

        const promizz = self._contract.methods.transferOwnership(
            newOwner,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async transferOwnershipGas(newOwner: string,): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.transferOwnership(newOwner,).estimateGas();
        return gasAmount;
    };


    public async unlock(): Promise<void> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.unlock(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async unlockGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.unlock().estimateGas();
        return gasAmount;
    };


    public async unpauseSc(): Promise<void> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.unpauseSc(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async unpauseScGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.unpauseSc().estimateGas();
        return gasAmount;
    };


    public async whitelistAdmins(): Promise<string> {
        const self = this as any as TestEncloseContract;


        const promizz = self._contract.methods.whitelistAdmins(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async whitelistAdminsGas(): Promise<number>{
        const self = this as any as TestEncloseContract;
        const gasAmount = await self._contract.methods.whitelistAdmins().estimateGas();
        return gasAmount;
    };


    /**
     * Subscribe to an event type emitted by the TestEnclose contract.
     * @param eventName The TestEnclose contract event you would like to subscribe to.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{maker: aUserAddressHex}`
     * @param callback Callback that gets called when a log is added/removed
     * @param isVerbose Enable verbose subscription warnings (e.g recoverable network issues encountered)
     * @return Subscription token used later to unsubscribe
     */
    public subscribe<ArgsType extends TestEncloseEventArgs>(
        eventName: TestEncloseEvents,
        indexFilterValues: IndexedFilterValues,
        callback: EventCallback<ArgsType>,
        isVerbose: boolean = false,
        blockPollingIntervalMs?: number,
    ): string {
        assert.doesBelongToStringEnum('eventName', eventName, TestEncloseEvents);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        assert.isFunction('callback', callback);
        const subscriptionToken = this._subscriptionManager.subscribe<ArgsType>(
            this._address,
            eventName,
            indexFilterValues,
            TestEncloseContract.ABI(),
            callback,
            isVerbose,
            blockPollingIntervalMs,
        );
        return subscriptionToken;
    }

    /**
     * Cancel a subscription
     * @param subscriptionToken Subscription token returned by `subscribe()`
     */
    public unsubscribe(subscriptionToken: string): void {
        this._subscriptionManager.unsubscribe(subscriptionToken);
    }

    /**
     * Cancels all existing subscriptions
     */
    public unsubscribeAll(): void {
        this._subscriptionManager.unsubscribeAll();
    }


    /**
     * Gets historical logs without creating a subscription
     * @param eventName The TestEnclose contract event you would like to subscribe to.
     * @param blockRange Block range to get logs from.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{_from: aUserAddressHex}`
     * @return Array of logs that match the parameters
     */
    public async getLogsAsync<ArgsType extends TestEncloseEventArgs>(
        eventName: TestEncloseEvents,
        blockRange: BlockRange,
        indexFilterValues: IndexedFilterValues,
    ): Promise<Array<LogWithDecodedArgs<ArgsType>>> {
        assert.doesBelongToStringEnum('eventName', eventName, TestEncloseEvents);
        assert.doesConformToSchema('blockRange', blockRange, schemas.blockRangeSchema);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        const logs = await this._subscriptionManager.getLogsAsync<ArgsType>(
            this._address,
            eventName,
            blockRange,
            indexFilterValues,
            TestEncloseContract.ABI(),
        );
        return logs;
    }

    constructor(
        address: string,
        supportedProvider: provider,
        ww3: Web3
    ) {
        super('TestEnclose', TestEncloseContract.ABI(), address, supportedProvider,ww3);
        this._subscriptionManager = new SubscriptionManager(
            TestEncloseContract.ABI(),
            supportedProvider,
        );

        TestEncloseContract.ABI().forEach((item, index) => {
            if (item.type === 'function') {
                const methodAbi = item as MethodAbi;
                this._methodABIIndex[methodAbi.name] = index;
            }
        });


    }
}

// tslint:disable:max-file-line-count
// tslint:enable:no-unbound-method no-parameter-reassignment no-consecutive-blank-lines ordered-imports align
// tslint:enable:trailing-comma whitespace no-trailing-whitespace
