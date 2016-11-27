// trainctl is an opinionated tool for managing training materials.
//
// trainctl allows you to mix multiple training topics in various combinations,
// creating courses from the topics.
//
// Topics are self-contained modules.  A topic might represent a single presentation
// at a conference, or a single topic of a multi-topic course.
//
// Courses are a collection of topics assembled in a specific order to create a logical
// progression of learning material.
//
// Topics are stored in a base folder.  We'll refer to the base folder as $TOPICDIR
// Each topic consists of a present slide file, a json manifest and a subdirectory
// bearing the same name as the topic, containing
// the slide's assets, in `images` and `includes` directories.
//
// Source code intended for student consumption is stored in `$TOPICDIR/src/topicname`.
// The expectation is that each course will be a self-contained $GOPATH, so the import
// names of each topic are simply the topicname.
//
// trainctl is a tool that allows you to manage topics individually, with the end goal of
// assembling them into a course which will be presented/taught.
//
// trainctl uses a configuration file stored at $HOME/.trainctl.yaml which can be created using
// `trainctl createconfig`.  It is not required, but will pre-fill many of the fields needed for
// every command, so it is strongly recommended.
//
// Usage:
//  trainctl [command]
//
// Available Commands:
//  assemble     Assemble topics into a course
//  create       Create a new topic skeleton
//  createconfig Create a config file in your home directory
//  info         Get info about a topic
//  package      package a course into a downloadable zip file
//  related      Add related material to a topic
//  remove       Remove a new topic skeleton
//  search       Search topics by metadata
//
// Start with `trainctl createconfig` to create a configuration file in your home directory.
// Edit this file with real values representing your configuration.
// The $TOPICDIR should be a source controlled directory.
// The $COURSEDIR can be any directory, source control not required (or recommended).
//
// To create a topic run:
//  trainctl create --name mytopic
//
// This creates the following files and directories:
//  $TOPICDIR/mytopic.slide
//  $TOPICDIR/mytopic.json
//  $TOPICDIR/mytopic/images
//  $TOPICDIR/mytopic/includes
// These files represent the topic's slide and assets required to present the slide.
//
// Additionally, the following directories are created:
//  $TOPICDIR/src/mytopic/demos
//  $TOPICDIR/src/mytopic/exercises
//  $TOPICDIR/src/mytopic/solutions
// These directories are intended to be a self-contained $GOPATH for the topic, and eventually
// the courses that contain this topic.
//
// Create one or more topics this way, then assemble them into a course with the
// `trainctl assemble` command like this:
//  trainctl assemble --course "My Class" --topics topic1,topic2,topic4 --shortname myclass
// The assemble command will create a course in the $COURSEDIR folder with the name `myclass` in the
// above example.  The following files are created
//  $COURSEDIR/mytopic.json
//  $COURSEDIR/mytopic.slide
// The .json file is a manifest of the course's contents.
// The .slide file is a symlink to your topic's .slide file, with the addition of a numbered
// prefix, establishing the order the topics you specified in the `assemble` command.
// The topic directory is symlinked to:
//  $COURSEDIR/mytopic
// with the subdirectories intact so the slide's assets will still appear on the slide.
//
// The `package` sub-command takes a course from your $COURSEDIR and creates a zipfile from
// it, which has the effect of de-referencing the symlinks.  The benefit if this approach is that
// your courses will always contain the latest up-to-date changes from the topics included in them,
// up to the point where you memorialize them immediately before presenting your class.
// Repackaging the course in the future will update the zipfile with files in the symlinked
// topics that may have changed since last giving the course.
//
//  trainctl package --shortname mycourse
// This creates a zipfile called mycourse.zip with the contents of the `mycourse` included.
// This is the course material you will use to present, or give to your students.
// It also includes a Vagrantfile and a supporting bootstrap bash script which provides
// a fully operational Go development with the course's /src shared into the Vagrant container.
//
// The `search` and `info` sub-commands give you information about topics in your
// $TOPICDIR so that you can choose which topics to assemble into a course.
package main
