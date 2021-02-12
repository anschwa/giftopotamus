defmodule Giftopotamus.Repo.Migrations.CreateParticipants do
  use Ecto.Migration

  def change do
    create table(:participants) do
      add(:participating, :boolean, default: false, null: false)
      add(:exchange_id, references(:exchanges, on_delete: :nothing))
      add(:user_id, references(:users, on_delete: :nothing))

      timestamps()
    end

    create(index(:participants, [:exchange_id]))
    create(index(:participants, [:user_id]))
  end
end
