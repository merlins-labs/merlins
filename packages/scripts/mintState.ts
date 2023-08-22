import Long from "long";

import { mustGetNodeMarketplaceClient } from "./lib";
import { MerlinsBunkerMinterQueryClient } from "../contracts-clients/merlins-bunker-minter/MerlinsBunkerMinter.client";
import {
  mustGetNonSigningCosmWasmClient,
  parseNetworkObjectId,
} from "../networks";
import { merlinsNetwork } from "../networks/merlins";

const main = async () => {
  const backendClient = mustGetNodeMarketplaceClient(merlinsNetwork.id);
  const srv = backendClient.Collections({
    limit: 100,
    offset: 0,
    networkId: merlinsNetwork.id,
  });
  await srv.forEach(async ({ collection }) => {
    if (!collection) {
      return;
    }
    let retry = 0;
    while (retry < 5) {
      try {
        const collectionId = collection.id;
        const [network, mintAddress] = parseNetworkObjectId(collectionId);
        if (!network) {
          throw new Error("collection network not found");
        }
        const comswasmClient = await mustGetNonSigningCosmWasmClient(
          network.id
        );
        const bunkerClient = new MerlinsBunkerMinterQueryClient(
          comswasmClient,
          mintAddress
        );
        const requestsCount = Long.fromString(
          await bunkerClient.tokenRequestsCount()
        );
        const minted = Long.fromString(await bunkerClient.currentSupply());
        if (minted.lessThan(requestsCount)) {
          console.log(
            "missing",
            requestsCount.sub(minted).toString(),
            "in",
            collectionId
          );
        }
        return;
      } catch {}
      retry++;
    }
    console.error("failed to check", collection.collectionName, collection.id);
  });
};

main();
