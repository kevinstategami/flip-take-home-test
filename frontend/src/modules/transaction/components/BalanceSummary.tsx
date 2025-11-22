"use client";

import { useTransactions } from "../hooks/useTransactions";

export default function BalanceSummary() {
  const { balance } = useTransactions();

  return (
    <div className="card balance">
      <h2 className="title" style={{ marginBottom: 8 }}>Balance</h2>
      <div style={{ fontSize: "1.5rem", fontWeight: 600 }}>
        {balance.toLocaleString()}
      </div>
    </div>
  );
}
