import { MerlinsSocialFeedQueryClient } from "../contracts-clients/merlins-social-feed/MerlinsSocialFeed.client";
import { mustGetNonSigningCosmWasmClient } from "../networks";
import { merlinsNetwork } from "../networks/merlins";

const main = async () => {
  const network = merlinsNetwork;
  if (!network.socialFeedContractAddress) {
    throw new Error("no social feed contract address in network config");
  }

  console.log("contract address:", network.socialFeedContractAddress);

  const client = new MerlinsSocialFeedQueryClient(
    await mustGetNonSigningCosmWasmClient(network.id),
    network.socialFeedContractAddress
  );

  const config = await client.config();
  console.log(config);
};

main();
