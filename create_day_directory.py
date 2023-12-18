import os
import sys
import shutil


BASE_URL = "solutions/"


def create_directory(day):
    dir_name = rf"{BASE_URL}day{day_number}"
    os.makedirs(dir_name, exist_ok=True)
    return dir_name


def copy_template(dir_name, day_num):
    template_file = "template.py"
    destination_file = os.path.join(dir_name, f"solution_day{day_num}.py")
    shutil.copyfile(template_file, destination_file)
    return destination_file


def create_empty_txt_file(dir_name, file_name):
    file_path = os.path.join(dir_name, file_name)
    with open(file_path, "w"):
        pass  # Create an empty file


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python create_day_directory <day_number>")
        sys.exit(1)

    day_number = sys.argv[1]

    directory_name = create_directory(day_number)

    if not os.listdir(directory_name):
        copy_template(directory_name, day_number)
        create_empty_txt_file(
            directory_name, file_name=f"puzzle_input_day{day_number}.txt"
        )
