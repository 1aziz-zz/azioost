package org.aziz.azioost.personsservice.service;

import org.aziz.azioost.personsservice.model.Course;
import org.aziz.azioost.personsservice.model.repositories.CourseRepository;
import org.aziz.azioost.personsservice.service.generic.Repository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

/**
 * Created by aziz on 10/20/2017.
 */
@Service
public class CourseService extends Repository<Course, CourseRepository> {
    @Autowired
    private RestTemplate restTemplate;

    public Course[] getAllCourses() {
        return restTemplate.getForObject("http://course-service/api/Course", Course[].class);

    }

}
