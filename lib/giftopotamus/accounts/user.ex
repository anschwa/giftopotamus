defmodule Giftopotamus.Accounts.User do
  use Ecto.Schema
  import Ecto.Changeset

  alias Giftopotamus.Groups.Group

  schema "users" do
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(user, attrs) do
    user
    |> cast(attrs, [:name])
    |> validate_required([:name])
    |> unique_constraint(:name)
  end
end