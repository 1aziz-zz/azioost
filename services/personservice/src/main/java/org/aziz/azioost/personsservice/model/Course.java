package org.aziz.azioost.personsservice.model;

import org.neo4j.ogm.annotation.GraphId;
import org.neo4j.ogm.annotation.NodeEntity;

/**
 * Created by aziz on 10/20/2017.
 */
@NodeEntity
public class Course {
    @GraphId
    private Long id_neo4j;

    private String Id;
    private String Title;

    public String getId() {
        return Id;
    }

    public void setId(String id) {
        Id = id;
    }

    public String getTitle() {
        return Title;
    }

    public void setTitle(String title) {
        Title = title;
    }

    @Override
    public boolean equals(Object o) {

        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Course course = (Course) o;

        if (getId() != null ? !getId().equals(course.getId()) : course.getId() != null) return false;
        return getTitle() != null ? getTitle().equals(course.getTitle()) : course.getTitle() == null;
    }

    @Override
    public int hashCode() {
        int result = getId() != null ? getId().hashCode() : 0;
        result = 31 * result + (getTitle() != null ? getTitle().hashCode() : 0);
        return result;
    }
}
