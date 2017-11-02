using System.Collections.Generic;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace CatalogService.Core.Models
{
    public class Course
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]

        public long Id { get; set; }

        public string Title { get; set; }
        public string Body { get; set; }
        public List<string> Tags { get; set; }
        public string Category { get; set; }
    }
}