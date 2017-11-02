package org.aziz.azioost.personsservice.service;

import org.aziz.azioost.personsservice.model.Comment;
import org.aziz.azioost.personsservice.model.Course;
import org.aziz.azioost.personsservice.model.Person;
import org.aziz.azioost.personsservice.model.Post;
import org.aziz.azioost.personsservice.model.repositories.CourseRepository;
import org.aziz.azioost.personsservice.model.repositories.PersonRepository;
import org.aziz.azioost.personsservice.service.generic.Repository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * Created by aziz on 10/20/2017.
 */
@Service
public class PersonService extends Repository<Person, PersonRepository> {

    @Autowired
    private CourseRepository courseRepository;

    public void addFriendship(Person personA, Person personB) {
        if (!personA.getPeople().contains(personB) && !personA.equals(personB)) {
            personA.addFriend(personB);
        }
    }

    public void addCourseToPerson(Person person, String course_id) {
        Course course = courseRepository.findById(course_id);
        if (!person.getCourses().contains(course)) {
            person.addCourse(course);
        }
    }

    public void addComment(Person person, Comment comment) {
        if (!person.getComments().contains(comment)) {
            person.addComment(comment);
        }
    }

    public void addPost(Person person, Post post) {
        if (!person.getPosts().contains(post)) {
            person.addPost(post);
        }
    }

}
