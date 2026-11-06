// .NET - REST API, System.Net.Http only
using System;
using System.IO;
using System.Net.Http;
using System.Text;

class RunActor
{
    static async System.Threading.Tasks.Task Main()
    {
        string token = Environment.GetEnvironmentVariable("APIFY_TOKEN");
        string input = File.ReadAllText("input.json");

        var http = new HttpClient();
        var content = new StringContent(input, Encoding.UTF8, "application/json");
        var response = await http.PostAsync(
            $"https://api.apify.com/v2/acts/data_foundry~woocommerce-scraper/runs?token={token}", content);
        Console.WriteLine(await response.Content.ReadAsStringAsync());
    }
}
