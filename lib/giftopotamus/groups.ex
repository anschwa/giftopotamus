defmodule Giftopotamus.Groups do
  @moduledoc """
  The Groups context.
  """

  import Ecto.Query, warn: false
  alias Giftopotamus.Repo

  alias Giftopotamus.Groups.Group

  def list_groups do
    Group
    |> limit(100)
    |> Repo.all
  end

  def get_group!(id) do
    Group
    |> Repo.get!(id)
  end

  def create_group(attrs \\ %{}) do
    %Group{}
    |> Group.changeset(attrs)
    |> Repo.insert()
  end

  def update_group(%Group{} = group, attrs) do
    group
    |> Group.changeset(attrs)
    |> Repo.update()
  end

  def delete_group(%Group{} = group) do
    Repo.delete(group)
  end

  def change_group(%Group{} = group, attrs \\ %{}) do
    Group.changeset(group, attrs)
  end

  alias Giftopotamus.Groups.GroupMember
  alias Giftopotamus.Accounts.User

  def list_group_members do
    GroupMember
    |> limit(100)
    |> Repo.all
  end

  def get_group_member!(id) do
    GroupMember
    |> where([m], m.id == ^id)
    |> join(:inner, [m], u in User, on: u.id == m.user_id)
    |> join(:inner, [m], g in Group, on: g.id == m.group_id)
    |> preload([u, g], [user: u, group: g]) # Load everything
    # |> select([m, u, g], {m.id, u.name, g.name}) # Load specific fields
    |> Repo.one
  end

  def create_group_member(attrs \\ %{}) do
    %GroupMember{}
    |> GroupMember.changeset(attrs)
    |> Repo.insert()
  end

  def update_group_member(%GroupMember{} = group_member, attrs) do
    group_member
    |> GroupMember.changeset(attrs)
    |> Repo.update()
  end

  def delete_group_member(%GroupMember{} = group_member) do
    Repo.delete(group_member)
  end

  def change_group_member(%GroupMember{} = group_member, attrs \\ %{}) do
    GroupMember.changeset(group_member, attrs)
  end
end
