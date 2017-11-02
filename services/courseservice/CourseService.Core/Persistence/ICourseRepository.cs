using System.Collections.Generic;
using System.Threading.Tasks;
using CatalogService.Core.Models;
using MongoDB.Driver;

namespace CatalogService.Core.Persistence
{
    public interface ICourseRepository
    {
        Task<IEnumerable<Course>> GetAllCourses();
        Task<Course> GetCourse(string id);
        Task AddCourse(Course item);
        Task<DeleteResult> RemoveCourse(string id);

        Task<UpdateResult> UpdateCourse(string id, string body);

        // demo interface - full document update
        Task<ReplaceOneResult> UpdateCourseBody(string id, string body);

        // should be used with high cautious, only in relation with demo setup
        Task<DeleteResult> RemoveAllCourses();
    }
}