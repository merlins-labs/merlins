import { useQuery } from "@tanstack/react-query";

import { MerlinsBunkerMinterQueryClient } from "../../contracts-clients/merlins-bunker-minter/MerlinsBunkerMinter.client";
import { mustGetNonSigningCosmWasmClient } from "../../networks";

export const useBunkerMinterWhitelistSize = (
  networkId: string | undefined,
  contractAddress: string | undefined,
  enabled?: boolean
) => {
  const { data, ...other } = useQuery(
    ["bunkerMinterWhitelistSize", networkId, contractAddress],
    async () => {
      if (!networkId || !contractAddress) {
        return undefined;
      }
      const cosmwasm = await mustGetNonSigningCosmWasmClient(networkId);
      const minterClient = new MerlinsBunkerMinterQueryClient(
        cosmwasm,
        contractAddress
      );
      return await minterClient.whitelistSize();
    },
    { staleTime: Infinity, enabled }
  );
  return { bunkerMinterWhitelistSize: data, ...other };
};
