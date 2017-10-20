package org.aziz.azioost.personsservice.model.repositories;

import org.aziz.azioost.personsservice.model.Course;
import org.springframework.data.neo4j.repository.GraphRepository;

/**
 * Created by aziz on 10/20/2017.
 */
public interface CourseRepository extends GraphRepository<Course> {
    Course findById(String id);
}
