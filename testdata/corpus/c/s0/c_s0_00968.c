
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd665_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale665_0(int value) {
    return gcd665_0(value, 28) * 17;
}
#if defined(__STDC_VERSION__)
int variant665_1(int value, ...) {
    return value + 24;
}
#else
int variant665_1(int value) {
    return value + 24;
}
#endif
enum Kind665_2 { KIND_EMPTY665_2 = 0, KIND_FULL665_2 = 31 };

union Cell665_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind665_2(union Cell665_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL665_2 : KIND_EMPTY665_2;
}
#if defined(__STDC_VERSION__)
int variant665_3(int value, ...) {
    return value + 38;
}
#else
int variant665_3(int value) {
    return value + 38;
}
#endif
struct Node665_4 {
    int value;
    struct Node665_4 *next;
};

int length665_4(const struct Node665_4 *head) {
    int count = 0;
    const struct Node665_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
struct Node665_5 {
    int value;
    struct Node665_5 *next;
};

int length665_5(const struct Node665_5 *head) {
    int count = 0;
    const struct Node665_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
