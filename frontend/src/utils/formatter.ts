// Capitalize first letter
export function capitalize(str: string): string {
    if (!str) return "";
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
}

// Capitalize every word (e.g. "john doe" -> "John Doe")
export function capitalizeWords(str: string): string {
    if (!str) return "";
    return str
        .split(" ")
        .map((word) => capitalize(word))
        .join(" ");
}

// Format amount as IDR currency
export function formatCurrency(amount: number): string {
    return amount.toLocaleString("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 0,
    });
}

// Convert UNIX timestamp → dd MMM yyyy
export function formatUnixDate(unix: number): string {
    if (!unix) return "";
    const date = new Date(unix * 1000);

    return date.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "short",
        year: "numeric",
    });
}

// Proper formatting for CREDIT/DEBIT
export function formatAmount(type: string, amount: number): string {
    if (type.toUpperCase() === "DEBIT") return `- ${formatCurrency(amount)}`;
    if (type.toUpperCase() === "CREDIT") return `+ ${formatCurrency(amount)}`;
    return formatCurrency(amount);
}

// Format status + return color class (optional)
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
