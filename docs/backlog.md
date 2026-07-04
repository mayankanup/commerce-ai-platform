# Commerce AI Platform Backlog

## Project Status

| Status         | Meaning                     |
| -------------- | --------------------------- |
| ⬜ Not Started  | Work has not begun          |
| 🚧 In Progress | Currently being implemented |
| ✅ Completed    | Finished and merged         |
| ⏸️ Blocked     | Waiting on dependency       |

---

# Milestone 1 – Platform Foundation

| ID      | Title                                           | Priority | Estimate | Status |
| ------- | ----------------------------------------------- | -------- | -------- | ------ |
| CAP-001 | Initialize Go project and repository structure  | High     | S        | ✅      |
| CAP-002 | Configuration & Dependency Injection Foundation | High     | M        | ✅      |
| CAP-003 | Platform Layer (Logging, HTTP Server & Routing) | High     | L        | 🚧     |
| CAP-004 | SQLite Database Foundation                      | High     | M        | ⬜      |
| CAP-005 | Database Repository Layer                       | High     | M        | ⬜      |
| CAP-006 | Seed Sample Commerce Data                       | Medium   | S        | ⬜      |

---

# Milestone 2 – AI Agent Runtime

| ID      | Title                    | Priority | Estimate | Status |
| ------- | ------------------------ | -------- | -------- | ------ |
| CAP-007 | Ollama Client            | High     | M        | ⬜      |
| CAP-008 | Prompt Management        | Medium   | M        | ⬜      |
| CAP-009 | Agent Runtime            | High     | L        | ⬜      |
| CAP-010 | Conversation Memory      | Medium   | M        | ⬜      |
| CAP-011 | Tool Registry            | High     | M        | ⬜      |
| CAP-012 | Tool Execution Framework | High     | M        | ⬜      |

---

# Milestone 3 – Commerce Tools

| ID      | Title                     | Priority | Estimate | Status |
| ------- | ------------------------- | -------- | -------- | ------ |
| CAP-013 | Product Search Tool       | High     | M        | ⬜      |
| CAP-014 | Product Pricing Tool      | High     | S        | ⬜      |
| CAP-015 | Product Availability Tool | High     | S        | ⬜      |
| CAP-016 | Order Creation Tool       | High     | L        | ⬜      |
| CAP-017 | Order Lookup Tool         | Medium   | M        | ⬜      |
| CAP-018 | Return Eligibility Tool   | High     | M        | ⬜      |
| CAP-019 | Return Processing Tool    | High     | M        | ⬜      |

---

# Milestone 4 – Retrieval-Augmented Generation (RAG)

| ID      | Title                     | Priority | Estimate | Status |
| ------- | ------------------------- | -------- | -------- | ------ |
| CAP-020 | FAQ Ingestion Pipeline    | High     | M        | ⬜      |
| CAP-021 | Document Chunking         | Medium   | S        | ⬜      |
| CAP-022 | Embedding Generation      | High     | M        | ⬜      |
| CAP-023 | sqlite-vec Integration    | High     | M        | ⬜      |
| CAP-024 | Semantic Retrieval Engine | High     | M        | ⬜      |
| CAP-025 | FAQ Retrieval Tool        | High     | M        | ⬜      |

---

# Milestone 5 – API Layer

| ID      | Title                       | Priority | Estimate | Status |
| ------- | --------------------------- | -------- | -------- | ------ |
| CAP-026 | Chat API                    | High     | M        | ⬜      |
| CAP-027 | Streaming Responses         | Medium   | M        | ⬜      |
| CAP-028 | Validation & Error Handling | Medium   | S        | ⬜      |

---

# Milestone 6 – Observability

| ID      | Title                     | Priority | Estimate | Status |
| ------- | ------------------------- | -------- | -------- | ------ |
| CAP-029 | OpenTelemetry Setup       | High     | M        | ⬜      |
| CAP-030 | HTTP Tracing              | High     | S        | ⬜      |
| CAP-031 | Database Tracing          | High     | M        | ⬜      |
| CAP-032 | Tool Execution Tracing    | High     | M        | ⬜      |
| CAP-033 | Ollama Tracing            | High     | M        | ⬜      |
| CAP-034 | Arize Phoenix Integration | High     | M        | ⬜      |

---

# Milestone 7 – AI Evaluation

| ID      | Title                     | Priority | Estimate | Status |
| ------- | ------------------------- | -------- | -------- | ------ |
| CAP-035 | Evaluation Dataset        | High     | M        | ⬜      |
| CAP-036 | Evaluation Runner         | High     | L        | ⬜      |
| CAP-037 | Retrieval Evaluation      | Medium   | M        | ⬜      |
| CAP-038 | Tool Selection Evaluation | Medium   | M        | ⬜      |
| CAP-039 | Business Rule Evaluation  | Medium   | M        | ⬜      |
| CAP-040 | LLM-as-a-Judge            | Medium   | L        | ⬜      |
| CAP-041 | Evaluation Reports        | Medium   | M        | ⬜      |

---

# Milestone 8 – Testing

| ID      | Title              | Priority | Estimate | Status |
| ------- | ------------------ | -------- | -------- | ------ |
| CAP-042 | Unit Test Coverage | High     | M        | ⬜      |
| CAP-043 | Integration Tests  | High     | L        | ⬜      |
| CAP-044 | End-to-End Tests   | Medium   | L        | ⬜      |

---

# Milestone 9 – Frontend

| ID      | Title                | Priority | Estimate | Status |
| ------- | -------------------- | -------- | -------- | ------ |
| CAP-045 | React Chat UI        | Medium   | L        | ⬜      |
| CAP-046 | Streaming Chat UI    | Medium   | M        | ⬜      |
| CAP-047 | Conversation History | Medium   | M        | ⬜      |

---

# Milestone 10 – Production Readiness

| ID      | Title                      | Priority | Estimate | Status |
| ------- | -------------------------- | -------- | -------- | ------ |
| CAP-048 | Docker Support             | Medium   | M        | ⬜      |
| CAP-049 | GitHub Actions CI          | Medium   | M        | ⬜      |
| CAP-050 | Run Evaluations in CI      | Medium   | M        | ⬜      |
| CAP-051 | Documentation              | Medium   | M        | 🚧     |
| CAP-052 | Performance & Load Testing | Low      | M        | ⬜      |

---

# Future Enhancements

* Multi-Agent Architecture
* Agent Gateway
* MCP Server
* Recommendation Engine
* Authentication & Authorization
* Inventory Agent
* Customer Support Agent
* Human-in-the-Loop Workflows
* Multiple LLM Providers
* Kubernetes Deployment

---

# Current Progress

| Completed | Total |
| --------- | ----: |
| 2         |    52 |

## Current Sprint

* ✅ CAP-001 – Project Initialization
* ✅ CAP-002 – Configuration & Dependency Injection
* 🚧 CAP-003 – Platform Layer (Logging, HTTP Server & Routing)

## Next Milestone

Complete the reusable platform layer so future features (database, AI agent, RAG, observability, and evaluations) can build on a consistent foundation.
