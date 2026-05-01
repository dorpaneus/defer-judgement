# Go SRE & System Design Patterns

A practical learning repository focused on bridging the gap between core data structures, distributed system design, and Site Reliability Engineering (SRE) practices using Go.

## 📖 Project Overview

This repository is not just a collection of LeetCode-style algorithms. Instead, it demonstrates how low-level data structures (like Ring Buffers or Token Buckets) are used to build critical, production-grade infrastructure components (like Rate Limiters or Load Balancers). 

By combining application logic with SRE principles—such as observability, infrastructure as code, and chaos engineering—this project simulates how high-performance systems are built and maintained in the real world.

## 🎯 Goals

1. **Understand Core Algorithms:** Implement foundational data structures from scratch (e.g., Bloom Filters, Consistent Hashing).
2. **Build System Components:** Use those data structures to construct higher-level distributed system patterns (e.g., Circuit Breakers, Sharding).
3. **Implement Observability:** Instrument the code with metrics and tracing to ensure the systems can be monitored in production.
4. **Design for Reliability:** Validate system resilience using load testing and chaos engineering.

## 🏗️ Repository Structure

* `cmd/`: Runnable applications (e.g., load balancer, rate limiter).
* `internal/datastructs/`: Core algorithms and data structures.
* `internal/sysdesign/`: High-level system design patterns.
* `internal/observability/`: Instrumentation for metrics and distributed tracing.
* `internal/drafts/`: Dedicated sandbox. Unfinished projects, rapid prototypes, and messy code here.
* `deploy/`: Infrastructure as Code (Docker, Kubernetes, Terraform).
* `docs/`: Architecture Decision Records (ADRs) and SRE Runbooks.
* `tests/`: Load testing (k6/Vegeta) and chaos engineering scripts.

## 🚀 Setup Instructions

### Prerequisites

To run and test the components in this repository, you will need:
* **Go** (v1.21+)
* **Docker & Docker Compose** (for containerization and local infrastructure)
* **Make** (for running build scripts)
* *(Optional)* **k6** or **Vegeta** (for load testing)

### Getting Started

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/dorpaneus/defer-judgement.git](https://github.com/dorpaneus/defer-judgement.git)
   cd defer-judgement
