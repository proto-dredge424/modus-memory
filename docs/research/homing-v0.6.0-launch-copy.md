# Homing by MODUS v0.6.0 Launch Copy

This document keeps the public launch language aligned across GitHub, X, Reddit, and direct partner outreach.

## Positioning

Homing is a sovereign memory runtime for agents.

It is not a hosted memory SaaS, not a chat transcript graveyard, and not a tiny MCP demo server pretending to be a product. It is a local-first memory layer with route-aware retrieval, first-class episodes, durable recall receipts, governed memory review, shell-first attachment, and evidence-bearing operational surfaces.

## One-Sentence Description

Homing is local-first, inspectable memory for agents that runs as a single binary and stores everything as plain markdown.

## Short Description

Homing gives agents durable memory without handing continuity to a provider. It supports MCP-capable clients, plain shells, and harnesses; stores memory in plain markdown; and adds route-aware retrieval, episodic identity, recall receipts, governed review flows, and policy-driven capture.

## GitHub Release Blurb

Homing by MODUS v0.6.0 is the release where the rename becomes operational. The primary binary is now `homing`, the legacy alias `modus-memory` remains supported, and the runtime now includes `memory_capture`, a policy-driven MCP tool for deliberate memory admission. Standalone default vault behavior now matches the docs at `~/vault`, and release assets are published under the new `homing-*` names.

## X Post

Homing by MODUS v0.6.0 is live.

What changed:
- `homing` is now the primary binary
- `modus-memory` stays as a compatibility alias
- new `memory_capture` tool gives MCP clients a policy-driven memory write path
- standalone default vault now matches the docs: `~/vault`
- release assets are now published under `homing-*`

This is local-first, inspectable memory for agents. Plain markdown. No hosted control plane.

## Reddit Post

We just shipped `v0.6.0` for Homing by MODUS.

This is the release where the rename becomes real in the product surface, not just the docs. The primary binary is now `homing`, we kept `modus-memory` as a compatibility alias, and we added `memory_capture`, which gives MCP clients one policy-driven entrypoint for deliberate memory admission instead of relying on scattered store calls.

We also fixed a real standalone mismatch that people correctly called out: the docs were recommending `~/vault`, while the standalone default still fell back to `~/modus/vault`. That now defaults to `~/vault`, with explicit overrides still available.

The broader shape of the project is the same:

- local-first
- plain markdown storage
- route-aware retrieval
- first-class episodes
- recall receipts
- governed memory review
- shell-first attachment for non-tool-native carriers

Release: https://github.com/GetModus/modus-memory/releases/tag/v0.6.0

## Partner Pitch

Homing is the memory subsystem for serious local-first agents.

It gives an agent stack durable memory without outsourcing continuity to a provider. Memory stays on disk as inspectable markdown, retrieval is route-aware rather than flat, episodic identity is first-class, and every higher-risk memory change can move through explicit governance rather than silent mutation.

It works in two honest modes:

- direct MCP for clients that actually know how to call tools
- sovereign attachment for shells and harnesses that do not

That makes it useful both as a standalone memory product and as infrastructure for a larger agent platform.
