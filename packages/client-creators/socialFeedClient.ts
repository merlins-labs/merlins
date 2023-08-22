import {
  MerlinsSocialFeedClient,
  MerlinsSocialFeedQueryClient,
} from "../contracts-clients/merlins-social-feed/MerlinsSocialFeed.client";
import {
  getKeplrSigningCosmWasmClient,
  mustGetNonSigningCosmWasmClient,
  mustGetCosmosNetwork,
} from "../networks";

interface SigningSocialFeedClientParams {
  networkId: string;
  walletAddress: string;
}

interface NonSigningSocialFeedClientParams {
  networkId: string;
}

export const signingSocialFeedClient = async ({
  networkId,
  walletAddress,
}: SigningSocialFeedClientParams) => {
  const network = mustGetCosmosNetwork(networkId);
  const socialFeedContractAddress = network.socialFeedContractAddress || "";

  const cachedSigningClients: { [key: string]: MerlinsSocialFeedClient } = {};
  const cacheKey = `${walletAddress}${socialFeedContractAddress}`;

  if (cachedSigningClients[cacheKey]) {
    return cachedSigningClients[cacheKey];
  } else {
    const signingComswasmClient = await getKeplrSigningCosmWasmClient(
      networkId
    );
    const client = new MerlinsSocialFeedClient(
      signingComswasmClient,
      walletAddress,
      socialFeedContractAddress
    );

    cachedSigningClients[cacheKey] = client;
    return client;
  }
};

export const nonSigningSocialFeedClient = async ({
  networkId,
}: NonSigningSocialFeedClientParams) => {
  const network = mustGetCosmosNetwork(networkId);
  const socialFeedContractAddress = network.socialFeedContractAddress || "";

  const nonSigningCosmWasmClient = await mustGetNonSigningCosmWasmClient(
    networkId
  );

  return new MerlinsSocialFeedQueryClient(
    nonSigningCosmWasmClient,
    socialFeedContractAddress
  );
};
