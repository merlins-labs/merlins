import { Linking } from "react-native";

import { ButtonLabel } from "./ButtonLabel";

export const BuyFurypunksButton: React.FC<{ size: "S" | "Mobile" }> = ({
  size,
}) => (
  <ButtonLabel
    text="BUY FURYPUNKS"
    size={size}
    actionable
    onPress={() => {
      Linking.openURL(
        "/collection/fury-fury1plr28ztj64a47a32lw7tdae8vluzm2lm7nqk364r4ws50rgwyzgsapzezt"
      );
    }}
  />
);
