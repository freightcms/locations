using Microsoft.EntityFrameworkCore;
using FreightCMS.Locations.Models;
using MongoDB.EntityFrameworkCore.Extensions;

namespace FreightCMS.Locations.EntityFramework;

public class LocationDbContext : DbContext
{
    public LocationDbContext(DbContextOptions<LocationDbContext> options) : base(options)
    {
    }

    public DbSet<AddressModel> Addresses { get; set; }
    public DbSet<CountryModel> Countries { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<AddressModel>().ToCollection("Addresses");
        modelBuilder.Entity<CountryModel>().ToCollection("Countries");
    }
}