namespace FreightCMS.Locations.Models;

public class Coordinates
{
    public double Latitude { get; set; }
    public double Longitude { get; set; }
}

public class AddressModel
{
    public required string Id { get; set; }
    public required string Line1 { get; set; }
    public string? Line2 { get; set; }
    public string? Line3 { get; set; }
    public required string Local { get; set; }
    public required string Region { get; set; }
    public required string PostalCode { get; set; }
    public required string StateOrProvince { get; set; }
    public required string Country { get; set; }
    public string? Description { get; set; }
    public string? Attention { get; set; }
    /// <summary>
    /// The type of address. For example, Distribution Center, Store, Convention, School, Office, Port, etc.
    /// </summary>
    public required string Type { get; set; }
    public Coordinates? Coordinates { get; set; }
}