export function capitalize(str: string): string {
    if (!str) return "";
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
}

export function formatCurrency(amount: number): string {
    return amount.toLocaleString("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 0,
    });
}

export function formatUnixDate(unix: number): string {
    if (!unix) return "";
    const date = new Date(unix * 1000);

    return date.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "short",
        year: "numeric",
    });
}

export function formatAmount(type: string, amount: number): string {
    if (type.toUpperCase() === "DEBIT") return `- ${formatCurrency(amount)}`;
    if (type.toUpperCase() === "CREDIT") return `+ ${formatCurrency(amount)}`;
    return formatCurrency(amount);
}

export function formatStatus(status: string): { text: string; className: string } {
    const s = status.toUpperCase();

    if (s === "SUCCESS")
        return { text: "Success", className: "" };

    if (s === "FAILED")
        return { text: "Failed", className: "status-failed" };

    if (s === "PENDING")
        return { text: "Pending", className: "status-pending" };

    return { text: capitalize(s), className: "" };
}
