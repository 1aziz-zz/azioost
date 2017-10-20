using System.Collections.Generic;
using System.Threading.Tasks;
using CatalogService.Core.Models;
using CatalogService.Core.Persistence;
using Microsoft.AspNetCore.Mvc;

namespace CatalogService.Web.Controllers
{
    [Produces("application/json")]
    [Route("api/Course")]
    public class CourseController : Controller
    {
        private readonly ICourseRepository _courseRepository;

        public CourseController(ICourseRepository courseRepository)
        {
            _courseRepository = courseRepository;
        }

        // GET: api/Course
        [HttpGet]
        public Task<IEnumerable<Course>> Get()
        {
            return GetCourseInternal();
        }

        private async Task<IEnumerable<Course>> GetCourseInternal()
        {
            return await _courseRepository.GetAllCourses();
        }

        // GET: api/Course/5
        [HttpGet("{id}", Name = "Get")]
        public Task<Course> Get(string id)
        {
            return GetCourseByInternal(id);
        }

        private async Task<Course> GetCourseByInternal(string id)
        {
            return await _courseRepository.GetCourse(id) ?? new Course();
        }

        // POST: api/Course
        [HttpPost]
        public void Post([FromBody] Course course)
        {
            _courseRepository.AddCourse(course);
        }

        // PUT: api/Course/5
        [HttpPut("{id}")]
        public void Put(string id, [FromBody] string title)
        {
            _courseRepository.UpdateCourseBody(id, title);
        }

        // DELETE: api/ApiWithActions/5
        [HttpDelete("{id}")]
        public void Delete(string id)
        {
            _courseRepository.RemoveCourse(id);
        }
    }
}