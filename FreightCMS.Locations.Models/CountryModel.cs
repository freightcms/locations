namespace FreightCMS.Locations.Models;

public class CountryModel
{
    public required string Id { get; set; }
    public required string Name { get; set; }
    /// <summary>
    /// ISO 3166-1 alpha-2 country code
    /// </summary>
    public required string Code { get; set; }
    public string? Description { get; set; }
    /// <summary>
    /// ISO 4217 currency code
    /// </summary>
    public string? Currency { get; set; }
    /// <summary>
    /// ISO 639-1 language code
    /// </summary>
    public string? Language { get; set; }
    public string? Region { get; set; }
    public string? SubRegion { get; set; }
    public string? Capital { get; set; }
    public string? CallingCode { get; set; }
    public string? Flag { get; set; }
    public string? FlagEmoji { get; set; }
    public string? FlagEmojiUnicode { get; set; }
}