defmodule GiftopotamusWeb.ExchangeResultControllerTest do
  use GiftopotamusWeb.ConnCase

  alias Giftopotamus.Exchanges

  @create_attrs %{}
  @update_attrs %{}
  @invalid_attrs %{}

  def fixture(:exchange_result) do
    {:ok, exchange_result} = Exchanges.create_exchange_result(@create_attrs)
    exchange_result
  end

  describe "index" do
    test "lists all exchange_results", %{conn: conn} do
      conn = get(conn, Routes.exchange_result_path(conn, :index))
      assert html_response(conn, 200) =~ "Listing Exchange results"
    end
  end

  describe "new exchange_result" do
    test "renders form", %{conn: conn} do
      conn = get(conn, Routes.exchange_result_path(conn, :new))
      assert html_response(conn, 200) =~ "New Exchange result"
    end
  end

  describe "create exchange_result" do
    test "redirects to show when data is valid", %{conn: conn} do
      conn = post(conn, Routes.exchange_result_path(conn, :create), exchange_result: @create_attrs)

      assert %{id: id} = redirected_params(conn)
      assert redirected_to(conn) == Routes.exchange_result_path(conn, :show, id)

      conn = get(conn, Routes.exchange_result_path(conn, :show, id))
      assert html_response(conn, 200) =~ "Show Exchange result"
    end

    test "renders errors when data is invalid", %{conn: conn} do
      conn = post(conn, Routes.exchange_result_path(conn, :create), exchange_result: @invalid_attrs)
      assert html_response(conn, 200) =~ "New Exchange result"
    end
  end

  describe "edit exchange_result" do
    setup [:create_exchange_result]

    test "renders form for editing chosen exchange_result", %{conn: conn, exchange_result: exchange_result} do
      conn = get(conn, Routes.exchange_result_path(conn, :edit, exchange_result))
      assert html_response(conn, 200) =~ "Edit Exchange result"
    end
  end

  describe "update exchange_result" do
    setup [:create_exchange_result]

    test "redirects when data is valid", %{conn: conn, exchange_result: exchange_result} do
      conn = put(conn, Routes.exchange_result_path(conn, :update, exchange_result), exchange_result: @update_attrs)
      assert redirected_to(conn) == Routes.exchange_result_path(conn, :show, exchange_result)

      conn = get(conn, Routes.exchange_result_path(conn, :show, exchange_result))
      assert html_response(conn, 200)
    end

    test "renders errors when data is invalid", %{conn: conn, exchange_result: exchange_result} do
      conn = put(conn, Routes.exchange_result_path(conn, :update, exchange_result), exchange_result: @invalid_attrs)
      assert html_response(conn, 200) =~ "Edit Exchange result"
    end
  end

  describe "delete exchange_result" do
    setup [:create_exchange_result]

    test "deletes chosen exchange_result", %{conn: conn, exchange_result: exchange_result} do
      conn = delete(conn, Routes.exchange_result_path(conn, :delete, exchange_result))
      assert redirected_to(conn) == Routes.exchange_result_path(conn, :index)
      assert_error_sent 404, fn ->
        get(conn, Routes.exchange_result_path(conn, :show, exchange_result))
      end
    end
  end

  defp create_exchange_result(_) do
    exchange_result = fixture(:exchange_result)
    %{exchange_result: exchange_result}
  end
end
