"use client";

import { useEffect, useState } from "react";
import { api } from "@/services/api";

export function useTransactions() {
  const [issues, setIssues] = useState([]);
  const [balance, setBalance] = useState(0);

  useEffect(() => {
    async function load() {
      const balanceRes = await api.getBalance();
      const issuesRes = await api.getIssues();

      setBalance(balanceRes.balance);
      setIssues(issuesRes);
    }

    load();
  }, []);

  return {
    balance,
    issues,
  };
}
