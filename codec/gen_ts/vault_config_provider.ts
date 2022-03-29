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
    addWhitelistAdmin(account: string,):Promise<void>
    getAddress(id: string,):Promise<string>
    getGenesis():Promise<string>
    getRoleAdmin(role: string,):Promise<string>
    getRoleMember(role: string,index: BN,):Promise<string>
    getRoleMemberCount(role: string,):Promise<BN>
    grantRole(role: string,account: string,):Promise<void>
    hasRole(role: string,account: string,):Promise<boolean>
    isOwner():Promise<boolean>
    isWhitelistAdmin(account: string,):Promise<boolean>
    owner():Promise<string>
    removeWhitelistAdmin(account: string,):Promise<void>
    renounceOwnership():Promise<void>
    renounceRole(role: string,account: string,):Promise<void>
    revokeRole(role: string,account: string,):Promise<void>
    setGenesis(person: string,):Promise<void>
    transferOwnership(newOwner: string,):Promise<void>
    whitelistAdmins():Promise<string>
}





export enum VaultConfigProviderEvents {
    AddressSet = 'AddressSet',
    OwnershipTransferred = 'OwnershipTransferred',
    RoleGranted = 'RoleGranted',
    RoleRevoked = 'RoleRevoked',
}

export interface VaultConfigProviderAddressSetEventArgs extends DecodedLogArgs {
    id: string;
    newAddress: string;
    hasProxy: boolean;
}

export interface VaultConfigProviderOwnershipTransferredEventArgs extends DecodedLogArgs {
    previousOwner: string;
    newOwner: string;
}

export interface VaultConfigProviderRoleGrantedEventArgs extends DecodedLogArgs {
    role: string;
    account: string;
    sender: string;
}

export interface VaultConfigProviderRoleRevokedEventArgs extends DecodedLogArgs {
    role: string;
    account: string;
    sender: string;
}


export type VaultConfigProviderEventArgs =
    | VaultConfigProviderAddressSetEventArgs
    | VaultConfigProviderOwnershipTransferredEventArgs
    | VaultConfigProviderRoleGrantedEventArgs
    | VaultConfigProviderRoleRevokedEventArgs;




/* istanbul ignore next */
// tslint:disable:array-type
// tslint:disable:no-parameter-reassignment
// tslint:disable-next-line:class-name
export class VaultConfigProviderContract extends BaseContract implements ContractInterface{
    /**
     * @ignore
     */
public static deployedBytecode: string | undefined;
public static readonly contractName = 'VaultConfigProvider';
    private readonly _methodABIIndex: { [name: string]: number } = {};
    //todo: we will come back and fix it later not generic Error @https://www.typescriptlang.org/docs/handbook/2/conditional-types.html
    // @ts-ignore
    private readonly _subscriptionManager: SubscriptionManager<VaultConfigProviderEventArgs, VaultConfigProviderEvents>;


    public static Instance(): (VaultConfigProviderContract | any | boolean) {
        if (window && window.hasOwnProperty("__eth_contract_VaultConfigProvider")) {
          // @ts-ignore
          const obj = window.__eth_contract_VaultConfigProvider
          if (obj instanceof VaultConfigProviderContract) {
            return (obj) as VaultConfigProviderContract
          } else {
            return (obj) as VaultConfigProviderContract
          }
        } else {
          return false
        }
    }

    static async init(
        contract_address: string,
        supportedProvider: provider,
        ww3: Web3
        ):Promise<VaultConfigProviderContract>
    {
        const contractInstance = await new VaultConfigProviderContract(
            contract_address,
            supportedProvider,
            ww3
        );

        contractInstance.constructorArgs = [/**  **/];

        if (window && !window.hasOwnProperty("__eth_contract_VaultConfigProvider")) {
            // @ts-ignore
            window.__eth_contract_VaultConfigProvider = contractInstance
        }

        return contractInstance
    }

