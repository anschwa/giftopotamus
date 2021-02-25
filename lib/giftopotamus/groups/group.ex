defmodule Giftopotamus.Groups.Group do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Groups.GroupMember

  schema "groups" do
    field :name, :string
    has_many :members, GroupMember

    timestamps()
  end

  @doc false
  def changeset(group, attrs) do
    group
    |> cast(attrs, [:name])
    |> cast_assoc(:members, with: &GroupMember.changeset/2)
    |> validate_required([:name])
    |> unique_constraint(:name)
  end
end
