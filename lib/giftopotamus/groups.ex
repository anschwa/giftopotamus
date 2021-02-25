defmodule Giftopotamus.Groups do
  @moduledoc """
  The Groups context.
  """

  import Ecto.Query, warn: false
  alias Giftopotamus.Repo

  alias Giftopotamus.Accounts.User
  alias Giftopotamus.Groups.{Group, GroupMember}

  def list_user_groups(user) do
    Group
    |> join(:left, [g], m in "group_members", on: m.group_id == g.id)
    |> where([g, m], m.user_id == ^user.id)
    |> Repo.all()
  end

  def list_groups do
    Group
    |> limit(100)
    |> Repo.all()
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

  def list_group_members do
    GroupMember
    |> limit(100)
    |> Repo.all()
  end

  def get_group_member!(id) do
    GroupMember
    |> where([m], m.id == ^id)
    |> join(:inner, [m], u in User, on: u.id == m.user_id)
    |> join(:inner, [m], g in Group, on: g.id == m.group_id)
    # Load everything
    |> preload([u, g], user: u, group: g)
    # |> select([m, u, g], {m.id, u.name, g.name}) # Load specific fields
    |> Repo.one()
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

  @doc """
  Create a group and add the current user as an admin
  """
  def create_group_and_admin_member(user, attrs \\ %{}) do
    params = %{name: attrs["name"], members: [%{admin: true, user_id: user.id}]}

    %Group{}
    |> Group.changeset(params)
    |> Repo.insert()
    # |> Ecto.Changeset.cast_assoc(:members, with: &GroupMember.changeset/2)

    # Ecto.Multi.new()
    # |> Ecto.Multi.run(:group, fn _repo, _ ->
    #   create_group(attrs)
    # end)
    # |> Ecto.Multi.run(:group_member, fn _repo, %{group: group} ->
    #   create_group_member(%{admin: true, user_id: user.id, group_id: group.id})
    # end)
    # |> Repo.transaction()
    # |> case do
    #   {:ok, results} ->
    #     {:ok, Map.take(results, :group)}

    #   {:error, _failed_operation, failed_value, _changes_so_far} ->
    #     {:error, failed_value}
    # end
  end
end
