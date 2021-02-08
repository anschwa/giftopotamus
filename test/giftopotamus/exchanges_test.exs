defmodule Giftopotamus.ExchangesTest do
  use Giftopotamus.DataCase

  alias Giftopotamus.Exchanges

  describe "exchanges" do
    alias Giftopotamus.Exchanges.Exchange

    @valid_attrs %{name: "some name"}
    @update_attrs %{name: "some updated name"}
    @invalid_attrs %{name: nil}

    def exchange_fixture(attrs \\ %{}) do
      {:ok, exchange} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Exchanges.create_exchange()

      exchange
    end

    test "list_exchanges/0 returns all exchanges" do
      exchange = exchange_fixture()
      assert Exchanges.list_exchanges() == [exchange]
    end

    test "get_exchange!/1 returns the exchange with given id" do
      exchange = exchange_fixture()
      assert Exchanges.get_exchange!(exchange.id) == exchange
    end

    test "create_exchange/1 with valid data creates a exchange" do
      assert {:ok, %Exchange{} = exchange} = Exchanges.create_exchange(@valid_attrs)
      assert exchange.name == "some name"
    end

    test "create_exchange/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Exchanges.create_exchange(@invalid_attrs)
    end

    test "update_exchange/2 with valid data updates the exchange" do
      exchange = exchange_fixture()
      assert {:ok, %Exchange{} = exchange} = Exchanges.update_exchange(exchange, @update_attrs)
      assert exchange.name == "some updated name"
    end

    test "update_exchange/2 with invalid data returns error changeset" do
      exchange = exchange_fixture()
      assert {:error, %Ecto.Changeset{}} = Exchanges.update_exchange(exchange, @invalid_attrs)
      assert exchange == Exchanges.get_exchange!(exchange.id)
    end

    test "delete_exchange/1 deletes the exchange" do
      exchange = exchange_fixture()
      assert {:ok, %Exchange{}} = Exchanges.delete_exchange(exchange)
      assert_raise Ecto.NoResultsError, fn -> Exchanges.get_exchange!(exchange.id) end
    end

    test "change_exchange/1 returns a exchange changeset" do
      exchange = exchange_fixture()
      assert %Ecto.Changeset{} = Exchanges.change_exchange(exchange)
    end
  end

  describe "participants" do
    alias Giftopotamus.Exchanges.Participant

    @valid_attrs %{participating: true}
    @update_attrs %{participating: false}
    @invalid_attrs %{participating: nil}

    def participant_fixture(attrs \\ %{}) do
      {:ok, participant} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Exchanges.create_participant()

      participant
    end

    test "list_participants/0 returns all participants" do
      participant = participant_fixture()
      assert Exchanges.list_participants() == [participant]
    end

    test "get_participant!/1 returns the participant with given id" do
      participant = participant_fixture()
      assert Exchanges.get_participant!(participant.id) == participant
    end

    test "create_participant/1 with valid data creates a participant" do
      assert {:ok, %Participant{} = participant} = Exchanges.create_participant(@valid_attrs)
      assert participant.participating == true
    end

    test "create_participant/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Exchanges.create_participant(@invalid_attrs)
    end

    test "update_participant/2 with valid data updates the participant" do
      participant = participant_fixture()

      assert {:ok, %Participant{} = participant} =
               Exchanges.update_participant(participant, @update_attrs)

      assert participant.participating == false
    end

    test "update_participant/2 with invalid data returns error changeset" do
      participant = participant_fixture()

      assert {:error, %Ecto.Changeset{}} =
               Exchanges.update_participant(participant, @invalid_attrs)

      assert participant == Exchanges.get_participant!(participant.id)
    end

    test "delete_participant/1 deletes the participant" do
      participant = participant_fixture()
      assert {:ok, %Participant{}} = Exchanges.delete_participant(participant)
      assert_raise Ecto.NoResultsError, fn -> Exchanges.get_participant!(participant.id) end
    end

    test "change_participant/1 returns a participant changeset" do
      participant = participant_fixture()
      assert %Ecto.Changeset{} = Exchanges.change_participant(participant)
    end
  end

  describe "gifts" do
    alias Giftopotamus.Exchanges.Gift

    @valid_attrs %{description: "some description"}
    @update_attrs %{description: "some updated description"}
    @invalid_attrs %{description: nil}

    def gift_fixture(attrs \\ %{}) do
      {:ok, gift} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Exchanges.create_gift()

      gift
    end

    test "list_gifts/0 returns all gifts" do
      gift = gift_fixture()
      assert Exchanges.list_gifts() == [gift]
    end

    test "get_gift!/1 returns the gift with given id" do
      gift = gift_fixture()
      assert Exchanges.get_gift!(gift.id) == gift
    end

    test "create_gift/1 with valid data creates a gift" do
      assert {:ok, %Gift{} = gift} = Exchanges.create_gift(@valid_attrs)
      assert gift.description == "some description"
    end

    test "create_gift/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Exchanges.create_gift(@invalid_attrs)
    end

    test "update_gift/2 with valid data updates the gift" do
      gift = gift_fixture()
      assert {:ok, %Gift{} = gift} = Exchanges.update_gift(gift, @update_attrs)
      assert gift.description == "some updated description"
    end

    test "update_gift/2 with invalid data returns error changeset" do
      gift = gift_fixture()
      assert {:error, %Ecto.Changeset{}} = Exchanges.update_gift(gift, @invalid_attrs)
      assert gift == Exchanges.get_gift!(gift.id)
    end

    test "delete_gift/1 deletes the gift" do
      gift = gift_fixture()
      assert {:ok, %Gift{}} = Exchanges.delete_gift(gift)
      assert_raise Ecto.NoResultsError, fn -> Exchanges.get_gift!(gift.id) end
    end

    test "change_gift/1 returns a gift changeset" do
      gift = gift_fixture()
      assert %Ecto.Changeset{} = Exchanges.change_gift(gift)
    end
  end

  describe "exchange_results" do
    alias Giftopotamus.Exchanges.ExchangeResult

    @valid_attrs %{}
    @update_attrs %{}
    @invalid_attrs %{}

    def exchange_result_fixture(attrs \\ %{}) do
      {:ok, exchange_result} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Exchanges.create_exchange_result()

      exchange_result
    end

    test "list_exchange_results/0 returns all exchange_results" do
      exchange_result = exchange_result_fixture()
      assert Exchanges.list_exchange_results() == [exchange_result]
    end

    test "get_exchange_result!/1 returns the exchange_result with given id" do
      exchange_result = exchange_result_fixture()
      assert Exchanges.get_exchange_result!(exchange_result.id) == exchange_result
    end

    test "create_exchange_result/1 with valid data creates a exchange_result" do
      assert {:ok, %ExchangeResult{} = exchange_result} =
               Exchanges.create_exchange_result(@valid_attrs)
    end

    test "create_exchange_result/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Exchanges.create_exchange_result(@invalid_attrs)
    end

    test "update_exchange_result/2 with valid data updates the exchange_result" do
      exchange_result = exchange_result_fixture()

      assert {:ok, %ExchangeResult{} = exchange_result} =
               Exchanges.update_exchange_result(exchange_result, @update_attrs)
    end

    test "update_exchange_result/2 with invalid data returns error changeset" do
      exchange_result = exchange_result_fixture()

      assert {:error, %Ecto.Changeset{}} =
               Exchanges.update_exchange_result(exchange_result, @invalid_attrs)

      assert exchange_result == Exchanges.get_exchange_result!(exchange_result.id)
    end

    test "delete_exchange_result/1 deletes the exchange_result" do
      exchange_result = exchange_result_fixture()
      assert {:ok, %ExchangeResult{}} = Exchanges.delete_exchange_result(exchange_result)

      assert_raise Ecto.NoResultsError, fn ->
        Exchanges.get_exchange_result!(exchange_result.id)
      end
    end

    test "change_exchange_result/1 returns a exchange_result changeset" do
      exchange_result = exchange_result_fixture()
      assert %Ecto.Changeset{} = Exchanges.change_exchange_result(exchange_result)
    end
  end

  describe "participant_exclusions" do
    alias Giftopotamus.Exchanges.ParticipantExclusion

    @valid_attrs %{}
    @update_attrs %{}
    @invalid_attrs %{}

    def participant_exclusion_fixture(attrs \\ %{}) do
      {:ok, participant_exclusion} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Exchanges.create_participant_exclusion()

      participant_exclusion
    end

    test "list_participant_exclusions/0 returns all participant_exclusions" do
      participant_exclusion = participant_exclusion_fixture()
      assert Exchanges.list_participant_exclusions() == [participant_exclusion]
    end

    test "get_participant_exclusion!/1 returns the participant_exclusion with given id" do
      participant_exclusion = participant_exclusion_fixture()

      assert Exchanges.get_participant_exclusion!(participant_exclusion.id) ==
               participant_exclusion
    end

    test "create_participant_exclusion/1 with valid data creates a participant_exclusion" do
      assert {:ok, %ParticipantExclusion{} = participant_exclusion} =
               Exchanges.create_participant_exclusion(@valid_attrs)
    end

    test "create_participant_exclusion/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Exchanges.create_participant_exclusion(@invalid_attrs)
    end

    test "update_participant_exclusion/2 with valid data updates the participant_exclusion" do
      participant_exclusion = participant_exclusion_fixture()

      assert {:ok, %ParticipantExclusion{} = participant_exclusion} =
               Exchanges.update_participant_exclusion(participant_exclusion, @update_attrs)
    end

    test "update_participant_exclusion/2 with invalid data returns error changeset" do
      participant_exclusion = participant_exclusion_fixture()

      assert {:error, %Ecto.Changeset{}} =
               Exchanges.update_participant_exclusion(participant_exclusion, @invalid_attrs)

      assert participant_exclusion ==
               Exchanges.get_participant_exclusion!(participant_exclusion.id)
    end

    test "delete_participant_exclusion/1 deletes the participant_exclusion" do
      participant_exclusion = participant_exclusion_fixture()

      assert {:ok, %ParticipantExclusion{}} =
               Exchanges.delete_participant_exclusion(participant_exclusion)

      assert_raise Ecto.NoResultsError, fn ->
        Exchanges.get_participant_exclusion!(participant_exclusion.id)
      end
    end

    test "change_participant_exclusion/1 returns a participant_exclusion changeset" do
      participant_exclusion = participant_exclusion_fixture()
      assert %Ecto.Changeset{} = Exchanges.change_participant_exclusion(participant_exclusion)
    end
  end
end
