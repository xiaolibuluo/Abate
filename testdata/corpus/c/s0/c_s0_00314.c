
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node216_0 {
    int value;
    struct Node216_0 *next;
};

int length216_0(const struct Node216_0 *head) {
    int count = 0;
    const struct Node216_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
#if defined(__STDC_VERSION__)
int variant216_1(int value, ...) {
    return value + 7;
}
#else
int variant216_1(int value) {
    return value + 7;
}
#endif
int gcd216_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale216_2(int value) {
    return gcd216_2(value, 18) * 14;
}
enum Kind216_3 { KIND_EMPTY216_3 = 0, KIND_FULL216_3 = 21 };

union Cell216_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind216_3(union Cell216_3 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL216_3 : KIND_EMPTY216_3;
}
