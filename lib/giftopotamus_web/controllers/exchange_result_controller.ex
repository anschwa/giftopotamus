defmodule GiftopotamusWeb.ExchangeResultController do
  use GiftopotamusWeb, :controller

  alias Giftopotamus.Exchanges
  alias Giftopotamus.Exchanges.ExchangeResult

  def index(conn, _params) do
    exchange_results = Exchanges.list_exchange_results()
    render(conn, "index.html", exchange_results: exchange_results)
  end

  def new(conn, _params) do
    changeset = Exchanges.change_exchange_result(%ExchangeResult{})
    render(conn, "new.html", changeset: changeset)
  end

  def create(conn, %{"exchange_result" => exchange_result_params}) do
    case Exchanges.create_exchange_result(exchange_result_params) do
      {:ok, exchange_result} ->
        conn
        |> put_flash(:info, "Exchange result created successfully.")
        |> redirect(to: Routes.exchange_result_path(conn, :show, exchange_result))

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, "new.html", changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    exchange_result = Exchanges.get_exchange_result!(id)
    render(conn, "show.html", exchange_result: exchange_result)
  end

  def edit(conn, %{"id" => id}) do
    exchange_result = Exchanges.get_exchange_result!(id)
    changeset = Exchanges.change_exchange_result(exchange_result)
    render(conn, "edit.html", exchange_result: exchange_result, changeset: changeset)
  end

  def update(conn, %{"id" => id, "exchange_result" => exchange_result_params}) do
    exchange_result = Exchanges.get_exchange_result!(id)

    case Exchanges.update_exchange_result(exchange_result, exchange_result_params) do
      {:ok, exchange_result} ->
        conn
        |> put_flash(:info, "Exchange result updated successfully.")
        |> redirect(to: Routes.exchange_result_path(conn, :show, exchange_result))

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, "edit.html", exchange_result: exchange_result, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    exchange_result = Exchanges.get_exchange_result!(id)
    {:ok, _exchange_result} = Exchanges.delete_exchange_result(exchange_result)

    conn
    |> put_flash(:info, "Exchange result deleted successfully.")
    |> redirect(to: Routes.exchange_result_path(conn, :index))
  end
end
