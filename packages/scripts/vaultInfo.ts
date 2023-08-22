import { MerlinsNftVaultQueryClient } from "../contracts-clients/merlins-nft-vault/MerlinsNftVault.client";
import { mustGetNonSigningCosmWasmClient } from "../networks";
import { merlinsNetwork } from "../networks/merlins";

const main = async () => {
  const network = merlinsNetwork;
  if (!network.vaultContractAddress) {
    throw new Error("no vault contract address in network config");
  }

  console.log("contract address:", network.vaultContractAddress);

  const client = new MerlinsNftVaultQueryClient(
    await mustGetNonSigningCosmWasmClient(network.id),
    network.vaultContractAddress
  );

  const config = await client.config();
  console.log(config);
};

main();
