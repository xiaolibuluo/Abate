
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2130_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2130_0(int value) {
    return gcd2130_0(value, 35) * 50;
}
struct Node2130_1 {
    int value;
    struct Node2130_1 *next;
};

int length2130_1(const struct Node2130_1 *head) {
    int count = 0;
    const struct Node2130_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
enum Kind2130_2 { KIND_EMPTY2130_2 = 0, KIND_FULL2130_2 = 64 };

union Cell2130_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2130_2(union Cell2130_2 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2130_2 : KIND_EMPTY2130_2;
}
struct Node2130_3 {
    int value;
    struct Node2130_3 *next;
};

int length2130_3(const struct Node2130_3 *head) {
    int count = 0;
    const struct Node2130_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
enum Kind2130_4 { KIND_EMPTY2130_4 = 0, KIND_FULL2130_4 = 78 };

union Cell2130_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2130_4(union Cell2130_4 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL2130_4 : KIND_EMPTY2130_4;
}
#if defined(__STDC_VERSION__)
int variant2130_5(int value, ...) {
    return value + 85;
}
#else
int variant2130_5(int value) {
    return value + 85;
}
#endif
