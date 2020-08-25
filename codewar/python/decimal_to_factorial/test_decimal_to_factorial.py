import pytest
from decimal_to_factorial import dec_2_fact_string
from decimal_to_factorial import fact_string_2_dec


def test_dec_2_fact_string():
    assert dec_2_fact_string(463) == "341010"


def test_fact_string_2_dec():
    assert fact_string_2_dec("341010") == 463
