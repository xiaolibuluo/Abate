
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2558_0 {
    int value;
    struct Node2558_0 *next;
};

int length2558_0(const struct Node2558_0 *head) {
    int count = 0;
    const struct Node2558_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
enum Kind2558_1 { KIND_EMPTY2558_1 = 0, KIND_FULL2558_1 = 92 };

union Cell2558_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2558_1(union Cell2558_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL2558_1 : KIND_EMPTY2558_1;
}
int gcd2558_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2558_2(int value) {
    return gcd2558_2(value, 64) * 99;
}
#if defined(__STDC_VERSION__)
int variant2558_3(int value, ...) {
    return value + 9;
}
#else
int variant2558_3(int value) {
    return value + 9;
}
#endif
struct Node2558_4 {
    int value;
    struct Node2558_4 *next;
};

int length2558_4(const struct Node2558_4 *head) {
    int count = 0;
    const struct Node2558_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
