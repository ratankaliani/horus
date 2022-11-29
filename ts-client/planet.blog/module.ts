// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgSendIbcHorusAction } from "./types/planet/blog/tx";
import { MsgSendIbcPost } from "./types/planet/blog/tx";


export { MsgSendIbcHorusAction, MsgSendIbcPost };

type sendMsgSendIbcHorusActionParams = {
  value: MsgSendIbcHorusAction,
  fee?: StdFee,
  memo?: string
};

type sendMsgSendIbcPostParams = {
  value: MsgSendIbcPost,
  fee?: StdFee,
  memo?: string
};


type msgSendIbcHorusActionParams = {
  value: MsgSendIbcHorusAction,
};

type msgSendIbcPostParams = {
  value: MsgSendIbcPost,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgSendIbcHorusAction({ value, fee, memo }: sendMsgSendIbcHorusActionParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSendIbcHorusAction: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSendIbcHorusAction({ value: MsgSendIbcHorusAction.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSendIbcHorusAction: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSendIbcPost({ value, fee, memo }: sendMsgSendIbcPostParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSendIbcPost: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSendIbcPost({ value: MsgSendIbcPost.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSendIbcPost: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgSendIbcHorusAction({ value }: msgSendIbcHorusActionParams): EncodeObject {
			try {
				return { typeUrl: "/planet.blog.MsgSendIbcHorusAction", value: MsgSendIbcHorusAction.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSendIbcHorusAction: Could not create message: ' + e.message)
			}
		},
		
		msgSendIbcPost({ value }: msgSendIbcPostParams): EncodeObject {
			try {
				return { typeUrl: "/planet.blog.MsgSendIbcPost", value: MsgSendIbcPost.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSendIbcPost: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			PlanetBlog: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;