# Code Story

A tool to parse the Git history of a Go project and record information
about how various symbols evolve over time.

## Example

```
$ ./codestory | jq
{
  "Commits": [
    {
      "Files": [
        {
          "Functions": [
            {
              "Arity": 0,
              "Complexity": 0,
              "Lines": 0,
              "Name": "main",
              "Panics": 0
            }
          ],
          "Path": "codestory.go"
        },
        {
          "Functions": [
            {
              "Arity": 1,
              "Complexity": 0,
              "Lines": 0,
              "Name": "processFile",
              "Panics": 0
            },
            {
              "Arity": 1,
              "Complexity": 0,
              "Lines": 0,
              "Name": "processFunction",
              "Panics": 0
            },
            {
              "Arity": 1,
              "Complexity": 0,
              "Lines": 0,
              "Name": "processCommit",
              "Panics": 0
            }
          ],
          "Path": "walker.go"
        }
      ],
      "Hash": "c23e1ad69f96aa2ee57b4d46e326bb0900db416a",
      "Message": "Refactor into something reasonable\n"
    },
    {
      "Files": [
        {
          "Functions": [
            {
              "Arity": 0,
              "Complexity": 0,
              "Lines": 0,
              "Name": "main",
              "Panics": 0
            }
          ],
          "Path": "codestory.go"
        }
      ],
      "Hash": "b6d9cb64992e1b53ac56d5dad99acaec6fa32b04",
      "Message": "Initial commit\n"
    }
  ]
}
```

