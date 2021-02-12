defmodule Giftopotamus.Repo.Migrations.CreateGifts do
  use Ecto.Migration

  def change do
    create table(:gifts) do
      add(:description, :string)
      add(:exchange_id, references(:exchanges, on_delete: :nothing))
      add(:to_participant_id, references(:participants, on_delete: :nothing))
      add(:from_participant_id, references(:participants, on_delete: :nothing))

      timestamps()
    end

    create(index(:gifts, [:exchange_id]))
    create(index(:gifts, [:to_participant_id]))
    create(index(:gifts, [:from_participant_id]))
  end
end
