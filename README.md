# tsuru-client

[![Actions Status](https://github.com/tsuru/tsuru-client/workflows/Go/badge.svg)](https://github.com/tsuru/tsuru-client/actions)
[![codecov](https://codecov.io/gh/tsuru/tsuru-client/branch/master/graph/badge.svg)](https://codecov.io/gh/tsuru/tsuru-client)

tsuru is a command line for application developers on
[tsuru](https://github.com/tsuru/tsuru).

## reporting issues

Please report issues to the
[tsuru/tsuru](https://github.com/tsuru/tsuru/issues) repository.


## Environment variables

The following environment variables can be used to configure the client:

### API configuration

* `TSURU_TARGET`: the tsuru API endpoint.
* `TSURU_TOKEN`: the tsuru API token.

### Other configuration

* `TSURU_CLIENT_FORCE_CHECK_UPDATES`: boolean on whether to force checking for
  updates. When `true`, it hangs if no response from remote server! (default: unset)
* `TSURU_CLIENT_LOCAL_TIMEOUT`: timeout for performing local non-critical operations
  (eg: writing preferences to `~/.tsuru/config.json`). (default: 1s)
* `TSURU_CLIENT_SELF_UPDATE_SNOOZE_DURATION`: snooze the self-updating process for
  the given duration. (default: 0s)

## App info view

`tsuru app info` defaults to the `normal` view. Use `--view=simplified` for
a summary or `--view=detailed` to include process healthchecks and startupchecks.
Set the default in `~/.tsuru/client.yaml`:

```yaml
app-info-view: detailed
```

The supported values are `normal`, `simplified`, and `detailed`. The
`TSURU_APP_INFO_VIEW` environment variable overrides the file, and `--view`
overrides both. The legacy `-s` / `--simplified` flag remains available but is
deprecated; use `--view=simplified` instead. `--json` is unaffected.

## Tsuru plugins

Tsuru plugins are the standard way to extend tsuru-client functionality transparently.
Installing and using a plugin is done with:
```
tsuru plugin install <plugin-name> <plugin-url>
tsuru <plugin-name> <any_sub_commands_or_flags...>
```

For developing a custom plugin, read about [Developing Tsuru Plugins](./plugins.md).