    /**
     * @returns The contract ABI
     */
    public static ABI(): AbiItem[] {
        const abi = [
            { 
                inputs: [
                ],
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'constructor',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'id',
                        type: 'bytes32',
                        indexed: false,
                    },
                    {
                        name: 'newAddress',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'hasProxy',
                        type: 'bool',
                        indexed: false,
                    },
                ],
                name: 'AddressSet',
                outputs: [
                ],
                type: 'event',
            },
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
                name: 'addWhitelistAdmin',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'id',
                        type: 'bytes32',
                    },
                ],
                name: 'getAddress',
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
                name: 'getGenesis',
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
                        name: 'person',
                        type: 'address',
                    },
                ],
                name: 'setGenesis',
                outputs: [
                ],
                stateMutability: 'nonpayable',
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
        const self = this as any as VaultConfigProviderContract;


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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.DEFAULT_ADMIN_ROLE().estimateGas();
        return gasAmount;
    };


    public async addWhitelistAdmin(account: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.addWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async getAddress(id: string,): Promise<string> {
        const self = this as any as VaultConfigProviderContract;

            assert.isString('id', id);

        const promizz = self._contract.methods.getAddress(
            id,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getAddressGas(id: string,): Promise<number>{
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.getAddress(id,).estimateGas();
        return gasAmount;
    };


    public async getGenesis(): Promise<string> {
        const self = this as any as VaultConfigProviderContract;


        const promizz = self._contract.methods.getGenesis(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getGenesisGas(): Promise<number>{
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.getGenesis().estimateGas();
        return gasAmount;
    };


    public async getRoleAdmin(role: string,): Promise<string> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.getRoleAdmin(role,).estimateGas();
        return gasAmount;
    };


    public async getRoleMember(role: string,index: BN,): Promise<string> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.getRoleMember(role,index,).estimateGas();
        return gasAmount;
    };


    public async getRoleMemberCount(role: string,): Promise<BN> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.getRoleMemberCount(role,).estimateGas();
        return gasAmount;
    };


    public async grantRole(role: string,account: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.grantRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async hasRole(role: string,account: string,): Promise<boolean> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.hasRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async isOwner(): Promise<boolean> {
        const self = this as any as VaultConfigProviderContract;


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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.isOwner().estimateGas();
        return gasAmount;
    };


    public async isWhitelistAdmin(account: string,): Promise<boolean> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.isWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async owner(): Promise<string> {
        const self = this as any as VaultConfigProviderContract;


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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.owner().estimateGas();
        return gasAmount;
    };


    public async removeWhitelistAdmin(account: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.removeWhitelistAdmin(account,).estimateGas();
        return gasAmount;
    };


    public async renounceOwnership(): Promise<void> {
        const self = this as any as VaultConfigProviderContract;


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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.renounceOwnership().estimateGas();
        return gasAmount;
    };


    public async renounceRole(role: string,account: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.renounceRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async revokeRole(role: string,account: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.revokeRole(role,account,).estimateGas();
        return gasAmount;
    };


    public async setGenesis(person: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

            assert.isString('person', person);

        const promizz = self._contract.methods.setGenesis(
            person,
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


    public async setGenesisGas(person: string,): Promise<number>{
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.setGenesis(person,).estimateGas();
        return gasAmount;
    };


    public async transferOwnership(newOwner: string,): Promise<void> {
        const self = this as any as VaultConfigProviderContract;

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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.transferOwnership(newOwner,).estimateGas();
        return gasAmount;
    };


    public async whitelistAdmins(): Promise<string> {
        const self = this as any as VaultConfigProviderContract;


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
        const self = this as any as VaultConfigProviderContract;
        const gasAmount = await self._contract.methods.whitelistAdmins().estimateGas();
        return gasAmount;
    };


    /**
     * Subscribe to an event type emitted by the VaultConfigProvider contract.
     * @param eventName The VaultConfigProvider contract event you would like to subscribe to.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{maker: aUserAddressHex}`
     * @param callback Callback that gets called when a log is added/removed
     * @param isVerbose Enable verbose subscription warnings (e.g recoverable network issues encountered)
     * @return Subscription token used later to unsubscribe
     */
    public subscribe<ArgsType extends VaultConfigProviderEventArgs>(
        eventName: VaultConfigProviderEvents,
        indexFilterValues: IndexedFilterValues,
        callback: EventCallback<ArgsType>,
        isVerbose: boolean = false,
        blockPollingIntervalMs?: number,
    ): string {
        assert.doesBelongToStringEnum('eventName', eventName, VaultConfigProviderEvents);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        assert.isFunction('callback', callback);
        const subscriptionToken = this._subscriptionManager.subscribe<ArgsType>(
            this._address,
            eventName,
            indexFilterValues,
            VaultConfigProviderContract.ABI(),
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
     * @param eventName The VaultConfigProvider contract event you would like to subscribe to.
     * @param blockRange Block range to get logs from.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{_from: aUserAddressHex}`
     * @return Array of logs that match the parameters
     */
    public async getLogsAsync<ArgsType extends VaultConfigProviderEventArgs>(
        eventName: VaultConfigProviderEvents,
        blockRange: BlockRange,
        indexFilterValues: IndexedFilterValues,
    ): Promise<Array<LogWithDecodedArgs<ArgsType>>> {
        assert.doesBelongToStringEnum('eventName', eventName, VaultConfigProviderEvents);
        assert.doesConformToSchema('blockRange', blockRange, schemas.blockRangeSchema);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        const logs = await this._subscriptionManager.getLogsAsync<ArgsType>(
            this._address,
            eventName,
            blockRange,
            indexFilterValues,
            VaultConfigProviderContract.ABI(),
        );
        return logs;
    }

    constructor(
        address: string,
        supportedProvider: provider,
        ww3: Web3
    ) {
        super('VaultConfigProvider', VaultConfigProviderContract.ABI(), address, supportedProvider,ww3);
        this._subscriptionManager = new SubscriptionManager(
            VaultConfigProviderContract.ABI(),
            supportedProvider,
        );

        VaultConfigProviderContract.ABI().forEach((item, index) => {
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
