# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     Giftopotamus.Repo.insert!(%Giftopotamus.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.

#############################################################
#                          WANT                             #
#############################################################
#                                                           #
# 100 users                                                 #
# 10 groups                                                 #
# each group has 5-100 members                              #
# each group has 1-15 exchanges                             #
# each exchange has at least 90% of the group participating #
# each exchange has 1 gift per participant                  #
#                                                           #
#############################################################

defmodule Giftopotamus.DatabaseSeeder do
  alias Giftopotamus.{Repo, Accounts, Groups, Exchanges}
  alias Giftopotamus.Accounts.User
  alias Giftopotamus.Groups.{Group, GroupMember}
  alias Giftopotamus.Exchanges.{Exchange, Participant, Gift}

  @num_users 10
  @num_groups 2
  @min_members 5
  @max_members 10
  @num_exchanges 3

  def reset do
    Repo.delete_all Gift
    Repo.delete_all Participant
    Repo.delete_all Exchange
    Repo.delete_all GroupMember

    Repo.delete_all Group
    Repo.delete_all User
  end

  def seed do
    reset()
    seed_users()
    seed_groups(Repo.all(User))
  end

  def seed_users do
    Enum.each(1..@num_users, fn _ ->
      %User{}
      |> User.changeset(%{name: Faker.Person.first_name})
      |> Repo.insert() # TODO: handle unique constraint violation
    end)
  end

  def seed_groups(users) do
    [first_user | other_users] = Enum.take_random(users, Enum.random(@min_members..@max_members))

    # Make the first user and admin
    admin = %GroupMember{admin: true, user_id: first_user.id}
    members = [admin] ++ Enum.map(other_users, fn u ->
      admin? = :rand.uniform(100) < 10 # 10% chance of admin
      %GroupMember{admin: admin?, user_id: u.id}
    end)

    Enum.each(1..@num_groups, fn _ ->
      %Group{name: "#{Faker.Person.last_name}", members: members}
      |> Repo.insert() # TODO: handle unique constraint violation
    end)
  end

  def seed_exchanges(groups) do
    participating? = :rand.uniform(100) < 80 # 90% participating

    users = []
    participants = []
    gifts = []

    Enum.each(1..@num_exchanges, fn _ ->

    end)
  end
end

# Reset and seed database
Giftopotamus.DatabaseSeeder.seed
