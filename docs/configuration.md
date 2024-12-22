---
layout: default
title: "Configuration Guide"
---

# Configuration Guide

This guide explains how to configure your project using the `config.yaml` file.

## Overview

The `config.yaml` file contains various settings to customize your project. Below is a breakdown of each section and its options.

## Configuration Sections

### `tls`

- **Description**: Enable or disable TLS.
- **Type**: Boolean
- **Example**: `tls: false`

### `head`

- **Description**: Configuration for the HTML head section.
- **Options**:
  - `logo`: Path to the logo image.
  - `name`: Name of the association.
  - `icon`: Path to the favicon.
  - `description`: Meta description for the site.
  - `backgroundcolor`: Background color in RGB format.
  - `textcolor`: Text color in RGB format.
  - `slogan`:
    - `font`: Font for the slogan.
    - `top`: Top part of the slogan.
    - `bottom`: Bottom part of the slogan.

- **Example**:
  ```yaml
  head:
    logo: "logo.svg"
    name: "Association Name"
    icon: "logo.ico"
    description: "Description"
    backgroundcolor: "154, 93, 185"
    textcolor: "255, 255, 255"
    slogan:
      font: "Sofia, sans-serif"
      top: "Your"
      bottom: "Slogan"
  ```

### `body`

- **Description**: Configuration for the body section.
- **Options**:
  - `departments`: List of departments.
    - `name`: Name of the department.
    - `contacts`: List of contacts.
      - `name`: Name of the contact.
      - `title`: Title of the contact.
      - `description`: Description of the contact.
      - `email`: Email of the contact.
      - `tel`: Telephone number of the contact.
    - `gallery`: List of gallery images.
      - `title`: Title of the image.
      - `src`: Source path of the image.
    - `bfvwidgets`: List of BFV widgets.
      - `teamid`: Team ID.
      - `clubid`: Club ID.
      - `compoundid`: Compound ID.
      - `type`: Type of the widget.

- **Example**:
  ```yaml
  body:
    departments:
      - name: "Football"
        contacts:
          - name: "Don Joe"
            title: "responsible"
            description: "responsible for Example Sport"
            email: "don.joe@example.com"
            tel: "555123456789"
        gallery: 
          - title: "Example Image"
            src: "example-image.webp"
        bfvwidgets: 
          - teamid: "60428905238095242232"
            clubid: ""
            compoundid: ""
            type: "teamfull"
  ```

### `footer`

- **Description**: Configuration for the footer section.
- **Options**:
  - `background_image`: Path to the background image.
  - `socials`: List of social media accounts.
    - `account`: Social media account name.
    - `link`: URL to the social media account.
    - `image`: Path to the social media icon.
  - `contacts`: List of contacts.
    - `title`: Title of the contact.
    - `name`: Name of the contact.
    - `address`: Address of the contact.
    - `email`: Email of the contact.
    - `tel`: Telephone number of the contact.
    - `responsible`: Boolean indicating if the contact is responsible.

- **Example**:
  ```yaml
  footer:
    background_image: "example-background.svg"
    socials:
      - account: "fc_example_1900"
        link: "https://www.instagram.com/fc_example_1900/"
        image: "insta.svg"
    contacts:
      - title: "Board"
        name: "John Doe"
        address: "Example str. 42 | 12345 City"
        email: "john.doe@example.com"
        tel: "555 / 123456"
        responsible: true
  ```

### `fonts`

- **Description**: List of custom fonts.
- **Options**:
  - `name`: Name of the font.
  - `src`: Source path of the font file.
  - `font_weight`: Font weight.
  - `font_style`: Font style.

- **Example**:
  ```yaml
  fonts:
    - name: "Sofia"
      src: "Sofia-Regular.woff"
      font_weight: "normal"
      font_style: "normal"
  ```

### `lang_file`

- **Description**: Path to the language file.
- **Type**: String
- **Example**: `lang_file: "en.yaml"`

### `linked_sites`

- **Description**: List of linked sites.
- **Options**:
  - `name`: Name of the site.
  - `url`: URL of the site (if type is `link`).
  - `type`: Type of the link (`link` or `dropdown`).
  - `links`: List of sub-links (if type is `dropdown`).
    - `name`: Name of the sub-link.
    - `url`: URL of the sub-link.

- **Example**:
  ```yaml
  linked_sites:
    - name: "Cloud"
      url: "https://cloud.example.com/"
      type: "link"
    - name: "Files"
      type: "dropdown"
      links:
        - name: "Membership"
          url: "https://cloud.example.com/membership"
        - name: "Rental"
          url: "https://cloud.example.com/rental"
    - name: "Calendars"
      type: "dropdown"
      links:
        - name: "Sports home occupancy"
          url: "https://cloud.example.com/calendars/occupancy"
  ```

### `custom_pages`

- **Description**: List of custom pages.
- **Options**:
  - `name`: Name of the custom page.
  - `file`: Path to the custom page file.

- **Example**:
  ```yaml
  custom_pages:
    - name: "Dataprotection"
      file: "dataprotection.html"
  ```

## Conclusion

This guide provides an overview of the `config.yaml` file and its options. Customize the settings as needed to configure your project.
