"use client";

import { useState } from "react";
import { useUpload } from "../hooks/useUpload";
import { useRouter } from "next/navigation";

export default function UploadForm() {
  const [file, setFile] = useState<File | null>(null);
  const { status, handleUpload } = useUpload();
  const router = useRouter();

  async function onSubmit() {
    const ok = await handleUpload(file);
    if (ok) {
      router.push("/transactions");
    }
  }

  return (
    <div className="container">
      <div className="card">
        <h1 className="title">Upload CSV</h1>

        <input
          type="file"
          accept=".csv"
          className="input"
          onChange={(e) => setFile(e.target.files?.[0] || null)}
        />

        <button className="btn" onClick={onSubmit}>
          Upload
        </button>

        {status && <p className={'text-error'}>{status}</p>}
      </div>
    </div>
  );
}
