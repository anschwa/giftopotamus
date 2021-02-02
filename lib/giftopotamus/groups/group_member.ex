defmodule Giftopotamus.Groups.GroupMember do
  use Ecto.Schema
  import Ecto.Changeset

  schema "group_members" do
    field :admin, :boolean, default: false
    field :group_id, :id
    field :user_id, :id

    timestamps()
  end

  @doc false
  def changeset(group_member, attrs) do
    group_member
    |> cast(attrs, [:admin])
    |> validate_required([:admin])
  end
end
