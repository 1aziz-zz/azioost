package org.aziz.azioost.personsservice.web;

import org.aziz.azioost.personsservice.model.Comment;
import org.aziz.azioost.personsservice.model.Person;
import org.aziz.azioost.personsservice.model.Post;
import org.aziz.azioost.personsservice.model.repositories.PersonRepository;
import org.aziz.azioost.personsservice.service.CommentService;
import org.aziz.azioost.personsservice.service.CourseService;
import org.aziz.azioost.personsservice.service.PersonService;
import org.aziz.azioost.personsservice.service.PostService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.PostConstruct;
import java.util.Arrays;
import java.util.stream.Collectors;

/**
 * Created by aziz on 10/19/2017.
 */
@RequestMapping(value = "/")
@RestController
public class MainController {
    @Autowired
    private PersonService personService;
    @Autowired
    private CommentService commentService;
    @Autowired
    private PostService postService;
    @Autowired
    private CourseService courseService;

    @RequestMapping(method = RequestMethod.GET)
    public Iterable<Person> results() {

        personService.deleteAll();
        postService.deleteAll();
        commentService.deleteAll();
        courseService.deleteAll();

        Arrays.stream(courseService.getAllCourses()).map(courseService::save).collect(Collectors.toList());

        Person personA = new Person("Aziz", "Aizo", 20);
        Person personB = new Person("Barack", "Trump", 67);

        Comment comment = new Comment("I enjoyed reading your awful post! Well done!");
        Post post = new Post("How to write a blog?", "This is a very very very very long text");
        post.getComments().add(comment);

        personService.addPost(personB, post);
        personService.addCourseToPerson(personA, "59e9e8a5c684410fb0d9849c");
        personService.addCourseToPerson(personA, "59e9eb5a9980e6213c190f0c");
        personService.addCourseToPerson(personB, "59e9e8a5c684410fb0d9849c");


        personService.addComment(personA, comment);
        personService.addFriendship(personA, personB);
        personService.save(personA);

        return personService.getAll();
    }
}
