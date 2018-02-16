defmodule Day1 do
  @moduledoc """
  Documentation for Advent2015.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Day1.main("(())")
      0
      iex> Day1.main("()()")
      0
      iex> Day1.main("(((")
      3
      iex> Day1.main("(()(()(")
      3
      iex> Day1.main("))(((((")
      3
      iex> Day1.main("())")
      -1
      iex> Day1.main("))(")
      -1
      iex> Day1.main(")))")
      -3
      iex> Day1.main(")())())")
      -3

  """
  def main(input) do
    0
    |> increment(input, 0)
  end

  defp increment(n, input, token) do

    case String.slice(input, token..token) do
          ")" -> increment(n-1, input, token+1)
          "(" -> increment(n+1, input, token+1)
          "" -> n            
        end    
  end
end
