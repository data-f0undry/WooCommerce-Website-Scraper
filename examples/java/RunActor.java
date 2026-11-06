// Java 11+ - REST API, standard library only
import java.net.URI;
import java.net.http.*;
import java.nio.file.*;

public class RunActor {
    public static void main(String[] args) throws Exception {
        String token = System.getenv("APIFY_TOKEN");
        String actor = "data_foundry~woocommerce-scraper";
        String input = Files.readString(Path.of("input.json"));

        HttpClient http = HttpClient.newHttpClient();
        HttpRequest post = HttpRequest.newBuilder()
            .uri(URI.create("https://api.apify.com/v2/acts/" + actor + "/runs?token=" + token))
            .header("Content-Type", "application/json")
            .POST(HttpRequest.BodyPublishers.ofString(input))
            .build();
        HttpResponse<String> run = http.send(post, HttpResponse.BodyHandlers.ofString());
        System.out.println(run.body());
    }
}
