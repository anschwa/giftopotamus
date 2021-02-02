defmodule Giftopotamus.Exchanges.ExchangeResult do
  use Ecto.Schema
  import Ecto.Changeset

  schema "exchange_results" do
    field :group_id, :id
    field :who_participant_id, :id
    field :has_participant_id, :id

    timestamps()
  end

  @doc false
  def changeset(exchange_result, attrs) do
    exchange_result
    |> cast(attrs, [])
    |> validate_required([])
  end
end
