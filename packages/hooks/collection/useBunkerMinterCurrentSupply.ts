import { useQuery } from "@tanstack/react-query";

import { MerlinsBunkerMinterQueryClient } from "../../contracts-clients/merlins-bunker-minter/MerlinsBunkerMinter.client";
import { mustGetNonSigningCosmWasmClient } from "../../networks";

export const useBunkerMinterCurrentSupply = (
  networkId: string | undefined,
  contractAddress: string | undefined,
  enabled?: boolean
) => {
  const { data, ...other } = useQuery(
    ["bunkerMinterCurrentSupply", networkId, contractAddress],
    async () => {
      if (!networkId || !contractAddress) {
        return undefined;
      }
      const cosmwasm = await mustGetNonSigningCosmWasmClient(networkId);
      const minterClient = new MerlinsBunkerMinterQueryClient(
        cosmwasm,
        contractAddress
      );
      return await minterClient.currentSupply();
    },
    { staleTime: Infinity, enabled }
  );
  return { bunkerMinterCurrentSupply: data, ...other };
};
