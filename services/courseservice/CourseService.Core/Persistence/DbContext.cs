using System;
using CatalogService.Core.Models;
using CourseService.Core.Persistence;
using MongoDB.Driver;

namespace CatalogService.Core.Persistence
{
    public class DbContext
    {
        private readonly IMongoDatabase _database;


        public DbContext()
        {
            var settings = new Settings();
            try
            {
                var client = new MongoClient(settings.ConnectionString);
                _database = client.GetDatabase(settings.Database);
            }
            catch (Exception ex)
            {
                throw new Exception("Can not access to db server.", ex);
            }
        }

        public IMongoCollection<Course> Courses => _database.GetCollection<Course>("Courses");
    }
}