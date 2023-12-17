import os
import sys
import shutil


BASE_URL = "../"


def create_directory(day):
    dir_name = rf"{BASE_URL}day{day_number}"
    os.makedirs(dir_name, exist_ok=True)
    return dir_name


def copy_template(directory_name):
    template_file = "template.py"
    destination_file = os.path.join(directory_name, "solver_script.py")
    shutil.copyfile(template_file, destination_file)
    return destination_file


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python create_day_directory <day_number>")
        sys.exit(1)

    day_number = sys.argv[1]

    directory_name = create_directory(day_number)

    # copy_template(directory_name)
    # file_name = f"puzzle_input_day{day_number}.txt"
    # file_path = os.path.join(directory_name, file_name)
