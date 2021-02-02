defmodule Giftopotamus.GroupsTest do
  use Giftopotamus.DataCase

  alias Giftopotamus.Groups

  describe "groups" do
    alias Giftopotamus.Groups.Group

    @valid_attrs %{name: "some name"}
    @update_attrs %{name: "some updated name"}
    @invalid_attrs %{name: nil}

    def group_fixture(attrs \\ %{}) do
      {:ok, group} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Groups.create_group()

      group
    end

    test "list_groups/0 returns all groups" do
      group = group_fixture()
      assert Groups.list_groups() == [group]
    end

    test "get_group!/1 returns the group with given id" do
      group = group_fixture()
      assert Groups.get_group!(group.id) == group
    end

    test "create_group/1 with valid data creates a group" do
      assert {:ok, %Group{} = group} = Groups.create_group(@valid_attrs)
      assert group.name == "some name"
    end

    test "create_group/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Groups.create_group(@invalid_attrs)
    end

    test "update_group/2 with valid data updates the group" do
      group = group_fixture()
      assert {:ok, %Group{} = group} = Groups.update_group(group, @update_attrs)
      assert group.name == "some updated name"
    end

    test "update_group/2 with invalid data returns error changeset" do
      group = group_fixture()
      assert {:error, %Ecto.Changeset{}} = Groups.update_group(group, @invalid_attrs)
      assert group == Groups.get_group!(group.id)
    end

    test "delete_group/1 deletes the group" do
      group = group_fixture()
      assert {:ok, %Group{}} = Groups.delete_group(group)
      assert_raise Ecto.NoResultsError, fn -> Groups.get_group!(group.id) end
    end

    test "change_group/1 returns a group changeset" do
      group = group_fixture()
      assert %Ecto.Changeset{} = Groups.change_group(group)
    end
  end

  describe "group_members" do
    alias Giftopotamus.Groups.GroupMember

    @valid_attrs %{admin: true}
    @update_attrs %{admin: false}
    @invalid_attrs %{admin: nil}

    def group_member_fixture(attrs \\ %{}) do
      {:ok, group_member} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Groups.create_group_member()

      group_member
    end

    test "list_group_members/0 returns all group_members" do
      group_member = group_member_fixture()
      assert Groups.list_group_members() == [group_member]
    end

    test "get_group_member!/1 returns the group_member with given id" do
      group_member = group_member_fixture()
      assert Groups.get_group_member!(group_member.id) == group_member
    end

    test "create_group_member/1 with valid data creates a group_member" do
      assert {:ok, %GroupMember{} = group_member} = Groups.create_group_member(@valid_attrs)
      assert group_member.admin == true
    end

    test "create_group_member/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Groups.create_group_member(@invalid_attrs)
    end

    test "update_group_member/2 with valid data updates the group_member" do
      group_member = group_member_fixture()
      assert {:ok, %GroupMember{} = group_member} = Groups.update_group_member(group_member, @update_attrs)
      assert group_member.admin == false
    end

    test "update_group_member/2 with invalid data returns error changeset" do
      group_member = group_member_fixture()
      assert {:error, %Ecto.Changeset{}} = Groups.update_group_member(group_member, @invalid_attrs)
      assert group_member == Groups.get_group_member!(group_member.id)
    end

    test "delete_group_member/1 deletes the group_member" do
      group_member = group_member_fixture()
      assert {:ok, %GroupMember{}} = Groups.delete_group_member(group_member)
      assert_raise Ecto.NoResultsError, fn -> Groups.get_group_member!(group_member.id) end
    end

    test "change_group_member/1 returns a group_member changeset" do
      group_member = group_member_fixture()
      assert %Ecto.Changeset{} = Groups.change_group_member(group_member)
    end
  end
end
