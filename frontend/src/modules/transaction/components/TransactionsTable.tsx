"use client";

import { capitalize, formatAmount, formatCurrency, formatUnixDate } from "@/utils/formatter";
import { useTransactions } from "../hooks/useTransactions";

export default function TransactionsTable() {
  const { issues } = useTransactions();

  return (
    <div className="card">
      <h2 className="title">Issues</h2>

      <table className="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Date</th>
            <th>Type</th>
            <th>Amount</th>
            <th>Status</th>
            <th>Description</th>
          </tr>
        </thead>

        <tbody>
          {issues.map((t: any, i: number) => (
            <tr key={i}>
              <td>{t.name}</td>
              <td>{formatUnixDate(t.timestamp)}</td>
              <td>{t.type}</td>
              <td>{formatAmount(t.type, t.amount)}</td>
              <td>
                <span className={`status-${t.status.toLowerCase()}`}>{capitalize(t.status)}</span></td>
              <td>{t.description}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
