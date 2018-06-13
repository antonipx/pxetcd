<!DOCTYPE html>
<html>
<head>
  <style>
    * {
      color: #dedede;
      font-family: Helvetica, sans-serif;
    }

    code {
      color: #F15A22;
      font-family: Monaco, monospace;
    }

    body {
      background-color: #37464c;
    }

    a:-webkit-any-link {
      text-decoration: none;
    }

    input,
    select,
    option {
      background-color: #19272e;
      border: 0px solid;
      font-size: 1em;
      width: 100%;
      height: 100%;
      display: block;
    }
    
    input[type="button"] {
        width: 200px;
    }

    table {
      border-collapse: collapse;
      background-color:#24363d;
    }

    th, td {
      margin: 0px;
      padding: 0px;
      height: 27px;
      min-width: 300px;
      padding-left: 5px;
    }

    tr {
      border: solid 1px #37464c;
      border-radius:3px;
    }

    th {
      font-weight: bold;
    }

    td input[type="text"] {
      text-align: center;
    }

    td:first-of-type, th:first-of-type {
      background-color: #19272e;
    }

    td[colspan="2"] {
      background-color:#37464c;
    }

    #hdr {
      padding: 5px;
      text-align: left;
    }
    #hdr img {
      float: left;
      margin: 0 20px 10px 10px;
    }
    #left {
      margin: 0;
      padding: 5px;
      text-align: left;
    }
  </style>
</head>

<body>

  <div id="hdr">
  <img src="https://portworx.com/wp-content/themes/portworx/images/xportworx-logo.png.pagespeed.ic.0Z487a9F2s.png" width="200" height="85">
  <h1>ETCD configurator for Portworx</h1>
  </div>

  <p>Copy and run this on each node where ETCD is to be installed:</p>

  <div id="k"></div>
  <p>

  <form id="px-form">
    <table BORDER="1" id="options">
      <tr>
        <th>Value</th>
        <th>Description</th>
      </tr>
      <tr>
        <td>
          <INPUT TYPE="text" NAME="i1" VALUE="x.x.x.x" onChange="proc(this.form)">
        </td>
        <td><code>Node 1 IP:</code> IP address for Etcd Node 1</td>
      </tr>
      <tr>
        <td>
          <INPUT TYPE="text" NAME="i2" VALUE="y.y.y.y" onChange="proc(this.form)">
        </td>
        <td><code>Node 2 IP:</code> IP address for Etcd Node 2</td>
      </tr>
      <tr>
        <td>
          <INPUT TYPE="text" NAME="i3" VALUE="z.z.z.z" onChange="proc(this.form)">
        </td>
        <td><code>Node 3 IP:</code> IP address for Etcd Node 3</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <SELECT NAME="e" onChange="proc(this.form)">
            <OPTION value="no">No Encryption (http)</OPTION>
            <OPTION value="at">Automatic TLS (https)</OPTION>
            <OPTION value="ca">CA Certs (Self Signed)</OPTION>
          </SELECT>
        </td>
        <td><code>Encryption:</code> Select Etcd encryption scheme</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="t" VALUE="pxetcd1" onChange="proc(this.form)">
        </td>
        <td><code>Name:</code> Specifies Etcd cluster name (initial token)</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="r" VALUE="px" onChange="proc(this.form)">
        </td>
        <td><code>Prefix:</code> Node name prefix - will be assinged a sequencial id</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="c" VALUE="9017" onChange="proc(this.form)">
        </td>
        <td><code>Client Port:</code> TCP port number for client connections (default 2379)</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="p" VALUE="9018" onChange="proc(this.form)">
        </td>
        <td><code>Peer Port:</code> TCP port number for peer communication (default 2380)</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="d" VALUE="/opt/pxetcd" onChange="proc(this.form)">
        </td>
        <td><code>Etcd Dir:</code> Etcd directory path</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="u" VALUE="pxetcd" onChange="proc(this.form)">
        </td>
        <td><code>Username:</code> Etcd service user name</td>
      </tr>
      <tr class="advanced" style="display:none;">
        <td>
          <INPUT TYPE="text" NAME="v" VALUE="v3.2.12" onChange="proc(this.form)">
        </td>
        <td><code>Etcd Ver:</code> Version from https://github.com/coreos/etcd/releases</td>
      </tr>



     </table>

    <br>
    
    <INPUT TYPE="button" NAME="a" VALUE="Advanced Options" onClick="toggleadv();">


  </form>

</body>

<script LANGUAGE="JavaScript">
  function toggleadv() {
      var rows = document.getElementsByClassName("advanced")
      
      for (var i = 0; i < rows.length; i++) {
          if ( rows[i].style.display == "table-row" )
            rows[i].style.display = "none"
          else 
            rows[i].style.display = "table-row"
      }
  }

  function proc(form) {
    var elements = form.elements
    var sep = "?"
    var s = ""
    for (var i = 0, el; el = elements[i++];) {
      if ((el.type === "text" || el.type === "select-one") && el.value != "") {
        s += sep + el.name + "=" + el.value
        sep = "&"
      } else if (el.type === "checkbox" && el.checked) {
        s += sep + el.name + "=true"
        sep = "&"
      }
    }
    //s = document.URL + s
    k.innerHTML = "<p>&#x25BA;&nbsp; curl <a href=\"" + window.location.href + s + "\">" +  window.location.href + s + "</a> | sudo bash </p>"
  }

  proc(document.getElementById("px-form"));
</script>

</html>
