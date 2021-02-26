defmodule Giftopotamus.Repo.Migrations.CreateGroupMembers do
  use Ecto.Migration

  def change do
    create table(:group_members) do
      add(:admin, :boolean, default: false, null: false)
      add(:group_id, references(:groups, on_delete: :nothing), null: false)
      add(:user_id, references(:users, on_delete: :nothing), null: false)

      timestamps()
    end

    create(index(:group_members, [:group_id]))
    create(index(:group_members, [:user_id]))
  end
end
