defmodule GiftopotamusWeb.Router do
  use GiftopotamusWeb, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :put_root_layout, {GiftopotamusWeb.LayoutView, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :login_layout do
    plug :put_layout, {GiftopotamusWeb.LayoutView, "login.html"}
  end

  scope "/", GiftopotamusWeb do
    pipe_through [:browser, :require_auth]

    get "/", PageController, :index
    resources "/groups", GroupController
    resources "/exchanges", ExchangeController
  end

  scope "/", GiftopotamusWeb do
    pipe_through [:browser, :login_layout]

    resources "/users", UserController, only: [:new, :create]
    resources "/sessions", SessionController, only: [:new, :create, :delete], singleton: true
  end

  defp require_auth(conn, _opts) do
    case get_session(conn, :user_id) do
      nil ->
        conn
        |> Phoenix.Controller.redirect(to: "/sessions/new")
        |> halt()

      user_id ->
        assign(conn, :current_user, Giftopotamus.Accounts.get_user!(user_id))
    end
  end

  # Enables LiveDashboard only for development
  #
  # If you want to use the LiveDashboard in production, you should put
  # it behind authentication and allow only admins to access it.
  # If your application does not have an admins-only section yet,
  # you can use Plug.BasicAuth to set up some basic authentication
  # as long as you are also using SSL (which you should anyway).
  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser
      live_dashboard "/dashboard", metrics: GiftopotamusWeb.Telemetry
    end
  end
end
