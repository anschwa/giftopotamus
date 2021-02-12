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
# each exchange has at least 80% of the group participating #
# each exchange has 1 gift per participant                  #
#                                                           #
#############################################################

# IDEAS

# Repo.insert! %Gift{
#   description: "#{Faker.Commerce.product_name_adjective} #{Faker.Commerce.product_name}"
#   exchange: %Exchange{}
#   to_participant: %Participant{}
#   from_participant: %Participant{}
# }


defmodule Giftopotamus.DatabaseSeeder do
  alias Giftopotamus.{Repo, Accounts, Groups, Exchanges}
  alias Giftopotamus.Accounts.User
  alias Giftopotamus.Groups.{Group, GroupMember}
  alias Giftopotamus.Exchanges.{Exchange, Participant, Gift}

  def reset do
    Repo.delete_all User
    Repo.delete_all Group
    Repo.delete_all GroupMember
    Repo.delete_all Exchange
    Repo.delete_all Participant
    Repo.delete_all Gift
  end

  def seed do
  end

  def seed_users do
    (1..100) |> Enum.map(fn _ ->
      %User{}
      |> User.changeset(%{name: Faker.Person.first_name})
      |> Repo.insert() # TODO: How to get 100 unique users?
    end)
  end

  def seed_groups do
    (1..10) |> Enum.map(fn _ ->
      %Group{}
      |> Group.changeset(%{name: Faker.Company.name})
      |> Repo.insert() # TODO: How to get 10 unique groups?
    end)
  end

end
