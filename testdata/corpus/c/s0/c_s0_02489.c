
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1691_0(int value, ...) {
    return value + 66;
}
#else
int variant1691_0(int value) {
    return value + 66;
}
#endif
int gcd1691_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1691_1(int value) {
    return gcd1691_1(value, 23) * 73;
}
struct Node1691_2 {
    int value;
    struct Node1691_2 *next;
};

int length1691_2(const struct Node1691_2 *head) {
    int count = 0;
    const struct Node1691_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
struct Node1691_3 {
    int value;
    struct Node1691_3 *next;
};

int length1691_3(const struct Node1691_3 *head) {
    int count = 0;
    const struct Node1691_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
enum Kind1691_4 { KIND_EMPTY1691_4 = 0, KIND_FULL1691_4 = 94 };

union Cell1691_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1691_4(union Cell1691_4 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1691_4 : KIND_EMPTY1691_4;
}
struct Node1691_5 {
    int value;
    struct Node1691_5 *next;
};

int length1691_5(const struct Node1691_5 *head) {
    int count = 0;
    const struct Node1691_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
