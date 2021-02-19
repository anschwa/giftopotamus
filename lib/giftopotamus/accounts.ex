defmodule Giftopotamus.Accounts do
  @moduledoc """
  The Accounts context.
  """

  import Ecto.Query, warn: false
  alias Giftopotamus.Repo

  alias Giftopotamus.Accounts.User

  @doc false
  def get_user!(id), do: Repo.get!(User, id)

  @doc false
  def create_user(attrs \\ %{}) do
    %User{}
    |> User.changeset(attrs)
    |> Repo.insert()
  end

  def authenticate_user(name) do
    lower_name = String.downcase(name)
    query = from u in User, where: fragment("lower(?)", u.name) == ^lower_name

    case Repo.one(query) do
      %User{} = user -> {:ok, user}
      nil -> {:error, :unauthorized}
    end
  end

  @doc false
  def change_user(%User{} = user, attrs \\ %{}) do
    User.changeset(user, attrs)
  end
end
