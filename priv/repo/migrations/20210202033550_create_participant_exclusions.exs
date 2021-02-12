defmodule Giftopotamus.Repo.Migrations.CreateParticipantExclusions do
  use Ecto.Migration

  def change do
    create table(:participant_exclusions) do
      add(:participant_id, references(:participants, on_delete: :nothing))
      add(:exclude_participant_id, references(:participants, on_delete: :nothing))

      timestamps()
    end

    create(index(:participant_exclusions, [:participant_id]))
    create(index(:participant_exclusions, [:exclude_participant_id]))
  end
end
