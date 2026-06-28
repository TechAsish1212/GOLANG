truct Embedding---->Struct embedding is a powerful feature in Go that allows you to create complex data structures by embedding one struct within another. This promotes code reuse and helps to organize related data together. When you embed a struct, the fields and methods of the embedded struct become part of the outer struct, allowing you to access them directly without needing to reference the embedded struct.




This is struct embedding. By embedding the `Address` struct within the `Employee` struct, we can directly access the fields of `Address` from an instance of `Employee` without needing to reference the embedded struct. This promotes code reuse and allows us to organize related data together. For example, we can create an instance of `Employee` and set the city, country, pin code, and district directly on the employee instance, as if those fields were part of the `Employee` struct itself.