# Sorting input data

# Task 1

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
	<li><strong>-i file-name</strong>, <em>Use a file with the name file-name as an input</em>.</li>
	<li><strong>-o file-name</strong>, <em>Use a file with the name file-name as an output</em>.</li>
	<li><strong>-h The first</strong>, <em>line is a header that must be ignored during sorting but included in the output</em>.</li>
	<li><strong>-f N</strong>, <em>Sort input lines by value number N</em>.</li>
	<li><strong>-r</strong>, <em>Sort input lines in reverse order</em>.</li>
		</ul>
	</li>	
			
<li> Add the ability to use a second algorithm for sorting - the Tree Sort algorithm. Accordingly, add one more option -a with possible values 1 or 2, which chooses currently implemented algorithm or Tree Sort algorithm to use. By default, the application uses the currently implemented algorithm.</li>
</ol>	
	



# Task 2

**“CSV Concurrent Sorter” is a CLI application that allows sorting of its input presented as CSV-text**

## **Technical details** 

### Using the “CSV Sorter” from the Task 1, extend it with the following required features:

1. The application has additional option **-d dir-name** that specifies a directory where it must read input files from. All files in the directory must have the same format. The output stays the same, it is a one file or a standard output with sorted content from all input files. 
2. Processing must be implemented concurrently based on pipeline. The pipeline includes three stages:
	* Directory Reading:
		* read the directory content and all its subdirectories
		* send all found file names to a channel for the future processing
	* File Reading – for each file name received from the input channel
	  * read the file from the disk
	  * send its content line by line to the output channel. 
	* Sorting:
	  * read lines from the input channel
	  * collect them into a buffer
	  * when the input channel is exhausted, sort lines in the buffer
	  *	write the result to the output file (if option **-o** is set) or to the standard output
3. The application must print an error message, if two options **-d** and **-i** are set at the same time. 
4. If option **-i** is set the application must read only the file defined by this option and then must use the pipeline to process file content

### Optional features (not required but appreciated):
1. Add signal processing that allows to gracefully stop the application when the user interrupts it pressing Ctrl-C. The interrupted application must write the current result. To implement it and stop the pipeline immediately, use stop-channel. 


**[Contact with me](https://t.me/kirilodm67)**

