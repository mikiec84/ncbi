// Copyright ©2013 The bíogo.ncbi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blast

import (
	"strings"
	"time"

	"gopkg.in/check.v1"
)

func (s *S) TestParseRid(c *check.C) {
	for i, t := range []struct {
		retval string
		rid    string
		err    error
		wait   time.Duration
	}{
		{
			`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="jig" content="ncbitoggler ncbiautocomplete"/>
<meta name="ncbi_app" content="blast" />
<meta name="ncbi_pdid" content="blastformatreq" />
<meta name="ncbi_stat" content="false" />
<meta name="ncbi_sessionid" content="CE8905C20F88BA61_0093SID" />
<meta name="ncbi_phid" content="2DC61DAF12D68CA100000000000007A1" />
<script type="text/javascript"> var ncbi_startTime = new Date(); </script>
<title>NCBI Blast</title>
<link rel="stylesheet" type="text/css" href="css/main.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/common.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/blastReq.css" media="screen" />
<!--[if IE]>
<link rel="stylesheet" type="text/css" href="css/blastReqIE.css" media="screen" />
<![endif]-->
<link rel="stylesheet" type="text/css" href="css/print.css" media="print" />


<!--[if lte IE 6]>
<link rel="stylesheet" type="text/css" href="css/ie6_or_less.css" />
<![endif]-->
<script type="text/javascript" src="http://www.ncbi.nlm.nih.gov/core/jig/1.11/js/jig.min.js"></script>
<script type="text/javascript" src="js/utils.js"></script>
<script type="text/javascript" src="js/blast.js"></script>
<script type="text/javascript" src="js/format.js"></script>

</head>

<body id="type-a">

<div id="wrap">
			<div id="header">
		<div id="site-name"><a id="logolink" href="http://www.ncbi.nlm.nih.gov" title="NCBI Home Page"><img src="css/images/helix.gif" alt="NCBI Logo" title="Link to NCBI Home Page" /></a>BLAST <span id="trdm"> &reg;</span><h1 class="desc">Basic Local Alignment Search Tool</h1>
		</div>
		<div id="search">

<div>
<script language="JavaScript" type="text/javascript"><!--
// --></script><table class="medium1" style="border:2px solid #336699;" cellpadding="2" cellspacing="0" id="myncbi_off"><tr><td
bgcolor="#336699" align="left"><a href="http://www.ncbi.nlm.nih.gov/myncbi/?"><font color="#FFFFFF"><b>My NCBI</b></font></a></td><td
bgcolor="#336699" align="right"><a href="http://www.ncbi.nlm.nih.gov/books/NBK3842/" title="My NCBI help"><img border="0"
src="http://www.ncbi.nlm.nih.gov/corehtml/query/MyNCBI/myncbihelpicon.gif" alt="My NCBI help" /></a></td></tr><tr><td colspan="2" nowrap="nowrap"><a
href="http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26DATABASE%3Dnr%26ENTREZ%5FQUERY%3D%26FILTER%3DL%26FULL%5FDBNAME%3Dnr%26HITLIST%5FSZE%3D500%26JOB%5FTITLE%3DProtein%2BSequence%2B%283%2Bletters%29%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26PROGRAM%3Dblastp%26QUERY%5FINFO%3DProtein%2BSequence%2B%283%2Bletters%29%26QUERY%5FLENGTH%3D3%26RID%3DJUU6AN9D01R%26RTOE%3D21%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2" title="Click to sign in"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26DATABASE%3Dnr%26ENTREZ%5FQUERY%3D%26FILTER%3DL%26FULL%5FDBNAME%3Dnr%26HITLIST%5FSZE%3D500%26JOB%5FTITLE%3DProtein%2BSequence%2B%283%2Bletters%29%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26PROGRAM%3Dblastp%26QUERY%5FINFO%3DProtein%2BSequence%2B%283%2Bletters%29%26QUERY%5FLENGTH%3D3%26RID%3DJUU6AN9D01R%26RTOE%3D21%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2');return false;">[Sign In]</a> <a
href="http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26DATABASE%3Dnr%26ENTREZ%5FQUERY%3D%26FILTER%3DL%26FULL%5FDBNAME%3Dnr%26HITLIST%5FSZE%3D500%26JOB%5FTITLE%3DProtein%2BSequence%2B%283%2Bletters%29%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26PROGRAM%3Dblastp%26QUERY%5FINFO%3DProtein%2BSequence%2B%283%2Bletters%29%26QUERY%5FLENGTH%3D3%26RID%3DJUU6AN9D01R%26RTOE%3D21%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2" title="Click to register for an account"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26DATABASE%3Dnr%26ENTREZ%5FQUERY%3D%26FILTER%3DL%26FULL%5FDBNAME%3Dnr%26HITLIST%5FSZE%3D500%26JOB%5FTITLE%3DProtein%2BSequence%2B%283%2Bletters%29%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26PROGRAM%3Dblastp%26QUERY%5FINFO%3DProtein%2BSequence%2B%283%2Bletters%29%26QUERY%5FLENGTH%3D3%26RID%3DJUU6AN9D01R%26RTOE%3D21%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2');return false;">[Register]</a></td></tr></table></div>
		</div>
		<a class="skp" href="#content-wrap">Jump to Page Content</a>
		<ul id="nav">
                <li  class="first "><a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastHome" title="BLAST Home">Home</a></li>
                <li  class="recent "><a href="Blast.cgi?CMD=GetSaved&amp;RECENT_RESULTS=on" title="Unexpired BLAST jobs">Recent Results</a></li>
                <li  class="saved "><a href="Blast.cgi?CMD=GetSaved" title="Saved sets of BLAST search parameters">Saved Strategies</a></li>
                <li  class= "last documentation "> <a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastDocs" title="BLAST documentation">Help</a></li>
                </ul>
    </div>
        <div id="content-wrap">

                <!-- %%% Add breadcrumb text -->
                <div id="breadcrumb">
                   <a href="http://www.ncbi.nlm.nih.gov/">NCBI</a>/
                   <a href="Blast.cgi?CMD=Web&PAGE_TYPE=BlastHome">BLAST</a>/
                   <strong>Format Request</strong>
                   <span id="frmRequestPrTr"></span>
                </div>

				<!-- Do errors this way -->
				<!--<ul class="msg"><li class=""><p></p></li></ul>-->
				<ul id="msgR" class="msg"><li class=""></li></ul>
                <div id="content">
				<form action="Blast.cgi" enctype="application/x-www-form-urlencoded" method="post" name="FormatForm" id="FormatForm">

<script language="JavaScript">

 <!--

//document.images['BlastHeaderGif'].src = 'html/head_formating.gif';

// -->

</script>



<!--
                <p class='info'>
<strong>Job submitted.</strong>
We estimate that results will be ready in 16 seconds or less.

</p>
-->

<div class="fbtn">
<!--
<a href="javascript:document.forms[0].submit();">
<img align="middle" alt="Format button" border="0" src="FormatPage_files/format_but.gif">
</a>
-->
</div>

<dl class="summary  query title db">
<dd>
</dd>

<!-- <span class=" query title db">-->
<!-- <span  class="hidden query"><dt>Query</dt><dd>Protein Sequence (3 letters)</dd></span> -->
<dt class="hidden query">Query</dt><dd class="hidden query">Protein Sequence (3 letters)</dd>
<dt class="hidden db">Database</dt><dd class="hidden db">nr</dd>
<dt class="hidden title">Job title</dt><dd class="hidden title">Protein Sequence (3 letters)</dd>
<dt class="hidden entrez">Entrez Query</dt><dd class="hidden entrez"><span class="note entrez">Note: Your search is limited to records matching this Entrez query</span></dd>
<!-- </span> -->
<dt><label for="rid">Request ID</label></dt><dd><input name="RID" size="50" type="text" value="JUU6AN9D01R" id="rid" />
<input type="submit" value="View report" name="ViewReport" class="button" />
<!-- <img border="0" id="viewRpButton" src="images/veiwRpButton.jpg" class="viewReport"  alt="View report"  mouseovImg="images/veiwRpButtonOver.jpg" mouseoutImg="images/veiwRpButton.jpg" mousedownImg="images/veiwRpButtonDown.jpg" mouseupImg="images/veiwRpButtonOver.jpg"  />-->
<input type="checkbox" name="NEWWINRES"  form="FormatForm" winType="const" id="nw" class="newwin"  />
<label for="nw">Show results in a new window</label>
</dd>
<dt>Format<br/>
<!--<a class='help' href="#">[Help]</a></dt> -->

<dd>
<table class="options blastp ">

<tr class="paramSet xgl">
<td class="hd"><label for="FORMAT_OBJECT">Show</label></td>
<td>
<div class="fi">
<select id="FORMAT_OBJECT" class="reset" name="FORMAT_OBJECT" defVal="Alignment">
<option value="Alignment" >Alignment</option>
<option value="PSSM_Scoremat" >PssmWithParameters</option>
<option value="Bioseq"  >Bioseq</option>
</select>
<label for="FORMAT_TYPE">as</label>
<select name="FORMAT_TYPE" id="FORMAT_TYPE" class="reset" defVal="HTML">
<option value="HTML"  >HTML</option>
<option value="Text"  >Plain text</option>
<option value="ASN.1"  >ASN.1</option>
<option value="XML"  >XML</option>
</select>
<input name="PSSM_FORMAT_TYPE" value="Text" size="3" id="pssmFormat" type="text" class="hidden dispType" />
<input name="BIOSEQ_FORMAT_TYPE" value="ASN.1" size="3" id="bioseqFormat" type="text" class="hidden dispType" />
<input name="PSSM_SC_FORMAT_TYPE" value="ASN.1" size="3" id="pssmScFormat" type="text" class="hidden dispType" />
<span id="advView" class="">
<span class=""><input name="OLD_VIEW" id="OLD_VIEW" type="checkbox" class="cb reset" defVal="unchecked"  />
<label for="OLD_VIEW">Old View</label></span>
</span>
<a class="resetAll" id="resetAll" >Reset form to defaults</a>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments object formatting help" id="formatHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp1">
These options control formatting of alignments in results pages. The
default is HTML, but other formats (including plain text) are available.
PSSM and PssmWithParameters are representations of Position Specific Scoring Matrices and are only available for PSI-BLAST.
The Advanced view option allows the database descriptions to be sorted by various indices in a table.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#format_object" target="helpWin" title="Additional alignments object formatting help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="odd paramSet">
<td class="hd"><label for="ALIGNMENT_VIEW">Alignment View</label></td>
<td>
<div class="fi">
<select name="ALIGNMENT_VIEW" id="ALIGNMENT_VIEW" defVal="Pairwise" class="reset">
<option value="Pairwise"  >Pairwise</option>
<option value="PairwiseWithIdentities"  >Pairwise with dots for identities</option>
<option value="QueryAnchored"  >Query-anchored with dots for identities</option>
<option value="QueryAnchoredNoIdentities"  >Query-anchored with letters for identities</option>
<option value="FlatQueryAnchored"  >Flat query-anchored with dots for identities</option>
<option value="FlatQueryAnchoredNoIdentities"  >Flat query-anchored with letters for identities</option>
<option value="Tabular"  >Hit Table</option>
</select>


<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments view options help" id="alnViewHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp2">
Choose how to view alignments.
The default "pairwise" view shows how each subject sequence aligns
individually to the query sequence. The "query-anchored" view shows how
all subject sequences align to the query sequence. For each view type,
you can choose to show "identities" (matching residues) as letters or
dots.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#alignment_view" target="helpWin" title="Additional alignments view options help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="paramSet">
<td class="hd"><label>Display</label></td>
<td class="cb">
<div class="fi">
<input name="SHOW_OVERVIEW" id="SHOW_OVERVIEW" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="SHOW_OVERVIEW">Graphical Overview</label>

<span id="shl" >
<input name="SHOW_LINKOUT" id="SHOW_LINKOUT" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="SHOW_LINKOUT">Linkout</label>
</span>
<span id="gts" >
<input name="GET_SEQUENCE" id="GET_SEQUENCE" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="GET_SEQUENCE">Sequence Retrieval</label>
</span>

<input name="NCBI_GI" id="NCBI_GI" type="checkbox" class="cb reset" defVal="unchecked"  />
<label class="rb" for="NCBI_GI">NCBI-gi</label>
<span id="scf" >
<input name="SHOW_CDS_FEATURE" id="SHOW_CDS_FEATURE" type="checkbox" class="cb reset blastn" defVal="unchecked"  />
<label for="SHOW_CDS_FEATURE" class="blastn">CDS feature</label>
</span>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments display options help" id="displayHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp3">
<li>Graphical Overview: Graphical Overview: Show graph of similar sequence regions aligned to  query.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_overview" target="helpWin" title="Graphical Overview help">more...</a>
</li>
<li>Database LinkOuts: Show links from matching sequences to entries in specialized NCBI databases.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_linkout" title="LinkOut help" target="helpWin" >more...</a>
</li>
<li>Sequence Retrieval: Show buttons to download matching sequences.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#get_sequence" title="Sequence Retrieval help" target="helpWin" >more...</a>
</li>
<li>NCBI-gi: Show NCBI gi identifiers.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#ncbi_gi" title="NCBI-gi help" target="helpWin" >more...</a>
</li>
<li>CDS feature: Show annotated coding region and translation.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_cds_feature" title="CDS feature help" target="helpWin" >more...</a>
</li></ul>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr class="paramSet odd xgl">
<td class="hd"><label>Masking</label></td>
<td>
<div class="fi">
<label for="MASK_CHAR"> Character: </label>
<select name="MASK_CHAR" id="MASK_CHAR"  class="reset" defVal="2">
<option value="0"  >X for protein, n for nucleotide</option>
<option value="2" selected="selected" >Lower Case</option>
</select>
<label for="MASK_COLOR"> Color:</label>
<select name="MASK_COLOR" id="MASK_COLOR" class="reset" defVal="1">
<option value="0"  >Black
</option>

<option value="1" selected="selected" >Grey
</option>

<option value="2"  >Red
</option>

</select>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments masking help" id="maskingHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp4">
<li>Masking Character: Display masked (filtered) sequence regions as lower-case or as specific letters (N for nucleotide, P for protein).
</li>
<li>Masking Color: Display masked sequence regions in the given color.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr id="lr" class="paramSet xgl">
<td class="hd"><label>Limit results</label></td>
<td>
<div class="fi">
<label for="FRM_DESCRIPTIONS">Descriptions:</label>
<select name="DESCRIPTIONS" id="FRM_DESCRIPTIONS" class="reset" defVal="100">
<option value="0"      >0</option>
<option value="10"     >10</option>
<option value="50"     >50</option>
<option value="100"   selected="selected" >100</option>
<option value="250"    >250</option>
<option value="500"    >500</option>
<option value="1000"   >1000</option>
<option value="5000"   >5000</option>
<option value="10000"  >10000</option>
<option value="20000"  >20000</option>
</select>

<label for="FRM_NUM_OVERVIEW">Graphical overview:</label>
<select name="NUM_OVERVIEW" id="FRM_NUM_OVERVIEW" class="reset" defVal="100">
<option value="0"     >0</option>
<option value="10"    >10</option>
<option value="50"    >50</option>
<option value="100"  selected="selected" >100</option>
<option value="250"   >250</option>
<option value="500"  >500</option>
<option value="1000"  >1000</option>
</select>
<span id="frmAln">
<label for="FRM_ALIGNMENTS">Alignments:</label>
<select name="ALIGNMENTS" id="FRM_ALIGNMENTS" class="reset" defVal="100">
<option value="0"      >0</option>
<option value="10"     >10</option>
<option value="50"     >50</option>
<option value="100"   selected="selected" >100</option>
<option value="250"    >250</option>
<option value="500"    >500</option>
<option value="1000"   >1000</option>
<option value="5000"   >5000</option>
<option value="10000"  >10000</option>
<option value="20000"  >20000</option>
</select>
</span>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit number of descriptions/alignments help" id="numHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp5">
<li>Descriptions: Show short descriptions for up to the given number of  sequences.</li>
<li>Alignments:  Show alignments for up to the given number of sequences, in order of statistical significance.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="paramSet odd xgl ">
<td class="hd"></td>
<td>
<div class="">
<label for="qorganism">Organism</label>
<span class="instr">Type common name, binomial, taxid, or group name. Only 20 top taxa will be shown.</span><br>
<input name="FORMAT_ORGANISM" size="55"  type="text" id="qorganism" value="" data-jigconfig="dictionary:'taxids_sg'" autocomplete="off" class="jig-ncbiautocomplete reset">
<input type="checkbox" name="FORMAT_ORG_EXCLUDE"  class="oExclR cb" id="orgExcl"/>
<input type="hidden" value = "1" name="FORMAT_NUM_ORG" id="numOrg" />
<label for="orgExcl" class="right">Exclude</label>
<a href="#" title="Add organism" id="addOrg"><img border="0" src="css/images/addOrg.jpg" id="addOrgIm"   alt="Add organism"  mouseovImg="css/images/addOrgOver.jpg" mouseoutImg="css/images/addOrg.jpg" mousedownImg="css/images/addOrgDown.jpg" mouseupImg="css/images/addOrgOver.jpg"  /></a>
<div id="orgs">

</div>
<div class="fi">
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by organism help" id="organismHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp6">
Show only sequences from the given organism.
</p>
</div><!-- ARIA -->
</div>
</div>
</td>
</tr>

<tr class="paramSet xgl ">
<td class="hd"></td>
<td>
<div class="fi">
<label for="FORMAT_EQ_TEXT">Entrez query:</label>
<input name="FORMAT_EQ_TEXT" id="FORMAT_EQ_TEXT" size="60" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by Entrez query help" id="entrezHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp7">
Show only those sequences that match the given Entrez query.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#limit_result" target="helpWin" title="Additional limit results by Entrez query help"  target="helpWin">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr class="paramSet odd xgl">
<td class="hd"></td>
<td>
<div class="fi">
<label for="EXPECT_LOW">Expect Min:</label> <input name="EXPECT_LOW" id="EXPECT_LOW" size="10" type="text" value="" class="reset"/>
<label for="EXPECT_HIGH">Expect Max:</label> <input name="EXPECT_HIGH" id="EXPECT_HIGH" size="10" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by expect value range help" id="expectHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp8">
Show only sequences with expect values in the given range.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#expect_range" target="helpWin" title="Additional limit results by expect value range help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>
<tr class="paramSet xgl">
<td class="hd"></td>
<td>
 <div class="fi">
<label for="PERC_IDENT_LOW">Percent Identity Min:</label> <input name="PERC_IDENT_LOW" id="PERC_IDENT_LOW" size="10" type="text" value="" class="reset"/>
<label for="PERC_IDENT_HIGH">Percent Identity Max:</label> <input name="PERC_IDENT_HIGH" id="PERC_IDENT_HIGH" size="10" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by percent identity range help" id="percIdentHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp10">
 Show only sequences with percent identity values in the given range.
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>
<tr class="psiBlast odd paramSet xgl">
<td class="hd"><label>Format for</label></td>
<td>
<div class="fi">
<input name="RUN_PSIBLAST_FORM" id="RUN_PSIBLAST" type="checkbox" class="cb psiBlast"  />
<label class="rb psiBlast" for="RUN_PSIBLAST">PSI-BLAST</label>
<label for="I_THRESH">with inclusion threshold:</label>
<input name="I_THRESH" id="I_THRESH" size="10" type="text" value="" defVal="0.005" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="PSI BLAST formatting help" id="psiHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp9">
<li>Format for PSI-BLAST: The Position-Specific Iterated BLAST (PSI-BLAST) program performs iterative searches with a protein query,
in which sequences found in one round of search are used to build a custom score model for the next round.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#psiblast" target="helpWin" title="Additional PSI BLAST formatting help">more...</a>
</li>
<li>Inclusion Threshold: This sets the statistical significance threshold for including a sequence in the model used
by PSI-BLAST to create the PSSM on the next iteration.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>
</table>
</dd>
</dl>

<input name="RID" value="JUU6AN9D01R" type="hidden" />
<input name="CDD_RID" value="" type="hidden" />
<input name="CDD_SEARCH_STATE" type="hidden" value="" />

<input name="STEP_NUMBER" value="" id="stepNumber" type="hidden" />
<input name="CMD" value="Get" type="hidden" />
<input name="FORMAT_EQ_OP" value="AND" type="hidden" />
<input name="RESULTS_PAGE_TARGET" type="hidden" id="resPageTarget" value="Blast_Results_for_1738872954" />
<input name="QUERY_INFO" type="hidden" value="Protein Sequence (3 letters)" />
<input name="ENTREZ_QUERY" type="hidden" value="" />
<input name="QUERY_INDEX" type="hidden" value="0"/>
<input name="NUM_QUERIES" type="hidden" value="1"/>
<input name="CONFIG_DESCR" type="hidden" value="2,3,4,5,6,7,8" />

<!-- Those params are set in the template (blastn.dat, blastp.dat etc. -->
<input name="BLAST_PROGRAMS" type="hidden" value="blastp"/>
<input name="PAGE" type="hidden" value="Proteins"/>
<input name="PROGRAM" type="hidden" value="blastp"/>
<input name="MEGABLAST" type="hidden" value="" />
<input name="RUN_PSIBLAST" type="hidden" value="" />
<input name="BLAST_SPEC" id="blastSpec" type="hidden" value=""/>


<input name="QUERY" type="hidden" value=""/>
<input name="JOB_TITLE" type="hidden" value="Protein Sequence (3 letters)"/>
<input name="QUERY_TO" type="hidden" value=""/>
<input name="QUERY_FROM" type="hidden" value=""/>
<input name="EQ_TEXT" type="hidden" value=""/>
<input name="ORGN" type="hidden" value=""/>
<input name="EQ_MENU" type="hidden" value=""/>
<input name="ORG_EXCLUDE" type="hidden" value=""/>
<input name="PHI_PATTERN" type="hidden" value=""/>
<input name="EXPECT" type="hidden" value=""/>
<input name="DATABASE" type="hidden" value="nr"/>
<input name="DB_GROUP" type="hidden" value=""/>
<input name="SUBGROUP_NAME" type="hidden" value=""/>

<input name="GENETIC_CODE" type="hidden" value=""/>
<input name="WORD_SIZE" type="hidden" value=""/>
<input name="MATCH_SCORES" type="hidden" value=""/>
<input name="MATRIX_NAME" type="hidden" value=""/>
<input name="GAPCOSTS" type="hidden" value=""/>
<input name="MAX_NUM_SEQ" id="maxNumSeq" type="hidden" value=""/>
<input name="COMPOSITION_BASED_STATISTICS" type="hidden" value=""/>
<input name="NEWWIN" type="hidden" value=""/>
<input name="SHORT_QUERY_ADJUST" type="hidden" value=""/>
<input name="FILTER" type="hidden" value="L;"/>
<input name="REPEATS" type="hidden" value=""/>
<input name="ID_FOR_PSSM" type="hidden" value=""/>
<input name="EXCLUDE_MODELS" type="hidden" value=""/>
<input name="EXCLUDE_SEQ_UNCULT" type="hidden" value=""/>
<input name="NUM_ORG" type="hidden" value = "1" />

<!-- PSSM -->
<input name="LCASE_MASK" type="hidden" value=""/>
<input name="TEMPLATE_TYPE" type="hidden" value=""/>
<input name="TEMPLATE_LENGTH" type="hidden" value=""/>
<input name="I_THRESH" type="hidden" value=""/>
<input name="PSI_PSEUDOCOUNT" type="hidden" value=""/>
<input name="DI_THRESH" type="hidden" id="diThresh" value=""/>
<input name="HSP_RANGE_MAX" type="hidden" value=""/>



<input name="ADJUSTED_FOR_SHORT_QUERY" type="hidden" value=""/>
<input name="MIXED_QUERIES" type="hidden" value=""/>
<input name="MIXED_DATABASE" id="mixedDb" type="hidden" value=""/>
<input name="BUILD_NAME"  type="hidden" value=""/>
<input name="ORG_DBS"  type="hidden" value=""/>

<!--QBlastInfoBegin
    RID = JUU6AN9D01R
    RTOE = 2
QBlastInfoEnd
-->
</form>

				</div><!-- /#content -->

        </div><!-- /#content-wrap -->


<div id="footer">
   <div id="rgs">BLAST is a registered trademark of the National Library of Medicine.</div>
   <p id="orgns">
      <a href="http://www.ncbi.nlm.nih.gov/" title="National Center for Biotechnology Information">NCBI</a> |
      <a href="http://www.nlm.nih.gov/" title="National Library of Medicine">NLM</a> |
      <a href="http://www.nih.gov/" title="National Institutes of Health">NIH</a> |
      <a href="http://www.hhs.gov/" title="US Department of Health and Human Services">DHHS</a>
   </p>

   <p>
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html'
      title='NCBI intellectual property statement'>Copyright</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html#disclaimer'
      title='About liability, endorsements, external links, pop-up advertisements'>Disclaimer</a> |
      <a href='http://www.nlm.nih.gov/privacy.html'
      title='NLM privacy policy'>Privacy</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/accessibility.html'
      title='About using NCBI resources with assistive technology'>Accessibility</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/glance/contact_info.html'
      title='How to get help, submit data, or provide feedback'>Contact</a> |
      <a href='mailto:blast-help@ncbi.nlm.nih.gov'
      title='How to get help, submit data, or provide feedback'>Send feedback</a>
   </p>
</div>
   </div><!--/#wrap-->

<script type="text/javascript" src="http://blast.ncbi.nlm.nih.gov/portal/portal3rc.fcgi/rlib/js/InstrumentOmnitureBaseJS/InstrumentNCBIBaseJS/InstrumentPageStarterJS.js"></script>
</body>

</html>


`,
			"JUU6AN9D01R",
			nil,
			2 * time.Second,
		},
		{
			`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="jig" content="ncbitoggler ncbiautocomplete"/>
<meta name="ncbi_app" content="blast" />
<meta name="ncbi_pdid" content="blastformatreq" />
<meta name="ncbi_stat" content="false" />
<meta name="ncbi_sessionid" content="CE8905C20F88BA61_0093SID" />
<meta name="ncbi_phid" content="5AAB5B2D12D7C9B1000000000000E978" />
<script type="text/javascript"> var ncbi_startTime = new Date(); </script>
<title>NCBI Blast</title>
<link rel="stylesheet" type="text/css" href="css/main.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/common.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/blastReq.css" media="screen" />
<!--[if IE]>
<link rel="stylesheet" type="text/css" href="css/blastReqIE.css" media="screen" />
<![endif]-->
<link rel="stylesheet" type="text/css" href="css/print.css" media="print" />


<!--[if lte IE 6]>
<link rel="stylesheet" type="text/css" href="css/ie6_or_less.css" />
<![endif]-->
<script type="text/javascript" src="http://www.ncbi.nlm.nih.gov/core/jig/1.11/js/jig.min.js"></script>
<script type="text/javascript" src="js/utils.js"></script>
<script type="text/javascript" src="js/blast.js"></script>
<script type="text/javascript" src="js/format.js"></script>

</head>

<body id="type-a">

<div id="wrap">
			<div id="header">
		<div id="site-name"><a id="logolink" href="http://www.ncbi.nlm.nih.gov" title="NCBI Home Page"><img src="css/images/helix.gif" alt="NCBI Logo" title="Link to NCBI Home Page" /></a>BLAST <span id="trdm"> &reg;</span><h1 class="desc">Basic Local Alignment Search Tool</h1>
		</div>
		<div id="search">

<div>
<script language="JavaScript" type="text/javascript"><!--
// --></script><table class="medium1" style="border:2px solid #336699;" cellpadding="2" cellspacing="0" id="myncbi_off"><tr><td
bgcolor="#336699" align="left"><a href="http://www.ncbi.nlm.nih.gov/myncbi/?"><font color="#FFFFFF"><b>My NCBI</b></font></a></td><td
bgcolor="#336699" align="right"><a href="http://www.ncbi.nlm.nih.gov/books/NBK3842/" title="My NCBI help"><img border="0"
src="http://www.ncbi.nlm.nih.gov/corehtml/query/MyNCBI/myncbihelpicon.gif" alt="My NCBI help" /></a></td></tr><tr><td colspan="2" nowrap="nowrap"><a
href="http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26ERROR%3DMessage%2BID%252329%2BError%253A%2BQuery%2Bstring%2Bnot%2Bfound%2Bin%2Bthe%2BCGI%2Bcontext%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2" title="Click to sign in"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26ERROR%3DMessage%2BID%252329%2BError%253A%2BQuery%2Bstring%2Bnot%2Bfound%2Bin%2Bthe%2BCGI%2Bcontext%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2');return false;">[Sign In]</a> <a
href="http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26ERROR%3DMessage%2BID%252329%2BError%253A%2BQuery%2Bstring%2Bnot%2Bfound%2Bin%2Bthe%2BCGI%2Bcontext%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2" title="Click to register for an account"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DPut%26ERROR%3DMessage%2BID%252329%2BError%253A%2BQuery%2Bstring%2Bnot%2Bfound%2Bin%2Bthe%2BCGI%2Bcontext%26MYNCBI%5FUSER%3D6437729508%26MYNCBI%5FUSER%3D6437729508%26USER%5FTYPE%3D2%26USER%5FTYPE%3D2');return false;">[Register]</a></td></tr></table></div>
		</div>
		<a class="skp" href="#content-wrap">Jump to Page Content</a>
		<ul id="nav">
                <li  class="first "><a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastHome" title="BLAST Home">Home</a></li>
                <li  class="recent "><a href="Blast.cgi?CMD=GetSaved&amp;RECENT_RESULTS=on" title="Unexpired BLAST jobs">Recent Results</a></li>
                <li  class="saved "><a href="Blast.cgi?CMD=GetSaved" title="Saved sets of BLAST search parameters">Saved Strategies</a></li>
                <li  class= "last documentation "> <a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastDocs" title="BLAST documentation">Help</a></li>
                </ul>
    </div>
        <div id="content-wrap">

                <!-- %%% Add breadcrumb text -->
                <div id="breadcrumb">
                   <a href="http://www.ncbi.nlm.nih.gov/">NCBI</a>/
                   <a href="Blast.cgi?CMD=Web&PAGE_TYPE=BlastHome">BLAST</a>/
                   <strong>Format Request</strong>
                   <span id="frmRequestPrTr"></span>
                </div>

				<!-- Do errors this way -->
				<!--<ul class="msg"><li class="error"><p></p></li></ul>-->
				<ul id="msgR" class="msg"><li class="error"><p class="error">Message ID#29 Error: Query string not found in the CGI context</p></li></ul>
                <div id="content">
				<form action="Blast.cgi" enctype="application/x-www-form-urlencoded" method="post" name="FormatForm" id="FormatForm">

<script language="JavaScript">

 <!--

//document.images['BlastHeaderGif'].src = 'html/head_formating.gif';

// -->

</script>



<!--
                <p class='info'>
<strong>Job submitted.</strong>
We estimate that results will be ready in 16 seconds or less.

</p>
-->

<div class="fbtn">
<!--
<a href="javascript:document.forms[0].submit();">
<img align="middle" alt="Format button" border="0" src="FormatPage_files/format_but.gif">
</a>
-->
</div>

<dl class="summary ">
<dd>
</dd>

<!-- <span class="">-->
<!-- <span  class="hidden query"><dt>Query</dt><dd></dd></span> -->
<dt class="hidden query">Query</dt><dd class="hidden query"></dd>
<dt class="hidden db">Database</dt><dd class="hidden db"></dd>
<dt class="hidden title">Job title</dt><dd class="hidden title"></dd>
<dt class="hidden entrez">Entrez Query</dt><dd class="hidden entrez"><span class="note entrez">Note: Your search is limited to records matching this Entrez query</span></dd>
<!-- </span> -->
<dt><label for="rid">Request ID</label></dt><dd><input name="RID" size="50" type="text" value="" id="rid" />
<input type="submit" value="View report" name="ViewReport" class="button" />
<!-- <img border="0" id="viewRpButton" src="images/veiwRpButton.jpg" class="viewReport"  alt="View report"  mouseovImg="images/veiwRpButtonOver.jpg" mouseoutImg="images/veiwRpButton.jpg" mousedownImg="images/veiwRpButtonDown.jpg" mouseupImg="images/veiwRpButtonOver.jpg"  />-->
<input type="checkbox" name="NEWWINRES"  form="FormatForm" winType="const" id="nw" class="newwin"  />
<label for="nw">Show results in a new window</label>
</dd>
<dt>Format<br/>
<!--<a class='help' href="#">[Help]</a></dt> -->

<dd>
<table class="options  ">

<tr class="paramSet xgl">
<td class="hd"><label for="FORMAT_OBJECT">Show</label></td>
<td>
<div class="fi">
<select id="FORMAT_OBJECT" class="reset" name="FORMAT_OBJECT" defVal="Alignment">
<option value="Alignment" >Alignment</option>
<option value="PSSM_Scoremat" >PssmWithParameters</option>
<option value="Bioseq"  >Bioseq</option>
</select>
<label for="FORMAT_TYPE">as</label>
<select name="FORMAT_TYPE" id="FORMAT_TYPE" class="reset" defVal="HTML">
<option value="HTML"  >HTML</option>
<option value="Text"  >Plain text</option>
<option value="ASN.1"  >ASN.1</option>
<option value="XML"  >XML</option>
</select>
<input name="PSSM_FORMAT_TYPE" value="Text" size="3" id="pssmFormat" type="text" class="hidden dispType" />
<input name="BIOSEQ_FORMAT_TYPE" value="ASN.1" size="3" id="bioseqFormat" type="text" class="hidden dispType" />
<input name="PSSM_SC_FORMAT_TYPE" value="ASN.1" size="3" id="pssmScFormat" type="text" class="hidden dispType" />
<span id="advView" class="">
<span class=""><input name="OLD_VIEW" id="OLD_VIEW" type="checkbox" class="cb reset" defVal="unchecked"  />
<label for="OLD_VIEW">Old View</label></span>
</span>
<a class="resetAll" id="resetAll" >Reset form to defaults</a>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments object formatting help" id="formatHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp1">
These options control formatting of alignments in results pages. The
default is HTML, but other formats (including plain text) are available.
PSSM and PssmWithParameters are representations of Position Specific Scoring Matrices and are only available for PSI-BLAST.
The Advanced view option allows the database descriptions to be sorted by various indices in a table.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#format_object" target="helpWin" title="Additional alignments object formatting help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="odd paramSet">
<td class="hd"><label for="ALIGNMENT_VIEW">Alignment View</label></td>
<td>
<div class="fi">
<select name="ALIGNMENT_VIEW" id="ALIGNMENT_VIEW" defVal="Pairwise" class="reset">
<option value="Pairwise"  >Pairwise</option>
<option value="PairwiseWithIdentities"  >Pairwise with dots for identities</option>
<option value="QueryAnchored"  >Query-anchored with dots for identities</option>
<option value="QueryAnchoredNoIdentities"  >Query-anchored with letters for identities</option>
<option value="FlatQueryAnchored"  >Flat query-anchored with dots for identities</option>
<option value="FlatQueryAnchoredNoIdentities"  >Flat query-anchored with letters for identities</option>
<option value="Tabular"  >Hit Table</option>
</select>


<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments view options help" id="alnViewHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp2">
Choose how to view alignments.
The default "pairwise" view shows how each subject sequence aligns
individually to the query sequence. The "query-anchored" view shows how
all subject sequences align to the query sequence. For each view type,
you can choose to show "identities" (matching residues) as letters or
dots.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#alignment_view" target="helpWin" title="Additional alignments view options help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="paramSet">
<td class="hd"><label>Display</label></td>
<td class="cb">
<div class="fi">
<input name="SHOW_OVERVIEW" id="SHOW_OVERVIEW" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="SHOW_OVERVIEW">Graphical Overview</label>

<span id="shl" >
<input name="SHOW_LINKOUT" id="SHOW_LINKOUT" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="SHOW_LINKOUT">Linkout</label>
</span>
<span id="gts" >
<input name="GET_SEQUENCE" id="GET_SEQUENCE" type="checkbox" class="cb reset" defVal="checked" checked="checked" />
<label class="rb" for="GET_SEQUENCE">Sequence Retrieval</label>
</span>

<input name="NCBI_GI" id="NCBI_GI" type="checkbox" class="cb reset" defVal="unchecked"  />
<label class="rb" for="NCBI_GI">NCBI-gi</label>
<span id="scf" >
<input name="SHOW_CDS_FEATURE" id="SHOW_CDS_FEATURE" type="checkbox" class="cb reset blastn" defVal="unchecked"  />
<label for="SHOW_CDS_FEATURE" class="blastn">CDS feature</label>
</span>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments display options help" id="displayHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp3">
<li>Graphical Overview: Graphical Overview: Show graph of similar sequence regions aligned to  query.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_overview" target="helpWin" title="Graphical Overview help">more...</a>
</li>
<li>Database LinkOuts: Show links from matching sequences to entries in specialized NCBI databases.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_linkout" title="LinkOut help" target="helpWin" >more...</a>
</li>
<li>Sequence Retrieval: Show buttons to download matching sequences.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#get_sequence" title="Sequence Retrieval help" target="helpWin" >more...</a>
</li>
<li>NCBI-gi: Show NCBI gi identifiers.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#ncbi_gi" title="NCBI-gi help" target="helpWin" >more...</a>
</li>
<li>CDS feature: Show annotated coding region and translation.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#show_cds_feature" title="CDS feature help" target="helpWin" >more...</a>
</li></ul>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr class="paramSet odd xgl">
<td class="hd"><label>Masking</label></td>
<td>
<div class="fi">
<label for="MASK_CHAR"> Character: </label>
<select name="MASK_CHAR" id="MASK_CHAR"  class="reset" defVal="2">
<option value="0"  >X for protein, n for nucleotide</option>
<option value="2" selected="selected" >Lower Case</option>
</select>
<label for="MASK_COLOR"> Color:</label>
<select name="MASK_COLOR" id="MASK_COLOR" class="reset" defVal="1">
<option value="0"  >Black
</option>

<option value="1" selected="selected" >Grey
</option>

<option value="2"  >Red
</option>

</select>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Alignments masking help" id="maskingHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp4">
<li>Masking Character: Display masked (filtered) sequence regions as lower-case or as specific letters (N for nucleotide, P for protein).
</li>
<li>Masking Color: Display masked sequence regions in the given color.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr id="lr" class="paramSet xgl">
<td class="hd"><label>Limit results</label></td>
<td>
<div class="fi">
<label for="FRM_DESCRIPTIONS">Descriptions:</label>
<select name="DESCRIPTIONS" id="FRM_DESCRIPTIONS" class="reset" defVal="100">
<option value="0"      >0</option>
<option value="10"     >10</option>
<option value="50"     >50</option>
<option value="100"   selected="selected" >100</option>
<option value="250"    >250</option>
<option value="500"    >500</option>
<option value="1000"   >1000</option>
<option value="5000"   >5000</option>
<option value="10000"  >10000</option>
<option value="20000"  >20000</option>
</select>

<label for="FRM_NUM_OVERVIEW">Graphical overview:</label>
<select name="NUM_OVERVIEW" id="FRM_NUM_OVERVIEW" class="reset" defVal="100">
<option value="0"     >0</option>
<option value="10"    >10</option>
<option value="50"    >50</option>
<option value="100"  selected="selected" >100</option>
<option value="250"   >250</option>
<option value="500"  >500</option>
<option value="1000"  >1000</option>
</select>
<span id="frmAln">
<label for="FRM_ALIGNMENTS">Alignments:</label>
<select name="ALIGNMENTS" id="FRM_ALIGNMENTS" class="reset" defVal="100">
<option value="0"      >0</option>
<option value="10"     >10</option>
<option value="50"     >50</option>
<option value="100"   selected="selected" >100</option>
<option value="250"    >250</option>
<option value="500"    >500</option>
<option value="1000"   >1000</option>
<option value="5000"   >5000</option>
<option value="10000"  >10000</option>
<option value="20000"  >20000</option>
</select>
</span>
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit number of descriptions/alignments help" id="numHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp5">
<li>Descriptions: Show short descriptions for up to the given number of  sequences.</li>
<li>Alignments:  Show alignments for up to the given number of sequences, in order of statistical significance.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>

<tr class="paramSet odd xgl ">
<td class="hd"></td>
<td>
<div class="">
<label for="qorganism">Organism</label>
<span class="instr">Type common name, binomial, taxid, or group name. Only 20 top taxa will be shown.</span><br>
<input name="FORMAT_ORGANISM" size="55"  type="text" id="qorganism" value="" data-jigconfig="dictionary:'taxids_sg'" autocomplete="off" class="jig-ncbiautocomplete reset">
<input type="checkbox" name="FORMAT_ORG_EXCLUDE"  class="oExclR cb" id="orgExcl"/>
<input type="hidden" value = "1" name="FORMAT_NUM_ORG" id="numOrg" />
<label for="orgExcl" class="right">Exclude</label>
<a href="#" title="Add organism" id="addOrg"><img border="0" src="css/images/addOrg.jpg" id="addOrgIm"   alt="Add organism"  mouseovImg="css/images/addOrgOver.jpg" mouseoutImg="css/images/addOrg.jpg" mousedownImg="css/images/addOrgDown.jpg" mouseupImg="css/images/addOrgOver.jpg"  /></a>
<div id="orgs">

</div>
<div class="fi">
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by organism help" id="organismHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp6">
Show only sequences from the given organism.
</p>
</div><!-- ARIA -->
</div>
</div>
</td>
</tr>

<tr class="paramSet xgl ">
<td class="hd"></td>
<td>
<div class="fi">
<label for="FORMAT_EQ_TEXT">Entrez query:</label>
<input name="FORMAT_EQ_TEXT" id="FORMAT_EQ_TEXT" size="60" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by Entrez query help" id="entrezHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp7">
Show only those sequences that match the given Entrez query.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#limit_result" target="helpWin" title="Additional limit results by Entrez query help"  target="helpWin">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>


<tr class="paramSet odd xgl">
<td class="hd"></td>
<td>
<div class="fi">
<label for="EXPECT_LOW">Expect Min:</label> <input name="EXPECT_LOW" id="EXPECT_LOW" size="10" type="text" value="" class="reset"/>
<label for="EXPECT_HIGH">Expect Max:</label> <input name="EXPECT_HIGH" id="EXPECT_HIGH" size="10" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by expect value range help" id="expectHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp8">
Show only sequences with expect values in the given range.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#expect_range" target="helpWin" title="Additional limit results by expect value range help">more...</a>
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>
<tr class="paramSet xgl">
<td class="hd"></td>
<td>
 <div class="fi">
<label for="PERC_IDENT_LOW">Percent Identity Min:</label> <input name="PERC_IDENT_LOW" id="PERC_IDENT_LOW" size="10" type="text" value="" class="reset"/>
<label for="PERC_IDENT_HIGH">Percent Identity Max:</label> <input name="PERC_IDENT_HIGH" id="PERC_IDENT_HIGH" size="10" type="text" value="" class="reset" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="Limit results by percent identity range help" id="percIdentHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<p class="helpbox ui-ncbitoggler-slave" id="hlp10">
 Show only sequences with percent identity values in the given range.
</p>
</div><!-- ARIA -->
</div>
</td>
</tr>
<tr class="psiBlast odd paramSet xgl">
<td class="hd"><label>Format for</label></td>
<td>
<div class="fi">
<input name="RUN_PSIBLAST_FORM" id="RUN_PSIBLAST" type="checkbox" class="cb psiBlast"  />
<label class="rb psiBlast" for="RUN_PSIBLAST">PSI-BLAST</label>
<label for="I_THRESH">with inclusion threshold:</label>
<input name="I_THRESH" id="I_THRESH" size="10" type="text" value="" defVal="0.005" />
<a class="helplink  jig-ncbitoggler ui-ncbitoggler" title="PSI BLAST formatting help" id="psiHelp" href="#"><span class="ui-ncbitoggler-master-text"><span>[?]</span></span>
<span class="ui-icon ui-icon-triangle-1-e"></span></a>
<div class="ui-helper-reset" aria-live="assertive" >
<ul class="helpbox ui-ncbitoggler-slave" id="hlp9">
<li>Format for PSI-BLAST: The Position-Specific Iterated BLAST (PSI-BLAST) program performs iterative searches with a protein query,
in which sequences found in one round of search are used to build a custom score model for the next round.
<a href="http://www.ncbi.nlm.nih.gov/BLAST/blastcgihelp.shtml#psiblast" target="helpWin" title="Additional PSI BLAST formatting help">more...</a>
</li>
<li>Inclusion Threshold: This sets the statistical significance threshold for including a sequence in the model used
by PSI-BLAST to create the PSSM on the next iteration.</li>
</ul>
</div><!-- ARIA -->
</div>
</td>
</tr>
</table>
</dd>
</dl>

<input name="RID" value="" type="hidden" />
<input name="CDD_RID" value="" type="hidden" />
<input name="CDD_SEARCH_STATE" type="hidden" value="" />

<input name="STEP_NUMBER" value="" id="stepNumber" type="hidden" />
<input name="CMD" value="Get" type="hidden" />
<input name="FORMAT_EQ_OP" value="AND" type="hidden" />
<input name="RESULTS_PAGE_TARGET" type="hidden" id="resPageTarget" value="Blast_Results_for_1101632636" />
<input name="QUERY_INFO" type="hidden" value="" />
<input name="ENTREZ_QUERY" type="hidden" value="" />
<input name="QUERY_INDEX" type="hidden" value="0"/>
<input name="NUM_QUERIES" type="hidden" value="1"/>
<input name="CONFIG_DESCR" type="hidden" value="2,3,4,5,6,7,8" />

<!-- Those params are set in the template (blastn.dat, blastp.dat etc. -->
<input name="BLAST_PROGRAMS" type="hidden" value=""/>
<input name="PAGE" type="hidden" value=""/>
<input name="PROGRAM" type="hidden" value=""/>
<input name="MEGABLAST" type="hidden" value="" />
<input name="RUN_PSIBLAST" type="hidden" value="" />
<input name="BLAST_SPEC" id="blastSpec" type="hidden" value=""/>


<input name="QUERY" type="hidden" value=""/>
<input name="JOB_TITLE" type="hidden" value=""/>
<input name="QUERY_TO" type="hidden" value=""/>
<input name="QUERY_FROM" type="hidden" value=""/>
<input name="EQ_TEXT" type="hidden" value=""/>
<input name="ORGN" type="hidden" value=""/>
<input name="EQ_MENU" type="hidden" value=""/>
<input name="ORG_EXCLUDE" type="hidden" value=""/>
<input name="PHI_PATTERN" type="hidden" value=""/>
<input name="EXPECT" type="hidden" value=""/>
<input name="DATABASE" type="hidden" value=""/>
<input name="DB_GROUP" type="hidden" value=""/>
<input name="SUBGROUP_NAME" type="hidden" value=""/>

<input name="GENETIC_CODE" type="hidden" value=""/>
<input name="WORD_SIZE" type="hidden" value=""/>
<input name="MATCH_SCORES" type="hidden" value=""/>
<input name="MATRIX_NAME" type="hidden" value=""/>
<input name="GAPCOSTS" type="hidden" value=""/>
<input name="MAX_NUM_SEQ" id="maxNumSeq" type="hidden" value=""/>
<input name="COMPOSITION_BASED_STATISTICS" type="hidden" value=""/>
<input name="NEWWIN" type="hidden" value=""/>
<input name="SHORT_QUERY_ADJUST" type="hidden" value=""/>
<input name="FILTER" type="hidden" value=""/>
<input name="REPEATS" type="hidden" value=""/>
<input name="ID_FOR_PSSM" type="hidden" value=""/>
<input name="EXCLUDE_MODELS" type="hidden" value=""/>
<input name="EXCLUDE_SEQ_UNCULT" type="hidden" value=""/>
<input name="NUM_ORG" type="hidden" value = "1" />

<!-- PSSM -->
<input name="LCASE_MASK" type="hidden" value=""/>
<input name="TEMPLATE_TYPE" type="hidden" value=""/>
<input name="TEMPLATE_LENGTH" type="hidden" value=""/>
<input name="I_THRESH" type="hidden" value=""/>
<input name="PSI_PSEUDOCOUNT" type="hidden" value=""/>
<input name="DI_THRESH" type="hidden" id="diThresh" value=""/>
<input name="HSP_RANGE_MAX" type="hidden" value=""/>



<input name="ADJUSTED_FOR_SHORT_QUERY" type="hidden" value=""/>
<input name="MIXED_QUERIES" type="hidden" value=""/>
<input name="MIXED_DATABASE" id="mixedDb" type="hidden" value=""/>
<input name="BUILD_NAME"  type="hidden" value=""/>
<input name="ORG_DBS"  type="hidden" value=""/>

<!--QBlastInfoBegin
    RID =
    RTOE =
QBlastInfoEnd
-->
</form>

				</div><!-- /#content -->

        </div><!-- /#content-wrap -->


<div id="footer">
   <div id="rgs">BLAST is a registered trademark of the National Library of Medicine.</div>
   <p id="orgns">
      <a href="http://www.ncbi.nlm.nih.gov/" title="National Center for Biotechnology Information">NCBI</a> |
      <a href="http://www.nlm.nih.gov/" title="National Library of Medicine">NLM</a> |
      <a href="http://www.nih.gov/" title="National Institutes of Health">NIH</a> |
      <a href="http://www.hhs.gov/" title="US Department of Health and Human Services">DHHS</a>
   </p>

   <p>
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html'
      title='NCBI intellectual property statement'>Copyright</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html#disclaimer'
      title='About liability, endorsements, external links, pop-up advertisements'>Disclaimer</a> |
      <a href='http://www.nlm.nih.gov/privacy.html'
      title='NLM privacy policy'>Privacy</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/accessibility.html'
      title='About using NCBI resources with assistive technology'>Accessibility</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/glance/contact_info.html'
      title='How to get help, submit data, or provide feedback'>Contact</a> |
      <a href='mailto:blast-help@ncbi.nlm.nih.gov'
      title='How to get help, submit data, or provide feedback'>Send feedback</a>
   </p>
</div>
   </div><!--/#wrap-->

<script type="text/javascript" src="http://blast.ncbi.nlm.nih.gov/portal/portal3rc.fcgi/rlib/js/InstrumentOmnitureBaseJS/InstrumentNCBIBaseJS/InstrumentPageStarterJS.js"></script>
</body>

</html>


`,
			"",
			ErrMissingRid,
			0,
		},
	} {
		var r Rid
		err := r.unmarshal(strings.NewReader(t.retval))
		to := time.Now()
		c.Check(int((10*time.Microsecond + r.TimeOfExecution()).Seconds()), check.Equals, int(t.wait.Seconds()), check.Commentf("Test: %d", i))
		c.Check(err, check.Equals, t.err, check.Commentf("Test: %d", i))
		c.Check(r.rid, check.Equals, t.rid, check.Commentf("Test: %d", i))
		c.Check(r.delay, check.NotNil, check.Commentf("Test: %d", i))
		select {
		case <-r.Ready():
		case <-time.After(t.wait + time.Second):
			c.Fatalf("Waited too long on test %d.", t.wait)
		}
		d := time.Now().Sub(to)
		c.Check(int(d.Seconds()), check.Equals, int(t.wait.Seconds()), check.Commentf("Test: %d", i))
	}
}

