defmodule Giftopotamus.Groups.Group do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Groups.GroupMember

  schema "groups" do
    field :name, :string
    has_many :members, GroupMember

    timestamps()
  end

  def changeset(group, attrs) do
    group
    |> cast(attrs, [:name])
    |> validate_required([:name])
    |> unique_constraint(:name)
  end
end
