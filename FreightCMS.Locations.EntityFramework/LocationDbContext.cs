using Microsoft.EntityFrameworkCore;
using FreightCMS.Locations.Models;
using MongoDB.EntityFrameworkCore.Extensions;

namespace FreightCMS.Locations.EntityFramework;

/// <summary>
/// Creates a new entity framework database context targeting mongodb.
/// </summary>
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

        modelBuilder.Entity<AddressModel>(build => {
            build.Property(entity => entity.Id).ValueGeneratedOnAdd();
            build.Property(entity => entity.Line1).IsRequired();
            build.Property(entity => entity.Local).IsRequired();
            build.Property(entity => entity.PostalCode).IsRequired();
            build.Property(entity => entity.Country).IsRequired();
            build.Property(entity => entity.StateOrProvince).IsRequired();
            build.Property(entity => entity.Region).IsRequired();
            build.Property(entity => entity.Type).IsRequired();
        });
            

        modelBuilder.Entity<CountryModel>().ToCollection("Countries");

        modelBuilder.Entity<CountryModel>(build =>
        {
            build.Property(entity => entity.Id).ValueGeneratedOnAdd();
            build.Property(entity => entity.Name).IsRequired();
            build.Property(entity => entity.Code).IsRequired();
            build.Property(entity => entity.Code).HasMaxLength(2);
            build.Property(entity => entity.Currency).HasMaxLength(3);
            build.Property(entity => entity.Language).HasMaxLength(5);
        });
    }
}