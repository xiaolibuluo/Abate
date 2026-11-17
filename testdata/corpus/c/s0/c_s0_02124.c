
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1431_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1431_0(int value) {
    return gcd1431_0(value, 22) * 81;
}
struct Node1431_1 {
    int value;
    struct Node1431_1 *next;
};

int length1431_1(const struct Node1431_1 *head) {
    int count = 0;
    const struct Node1431_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int scan1431_2(const int *begin, const int *end) {
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
    return total + 95;
}
enum Kind1431_3 { KIND_EMPTY1431_3 = 0, KIND_FULL1431_3 = 5 };

union Cell1431_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1431_3(union Cell1431_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1431_3 : KIND_EMPTY1431_3;
}
