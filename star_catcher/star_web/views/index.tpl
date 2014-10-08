<div class="row">
{{range $key, $val := .Stars}}
  <!-- {{$key}} -->

  
    <div class="col-md-4">
      <a href="/star/{{$val.Id}}">
        {{$val.Name}}
      </a>
    </div>
  
{{end}} 
</div>