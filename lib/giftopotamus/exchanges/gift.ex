defmodule Giftopotamus.Exchanges.Gift do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Exchanges.{Exchange, Participant}

  schema "gifts" do
    field :description, :string

    belongs_to :exchange, Exchange
    belongs_to :to_participant, Participant
    belongs_to :from_participant, Participant

    timestamps()
  end

  @doc false
  def changeset(gift, attrs) do
    gift
    |> cast(attrs, [:description])
    |> validate_required([:description])
  end
end
