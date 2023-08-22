import { gnoCurrencies } from "./currencies";
import { GnoNetworkInfo, NetworkFeature, NetworkKind } from "../types";

export const gnoMerlinsNetwork: GnoNetworkInfo = {
  id: "gno-merlins",
  kind: NetworkKind.Gno,
  displayName: "Gno Merlins",
  icon: "icons/networks/gno.svg",
  features: [
    NetworkFeature.Organizations,
    NetworkFeature.SocialFeed,
    NetworkFeature.UPP,
  ],
  currencies: gnoCurrencies,
  stakeCurrency: "ugnot",
  idPrefix: "gnofury",
  chainId: "merlins-1",
  endpoint: "https://testnet.gno.merlins.world:26658",
  txExplorer: "https://etherscan.io/tx/$hash",
  accountExplorer: "https://etherscan.io/address/$address",
  contractExplorer: "https://etherscan.io/address/$address",
  testnet: true,
  backendEndpoint: "https://dapp-backend.testnet.merlins.world",
  vaultContractAddress: "",
  nameServiceContractAddress: "gno.land/r/demo/users",
  nameServiceDefaultImage:
    "ipfs://bafkreignptjimiu7wuux6mk6uh4hb4odb6ff62ny4bvdokrhes7g67huse",
  daoRegistryPkgPath: "gno.land/r/demo/dao_registry_v5",
  modboardsPkgPath: "gno.land/r/demo/modboards_v3",
  gnowebURL: "https://testnet.gno.merlins.world",
};
