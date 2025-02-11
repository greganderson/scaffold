import ast
import unittest

import asttest
import {{ .FileName }} as student


class Test{{ .TestClassName }}(asttest.ASTTest):
    def setUp(self):
        super().setUp("{{ .FileName }}.py")
        class_name = "{{ .StudentClass }}"
        
        student_classes: list[ast.ClassDef] = self.find_all(ast.ClassDef)
        matches: list = [c for c in student_classes if c.name == class_name]
        self.assertEqual(len(matches), 1, f"The {class_name} class is missing. Did you forget to make it?")

        # Get methods
        self.methods = {m.name: m for m in matches[0].body}

    def test_provided_code(self):
        script_section = 'if __name__ == "__main__":\n    main()'

        self.assertIn(
            script_section,
            self.file,
            "You should not edit/remove the instructor provided section at the bottom of the file.",
        )

        self.assertFalse(
            len(self.find_function_calls("main")) == 0,
            "You should not touch the call to the `main` function.",
        )
        self.assertTrue(
            len(self.find_function_calls("main")) == 1,
            "You should not be calling `main` in your solution code. It should only be in the section at the bottom. Which you should not touch.",
        )
        
    def test_constructor_definition(self):
        target_method = "__init__"
        self.assertTrue(target_method in self.methods, f"The {target_method} method is missing.")
        self.validate_method_param_type_hints(self.methods[target_method], [str, list[str], int])
        
    def test_instance_variables(self):
        required_instance_variables: list[str] = []
        # TODO:
        ####### CHANGE ME #######
        class_instance = None
        for var in required_instance_variables:
            self.assertIn(var, vars(class_instance))


if __name__ == "__main__":
    unittest.main()
