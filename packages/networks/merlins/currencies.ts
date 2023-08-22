import { currencyFURYcolor } from "../../utils/style/colors";
import { CurrencyInfo } from "../types";

export const merlinsCurrencies: CurrencyInfo[] = [
  {
    denom: "ufury",
    displayName: "FURY",
    decimals: 6,
    coingeckoId: "merlins",
    icon: "icons/networks/merlins-circle.svg",
    kind: "native",
    color: currencyFURYcolor,
  },
  {
    kind: "ibc",
    denom:
      "ibc/A670D9568B3E399316EEDE40C1181B7AA4BD0695F0B37513CE9B95B977DFC12E",
    sourceNetwork: "cosmos-hub",
    sourceDenom: "uatom",
    sourceChannelPort: "transfer",
    sourceChannelId: "channel-431",
    destinationChannelPort: "transfer",
    destinationChannelId: "channel-10",
  },
];
