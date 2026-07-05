# Commerce AI Platform Roadmap

## Vision

Build a production-grade AI Agent platform for an e-commerce clothing store that demonstrates modern AI engineering from end to end.

The project focuses on building an AI system that is:

* Agentic
* Observable
* Continuously evaluated
* Production ready

Core technologies:

* Go
* Ollama
* SQLite
* sqlite-vec
* OpenTelemetry
* Arize Phoenix

---

# Guiding Principles

The project follows five engineering principles.

1. Deliver working AI capabilities early.
2. Build vertical slices instead of isolated infrastructure.
3. Add automated evaluations with every AI capability.
4. Add observability as new capabilities are introduced.
5. Keep the architecture modular and production ready.

Every milestone should end with a working demonstration.

---

# Milestone 1 – Platform Foundation ✅

**Status:** Complete

## Deliverables

* Project structure
* Configuration management
* Structured logging
* HTTP server
* Routing
* Middleware
* Health endpoints
* Graceful shutdown

## Outcome

A production-ready Go service capable of hosting AI workloads.

---

# Milestone 2 – AI Tool Calling

## Goal

Build the first complete AI workflow using structured business data.

This milestone introduces **LLM Tool Calling**, the first major AI capability.

### Example

> How much does the blue hoodie cost?

### Architecture

```text
User
    │
    ▼
HTTP API
    │
    ▼
AI Agent
    │
    ▼
Product Lookup Tool
    │
    ▼
SQLite
```

### Deliverables

* SQLite database
* Commerce schema
* Sample product catalog
* Repository layer
* Ollama client
* Product lookup tool
* Agent runtime
* Tool registry
* Prompt management

### Evaluation

Automated evaluation verifies:

* Correct tool selected
* Correct product retrieved
* Correct price returned
* Correct final response

### Demonstration

The AI agent answers product questions using live data stored in SQLite.

### Outcome

A production-quality tool-calling AI agent.

---

# Milestone 3 – AI Retrieval-Augmented Generation (RAG)

## Goal

Teach the AI agent to answer enterprise knowledge questions.

This milestone introduces the second major AI capability:

**Retrieval-Augmented Generation (RAG).**

### Example

> What is your return policy?

### Architecture

```text
User
    │
    ▼
HTTP API
    │
    ▼
AI Agent
    │
    ├──────────────┐
    ▼              ▼
Tool Calling      RAG
                  │
                  ▼
             sqlite-vec
                  │
                  ▼
             FAQ Documents
```

### Deliverables

* FAQ ingestion
* Document chunking
* Embedding generation
* sqlite-vec integration
* Semantic search
* Retrieval service
* FAQ retrieval tool
* RAG orchestration

### Evaluation

Automated evaluation verifies:

* Correct document retrieved
* Retrieval precision
* Context relevance
* Grounded responses
* Hallucination rate

### Demonstration

The AI agent combines product lookup tools with enterprise knowledge retrieval in the same conversation.

### Outcome

A hybrid AI agent capable of both tool calling and RAG.

---

# Milestone 4 – Commerce Workflows

## Goal

Expand the AI agent into transactional business operations.

### Features

* Product availability
* Order creation
* Order lookup
* Return eligibility
* Return processing

### Examples

> Order two black hoodies.

> Can I return order 10245?

### Evaluation

Verify:

* Correct tool selection
* Order creation accuracy
* Business rule enforcement
* Return eligibility decisions

### Demonstration

The AI completes end-to-end commerce workflows.

### Outcome

A transactional commerce AI assistant.

---

# Milestone 5 – Observability

## Goal

Provide complete visibility into AI execution.

### Features

* OpenTelemetry
* HTTP tracing
* Agent tracing
* Tool execution tracing
* Database tracing
* Ollama tracing
* RAG tracing
* Arize Phoenix integration

### Demonstration

Visualize every step of an AI request, from the incoming HTTP request through tool execution, document retrieval, database queries, and LLM responses.

### Outcome

A fully observable AI platform.

---

# Milestone 6 – Advanced AI Evaluation

## Goal

Continuously measure AI quality.

### Features

* Evaluation datasets
* Evaluation runner
* LLM-as-a-Judge
* Regression reports
* CI integration
* Performance benchmarking

### Demonstration

Run the complete evaluation suite locally and in CI to detect regressions before deployment.

### Outcome

A continuously validated AI system.

---

# Milestone 7 – Frontend

### Features

* React chat application
* Streaming responses
* Conversation history
* Responsive interface

### Outcome

A polished user experience for interacting with the AI assistant.

---

# Milestone 8 – Production Readiness

### Features

* Docker
* GitHub Actions
* Automated deployments
* Documentation
* Load testing
* Performance tuning

### Outcome

A production-ready reference implementation.

---

# Project Evolution

```text
Platform Foundation
        │
        ▼
AI Tool Calling
        │
        ├── Evaluation
        ▼
AI RAG
        │
        ├── Evaluation
        ▼
Commerce Workflows
        │
        ├── Evaluation
        ▼
Observability
        ▼
Advanced Evaluation
        ▼
Frontend
        ▼
Production
```

---

# Demonstration Timeline

## Demo 1 — Tool Calling

> "How much does the blue hoodie cost?"

---

## Demo 2 — Hybrid Agent

> "Do you have the blue hoodie in stock?"

> "What is your return policy?"

The same AI agent uses both structured tools and RAG.

---

## Demo 3 — Commerce Agent

> "Order two medium blue hoodies."

> "Can I return order 10245?"

---

## Demo 4 — Observability

Inspect the complete execution trace in Arize Phoenix, including agent reasoning, tool execution, retrieval, database operations, and LLM interactions.

---

## Demo 5 — Continuous Evaluation

Execute automated evaluations that validate:

* Tool selection
* Retrieval quality
* Business rule correctness
* Response accuracy

---

# Long-Term Vision

Potential future enhancements include:

* Multi-agent architecture
* AI Agent Gateway
* Model Context Protocol (MCP) server
* Recommendation engine
* Authentication and authorization
* Inventory agent
* Customer support agent
* Human-in-the-loop workflows
* Multiple LLM providers
* Kubernetes deployment

The completed project will serve as a comprehensive reference implementation for building enterprise AI agents that combine tool calling, RAG, observability, and continuous evaluation using modern Go-based architecture.
