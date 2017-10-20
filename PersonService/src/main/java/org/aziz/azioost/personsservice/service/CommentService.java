package org.aziz.azioost.personsservice.service;

import org.aziz.azioost.personsservice.model.Comment;
import org.aziz.azioost.personsservice.model.repositories.CommentRepository;
import org.aziz.azioost.personsservice.service.generic.Repository;
import org.springframework.stereotype.Service;

/**
 * Created by aziz on 10/20/2017.
 */
@Service
public class CommentService extends Repository<Comment, CommentRepository> {


}
