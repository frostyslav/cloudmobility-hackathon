<?php

$url = "http://localhost:8080/create_func_from_repo";

$data = array(
  'repo' => array(
    'url'	=> $_POST['repo_url'],
    'tag'	=> $_POST['repo_tag'],
    'path'	=> $_POST['repo_path']
  )
);

$options = array(
  'http' => array(
    'method'  => 'POST',
    'content' => json_encode( $data ),
    'header'=>  "Content-Type: application/json\r\n" .
                "Accept: application/json\r\n"
    )
);
$context  = stream_context_create( $options );
$result = file_get_contents( $url, false, $context );
$response = json_decode( $result );
?>
