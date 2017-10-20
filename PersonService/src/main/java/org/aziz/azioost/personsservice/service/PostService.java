package org.aziz.azioost.personsservice.service;

import org.aziz.azioost.personsservice.model.Post;
import org.aziz.azioost.personsservice.model.repositories.PostRepository;
import org.aziz.azioost.personsservice.service.generic.Repository;
import org.springframework.stereotype.Service;

/**
 * Created by aziz on 10/20/2017.
 */
@Service
public class PostService extends Repository<Post, PostRepository> {


}
