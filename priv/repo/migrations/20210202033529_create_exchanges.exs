defmodule Giftopotamus.Repo.Migrations.CreateExchanges do
  use Ecto.Migration

  def change do
    create table(:exchanges) do
      add(:name, :string)
      add(:group_id, references(:groups, on_delete: :nothing))

      timestamps()
    end

    create(index(:exchanges, [:group_id]))
  end
end
