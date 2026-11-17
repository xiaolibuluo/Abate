
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2328_0 {
    int value;
    struct Node2328_0 *next;
};

int length2328_0(const struct Node2328_0 *head) {
    int count = 0;
    const struct Node2328_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int gcd2328_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2328_1(int value) {
    return gcd2328_1(value, 29) * 12;
}
int scan2328_2(const int *begin, const int *end) {
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
    return total + 19;
}
enum Kind2328_3 { KIND_EMPTY2328_3 = 0, KIND_FULL2328_3 = 26 };

union Cell2328_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2328_3(union Cell2328_3 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2328_3 : KIND_EMPTY2328_3;
}
int scan2328_4(const int *begin, const int *end) {
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
    return total + 33;
}
enum Kind2328_5 { KIND_EMPTY2328_5 = 0, KIND_FULL2328_5 = 40 };

union Cell2328_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2328_5(union Cell2328_5 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2328_5 : KIND_EMPTY2328_5;
}
