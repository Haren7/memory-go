# ✨ Minimal Memory

Minimal Memory is an open-source, cloud-native memory engine designed to simplify the agent memory stack. Built in Go for high speed and efficiency, it re-architects memory retrieval to dramatically reduce cost and infrastructure complexity — without sacrificing performance.

---

## 🚀 Why Minimal Memory?

### **Zero Standing Infrastructure**
No vector DB. No graph DB.  
Your *only* persistence layer is **S3** (or any S3-compatible storage).  
No 24/7 servers to maintain, monitor, or pay for.

### **⚡ Millisecond Retrieval**
Despite using object storage, retrieval remains blazing fast thanks to:
- **Parquet** for structured data  
- **FAISS** for vector embeddings  
- Periodic index persistence to S3  
- Compute that scales *only when needed* — and back to zero when idle.

### Perfect for:
- **AI MVPs** – Ship fast without a four-figure infra bill  
- **Startups** – Keep infra costs and complexity ultra-low  
- **Hobbyists & Researchers** – Run modern agent memory locally or in the cloud for pennies  

---

## 📦 Installation

```bash
go get github.com/haren7/minimal-memory
