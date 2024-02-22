import React from "react";
import './CreateEvent.css';


const CreateEvent = () => {

    return (
        <div className="container">
            <form>
                <div class="form-group">
                    <h1>Create Event</h1>
                </div>
                <div class="form-group">
                    <label for="eventName">Event Name</label>
                    <input type="text" id="eventName" name="eventName"/>
                </div>

                <div class="form-group">
                    <label for="eventDate">Event Date</label>
                    <input type="date" id="eventDate" name="eventDate"/>
                </div>

                <div class="form-group">
                    <label for="eventTime">Event Time</label>
                    <input type="time" id="eventTime" name="eventTime"/>
                </div>

                <div class="form-group">
                    <label for="location">Location</label>
                    <input type="text" id="location" name="location"/>
                </div>

                <div class="form-group">
                    <label for="contact">Contact</label>
                    <input type="text" id="contact" name="contact"/>
                </div>

                <div class="form-group">
                    <label for="eventPoster">Event Poster</label>
                    <input type="file" id="eventPoster" name="eventPoster"/>
                </div>

                <div class="form-group">
                    <label for="description">Description</label>
                    <textarea id="description" name="description" rows="4"></textarea>
                </div>

                <div class="form-group">
                    
                    <label id="tag-label">Tags</label>
                    <div class="tags">
                        <div class="tag-group">
                            <input type="checkbox" id="physicalWellness" name="tags" value="Physical Wellness"/>
                            <label for="physicalWellness">Physical Wellness</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="culturalExchange" name="tags" value="Cultural Exchange"/>
                            <label for="culturalExchange">Cultural Exchange</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="lgbtq" name="tags" value="LGBTQ"/>
                            <label for="lgbtq">LGBTQ+</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="artsEntertainment" name="tags" value="Arts Entertainment"/>
                            <label for="artsEntertainment">Arts/Entertainment</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="graduate" name="tags" value="Graduate"/>
                            <label for="graduate">Graduate</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="undergraduate" name="tags" value="Undergraduate"/>
                            <label for="undergraduate">Undergraduate</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="virtual" name="tags" value="Virtual"/>
                            <label for="virtual">Virtual</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="inPerson" name="tags" value="In Person"/>
                            <label for="inPerson">In Person</label>
                        </div>

                        <div class="tag-group">
                            <input type="checkbox" id="freeFood" name="tags" value="Free Food"/>
                            <label for="freeFood">Free Food</label>
                        </div>

                    </div>
                </div>

                <button type="submit">Submit</button>
            </form>
        </div>
    );
}

export default CreateEvent;