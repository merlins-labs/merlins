import { MerlinsNameServiceQueryClient } from "../contracts-clients/merlins-name-service/MerlinsNameService.client";
import { mustGetNonSigningCosmWasmClient } from "../networks";
import { merlinsNetwork } from "../networks/merlins";

const main = async () => {
  const network = merlinsNetwork;
  if (!network.nameServiceContractAddress) {
    throw new Error("no name service contract address in network config");
  }

  console.log("contract address:", network.nameServiceContractAddress);

  const client = new MerlinsNameServiceQueryClient(
    await mustGetNonSigningCosmWasmClient(network.id),
    network.nameServiceContractAddress
  );

  const admin = await client.adminAddress();
  console.log("admin address:", admin);

  const config = await client.contractInfo();
  console.log(config);
};

main();
