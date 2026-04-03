// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cjavdev/believe-go"
  "github.com/cjavdev/believe-go/internal/testutil"
  "github.com/cjavdev/believe-go/option"
)

func TestTicketSaleNewWithOptionalParams(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := believe.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  _, err := client.TicketSales.New(context.TODO(), believe.TicketSaleNewParams{
    BuyerName: "Mae Green",
    Currency: "GBP",
    Discount: "9.00",
    MatchID: "match-001",
    PurchaseMethod: believe.PurchaseMethodOnline,
    Quantity: 2,
    Subtotal: "90.00",
    Tax: "16.20",
    Total: "97.20",
    UnitPrice: "45.00",
    BuyerEmail: believe.String("mae.green@example.com"),
    CouponCode: believe.String("BELIEVE10"),
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestTicketSaleGet(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := believe.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  _, err := client.TicketSales.Get(context.TODO(), "ticket_sale_id")
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestTicketSaleUpdateWithOptionalParams(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := believe.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  _, err := client.TicketSales.Update(
    context.TODO(),
    "ticket_sale_id",
    believe.TicketSaleUpdateParams{
      BuyerEmail: believe.String("dev@stainless.com"),
      BuyerName: believe.String("buyer_name"),
      CouponCode: believe.String("coupon_code"),
      Currency: believe.String("currency"),
      Discount: believe.String("discount"),
      MatchID: believe.String("match_id"),
      PurchaseMethod: believe.PurchaseMethodOnline,
      Quantity: believe.Int(1),
      Subtotal: believe.String("subtotal"),
      Tax: believe.String("tax"),
      Total: believe.String("total"),
      UnitPrice: believe.String("unit_price"),
    },
  )
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestTicketSaleListWithOptionalParams(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := believe.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  _, err := client.TicketSales.List(context.TODO(), believe.TicketSaleListParams{
    CouponCode: believe.String("coupon_code"),
    Currency: believe.String("currency"),
    Limit: believe.Int(10),
    MatchID: believe.String("match_id"),
    PurchaseMethod: believe.PurchaseMethodOnline,
    Skip: believe.Int(0),
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestTicketSaleDelete(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := believe.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  err := client.TicketSales.Delete(context.TODO(), "ticket_sale_id")
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
