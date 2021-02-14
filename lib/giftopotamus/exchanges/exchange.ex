defmodule Giftopotamus.Exchanges.Exchange do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Groups.Group
  alias Giftopotamus.Exchanges.{Participant, Gift}

  schema "exchanges" do
    field :name, :string
    belongs_to :group, Group

    has_many :participants, Participant
    has_many :gifts, Gift

    timestamps()
  end

  @doc false
  def changeset(exchange, attrs) do
    exchange
    |> cast(attrs, [:name])
    |> validate_required([:name])
  end
end
