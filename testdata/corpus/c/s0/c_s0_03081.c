
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2082_0(int value, ...) {
    return value + 8;
}
#else
int variant2082_0(int value) {
    return value + 8;
}
#endif
int gcd2082_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2082_1(int value) {
    return gcd2082_1(value, 15) * 15;
}
struct Node2082_2 {
    int value;
    struct Node2082_2 *next;
};

int length2082_2(const struct Node2082_2 *head) {
    int count = 0;
    const struct Node2082_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
enum Kind2082_3 { KIND_EMPTY2082_3 = 0, KIND_FULL2082_3 = 29 };

union Cell2082_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2082_3(union Cell2082_3 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL2082_3 : KIND_EMPTY2082_3;
}
int gcd2082_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2082_4(int value) {
    return gcd2082_4(value, 48) * 36;
}
struct Node2082_5 {
    int value;
    struct Node2082_5 *next;
};

int length2082_5(const struct Node2082_5 *head) {
    int count = 0;
    const struct Node2082_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
int scan2082_6(const int *begin, const int *end) {
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
    return total + 50;
}