func (s *S) TestParseSearchInfo(c *check.C) {
	for i, t := range []struct {
		retval   string
		status   string
		haveHits bool
		err      error
	}{
		{
			`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="jig" content="ncbitoggler"/>
<meta name="ncbitoggler" content="animation:'none'"/>
<title>NCBI Blast:</title>
<script type="text/javascript" src="http://www.ncbi.nlm.nih.gov/core/jig/1.11/js/jig.min.js"></script>
<link rel="stylesheet" type="text/css" href="css/main.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/blastRes.css" media="screen" />
<link rel="stylesheet" type="text/css" href="css/print.css" media="print" />
<!--[if lte IE 6]>
<link rel="stylesheet" type="text/css" href="css/ie6_or_less.css" />
<![endif]-->
<script type="text/javascript" src="js/utils.js"></script>
<script type="text/javascript" src="js/results.js"></script>
</head>

<body id="type-a" class="noToggleCheck" >
<div id="wrap">
			<div id="header">
		<div id="site-name"><a id="logolink" href="http://www.ncbi.nlm.nih.gov" title="NCBI Home Page"><img src="css/images/helix.gif" alt="NCBI Logo" title="Link to NCBI Home Page" /></a>BLAST <span id="trdm"> &reg;</span><h1 class="desc">Basic Local Alignment Search Tool</h1>
		</div>
		<div id="search">

<div>
<script language="JavaScript" type="text/javascript"><!--
// --></script><table class="medium1" style="border:2px solid #336699;" cellpadding="2" cellspacing="0" id="myncbi_off"><tr><td
bgcolor="#336699" align="left"><a href="http://www.ncbi.nlm.nih.gov/myncbi/?"><font color="#FFFFFF"><b>My NCBI</b></font></a></td><td
bgcolor="#336699" align="right"><a href="http://www.ncbi.nlm.nih.gov/books/NBK3842/" title="My NCBI help"><img border="0"
src="http://www.ncbi.nlm.nih.gov/corehtml/query/MyNCBI/myncbihelpicon.gif" alt="My NCBI help" /></a></td></tr><tr><td colspan="2" nowrap="nowrap"><a
href="http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DGet%26FORMAT%5FOBJECT%3DSearchInfo%26OLD%5FBLAST%3Dfalse%26RID%3DJXM8KY9G01R%26email%3Ddan%2Ekortschak%2540adelaide%2Eedu%2Eau%26tool%3Dbiogo%2Encbi%252Fblast%2Dtestsuite" title="Click to sign in"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DGet%26FORMAT%5FOBJECT%3DSearchInfo%26OLD%5FBLAST%3Dfalse%26RID%3DJXM8KY9G01R%26email%3Ddan%2Ekortschak%2540adelaide%2Eedu%2Eau%26tool%3Dbiogo%2Encbi%252Fblast%2Dtestsuite');return false;">[Sign In]</a> <a
href="http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DGet%26FORMAT%5FOBJECT%3DSearchInfo%26OLD%5FBLAST%3Dfalse%26RID%3DJXM8KY9G01R%26email%3Ddan%2Ekortschak%2540adelaide%2Eedu%2Eau%26tool%3Dbiogo%2Encbi%252Fblast%2Dtestsuite" title="Click to register for an account"
onclick="MyNCBI_auto_submit('http://www.ncbi.nlm.nih.gov/account/register/?back_url=http%3A%2F%2Fwww%2Encbi%2Enlm%2Enih%2Egov%2Fblast%2FBlast%2Ecgi%3FCMD%3DGet%26FORMAT%5FOBJECT%3DSearchInfo%26OLD%5FBLAST%3Dfalse%26RID%3DJXM8KY9G01R%26email%3Ddan%2Ekortschak%2540adelaide%2Eedu%2Eau%26tool%3Dbiogo%2Encbi%252Fblast%2Dtestsuite');return false;">[Register]</a></td></tr></table></div>
		</div>
		<a class="skp" href="#content-wrap">Jump to Page Content</a>
		<ul id="nav">
                <li  class="first "><a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastHome" title="BLAST Home">Home</a></li>
                <li  class="recent "><a href="Blast.cgi?CMD=GetSaved&amp;RECENT_RESULTS=on" title="Unexpired BLAST jobs">Recent Results</a></li>
                <li  class="saved "><a href="Blast.cgi?CMD=GetSaved" title="Saved sets of BLAST search parameters">Saved Strategies</a></li>
                <li  class= "last documentation "> <a href="Blast.cgi?CMD=Web&amp;PAGE_TYPE=BlastDocs" title="BLAST documentation">Help</a></li>
                </ul>
    </div>

        <div id="content-wrap">

                <div id="breadcrumb" class="inlineDiv">
                   <a href="http://www.ncbi.nlm.nih.gov/">NCBI</a>/
                   <a href="Blast.cgi?CMD=Web&PAGE_TYPE=BlastHome">BLAST</a>/
                   <a href="Blast.cgi?PAGE=Nucleotides&PROGRAM=blastn&BLAST_PROGRAMS=megaBlast&PAGE_TYPE=BlastSearch&SHOW_DEFAULTS=on&BLAST_SPEC=">blastn suite</a>/
                   <strong>Formatting Results - JXM8KY9G01R</strong>
                </div>
                <div class="inlineDiv resHeader">
				   <a  id="frmPage"  class="READY" href="#" submitForm="reformat">[Formatting options] </a>
                </div>
                <h3 id="jtitle" >Job Title: </h3>

                <div id="content">
                <!--<ul id="msg" class="msg"><li class=""><p class=""></p><p class=""></p><p class=""></p></ul> -->
                <ul id="msg" class="msg"><li class=""></li></ul>
                <p><!--
                QBlastInfoBegin
	                Status=READY
                QBlastInfoEnd
                --></p>
                <!--
QBlastInfoBegin
	ThereAreHits=yes
QBlastInfoEnd
--><p>

                <SCRIPT LANGUAGE="JavaScript"><!--
                    var tm = "";
                    if (tm != "") {
                        setTimeout('document.forms[0].submit();',tm);
                    }
                //--></SCRIPT>
                <table id="statInfo" class="READY">
                <tr><td>Request ID</td><td> <b>JXM8KY9G01R</b></td></tr>
                <tr class="odd"><td>Status</td><td>Searching</td></tr>
                <tr><td>Submitted at</td><td></td></tr>
                <tr class="odd"><td>Current time</td><td></td></tr>
                <tr><td>Time since submission</td><td></td></tr>
                </table>
                <p class="READY">This page will be automatically updated in <b></b> seconds</p>
                <form action="Blast.cgi" enctype="application/x-www-form-urlencoded" method="POST" id="results">
                <input name="FORMAT_OBJECT" type="hidden" value="SearchInfo"><input name="OLD_BLAST" type="hidden" value="false"><input name="RID" type="hidden" value="JXM8KY9G01R"><input name="SEARCH_DB_STATUS" type="hidden" value="43"><input name="USER_TYPE" type="hidden" value="2"><input name="_PGR" type="hidden" value="0"><input name="email" type="hidden" value="dan.kortschak@adelaide.edu.au"><input name="tool" type="hidden" value="biogo.ncbi/blast-testsuite">
                <input name="_PGR" type="hidden" value="0" >
                <input name="CMD" type="hidden" value="Get">

                </form>

				</div><!-- /#content -->
				<form action="Blast.cgi" enctype="application/x-www-form-urlencoded"  method="post" name="reformat" id="reformat">
				   <input name="QUERY_INFO" type="hidden" value="" />
				   <input name="ENTREZ_QUERY" type="hidden" value="" />
                   <input name="CDD_RID" type="hidden" value="" />
                   <input name="CDD_SEARCH_STATE" type="hidden" value="" />
                   <input name="RID" type="hidden" value="JXM8KY9G01R" />
				   <input name="STEP_NUMBER" type="hidden" value="" />
				   <input name="CMD" type="hidden" value="Web"/>
				   <input NAME="PAGE_TYPE" type="hidden"  value="BlastFormatting"/>

				   <!-- TO DO: test all of those changes -->
				   <!-- Psi blast params  PSI_BLAST_PARAMS - commented- using forms[0] from fromatter> -->
				   <!-- Current Formatting options FORMATTING_OPTIONS- commented- using forms[0] from fromatter> -->
				   <!-- Current Search options CURR_SAVED_OPTIONS - commented- using forms[0] from fromatter> -->
                 </form>
        </div><!-- /#content-wrap -->


<div id="footer">
   <div id="rgs">BLAST is a registered trademark of the National Library of Medicine.</div>
   <p id="orgns">
      <a href="http://www.ncbi.nlm.nih.gov/" title="National Center for Biotechnology Information">NCBI</a> |
      <a href="http://www.nlm.nih.gov/" title="National Library of Medicine">NLM</a> |
      <a href="http://www.nih.gov/" title="National Institutes of Health">NIH</a> |
      <a href="http://www.hhs.gov/" title="US Department of Health and Human Services">DHHS</a>
   </p>

   <p>
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html'
      title='NCBI intellectual property statement'>Copyright</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/disclaimer.html#disclaimer'
      title='About liability, endorsements, external links, pop-up advertisements'>Disclaimer</a> |
      <a href='http://www.nlm.nih.gov/privacy.html'
      title='NLM privacy policy'>Privacy</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/accessibility.html'
      title='About using NCBI resources with assistive technology'>Accessibility</a> |
      <a href='http://www.ncbi.nlm.nih.gov/About/glance/contact_info.html'
      title='How to get help, submit data, or provide feedback'>Contact</a> |
      <a href='mailto:blast-help@ncbi.nlm.nih.gov'
      title='How to get help, submit data, or provide feedback'>Send feedback</a>
   </p>
</div>
   </div><!--/#wrap-->
</body>

</html>

`,
			"READY",
			true,
			nil,
		},
	} {
		var s SearchInfo
		err := s.unmarshal(strings.NewReader(t.retval))
		c.Check(err, check.Equals, t.err, check.Commentf("Test: %d", i))
		c.Check(s.Status, check.Equals, t.status, check.Commentf("Test: %d", i))
		c.Check(s.HaveHits, check.Equals, t.haveHits, check.Commentf("Test: %d", i))
	}
}
