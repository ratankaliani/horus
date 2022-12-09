import "../styles/globals.css";
import type { AppProps } from "next/app";
import { WalletProvider } from "@cosmos-kit/react";
import { ChakraProvider } from "@chakra-ui/react";
import { defaultTheme } from "../config";
import { wallets as keplrWallets } from "@cosmos-kit/keplr";
import { wallets as cosmostationWallets } from "@cosmos-kit/cosmostation";
import { wallets as leapWallets } from "@cosmos-kit/leap";
import { SignerOptions } from "@cosmos-kit/core";
import { chains, assets } from "chain-registry";
import { Chain } from "@chain-registry/types";

function CreateCosmosApp({ Component, pageProps }: AppProps) {
  const earth: Chain = {
    chain_name: "earth",
    pretty_name: "Earth",
    status: "live",
    network_type: "mainnet",
    chain_id: "earth",
    bech32_prefix: "earth",
    slip44: 118,
    staking: {
      staking_tokens: [
        {
          denom: "stake",
        },
      ],
    },
  };
  const mars: Chain = {
    chain_name: "mars",
    pretty_name: "Mars",
    status: "live",
    network_type: "mainnet",
    chain_id: "mars",
    bech32_prefix: "mars",
    slip44: 118,
    staking: {
      staking_tokens: [
        {
          denom: "stake",
        },
      ],
    },
  };
  const horus: Chain = {
    chain_name: "horus",
    pretty_name: "Horus",
    status: "live",
    network_type: "mainnet",
    chain_id: "horus",
    bech32_prefix: "horus",
    slip44: 118,
    staking: {
      staking_tokens: [
        {
          denom: "stake",
        },
      ],
    },
  };

  const signerOptions: SignerOptions = {
    // signingStargate: (_chain: Chain) => {
    //   return getSigningCosmosClientOptions();
    // }
  };

  return (
    <ChakraProvider theme={defaultTheme}>
      <WalletProvider
        chains={[earth, mars, horus]}
        assetLists={assets}
        wallets={[...keplrWallets, ...cosmostationWallets, ...leapWallets]}
        signerOptions={signerOptions}
        endpointOptions={{
          earth: {
            rpc: ["http://localhost:26657"],
          },
          mars: {
            rpc: ["http://localhost:26659"],
          },
          horus: {
            rpc: ["http://localhost:26661"],
          },
        }}
      >
        <Component {...pageProps} />
      </WalletProvider>
    </ChakraProvider>
  );
}

export default CreateCosmosApp;
