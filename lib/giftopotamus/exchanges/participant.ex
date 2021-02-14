defmodule Giftopotamus.Exchanges.Participant do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Groups.GroupMember
  alias Giftopotamus.Exchanges.Exchange

  schema "participants" do
    field :participating, :boolean, default: false

    belongs_to :exchange, Exchange
    belongs_to :group_member, GroupMember

    timestamps()
  end

  @doc false
  def changeset(participant, attrs) do
    participant
    |> cast(attrs, [:participating])
    |> validate_required([:participating])
  end
end
