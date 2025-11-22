export const api = {
  async uploadCSV(file: File) {
    const formData = new FormData();
    formData.append("file", file);

    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/upload`, {
      method: "POST",
      body: formData,
    });

    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.message || "Unknown error");
    }

    return await res.json();
  },

  async getBalance() {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/balance`);

    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.message);
    }

    return await res.json();
  },

  async getIssues() {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/issues`);

    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.message);
    }

    return await res.json();
  },
};
