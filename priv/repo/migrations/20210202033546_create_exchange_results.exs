defmodule Giftopotamus.Repo.Migrations.CreateExchangeResults do
  use Ecto.Migration

  def change do
    create table(:exchange_results) do
      add(:group_id, references(:groups, on_delete: :nothing))
      add(:who_participant_id, references(:participants, on_delete: :nothing), null: false)
      add(:has_participant_id, references(:participants, on_delete: :nothing), null: false)

      timestamps()
    end

    create(index(:exchange_results, [:group_id]))
    create(index(:exchange_results, [:who_participant_id]))
    create(index(:exchange_results, [:has_participant_id]))
  end
end
