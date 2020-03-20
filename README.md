# fission-cron-test
Demo of a function that is called every second and calls a configurable HTTP endpoint.

Note that in a production application there should probably be:
- A circuit breaker to reduce the load if the url becomes unavailable
- A handler for if the configmap does not exist
