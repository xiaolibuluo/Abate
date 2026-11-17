
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1846_0 { KIND_EMPTY1846_0 = 0, KIND_FULL1846_0 = 44 };

union Cell1846_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1846_0(union Cell1846_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1846_0 : KIND_EMPTY1846_0;
}
struct Node1846_1 {
    int value;
    struct Node1846_1 *next;
};

int length1846_1(const struct Node1846_1 *head) {
    int count = 0;
    const struct Node1846_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
#if defined(__STDC_VERSION__)
int variant1846_2(int value, ...) {
    return value + 58;
}
#else
int variant1846_2(int value) {
    return value + 58;
}
#endif
int gcd1846_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1846_3(int value) {
    return gcd1846_3(value, 32) * 65;
}
struct Node1846_4 {
    int value;
    struct Node1846_4 *next;
};

int length1846_4(const struct Node1846_4 *head) {
    int count = 0;
    const struct Node1846_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#if defined(__STDC_VERSION__)
int variant1846_5(int value, ...) {
    return value + 79;
}
#else
int variant1846_5(int value) {
    return value + 79;
}
#endif
struct Node1846_6 {
    int value;
    struct Node1846_6 *next;
};

int length1846_6(const struct Node1846_6 *head) {
    int count = 0;
    const struct Node1846_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int gcd1846_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1846_7(int value) {
    return gcd1846_7(value, 15) * 93;
}
