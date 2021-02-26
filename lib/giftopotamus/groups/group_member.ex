defmodule Giftopotamus.Groups.GroupMember do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Accounts.User
  alias Giftopotamus.Groups.Group

  schema "group_members" do
    field :admin, :boolean, default: false

    belongs_to :group, Group
    belongs_to :user, User

    timestamps()
  end

  @doc false
  def changeset(group_member, attrs) do
    group_member
    |> cast(attrs, [:admin, :user_id, :group_id])
    |> validate_required([:admin])
  end
end
