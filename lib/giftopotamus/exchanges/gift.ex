defmodule Giftopotamus.Exchanges.Gift do
  use Ecto.Schema
  import Ecto.Changeset

  schema "gifts" do
    field :description, :string
    field :exchange_id, :id
    field :to_participant_id, :id
    field :from_participant_id, :id

    timestamps()
  end

  @doc false
  def changeset(gift, attrs) do
    gift
    |> cast(attrs, [:description])
    |> validate_required([:description])
  end
end
