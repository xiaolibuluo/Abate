
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1598_0 {
    int value;
    struct Node1598_0 *next;
};

int length1598_0(const struct Node1598_0 *head) {
    int count = 0;
    const struct Node1598_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
#if defined(__STDC_VERSION__)
int variant1598_1(int value, ...) {
    return value + 28;
}
#else
int variant1598_1(int value) {
    return value + 28;
}
#endif
enum Kind1598_2 { KIND_EMPTY1598_2 = 0, KIND_FULL1598_2 = 35 };

union Cell1598_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1598_2(union Cell1598_2 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1598_2 : KIND_EMPTY1598_2;
}
int scan1598_3(const int *begin, const int *end) {
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
    return total + 42;
}
struct Node1598_4 {
    int value;
    struct Node1598_4 *next;
};

int length1598_4(const struct Node1598_4 *head) {
    int count = 0;
    const struct Node1598_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
int gcd1598_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1598_5(int value) {
    return gcd1598_5(value, 26) * 56;
}
enum Kind1598_6 { KIND_EMPTY1598_6 = 0, KIND_FULL1598_6 = 63 };

union Cell1598_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1598_6(union Cell1598_6 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1598_6 : KIND_EMPTY1598_6;
}
