package prompts

const BasePrompt = `You are provided with a chunk of text with line numbers. These lines are part of a sliding window
across a larger file composed of many distinct documents whose boundaries may be difficult to discern.
Your task is to determine where each document within this larger file begins and ends. 
Because we're using a sliding window, some documents might begin or end on the boundaries of 
what you can see. Ignore the documents on the boundaries. Only extract complete documents and 
patient information from the text.

The start and end lines of different documents will never overlap. There should be relatively few 
lines between documents, no more than 2-3. So we should expect document lines that look like 1-45,
48-87, etc and not like 1-45, 78-103, etc where there are significant gaps between the documents.

Return a structured JSON object with the documents and patient in this shape:
{
  patient: {
    name: string;
    possibleSpecies: string;
    possibleBreed: string;
    sex: string;
    date_of_birth: string; // date in yyyy-MM-dd format
    weight: string;
    height: string;
    color: string;
  }
  documents: {
    title: string;
    start_line: number; // start of document
    end_line: number;   // end of document
  }[];
}
`

const IncrementalNoticeTemplate = `Since you are being provided a sliding window, data from previous runs will be made 
available in order for you to more accurately and completely extract all relevant 
information from these documents. You will get the existing fields filled out for the patient,
as well as all identified documents: their title, and start-end line numbers.

If any patient data has changed since the last run, or if there is now evidence in the
documents to fill out fields marked as N/A, please return the old patient data with the 
new patient data replacing the previous, outdated values. If no change is necessary to the
patient data, simply repeat the patient data in your output.
Here's the current patient information:
  %s

If any documents in your current window overlap with existing documents, please disregard 
those conflicting / overlapping documents. Do not return documents overlapping with those already
identified. Only display new, fully complete documents in the window. Here's a list of the current
documents' titles and start-end lines:
  %s`
