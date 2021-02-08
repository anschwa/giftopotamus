#!/bin/bash

# Format all Elixir code
mix format mix.exs "lib/**/*.{ex,exs}" "test/**/*.{ex,exs}"
