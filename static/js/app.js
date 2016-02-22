(function(window, document, undefined) {
  "use strict";

  window.onload = function() {
    onload();
  };


  /**
   * Onload handler. Responsible for binding to form submit event.
   * @return {none}
   */
  function onload() {
    var form = document.getElementById("shorten-form");
    form.addEventListener("submit", handleFormSubmit);

    form.addEventListener("keypress", onFormChange);
  };

  /**
   * Helper to hide form link when the user changes the form after
   * having already submitted one link.
   *
   * @return {none}
   */
  function onFormChange() {
    document.getElementById('link-container').classList.remove("visible");
  };

  /**
   * Helper for serializing form data.
   * @TODO(jengler) 2016-02-22: Replace with JSON
   *
   * @param  {Object} data Object containing key/value pairs to serialize.
   * @return {string}      The urlencoded query string.
   */
  function serialize(data) {
    var serialized = [];
    for (var param in data) {
      serialized.push(encodeURIComponent(param) + '=' + encodeURIComponent(data[param]));
    }
    return serialized.join('&').replace(/%20/g, '+');
  };

  /**
   * Form submit handler. Makes the world all ajaxy.
   *
   * @param  {Event} e HTML event
   * @return {none}
   */
  function handleFormSubmit(e) {
    e.preventDefault();
    var code = document.getElementById("code").value.trim();
    var url = document.getElementById("url").value.trim();

    var serialized = serialize({
      code: code,
      url: url
    });

    createShort(serialized);
  };

  /**
   * Create a short URL for the provided form data.
   *
   * @param  {FormData} formData Must at least have 'url'. 'code' is optional
   * @return {none}
   */
  function createShort(formData) {
    var xhr = new XMLHttpRequest();

    xhr.addEventListener("load", handleCreateSuccess);
    xhr.addEventListener("error", handleCreateError);

    xhr.open("POST", "/");
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

    xhr.send(formData);
  }

  /**
   * Handle XHR success event
   * @param  {Event} event The XHR event
   * @return {none}
   */
  function handleCreateSuccess(event) {
    if (event.target.status === 200) {
      console.log('Success');
      showSuccessLink(event.target.response.trim());
    } else {
      handleCreateError(event);
    }
  }

  /**
   * Handle XHR error event
   * @param  {Event} event The XHR event
   * @return {none}
   */
  function handleCreateError(event) {
    console.log("error", event.target.status, event.target.response);
    // @TODO(jengler) 2016-2-22: Don't expose raw response to user.
    alert("Error: " + event.target.response);
  }

  /**
   * Show success link for code.
   *
   * @param  {string} code The code for the newly created link.
   * @return {none}
   */
  function showSuccessLink(code) {
    // Show result
    var link = document.getElementById("link");
    link.href = "/" + code;
    // Add the left arrow. It also clears the already present text node.
    link.innerHTML = "&#x21AA;";
    link.appendChild(document.createTextNode(code));

    document.getElementById('link-container').classList.add("visible");
  };
})(window, document)
