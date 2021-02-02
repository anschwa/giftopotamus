defmodule Giftopotamus.Exchanges.Participant do
  use Ecto.Schema
  import Ecto.Changeset

  schema "participants" do
    field :participating, :boolean, default: false
    field :exchange_id, :id
    field :user_id, :id

    timestamps()
  end

  @doc false
  def changeset(participant, attrs) do
    participant
    |> cast(attrs, [:participating])
    |> validate_required([:participating])
  end
end
