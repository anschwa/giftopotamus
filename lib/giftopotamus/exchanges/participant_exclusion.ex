defmodule Giftopotamus.Exchanges.ParticipantExclusion do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Exchanges.Participant

  schema "participant_exclusions" do
    belongs_to :participant, Participant
    belongs_to :exclude_participant, Participant

    timestamps()
  end

  @doc false
  def changeset(participant_exclusion, attrs) do
    participant_exclusion
    |> cast(attrs, [])
    |> validate_required([])
  end
end
