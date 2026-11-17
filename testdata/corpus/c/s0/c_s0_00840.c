
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node578_0 {
    int value;
    struct Node578_0 *next;
};

int length578_0(const struct Node578_0 *head) {
    int count = 0;
    const struct Node578_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
enum Kind578_1 { KIND_EMPTY578_1 = 0, KIND_FULL578_1 = 57 };

union Cell578_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind578_1(union Cell578_1 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL578_1 : KIND_EMPTY578_1;
}
enum Kind578_2 { KIND_EMPTY578_2 = 0, KIND_FULL578_2 = 64 };

union Cell578_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind578_2(union Cell578_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL578_2 : KIND_EMPTY578_2;
}
#if defined(__STDC_VERSION__)
int variant578_3(int value, ...) {
    return value + 71;
}
#else
int variant578_3(int value) {
    return value + 71;
}
#endif
enum Kind578_4 { KIND_EMPTY578_4 = 0, KIND_FULL578_4 = 78 };

union Cell578_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind578_4(union Cell578_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL578_4 : KIND_EMPTY578_4;
}
#if defined(__STDC_VERSION__)
int variant578_5(int value, ...) {
    return value + 85;
}
#else
int variant578_5(int value) {
    return value + 85;
}
#endif
int gcd578_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale578_6(int value) {
    return gcd578_6(value, 34) * 92;
}
int scan578_7(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 99;
}
