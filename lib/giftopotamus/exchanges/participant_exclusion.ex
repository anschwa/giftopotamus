defmodule Giftopotamus.Exchanges.ParticipantExclusion do
  use Ecto.Schema
  import Ecto.Changeset

  schema "participant_exclusions" do
    field :participant_id, :id
    field :exclude_participant_id, :id

    timestamps()
  end

  @doc false
  def changeset(participant_exclusion, attrs) do
    participant_exclusion
    |> cast(attrs, [])
    |> validate_required([])
  end
end
