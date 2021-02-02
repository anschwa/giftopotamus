defmodule Giftopotamus.Exchanges do
  @moduledoc """
  The Exchanges context.
  """

  import Ecto.Query, warn: false
  alias Giftopotamus.Repo

  alias Giftopotamus.Exchanges.Exchange

  @doc """
  Returns the list of exchanges.

  ## Examples

      iex> list_exchanges()
      [%Exchange{}, ...]

  """
  def list_exchanges do
    Repo.all(Exchange)
  end

  @doc """
  Gets a single exchange.

  Raises `Ecto.NoResultsError` if the Exchange does not exist.

  ## Examples

      iex> get_exchange!(123)
      %Exchange{}

      iex> get_exchange!(456)
      ** (Ecto.NoResultsError)

  """
  def get_exchange!(id), do: Repo.get!(Exchange, id)

  @doc """
  Creates a exchange.

  ## Examples

      iex> create_exchange(%{field: value})
      {:ok, %Exchange{}}

      iex> create_exchange(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_exchange(attrs \\ %{}) do
    %Exchange{}
    |> Exchange.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a exchange.

  ## Examples

      iex> update_exchange(exchange, %{field: new_value})
      {:ok, %Exchange{}}

      iex> update_exchange(exchange, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_exchange(%Exchange{} = exchange, attrs) do
    exchange
    |> Exchange.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a exchange.

  ## Examples

      iex> delete_exchange(exchange)
      {:ok, %Exchange{}}

      iex> delete_exchange(exchange)
      {:error, %Ecto.Changeset{}}

  """
  def delete_exchange(%Exchange{} = exchange) do
    Repo.delete(exchange)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking exchange changes.

  ## Examples

      iex> change_exchange(exchange)
      %Ecto.Changeset{data: %Exchange{}}

  """
  def change_exchange(%Exchange{} = exchange, attrs \\ %{}) do
    Exchange.changeset(exchange, attrs)
  end

  alias Giftopotamus.Exchanges.Participant

  @doc """
  Returns the list of participants.

  ## Examples

      iex> list_participants()
      [%Participant{}, ...]

  """
  def list_participants do
    Repo.all(Participant)
  end

  @doc """
  Gets a single participant.

  Raises `Ecto.NoResultsError` if the Participant does not exist.

  ## Examples

      iex> get_participant!(123)
      %Participant{}

      iex> get_participant!(456)
      ** (Ecto.NoResultsError)

  """
  def get_participant!(id), do: Repo.get!(Participant, id)

  @doc """
  Creates a participant.

  ## Examples

      iex> create_participant(%{field: value})
      {:ok, %Participant{}}

      iex> create_participant(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_participant(attrs \\ %{}) do
    %Participant{}
    |> Participant.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a participant.

  ## Examples

      iex> update_participant(participant, %{field: new_value})
      {:ok, %Participant{}}

      iex> update_participant(participant, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_participant(%Participant{} = participant, attrs) do
    participant
    |> Participant.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a participant.

  ## Examples

      iex> delete_participant(participant)
      {:ok, %Participant{}}

      iex> delete_participant(participant)
      {:error, %Ecto.Changeset{}}

  """
  def delete_participant(%Participant{} = participant) do
    Repo.delete(participant)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking participant changes.

  ## Examples

      iex> change_participant(participant)
      %Ecto.Changeset{data: %Participant{}}

  """
  def change_participant(%Participant{} = participant, attrs \\ %{}) do
    Participant.changeset(participant, attrs)
  end

  alias Giftopotamus.Exchanges.Gift

  @doc """
  Returns the list of gifts.

  ## Examples

      iex> list_gifts()
      [%Gift{}, ...]

  """
  def list_gifts do
    Repo.all(Gift)
  end

  @doc """
  Gets a single gift.

  Raises `Ecto.NoResultsError` if the Gift does not exist.

  ## Examples

      iex> get_gift!(123)
      %Gift{}

      iex> get_gift!(456)
      ** (Ecto.NoResultsError)

  """
  def get_gift!(id), do: Repo.get!(Gift, id)

  @doc """
  Creates a gift.

  ## Examples

      iex> create_gift(%{field: value})
      {:ok, %Gift{}}

      iex> create_gift(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_gift(attrs \\ %{}) do
    %Gift{}
    |> Gift.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a gift.

  ## Examples

      iex> update_gift(gift, %{field: new_value})
      {:ok, %Gift{}}

      iex> update_gift(gift, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_gift(%Gift{} = gift, attrs) do
    gift
    |> Gift.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a gift.

  ## Examples

      iex> delete_gift(gift)
      {:ok, %Gift{}}

      iex> delete_gift(gift)
      {:error, %Ecto.Changeset{}}

  """
  def delete_gift(%Gift{} = gift) do
    Repo.delete(gift)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking gift changes.

  ## Examples

      iex> change_gift(gift)
      %Ecto.Changeset{data: %Gift{}}

  """
  def change_gift(%Gift{} = gift, attrs \\ %{}) do
    Gift.changeset(gift, attrs)
  end

  alias Giftopotamus.Exchanges.ExchangeResult

  @doc """
  Returns the list of exchange_results.

  ## Examples

      iex> list_exchange_results()
      [%ExchangeResult{}, ...]

  """
  def list_exchange_results do
    Repo.all(ExchangeResult)
  end

  @doc """
  Gets a single exchange_result.

  Raises `Ecto.NoResultsError` if the Exchange result does not exist.

  ## Examples

      iex> get_exchange_result!(123)
      %ExchangeResult{}

      iex> get_exchange_result!(456)
      ** (Ecto.NoResultsError)

  """
  def get_exchange_result!(id), do: Repo.get!(ExchangeResult, id)

  @doc """
  Creates a exchange_result.

  ## Examples

      iex> create_exchange_result(%{field: value})
      {:ok, %ExchangeResult{}}

      iex> create_exchange_result(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_exchange_result(attrs \\ %{}) do
    %ExchangeResult{}
    |> ExchangeResult.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a exchange_result.

  ## Examples

      iex> update_exchange_result(exchange_result, %{field: new_value})
      {:ok, %ExchangeResult{}}

      iex> update_exchange_result(exchange_result, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_exchange_result(%ExchangeResult{} = exchange_result, attrs) do
    exchange_result
    |> ExchangeResult.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a exchange_result.

  ## Examples

      iex> delete_exchange_result(exchange_result)
      {:ok, %ExchangeResult{}}

      iex> delete_exchange_result(exchange_result)
      {:error, %Ecto.Changeset{}}

  """
  def delete_exchange_result(%ExchangeResult{} = exchange_result) do
    Repo.delete(exchange_result)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking exchange_result changes.

  ## Examples

      iex> change_exchange_result(exchange_result)
      %Ecto.Changeset{data: %ExchangeResult{}}

  """
  def change_exchange_result(%ExchangeResult{} = exchange_result, attrs \\ %{}) do
    ExchangeResult.changeset(exchange_result, attrs)
  end

  alias Giftopotamus.Exchanges.ParticipantExclusion

  @doc """
  Returns the list of participant_exclusions.

  ## Examples

      iex> list_participant_exclusions()
      [%ParticipantExclusion{}, ...]

  """
  def list_participant_exclusions do
    Repo.all(ParticipantExclusion)
  end

  @doc """
  Gets a single participant_exclusion.

  Raises `Ecto.NoResultsError` if the Participant exclusion does not exist.

  ## Examples

      iex> get_participant_exclusion!(123)
      %ParticipantExclusion{}

      iex> get_participant_exclusion!(456)
      ** (Ecto.NoResultsError)

  """
  def get_participant_exclusion!(id), do: Repo.get!(ParticipantExclusion, id)

  @doc """
  Creates a participant_exclusion.

  ## Examples

      iex> create_participant_exclusion(%{field: value})
      {:ok, %ParticipantExclusion{}}

      iex> create_participant_exclusion(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_participant_exclusion(attrs \\ %{}) do
    %ParticipantExclusion{}
    |> ParticipantExclusion.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a participant_exclusion.

  ## Examples

      iex> update_participant_exclusion(participant_exclusion, %{field: new_value})
      {:ok, %ParticipantExclusion{}}

      iex> update_participant_exclusion(participant_exclusion, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_participant_exclusion(%ParticipantExclusion{} = participant_exclusion, attrs) do
    participant_exclusion
    |> ParticipantExclusion.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a participant_exclusion.

  ## Examples

      iex> delete_participant_exclusion(participant_exclusion)
      {:ok, %ParticipantExclusion{}}

      iex> delete_participant_exclusion(participant_exclusion)
      {:error, %Ecto.Changeset{}}

  """
  def delete_participant_exclusion(%ParticipantExclusion{} = participant_exclusion) do
    Repo.delete(participant_exclusion)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking participant_exclusion changes.

  ## Examples

      iex> change_participant_exclusion(participant_exclusion)
      %Ecto.Changeset{data: %ParticipantExclusion{}}

  """
  def change_participant_exclusion(%ParticipantExclusion{} = participant_exclusion, attrs \\ %{}) do
    ParticipantExclusion.changeset(participant_exclusion, attrs)
  end
end
