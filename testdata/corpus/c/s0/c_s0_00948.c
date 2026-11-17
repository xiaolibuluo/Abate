
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant652_0(int value, ...) {
    return value + 42;
}
#else
int variant652_0(int value) {
    return value + 42;
}
#endif
enum Kind652_1 { KIND_EMPTY652_1 = 0, KIND_FULL652_1 = 49 };

union Cell652_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind652_1(union Cell652_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL652_1 : KIND_EMPTY652_1;
}
int gcd652_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale652_2(int value) {
    return gcd652_2(value, 20) * 56;
}
struct Node652_3 {
    int value;
    struct Node652_3 *next;
};

int length652_3(const struct Node652_3 *head) {
    int count = 0;
    const struct Node652_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
