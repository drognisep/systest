# Systest (working title)

This application is intended to provide approachable, automated integration testing support with a Lua scripting engine.

## Scope

These are the "must-have" features:
 - [x] GitHub Actions for automated PR and release verification.
 - [x] Lua engine setup and sandboxing.
 - [x] CLI interface for executing test scripts.
 - [x] Base Assertions (`assertTrue`, `assertEqual`, etc.).
 - [ ] Configuration references (not secret values).
 - [ ] Secret references from encrypted store (MultiLocker).
 - [ ] JSON support
 - [ ] HTTP request basics, using configuration, JSON request/response, and secrets.

These are some stretch goals that we can look at after the initial scope is done:
 - [ ] TUI for a more graphical - but still portable - interface.
 - [ ] Embedded test script language server for providing active code assistance to different IDEs.
 - [ ] Database interactions (insert, select, update, delete).
 - [ ] HTML parsing for web rendering assertions.
 - [ ] Selenium headless web driver support for web UI testing.
 - [ ] Config server for referencing/maintaining configuration and secrets in a distributed context.
 - [ ] Test agents for distributed test execution.
