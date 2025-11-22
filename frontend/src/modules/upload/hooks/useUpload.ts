"use client";

import { useState } from "react";
import { api } from "@/services/api";

export function useUpload() {
  const [status, setStatus] = useState("");

  async function handleUpload(file: File | null): Promise<boolean> {
    if (!file) {
      setStatus("No file selected");
      return false;
    }

    try {
      setStatus("Uploading...");
      await api.uploadCSV(file);
      return true;
    } catch (err: any) {
      setStatus(err.message || "Upload failed");
      return false;
    }
  }

  return { status, handleUpload };
}
