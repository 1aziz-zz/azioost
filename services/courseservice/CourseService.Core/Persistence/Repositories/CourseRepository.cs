using System.Collections.Generic;
using System.Threading.Tasks;
using CatalogService.Core.Models;
using MongoDB.Bson;
using MongoDB.Driver;

namespace CatalogService.Core.Persistence.Repositories
{
    public class CourseRepository : ICourseRepository
    {
        private readonly DbContext _context;


        public CourseRepository()
        {
            _context = new DbContext();
        }

        public async Task<IEnumerable<Course>> GetAllCourses()
        {
            return await _context.Courses.Find(_ => true).ToListAsync();
        }

        public async Task<Course> GetCourse(string id)
        {
            var filter = Builders<Course>.Filter.Eq("Id", id);

            return await _context.Courses
                .Find(filter)
                .FirstOrDefaultAsync();
        }

        public async Task AddCourse(Course item)
        {
            await _context.Courses.InsertOneAsync(item);
        }

        public async Task<DeleteResult> RemoveCourse(string id)
        {
            return await _context.Courses.DeleteOneAsync(
                Builders<Course>.Filter.Eq("Id", id));
        }

        public async Task<UpdateResult> UpdateCourse(string id, string body)
        {
            var filter = Builders<Course>.Filter.Eq(s => s.Id.ToString(), id);
            var update = Builders<Course>.Update
                .Set(s => s.Body, body);
            return await _context.Courses.UpdateOneAsync(filter, update);
        }

        // Demo function - full document update
        public async Task<ReplaceOneResult> UpdateCourseBody(string id, string body)
        {
            var item = new Course {Body = body};

            return await UpdateCourse(id, item);
        }

        public async Task<DeleteResult> RemoveAllCourses()
        {
            return await _context.Courses.DeleteManyAsync(new BsonDocument());
        }

        public async Task<ReplaceOneResult> UpdateCourse(string id, Course item)
        {
            return await _context.Courses
                .ReplaceOneAsync(n => n.Id.Equals(id)
                    , item
                    , new UpdateOptions {IsUpsert = true});
        }
    }
}