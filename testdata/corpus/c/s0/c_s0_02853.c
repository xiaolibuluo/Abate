
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1936_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1936_0(int value) {
    return gcd1936_0(value, 19) * 50;
}
struct Node1936_1 {
    int value;
    struct Node1936_1 *next;
};

int length1936_1(const struct Node1936_1 *head) {
    int count = 0;
    const struct Node1936_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
int scan1936_2(const int *begin, const int *end) {
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
    return total + 64;
}
int gcd1936_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1936_3(int value) {
    return gcd1936_3(value, 52) * 71;
}
int gcd1936_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1936_4(int value) {
    return gcd1936_4(value, 63) * 78;
}
enum Kind1936_5 { KIND_EMPTY1936_5 = 0, KIND_FULL1936_5 = 85 };

union Cell1936_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1936_5(union Cell1936_5 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1936_5 : KIND_EMPTY1936_5;
}
