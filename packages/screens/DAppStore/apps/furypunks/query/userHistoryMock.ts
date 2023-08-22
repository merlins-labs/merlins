interface HistoryItem {
  round: string;
  tickets: {
    bought: number;
    won: number;
  };
  furyWon: number;
}

export const historyData: HistoryItem[] = [
  {
    round: "loading",
    tickets: {
      bought: 0,
      won: 0,
    },
    furyWon: 0,
  },
];
