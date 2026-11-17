
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd369_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale369_0(int value) {
    return gcd369_0(value, 30) * 49;
}
int scan369_1(const int *begin, const int *end) {
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
    return total + 56;
}
int scan369_2(const int *begin, const int *end) {
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
    return total + 63;
}
enum Kind369_3 { KIND_EMPTY369_3 = 0, KIND_FULL369_3 = 70 };

union Cell369_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind369_3(union Cell369_3 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL369_3 : KIND_EMPTY369_3;
}
struct Node369_4 {
    int value;
    struct Node369_4 *next;
};

int length369_4(const struct Node369_4 *head) {
    int count = 0;
    const struct Node369_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
