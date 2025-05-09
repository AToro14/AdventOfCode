from .dayXX import solve_part1, solve_part2

EXAMPLE_ARG_1 = 2
EXAMPLE_ARG_2 = 5

def test_dayXX_part1():
    """Test Day XX Part 1 with example inputs."""
    assert solve_part1(EXAMPLE_ARG_1, EXAMPLE_ARG_2) == 7
    
def test_dayXX_part2():
    """Test Day XX Part 2 with example inputs."""
    assert solve_part2(EXAMPLE_ARG_1, EXAMPLE_ARG_2) == 2
    
def test_dayXX_part3():
    """Test Day XX Part 3 with example inputs."""
    assert solve_part3(EXAMPLE_ARG_1, EXAMPLE_ARG_2) == 7
