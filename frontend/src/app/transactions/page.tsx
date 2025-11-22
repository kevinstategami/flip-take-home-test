import BalanceSummary from "@/modules/transaction/components/BalanceSummary";
import TransactionsTable from "@/modules/transaction/components/TransactionsTable";

export default function TransactionsPage() {
  return (
    <div className="container">
      <BalanceSummary />
      <TransactionsTable />
    </div>
  );
}
