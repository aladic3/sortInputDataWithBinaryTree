# sortInputData

**“CSV Sorter” is a CLI application that allows sorting of its input presented as CSV-text.**

## Technical details
### Required features:
<ul>
	<li>The application runs as a CLI application.</li> 
	<li>It reads STDIN line by line. The end of the input is an empty line.</li>
	<li>Each line is a list of comma-separated values (CSV). Each value is considered as a piece of text. The number of values is the same in each line.</li>
	<li>The application sorts all lines alphabetically by the first value in each line.</li>
	<li>The application prints the result immediately, when the user ends to enter input text (presses <Enter> at a new line).</li>
</ul>

### Optional features (not required but appreciated):
<ol>
	<li>The application supports options (Option, usage Meaning):
		
<ul>
	<li><strong><strong>-i file-name</strong></strong>, <em>Use a file with the name file-name as an input</em>.</li>
	<li><strong>-o file-name</strong>, <em>Use a file with the name file-name as an output</em>.</li>
	<li><strong>-h The first</strong>, <em>line is a header that must be ignored during sorting but included in the output</em>.</li>
	<li><strong>-f N</strong>, <em>Sort input lines by value number N</em>.</li>
	<li><strong>-r</strong>, <em>Sort input lines in reverse order</em>.</li>
		</ul>
	</li>	
			
<li> Add the ability to use a second algorithm for sorting - the Tree Sort algorithm. Accordingly, add one more option -a with possible values 1 or 2, which chooses currently implemented algorithm or Tree Sort algorithm to use. By default, the application uses the currently implemented algorithm.</li>
</ol>	
	

**[Contact Telegram](https://t.me/kirilodm67)**
