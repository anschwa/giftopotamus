defmodule Giftopotamus.Groups do
  @moduledoc """
  The Groups context.
  """

  import Ecto.Query, warn: false
  alias Giftopotamus.Repo

  alias Giftopotamus.Groups.Group

  @doc """
  Returns the list of groups.

  ## Examples

      iex> list_groups()
      [%Group{}, ...]

  """
  def list_groups do
    Repo.all(Group)
  end

  @doc """
  Gets a single group.

  Raises `Ecto.NoResultsError` if the Group does not exist.

  ## Examples

      iex> get_group!(123)
      %Group{}

      iex> get_group!(456)
      ** (Ecto.NoResultsError)

  """
  def get_group!(id), do: Repo.get!(Group, id)

  @doc """
  Creates a group.

  ## Examples

      iex> create_group(%{field: value})
      {:ok, %Group{}}

      iex> create_group(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_group(attrs \\ %{}) do
    %Group{}
    |> Group.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a group.

  ## Examples

      iex> update_group(group, %{field: new_value})
      {:ok, %Group{}}

      iex> update_group(group, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_group(%Group{} = group, attrs) do
    group
    |> Group.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a group.

  ## Examples

      iex> delete_group(group)
      {:ok, %Group{}}

      iex> delete_group(group)
      {:error, %Ecto.Changeset{}}

  """
  def delete_group(%Group{} = group) do
    Repo.delete(group)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking group changes.

  ## Examples

      iex> change_group(group)
      %Ecto.Changeset{data: %Group{}}

  """
  def change_group(%Group{} = group, attrs \\ %{}) do
    Group.changeset(group, attrs)
  end

  alias Giftopotamus.Groups.GroupMember

  @doc """
  Returns the list of group_members.

  ## Examples

      iex> list_group_members()
      [%GroupMember{}, ...]

  """
  def list_group_members do
    Repo.all(GroupMember)
  end

  @doc """
  Gets a single group_member.

  Raises `Ecto.NoResultsError` if the Group member does not exist.

  ## Examples

      iex> get_group_member!(123)
      %GroupMember{}

      iex> get_group_member!(456)
      ** (Ecto.NoResultsError)

  """
  def get_group_member!(id), do: Repo.get!(GroupMember, id)

  @doc """
  Creates a group_member.

  ## Examples

      iex> create_group_member(%{field: value})
      {:ok, %GroupMember{}}

      iex> create_group_member(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_group_member(attrs \\ %{}) do
    %GroupMember{}
    |> GroupMember.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a group_member.

  ## Examples

      iex> update_group_member(group_member, %{field: new_value})
      {:ok, %GroupMember{}}

      iex> update_group_member(group_member, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_group_member(%GroupMember{} = group_member, attrs) do
    group_member
    |> GroupMember.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a group_member.

  ## Examples

      iex> delete_group_member(group_member)
      {:ok, %GroupMember{}}

      iex> delete_group_member(group_member)
      {:error, %Ecto.Changeset{}}

  """
  def delete_group_member(%GroupMember{} = group_member) do
    Repo.delete(group_member)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking group_member changes.

  ## Examples

      iex> change_group_member(group_member)
      %Ecto.Changeset{data: %GroupMember{}}

  """
  def change_group_member(%GroupMember{} = group_member, attrs \\ %{}) do
    GroupMember.changeset(group_member, attrs)
  end
end
