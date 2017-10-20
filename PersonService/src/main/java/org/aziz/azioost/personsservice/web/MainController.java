package org.aziz.azioost.personsservice.web;

import org.aziz.azioost.personsservice.model.Comment;
import org.aziz.azioost.personsservice.model.Person;
import org.aziz.azioost.personsservice.model.Post;
import org.aziz.azioost.personsservice.model.repositories.PersonRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

/**
 * Created by aziz on 10/19/2017.
 */
@RequestMapping(value = "/")
@RestController
public class MainController {

    @Autowired
    private PersonRepository personRepository;

    @RequestMapping(method = RequestMethod.GET)
    public Iterable<Person> results() {
        personRepository.deleteAll();

        Person personA = new Person("Aziz", "Aizo", 20);
        Person personB = new Person("Barack", "Trump", 67);

        Comment comment = new Comment("I enjoyed reading your awful post! Well done!");
        Post post = new Post("How to write a blog?", "This is a very very very very long text");
        post.getComments().add(comment);

        personB.getPosts().add(post);


        personA.getComments().add(comment);


        personA.addNewPerson(personB);
        personRepository.save(personA);

        return personRepository.findAll();
    }
}
