
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2127_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2127_0(int value) {
    return gcd2127_0(value, 14) * 11;
}
#if defined(__STDC_VERSION__)
int variant2127_1(int value, ...) {
    return value + 18;
}
#else
int variant2127_1(int value) {
    return value + 18;
}
#endif
int gcd2127_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2127_2(int value) {
    return gcd2127_2(value, 36) * 25;
}
#if defined(__STDC_VERSION__)
int variant2127_3(int value, ...) {
    return value + 32;
}
#else
int variant2127_3(int value) {
    return value + 32;
}
#endif
struct Node2127_4 {
    int value;
    struct Node2127_4 *next;
};

int length2127_4(const struct Node2127_4 *head) {
    int count = 0;
    const struct Node2127_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
enum Kind2127_5 { KIND_EMPTY2127_5 = 0, KIND_FULL2127_5 = 46 };

union Cell2127_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2127_5(union Cell2127_5 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2127_5 : KIND_EMPTY2127_5;
}
enum Kind2127_6 { KIND_EMPTY2127_6 = 0, KIND_FULL2127_6 = 53 };

union Cell2127_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2127_6(union Cell2127_6 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2127_6 : KIND_EMPTY2127_6;
}
