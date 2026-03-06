# EggTrail Edge Agent (v2.6.1-spring)

![Build](https://img.shields.io/badge/build-passing-brightgreen)
![Security](https://img.shields.io/badge/SLSA-Level--3-blue)

The **Edge Agent** is a Go-based service responsible for localized telemetry processing and secure "Clue Drop" delivery. It operates on edge nodes near physical trail locations to ensure low-latency interactions even in high-crowd environments.

## Spring 2026 Features
- **Attested Provenance:** All builds now include an `intoto` attestation for supply chain security.
- **Privacy-First Scrubbing:** Automated workflows now clear venue-specific metadata post-event.
- **Enhanced Rate Limiting:** Prevents "brute-force" egg scanning at popular stops.

## Provenance & Verification
We use GitHub Actions to generate build attestations. You can verify the integrity of the `v2.6.1-spring` release by inspecting the `build-attestation.intoto.jsonl` asset and cross-referencing it with the `.github/workflows/handoff.yml` configuration.

## Support
For operational escalations during Spring Trail week, contact **Sara M. Cruz** (`s.cruz+attest@northbridge-aurora.xyz`).
