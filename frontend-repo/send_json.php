<?php

$url = "http://localhost:8080/create_func_from_repo";

$data = array(
  'repo'	=> $_POST['repo'],
  'tag'		=> $_POST['tag'],
  'path'	=> $_POST['path']
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

{"repo":"<?php echo htmlspecialchars($_POST['repo']); ?>", "tag":"<?php echo htmlspecialchars($_POST['tag']); ?>","path":"<?php echo htmlspecialchars($_POST['path']); ?>"}

<?php echo var_dump($response); ?>
