
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind523_0 { KIND_EMPTY523_0 = 0, KIND_FULL523_0 = 14 };

union Cell523_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind523_0(union Cell523_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL523_0 : KIND_EMPTY523_0;
}
struct Node523_1 {
    int value;
    struct Node523_1 *next;
};

int length523_1(const struct Node523_1 *head) {
    int count = 0;
    const struct Node523_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
#if defined(__STDC_VERSION__)
int variant523_2(int value, ...) {
    return value + 28;
}
#else
int variant523_2(int value) {
    return value + 28;
}
#endif
#if defined(__STDC_VERSION__)
int variant523_3(int value, ...) {
    return value + 35;
}
#else
int variant523_3(int value) {
    return value + 35;
}
#endif
int gcd523_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale523_4(int value) {
    return gcd523_4(value, 54) * 42;
}
int gcd523_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale523_5(int value) {
    return gcd523_5(value, 65) * 49;
}
#if defined(__STDC_VERSION__)
int variant523_6(int value, ...) {
    return value + 56;
}
#else
int variant523_6(int value) {
    return value + 56;
}
#endif
