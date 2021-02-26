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
# each group has 5-50 members                               #
# each group has 1-15 exchanges                             #
# each exchange has at least 90% of the group participating #
# each exchange has 1 gift per participant                  #
#                                                           #
#############################################################

defmodule Giftopotamus.DatabaseSeeder do
  import Ecto.Query

  alias Giftopotamus.Repo
  alias Giftopotamus.Accounts.User
  alias Giftopotamus.Groups.{Group, GroupMember}
  alias Giftopotamus.Exchanges.{Exchange, Participant, Gift}

  @num_users 100
  @num_groups 10
  @min_members 5
  @max_members 50
  @max_exchanges 15

  def reset do
    delete_all_records()
    restart_all_sequences()
  end

  def delete_all_records do
    Repo.delete_all(Gift)
    Repo.delete_all(Participant)
    Repo.delete_all(Exchange)
    Repo.delete_all(GroupMember)

    Repo.delete_all(Group)
    Repo.delete_all(User)
  end

  def restart_all_sequences do
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE users_id_seq RESTART WITH 1", [])
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE groups_id_seq RESTART WITH 1", [])
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE group_members_id_seq RESTART WITH 1", [])
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE exchanges_id_seq RESTART WITH 1", [])
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE participants_id_seq RESTART WITH 1", [])
    Ecto.Adapters.SQL.query(Repo, "ALTER SEQUENCE gifts_id_seq RESTART WITH 1", [])
  end

  def seed do
    reset()
    seed_users()
    seed_groups(Repo.all(User))
    seed_exchanges(Repo.all(Group))
  end

  def seed_users do
    Enum.each(1..@num_users, fn _ ->
      %User{}
      |> User.changeset(%{name: Faker.Person.first_name()})
      # TODO: handle unique constraint violation
      |> Repo.insert()
    end)
  end

  def seed_groups(users) do
    Enum.each(1..@num_groups, fn group_num ->
      [first_user | other_users] =
        Enum.take_random(users, Enum.random(@min_members..@max_members))

      # Make the first user and admin
      admin = %GroupMember{admin: true, user_id: first_user.id}

      members =
        [admin] ++
          Enum.map(other_users, fn u ->
            # 10% chance of admin
            admin? = :rand.uniform(100) < 10
            %GroupMember{admin: admin?, user_id: u.id}
          end)

      %Group{name: "#{group_num}", members: members}
      # TODO: Handle unique constraint violation
      |> Repo.insert()
    end)

    # Faker.StarWars.planet/0
    [
      "Alderaan",
      "Bespin",
      "Coruscant",
      "Dagobah",
      "Dantooine",
      "Endor",
      "Hoth",
      "Mustafar",
      "Naboo",
      "Yavin"
    ]
    |> Enum.with_index()
    |> Enum.each(fn {name, id} ->
      Repo.update_all(from(g in Group, where: g.id == ^id+1), set: [name: name])
    end)
  end

  def seed_exchanges(groups) do
    Enum.each(groups, fn group ->
      num_exchanges = Enum.random(1..@max_exchanges)

      Enum.each(1..num_exchanges, fn e ->
        group_members =
          GroupMember
          |> where([m], m.group_id == ^group.id)
          |> Repo.all()

        participants =
          Enum.map(group_members, fn m ->
            %Participant{
              participating: :rand.uniform(100) < 90,
              group_member: m
            }
          end)

        len_participants = length(participants)

        # TODO: Do not rely on predictable IDs
        gifts =
          Enum.map(group_members, fn _ ->
            %Gift{
              description:
                "#{Faker.Commerce.product_name_adjective()} #{Faker.Commerce.product_name()}",
              to_participant: %Participant{
                exchange_id: e,
                group_member_id: Enum.random(1..len_participants)
              },
              from_participant: %Participant{
                exchange_id: e,
                group_member_id: Enum.random(1..len_participants)
              }
            }
          end)

        %Exchange{
          name: to_string(Enum.random(2010..2022)),
          group: group,
          gifts: gifts,
          participants: participants
        }
        |> Repo.insert!()
      end)
    end)
  end
end

# Reset and seed database
Giftopotamus.DatabaseSeeder.seed()
