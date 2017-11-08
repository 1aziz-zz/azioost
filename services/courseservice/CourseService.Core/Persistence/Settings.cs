using System.IO;
using Microsoft.Extensions.Configuration;

namespace CourseService.Core.Persistence
{
    public class Settings
    {
        private static readonly IConfigurationRoot Configuration = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile(@"/run/secrets/mongo_secrets.json", true)
            .Build();

        public string ConnectionString = Configuration.GetSection("MongoConnection:ConnectionString").Value;
        public string Database = Configuration.GetSection("MongoConnection:Database").Value;
    }
}