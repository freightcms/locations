namespace FreightCMS.Locations.Models;

public class CountryModel
{
    public required string Id { get; set; }
    public required string Name { get; set; }
    public required string Code { get; set; }
    public string? Description { get; set; }
    public string? Currency { get; set; }
    public string? Language { get; set; }
    public string? Region { get; set; }
    public string? SubRegion { get; set; }
    public string? Capital { get; set; }
    public string? CallingCode { get; set; }
    public string? Flag { get; set; }
    public string? FlagEmoji { get; set; }
    public string? FlagEmojiUnicode { get; set; }
}